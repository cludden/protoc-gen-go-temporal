package plugin

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	j "github.com/dave/jennifer/jen"
	"github.com/hako/durafmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// genActivitiesInterface generates an <Service>Activities interface
func (m *Manifest) genActivitiesInterface(f *j.File) {
	iface := m.Names().activitiesIface(m.Service)
	f.Commentf("%s describes available worker activities", iface)
	f.Type().Id(iface).InterfaceFunc(func(g *j.Group) {
		for _, activity := range m.activitiesOrdered {
			if m.methods[activity].Desc.Parent() != m.Service.Desc {
				continue
			}
			method := m.methods[activity]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			commentWithDefaultf(g, methodSet(method), "%s implements a(n) %s activity definition", activity, m.fqnForActivity(activity))
			g.Id(m.methods[activity].GoName).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					if hasInput {
						g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(g *j.Group) {
					if hasOutput {
						g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					g.Error()
				}).
				Line()
		}
	})
}

// genActivityFunction generates a public <Activity>[Local] function
func (m *Manifest) genActivityFunction(f *j.File, activity protoreflect.FullName, local, async bool) {
	method := m.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	methodName := m.methods[activity].GoName
	var annotations []string
	if local {
		methodName = m.toCamel("%sLocal", methodName)
		annotations = append(annotations, "locally")
	}
	if async {
		methodName = m.toCamel("%sAsync", methodName)
		annotations = append(annotations, "asynchronously")
	}
	slices.Sort(annotations)

	desc := method.Comments.Leading.String()
	if desc != "" {
		desc = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
	} else {
		desc = fmt.Sprintf("%s executes a(n) %s activity", methodName, m.fqnForActivity(activity))
	}
	if len(annotations) > 0 {
		desc = fmt.Sprintf("%s (%s)", desc, strings.Join(annotations, ", "))
	}

	commentWithDefaultf(f, methodSet(method), desc)
	f.Func().
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if local {
				g.Id("options").Op("...").Op("*").Id(m.toCamel("%sLocalActivityOptions", activity))
			} else {
				g.Id("options").Op("...").Op("*").Id(m.toCamel("%sActivityOptions", activity))
			}
		}).
		ParamsFunc(func(g *j.Group) {
			if async {
				g.Op("*").Id(m.Names().activityFuture(activity))
			} else {
				if hasOutput {
					g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				g.Error()
			}
		}).
		BlockFunc(func(g *j.Group) {
			if !async {
				g.Return(
					j.Id(m.toCamel("%sAsync", methodName)).CallFunc(func(g *j.Group) {
						g.Id("ctx")
						if hasInput {
							g.Id("req")
						}
						g.Id("options").Op("...")
					}).Dot("Get").Call(j.Id("ctx")),
				)
				return
			}
			if isDeprecated(method) {
				g.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated activity detected"), j.Lit("activity"), j.Id(m.toCamel("%sActivityName", activity))).Line()
			}

			// initialize options
			if local {
				g.Var().Id("o").Op("*").Id(m.toCamel("%sLocalActivityOptions", activity))
			} else {
				g.Var().Id("o").Op("*").Id(m.toCamel("%sActivityOptions", activity))
			}
			g.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).
				Block(
					j.Id("o").Op("=").Id("options").Index(j.Lit(0)),
				).
				Else().
				BlockFunc(func(g *j.Group) {
					if local {
						g.Id("o").Op("=").Id(m.toCamel("New%sLocalActivityOptions", activity)).Call()
					} else {
						g.Id("o").Op("=").Id(m.toCamel("New%sActivityOptions", activity)).Call()
					}
				})

			// initialize activity options
			g.Var().Err().Error()
			g.If(
				j.List(j.Id("ctx"), j.Err()).Op("=").Id("o").Dot("Build").Call(j.Id("ctx")),
				j.Err().Op("!=").Nil(),
			).BlockFunc(func(g *j.Group) {
				if async {
					g.List(j.Id("errF"), j.Id("errS")).Op(":=").Qual(workflowPkg, "NewFuture").Call(j.Id("ctx"))
					g.Id("errS").Dot("SetError").Call(j.Err())
				}
				g.ReturnFunc(func(g *j.Group) {
					if async {
						g.Op("&").Id(m.Names().activityFuture(activity)).Values(
							j.Id("Future").Op(":").Id("errF"),
						)
					} else {
						if hasOutput {
							g.Nil()
						}
						g.Qual("fmt", "Errorf").Call(j.Lit("error initializing activity options: %w"), j.Err())
					}

				})
			})

			// initialize activity reference
			if local {
				g.Var().Id("activity").Any()
				g.If(j.Id("o").Dot("fn").Op("!=").Nil()).
					Block(
						j.Id("activity").Op("=").Id("o").Dot("fn"),
					).
					Else().
					Block(
						j.Id("activity").Op("=").Id(m.toCamel("%sActivityName", activity)),
					)
			} else {
				g.Id("activity").Op(":=").Id(m.toCamel("%sActivityName", activity))
			}

			g.If(j.Id("o").Dot("dc").Op("!=").Nil()).
				Block(
					j.Id("ctx").Op("=").Qual(workflowPkg, "WithDataConverter").Call(
						j.Id("ctx"),
						j.Id("o").Dot("dc"),
					),
				)

			// initialize activity future
			g.Id("future").Op(":=").Op("&").Id(m.Names().activityFuture(activity)).ValuesFunc(func(g *j.Group) {
				methodName := "ExecuteActivity"
				if local {
					methodName = "ExecuteLocalActivity"
				}

				g.Id("Future").Op(":").Qual(workflowPkg, methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("activity")
					if hasInput {
						g.Id("req")
					}
				})
			})

			g.ReturnFunc(func(g *j.Group) {
				if async {
					g.Add(j.Id("future"))
				} else {
					g.Add(j.Id("future").Dot("Get").Call(j.Id("ctx")))
				}
			})
		})
}

// genActivityFuture generates a <Activity>Future struct
func (m *Manifest) genActivityFuture(f *j.File, activity protoreflect.FullName) {
	future := m.Names().activityFuture(activity)

	f.Commentf("%s describes a(n) %s activity execution", future, m.fqnForActivity(activity))
	f.Type().Id(future).Struct(
		j.Id("Future").Qual(workflowPkg, "Future"),
	)
}

// genActivityFutureGetMethod generates a <Workflow>Future's Get method
func (m *Manifest) genActivityFutureGetMethod(f *j.File, activity protoreflect.FullName) {
	method := m.methods[activity]
	hasOutput := !isEmpty(method.Output)
	future := m.Names().activityFuture(activity)

	f.Comment("Get blocks on the activity's completion, returning the response")
	f.Func().
		Params(j.Id("f").Op("*").Id(future)).
		Id("Get").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			if hasOutput {
				g.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				g.If(
					j.Err().Op(":=").Id("f").Dot("Future").Dot("Get").Call(
						j.Id("ctx"), j.Op("&").Id("resp"),
					),
					j.Err().Op("!=").Nil(),
				).Block(
					j.Return(j.Nil(), j.Err()),
				)
				g.Return(j.Op("&").Id("resp"), j.Nil())
			} else {
				g.Return(j.Id("f").Dot("Future").Dot("Get").Call(
					j.Id("ctx"), j.Nil(),
				))
			}
		})
}

// genActivityFutureSelectMethod generates a <Workflow>Future's Select method
func (m *Manifest) genActivityFutureSelectMethod(f *j.File, activity protoreflect.FullName) {
	future := m.Names().activityFuture(activity)

	f.Comment("Select adds the activity's completion to the selector, callback can be nil")
	f.Func().
		Params(j.Id("f").Op("*").Id(future)).
		Id("Select").
		Params(
			j.Id("sel").Qual(workflowPkg, "Selector"),
			j.Id("fn").Func().Params(j.Op("*").Id(future)),
		).
		Params(
			j.Qual(workflowPkg, "Selector"),
		).
		Block(
			j.Return(
				j.Id("sel").Dot("AddFuture").Call(
					j.Id("f").Dot("Future"),
					j.Func().
						Params(j.Qual(workflowPkg, "Future")).
						Block(
							j.If(j.Id("fn").Op("!=").Nil()).Block(
								j.Id("fn").Call(j.Id("f")),
							),
						),
				),
			),
		)
}

// genActivityRegisterAllFunction generates a Register<Service>Activities public function
func (m *Manifest) genActivityRegisterAllFunction(f *j.File) {
	iface := m.Names().activitiesIface(m.Service)
	fn := fmt.Sprintf("Register%s", iface)
	f.Commentf("%s registers activities with a worker", fn)
	f.Func().Id(fn).
		Params(
			j.Id("r").Qual(workerPkg, "ActivityRegistry"),
			j.Id("activities").Id(iface),
		).
		BlockFunc(func(g *j.Group) {
			for _, activity := range m.activitiesOrdered {
				if m.methods[activity].Desc.Parent() != m.Service.Desc {
					continue
				}
				g.Id(fmt.Sprintf("Register%sActivity", m.methods[activity].GoName)).Call(
					j.Id("r"), j.Id("activities").Dot(m.methods[activity].GoName),
				)
			}
		})
}

// genActivityRegisterOneFunction generates a Register<Activity> public function
func (m *Manifest) genActivityRegisterOneFunction(f *j.File, activity protoreflect.FullName) {
	method := m.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	f.Commentf("Register%sActivity registers a %s activity", m.methods[activity].GoName, m.fqnForActivity(activity))
	f.Func().Id(fmt.Sprintf("Register%sActivity", m.methods[activity].GoName)).
		Params(
			j.Id("r").Qual(workerPkg, "ActivityRegistry"),
			j.Id("fn").Func().
				ParamsFunc(func(g *j.Group) {
					g.Qual("context", "Context")
					if hasInput {
						g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(g *j.Group) {
					if hasOutput {
						g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					g.Error()
				}),
		).
		Block(
			j.Id("r").Dot("RegisterActivityWithOptions").Call(
				j.Id("fn"), j.Qual(activityPkg, "RegisterOptions").Block(
					j.Id("Name").Op(":").Id(m.toCamel("%sActivityName", activity)).Op(","),
				),
			),
		)
}

// genActivityOptions generates an <Activity>ActivityOptions struct
func (m *Manifest) genActivityOptions(f *j.File, activity protoreflect.FullName, local bool) {
	optionType := "ActivityOptions"
	typeName := m.toCamel("%sActivityOptions", activity)
	if local {
		optionType, typeName = "LocalActivityOptions", m.toCamel("%sLocalActivityOptions", activity)
	}
	method := m.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	opts := m.activities[activity]

	// generate type definition
	f.Commentf("%s provides configuration for a(n) %s activity", typeName, m.fqnForActivity(activity))
	f.Type().Id(typeName).StructFunc(func(g *j.Group) {
		g.Id("options").Qual(workflowPkg, optionType)
		g.Id("retryPolicy").Op("*").Qual(temporalPkg, "RetryPolicy")
		g.Id("scheduleToCloseTimeout").Op("*").Qual("time", "Duration")
		g.Id("startToCloseTimeout").Op("*").Qual("time", "Duration")
		g.Id("dc").Qual(converterPkg, "DataConverter")
		if local {
			g.Id("fn").
				Func().
				ParamsFunc(func(g *j.Group) {
					g.Qual("context", "Context")
					if hasInput {
						g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(g *j.Group) {
					if hasOutput {
						g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					g.Error()
				})
		} else {
			g.Id("heartbeatTimeout").Op("*").Qual("time", "Duration")
			g.Id("scheduleToStartTimeout").Op("*").Qual("time", "Duration")
			g.Id("taskQueue").Op("*").String()
			g.Id("waitForCancellation").Op("*").Bool()
		}
	})

	// generate constructor
	ctorName := "New" + typeName
	f.Commentf("%s initializes a new %s value", ctorName, typeName)
	f.Func().
		Id(ctorName).
		Params().
		Op("*").Id(typeName).
		Block(
			j.Return(j.Op("&").Id(typeName).Values()),
		)

	// generate Build method
	f.Commentf("Build initializes a workflow.Context with appropriate %s values derived from schema defaults and any user-defined overrides", optionType)
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("Build").
		Params(
			j.Id("ctx").Qual(workflowPkg, "Context"),
		).
		Params(
			j.Qual(workflowPkg, "Context"),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			g.Id("opts").Op(":=").Id("o").Dot("options")

			// set HeartbeatTimeout
			if !local {
				heartbeatTimeout := g.If(j.Id("v").Op(":=").Id("o").Dot("heartbeatTimeout"), j.Id("v").Op("!=").Nil()).
					Block(
						j.Id("opts").Dot("HeartbeatTimeout").Op("=").Op("*").Id("v"),
					)
				if d := opts.GetHeartbeatTimeout().AsDuration(); d > 0 {
					heartbeatTimeout.Else().If(j.Id("opts").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
						j.Id("opts").Dot("HeartbeatTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
					)
				}
			}

			// set RetryPolicy
			retryPolicy := g.If(j.Id("v").Op(":=").Id("o").Dot("retryPolicy"), j.Id("v").Op("!=").Nil()).
				Block(
					j.Id("opts").Dot("RetryPolicy").Op("=").Id("v"),
				)
			if policy := opts.GetRetryPolicy(); policy != nil {
				retryPolicy.Else().If(j.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
					j.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").ValuesFunc(func(g *j.Group) {
						if d := policy.GetInitialInterval(); d.IsValid() {
							g.Id("InitialInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
						}
						if d := policy.GetMaxInterval(); d.IsValid() {
							g.Id("MaximumInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
						}
						if n := policy.GetBackoffCoefficient(); n != 0 {
							g.Id("BackoffCoefficient").Op(":").Lit(n)
						}
						if n := policy.GetMaxAttempts(); n != 0 {
							g.Id("MaximumAttempts").Op(":").Lit(n)
						}
						if errs := policy.GetNonRetryableErrorTypes(); len(errs) > 0 {
							g.Id("NonRetryableErrorTypes").Op(":").Index().String().CustomFunc(multiLineValues, func(g *j.Group) {
								for _, err := range errs {
									g.Lit(err)
								}
							})
						}
					}),
				)
			}

			// set ScheduleToCloseTimeout
			scheduleToCloseTimeout := g.If(j.Id("v").Op(":=").Id("o").Dot("scheduleToCloseTimeout"), j.Id("v").Op("!=").Nil()).
				Block(
					j.Id("opts").Dot("ScheduleToCloseTimeout").Op("=").Op("*").Id("v"),
				)
			if d := opts.GetScheduleToCloseTimeout().AsDuration(); d > 0 {
				scheduleToCloseTimeout.Else().If(j.Id("opts").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set ScheduleToStartTimeout
			if !local {
				scheduleToStartTimeout := g.If(j.Id("v").Op(":=").Id("o").Dot("scheduleToStartTimeout"), j.Id("v").Op("!=").Nil()).
					Block(
						j.Id("opts").Dot("ScheduleToStartTimeout").Op("=").Op("*").Id("v"),
					)
				if d := opts.GetScheduleToStartTimeout().AsDuration(); d > 0 {
					scheduleToStartTimeout.Else().If(j.Id("opts").Dot("ScheduleToStartTimeout").Op("==").Lit(0)).Block(
						j.Id("opts").Dot("ScheduleToStartTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
					)
				}
			}

			// set StartToCloseTimeout
			startToCloseTimeout := g.If(j.Id("v").Op(":=").Id("o").Dot("startToCloseTimeout"), j.Id("v").Op("!=").Nil()).
				Block(
					j.Id("opts").Dot("StartToCloseTimeout").Op("=").Op("*").Id("v"),
				)
			if d := opts.GetStartToCloseTimeout().AsDuration(); d > 0 {
				startToCloseTimeout.Else().If(j.Id("opts").Dot("StartToCloseTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("StartToCloseTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set TaskQueue
			if !local {
				g.If(j.Id("v").Op(":=").Id("o").Dot("taskQueue"), j.Id("v").Op("!=").Nil()).
					Block(
						j.Id("opts").Dot("TaskQueue").Op("=").Op("*").Id("v"),
					).Else().
					If(j.Id("opts").Dot("TaskQueue").Op("==").Lit("")).
					BlockFunc(func(g *j.Group) {
						var taskQueue j.Code
						if tq := opts.GetTaskQueue(); tq != "" {
							taskQueue = j.Lit(tq)
						}
						if tq := m.opts.GetTaskQueue(); taskQueue == nil && tq != "" {
							taskQueue = j.Id(m.toCamel("%sTaskQueue", m.GoName))
						}
						if taskQueue != nil {
							g.Id("opts").Dot("TaskQueue").Op("=").Add(taskQueue)
						} else {
							g.Id("opts").Dot("TaskQueue").Op("=").Qual(workflowPkg, "GetInfo").Call(j.Id("ctx")).Dot("TaskQueueName")
						}
					})
			}

			// set WaitForCancellation
			if !local {
				waitForCancellation := g.If(j.Id("v").Op(":=").Id("o").Dot("waitForCancellation"), j.Id("v").Op("!=").Nil()).
					Block(
						j.Id("opts").Dot("WaitForCancellation").Op("=").Op("*").Id("v"),
					)
				if opts.GetWaitForCancellation() {
					waitForCancellation.Else().If(j.Op("!").Id("opts").Dot("WaitForCancellation")).Block(
						j.Id("opts").Dot("WaitForCancellation").Op("=").Lit(opts.GetWaitForCancellation()),
					)
				}
			}

			g.Return(j.Qual(workflowPkg, fmt.Sprintf("With%s", optionType)).Call(j.Id("ctx"), j.Id("opts")), j.Nil())
		})

	if local {
		f.Commentf("Local specifies a custom %s implementation", m.fqnForActivity(activity))
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("Local").
			Params(
				j.Id("fn").
					Func().
					ParamsFunc(func(g *j.Group) {
						g.Qual("context", "Context")
						if hasInput {
							g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(g *j.Group) {
						if hasOutput {
							g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
						}
						g.Error()
					}),
			).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("fn").Op("=").Id("fn"),
				j.Return(j.Id("o")),
			)
	}

	f.Commentf("%s specifies an initial %s value to which defaults will be applied", fmt.Sprintf("With%s", optionType), optionType)
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id(fmt.Sprintf("With%s", optionType)).
		Params(
			j.Id("options").Qual(workflowPkg, optionType),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("options").Op("=").Id("options"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithDataConverter registers a DataConverter for the (local) activity")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithDataConverter").
		Params(
			j.Id("dc").Qual(converterPkg, "DataConverter"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("dc").Op("=").Id("dc"),
			j.Return(j.Id("o")),
		)

	if !local {
		f.Comment("WithHeartbeatTimeout sets the HeartbeatTimeout value")
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("WithHeartbeatTimeout").
			Params(
				j.Id("d").Qual("time", "Duration"),
			).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("heartbeatTimeout").Op("=").Op("&").Id("d"),
				j.Return(j.Id("o")),
			)
	}

	f.Comment("WithRetryPolicy sets the RetryPolicy value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithRetryPolicy").
		Params(
			j.Id("policy").Op("*").Qual(temporalPkg, "RetryPolicy"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("retryPolicy").Op("=").Id("policy"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithScheduleToCloseTimeout sets the ScheduleToCloseTimeout value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithScheduleToCloseTimeout").
		Params(
			j.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("scheduleToCloseTimeout").Op("=").Op("&").Id("d"),
			j.Return(j.Id("o")),
		)

	if !local {
		f.Comment("WithScheduleToStartTimeout sets the ScheduleToStartTimeout value")
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("WithScheduleToStartTimeout").
			Params(
				j.Id("d").Qual("time", "Duration"),
			).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("scheduleToStartTimeout").Op("=").Op("&").Id("d"),
				j.Return(j.Id("o")),
			)
	}

	f.Comment("WithStartToCloseTimeout sets the StartToCloseTimeout value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithStartToCloseTimeout").
		Params(
			j.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("startToCloseTimeout").Op("=").Op("&").Id("d"),
			j.Return(j.Id("o")),
		)

	if !local {
		f.Comment("WithTaskQueue sets the TaskQueue value")
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("WithTaskQueue").
			Params(
				j.Id("tq").String(),
			).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("taskQueue").Op("=").Op("&").Id("tq"),
				j.Return(j.Id("o")),
			)
	}

	if !local {
		f.Comment("WithWaitForCancellation sets the WaitForCancellation value")
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("WithWaitForCancellation").
			Params(
				j.Id("wait").Bool(),
			).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("waitForCancellation").Op("=").Op("&").Id("wait"),
				j.Return(j.Id("o")),
			)
	}
}

func (n *names) activitiesIface(service *protogen.Service) string {
	return n.caser.ToCamel(fmt.Sprintf("%sActivities", service.GoName))
}

func (n *names) activityFuture(activity protoreflect.FullName) string {
	return n.toCamel("%sFuture", activity)
}
