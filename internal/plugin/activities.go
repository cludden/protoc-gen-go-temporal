package plugin

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	g "github.com/dave/jennifer/jen"
	"github.com/hako/durafmt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// genActivitiesInterface generates an <Service>Activities interface
func (svc *Manifest) genActivitiesInterface(f *g.File) {
	f.Commentf("%sActivities describes available worker activities", svc.Service.GoName)
	f.Type().Id(fmt.Sprintf("%sActivities", svc.Service.GoName)).InterfaceFunc(func(methods *g.Group) {
		for _, activity := range svc.activitiesOrdered {
			if svc.methods[activity].Desc.Parent() != svc.Service.Desc {
				continue
			}
			method := svc.methods[activity]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			commentWithDefaultf(methods, methodSet(method), "%s implements a(n) %s activity definition", activity, svc.fqnForActivity(activity))
			methods.Id(svc.methods[activity].GoName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
					}
					returnVals.Error()
				}).
				Line()
		}
	})
}

// genActivityFunction generates a public <Activity>[Local] function
func (svc *Manifest) genActivityFunction(f *g.File, activity protoreflect.FullName, local, async bool) {
	method := svc.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	methodName := svc.methods[activity].GoName
	var annotations []string
	if local {
		methodName = svc.toCamel("%sLocal", methodName)
		annotations = append(annotations, "locally")
	}
	if async {
		methodName = svc.toCamel("%sAsync", methodName)
		annotations = append(annotations, "asynchronously")
	}
	sort.Slice(annotations, func(i, j int) bool {
		return annotations[i] < annotations[j]
	})

	desc := method.Comments.Leading.String()
	if desc != "" {
		desc = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
	} else {
		desc = fmt.Sprintf("%s executes a(n) %s activity", methodName, svc.fqnForActivity(activity))
	}
	if len(annotations) > 0 {
		desc = fmt.Sprintf("%s (%s)", desc, strings.Join(annotations, ", "))
	}

	commentWithDefaultf(f, methodSet(method), desc)
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			if local {
				args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sLocalActivityOptions", activity))
			} else {
				args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sActivityOptions", activity))
			}
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if async {
				returnVals.Op("*").Id(fmt.Sprintf("%sFuture", method.GoName))
			} else {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			}
		}).
		BlockFunc(func(fn *g.Group) {
			if !async {
				fn.Return(
					g.Id(svc.toCamel("%sAsync", methodName)).CallFunc(func(args *g.Group) {
						args.Id("ctx")
						if hasInput {
							args.Id("req")
						}
						args.Id("options").Op("...")
					}).Dot("Get").Call(g.Id("ctx")),
				)
				return
			}
			if isDeprecated(method) {
				fn.Qual(workflowPkg, "GetLogger").Call(g.Id("ctx")).Dot("Warn").Call(g.Lit("use of deprecated activity detected"), g.Lit("activity"), g.Id(svc.toCamel("%sActivityName", activity))).Line()
			}

			// initialize options
			if local {
				fn.Var().Id("o").Op("*").Id(svc.toCamel("%sLocalActivityOptions", activity))
			} else {
				fn.Var().Id("o").Op("*").Id(svc.toCamel("%sActivityOptions", activity))
			}
			fn.If(g.Len(g.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(g.Lit(0)).Op("!=").Nil()).
				Block(
					g.Id("o").Op("=").Id("options").Index(g.Lit(0)),
				).
				Else().
				BlockFunc(func(bl *g.Group) {
					if local {
						bl.Id("o").Op("=").Id(svc.toCamel("New%sLocalActivityOptions", activity)).Call()
					} else {
						bl.Id("o").Op("=").Id(svc.toCamel("New%sActivityOptions", activity)).Call()
					}
				})

			// initialize activity options
			fn.Var().Err().Error()
			fn.If(
				g.List(g.Id("ctx"), g.Err()).Op("=").Id("o").Dot("Build").Call(g.Id("ctx")),
				g.Err().Op("!=").Nil(),
			).BlockFunc(func(bl *g.Group) {
				if async {
					bl.List(g.Id("errF"), g.Id("errS")).Op(":=").Qual(workflowPkg, "NewFuture").Call(g.Id("ctx"))
					bl.Id("errS").Dot("SetError").Call(g.Err())
				}
				bl.ReturnFunc(func(returnVals *g.Group) {
					if async {
						returnVals.Op("&").Id(svc.toCamel("%sFuture", activity)).Values(
							g.Id("Future").Op(":").Id("errF"),
						)
					} else {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Qual("fmt", "Errorf").Call(g.Lit("error initializing activity options: %w"), g.Err())
					}

				})
			})

			// initialize activity reference
			if local {
				fn.Var().Id("activity").Any()
				fn.If(g.Id("o").Dot("fn").Op("!=").Nil()).
					Block(
						g.Id("activity").Op("=").Id("o").Dot("fn"),
					).
					Else().
					Block(
						g.Id("activity").Op("=").Id(svc.toCamel("%sActivityName", activity)),
					)
			} else {
				fn.Id("activity").Op(":=").Id(svc.toCamel("%sActivityName", activity))
			}

			// initialize activity future
			fn.Id("future").Op(":=").Op("&").Id(svc.toCamel("%sFuture", activity)).ValuesFunc(func(values *g.Group) {
				methodName := "ExecuteActivity"
				if local {
					methodName = "ExecuteLocalActivity"
				}

				values.Id("Future").Op(":").Qual(workflowPkg, methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("activity")
					if hasInput {
						args.Id("req")
					}
				})
			})

			fn.ReturnFunc(func(returnVals *g.Group) {
				if async {
					returnVals.Add(g.Id("future"))
				} else {
					returnVals.Add(g.Id("future").Dot("Get").Call(g.Id("ctx")))
				}
			})
		})
}

// genActivityFuture generates a <Activity>Future struct
func (svc *Manifest) genActivityFuture(f *g.File, activity protoreflect.FullName) {
	future := svc.toCamel("%sFuture", activity)

	f.Commentf("%s describes a(n) %s activity execution", future, svc.fqnForActivity(activity))
	f.Type().Id(future).Struct(
		g.Id("Future").Qual(workflowPkg, "Future"),
	)
}

// genActivityFutureGetMethod generates a <Workflow>Future's Get method
func (svc *Manifest) genActivityFutureGetMethod(f *g.File, activity protoreflect.FullName) {
	method := svc.methods[activity]
	hasOutput := !isEmpty(method.Output)
	future := svc.toCamel("%sFuture", activity)

	f.Comment("Get blocks on the activity's completion, returning the response")
	f.Func().
		Params(g.Id("f").Op("*").Id(future)).
		Id("Get").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				fn.If(
					g.Err().Op(":=").Id("f").Dot("Future").Dot("Get").Call(
						g.Id("ctx"), g.Op("&").Id("resp"),
					),
					g.Err().Op("!=").Nil(),
				).Block(
					g.Return(g.Nil(), g.Err()),
				)
				fn.Return(g.Op("&").Id("resp"), g.Nil())
			} else {
				fn.Return(g.Id("f").Dot("Future").Dot("Get").Call(
					g.Id("ctx"), g.Nil(),
				))
			}
		})
}

// genActivityFutureSelectMethod generates a <Workflow>Future's Select method
func (svc *Manifest) genActivityFutureSelectMethod(f *g.File, activity protoreflect.FullName) {
	future := svc.toCamel("%sFuture", activity)

	f.Comment("Select adds the activity's completion to the selector, callback can be nil")
	f.Func().
		Params(g.Id("f").Op("*").Id(future)).
		Id("Select").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().Params(g.Op("*").Id(future)),
		).
		Params(
			g.Qual(workflowPkg, "Selector"),
		).
		Block(
			g.Return(
				g.Id("sel").Dot("AddFuture").Call(
					g.Id("f").Dot("Future"),
					g.Func().
						Params(g.Qual(workflowPkg, "Future")).
						Block(
							g.If(g.Id("fn").Op("!=").Nil()).Block(
								g.Id("fn").Call(g.Id("f")),
							),
						),
				),
			),
		)
}

// genActivityRegisterAllFunction generates a Register<Service>Activities public function
func (svc *Manifest) genActivityRegisterAllFunction(f *g.File) {
	f.Commentf("Register%sActivities registers activities with a worker", svc.Service.GoName)
	f.Func().Id(fmt.Sprintf("Register%sActivities", svc.Service.GoName)).
		Params(
			g.Id("r").Qual(workerPkg, "ActivityRegistry"),
			g.Id("activities").Id(svc.toCamel("%sActivities", svc.Service.GoName)),
		).
		BlockFunc(func(fn *g.Group) {
			for _, activity := range svc.activitiesOrdered {
				if svc.methods[activity].Desc.Parent() != svc.Service.Desc {
					continue
				}
				fn.Id(fmt.Sprintf("Register%sActivity", svc.methods[activity].GoName)).Call(
					g.Id("r"), g.Id("activities").Dot(svc.methods[activity].GoName),
				)
			}
		})
}

// genActivityRegisterOneFunction generates a Register<Activity> public function
func (svc *Manifest) genActivityRegisterOneFunction(f *g.File, activity protoreflect.FullName) {
	method := svc.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	f.Commentf("Register%sActivity registers a %s activity", svc.methods[activity].GoName, svc.fqnForActivity(activity))
	f.Func().Id(fmt.Sprintf("Register%sActivity", svc.methods[activity].GoName)).
		Params(
			g.Id("r").Qual(workerPkg, "ActivityRegistry"),
			g.Id("fn").Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual("context", "Context")
					if hasInput {
						args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
					}
					returnVals.Error()
				}),
		).
		Block(
			g.Id("r").Dot("RegisterActivityWithOptions").Call(
				g.Id("fn"), g.Qual(activityPkg, "RegisterOptions").Block(
					g.Id("Name").Op(":").Id(svc.toCamel("%sActivityName", activity)).Op(","),
				),
			),
		)
}

// genActivityOptions generates an <Activity>ActivityOptions struct
func (svc *Manifest) genActivityOptions(f *g.File, activity protoreflect.FullName, local bool) {
	optionType := "ActivityOptions"
	typeName := svc.toCamel("%sActivityOptions", activity)
	if local {
		optionType, typeName = "LocalActivityOptions", svc.toCamel("%sLocalActivityOptions", activity)
	}
	method := svc.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	opts := svc.activities[activity]

	// generate type definition
	f.Commentf("%s provides configuration for a(n) %s activity", typeName, svc.fqnForActivity(activity))
	f.Type().Id(typeName).StructFunc(func(values *g.Group) {
		values.Id("options").Qual(workflowPkg, optionType)
		values.Id("retryPolicy").Op("*").Qual(temporalPkg, "RetryPolicy")
		values.Id("scheduleToCloseTimeout").Op("*").Qual("time", "Duration")
		values.Id("startToCloseTimeout").Op("*").Qual("time", "Duration")
		if local {
			values.Id("fn").
				Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual("context", "Context")
					if hasInput {
						args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
					}
					returnVals.Error()
				})
		} else {
			values.Id("heartbeatTimeout").Op("*").Qual("time", "Duration")
			values.Id("scheduleToStartTimeout").Op("*").Qual("time", "Duration")
			values.Id("taskQueue").Op("*").String()
			values.Id("waitForCancellation").Op("*").Bool()
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
			g.Return(g.Op("&").Id(typeName).Values()),
		)

	// generate Build method
	f.Commentf("Build initializes a workflow.Context with appropriate %s values derived from schema defaults and any user-defined overrides", optionType)
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("Build").
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
		).
		Params(
			g.Qual(workflowPkg, "Context"),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			fn.Id("opts").Op(":=").Id("o").Dot("options")

			// set HeartbeatTimeout
			if !local {
				heartbeatTimeout := fn.If(g.Id("v").Op(":=").Id("o").Dot("heartbeatTimeout"), g.Id("v").Op("!=").Nil()).
					Block(
						g.Id("opts").Dot("HeartbeatTimeout").Op("=").Op("*").Id("v"),
					)
				if d := opts.GetHeartbeatTimeout().AsDuration(); d > 0 {
					heartbeatTimeout.Else().If(g.Id("opts").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
						g.Id("opts").Dot("HeartbeatTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
					)
				}
			}

			// set RetryPolicy
			retryPolicy := fn.If(g.Id("v").Op(":=").Id("o").Dot("retryPolicy"), g.Id("v").Op("!=").Nil()).
				Block(
					g.Id("opts").Dot("RetryPolicy").Op("=").Id("v"),
				)
			if policy := opts.GetRetryPolicy(); policy != nil {
				retryPolicy.Else().If(g.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
					g.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").ValuesFunc(func(fields *g.Group) {
						if d := policy.GetInitialInterval(); d.IsValid() {
							fields.Id("InitialInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
						}
						if d := policy.GetMaxInterval(); d.IsValid() {
							fields.Id("MaximumInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
						}
						if n := policy.GetBackoffCoefficient(); n != 0 {
							fields.Id("BackoffCoefficient").Op(":").Lit(n)
						}
						if n := policy.GetMaxAttempts(); n != 0 {
							fields.Id("MaximumAttempts").Op(":").Lit(n)
						}
						if errs := policy.GetNonRetryableErrorTypes(); len(errs) > 0 {
							fields.Id("NonRetryableErrorTypes").Op(":").Index().String().CustomFunc(multiLineValues, func(vals *g.Group) {
								for _, err := range errs {
									vals.Lit(err)
								}
							})
						}
					}),
				)
			}

			// set ScheduleToCloseTimeout
			scheduleToCloseTimeout := fn.If(g.Id("v").Op(":=").Id("o").Dot("scheduleToCloseTimeout"), g.Id("v").Op("!=").Nil()).
				Block(
					g.Id("opts").Dot("ScheduleToCloseTimeout").Op("=").Op("*").Id("v"),
				)
			if d := opts.GetScheduleToCloseTimeout().AsDuration(); d > 0 {
				scheduleToCloseTimeout.Else().If(g.Id("opts").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set ScheduleToStartTimeout
			if !local {
				scheduleToStartTimeout := fn.If(g.Id("v").Op(":=").Id("o").Dot("scheduleToStartTimeout"), g.Id("v").Op("!=").Nil()).
					Block(
						g.Id("opts").Dot("ScheduleToStartTimeout").Op("=").Op("*").Id("v"),
					)
				if d := opts.GetScheduleToStartTimeout().AsDuration(); d > 0 {
					scheduleToStartTimeout.Else().If(g.Id("opts").Dot("ScheduleToStartTimeout").Op("==").Lit(0)).Block(
						g.Id("opts").Dot("ScheduleToStartTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
					)
				}
			}

			// set StartToCloseTimeout
			startToCloseTimeout := fn.If(g.Id("v").Op(":=").Id("o").Dot("startToCloseTimeout"), g.Id("v").Op("!=").Nil()).
				Block(
					g.Id("opts").Dot("StartToCloseTimeout").Op("=").Op("*").Id("v"),
				)
			if d := opts.GetStartToCloseTimeout().AsDuration(); d > 0 {
				startToCloseTimeout.Else().If(g.Id("opts").Dot("StartToCloseTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("StartToCloseTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set TaskQueue
			if !local {
				fn.If(g.Id("v").Op(":=").Id("o").Dot("taskQueue"), g.Id("v").Op("!=").Nil()).
					Block(
						g.Id("opts").Dot("TaskQueue").Op("=").Op("*").Id("v"),
					).Else().
					If(g.Id("opts").Dot("TaskQueue").Op("==").Lit("")).
					BlockFunc(func(bl *g.Group) {
						var taskQueue g.Code
						if tq := opts.GetTaskQueue(); tq != "" {
							taskQueue = g.Lit(tq)
						}
						if tq := svc.opts.GetTaskQueue(); taskQueue == nil && tq != "" {
							taskQueue = g.Id(svc.toCamel("%sTaskQueue", svc.GoName))
						}
						if taskQueue != nil {
							bl.Id("opts").Dot("TaskQueue").Op("=").Add(taskQueue)
						} else {
							bl.Id("opts").Dot("TaskQueue").Op("=").Qual(workflowPkg, "GetInfo").Call(g.Id("ctx")).Dot("TaskQueueName")
						}
					})
			}

			// set WaitForCancellation
			if !local {
				waitForCancellation := fn.If(g.Id("v").Op(":=").Id("o").Dot("waitForCancellation"), g.Id("v").Op("!=").Nil()).
					Block(
						g.Id("opts").Dot("WaitForCancellation").Op("=").Op("*").Id("v"),
					)
				if opts.GetWaitForCancellation() {
					waitForCancellation.Else().If(g.Op("!").Id("opts").Dot("WaitForCancellation")).Block(
						g.Id("opts").Dot("WaitForCancellation").Op("=").Lit(opts.GetWaitForCancellation()),
					)
				}
			}

			fn.Return(g.Qual(workflowPkg, fmt.Sprintf("With%s", optionType)).Call(g.Id("ctx"), g.Id("opts")), g.Nil())
		})

	if local {
		f.Commentf("Local specifies a custom %s implementation", svc.fqnForActivity(activity))
		f.Func().
			Params(g.Id("o").Op("*").Id(typeName)).
			Id("Local").
			Params(
				g.Id("fn").
					Func().
					ParamsFunc(func(args *g.Group) {
						args.Qual("context", "Context")
						if hasInput {
							args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
						}
						returnVals.Error()
					}),
			).
			Op("*").Id(typeName).
			Block(
				g.Id("o").Dot("fn").Op("=").Id("fn"),
				g.Return(g.Id("o")),
			)
	}

	f.Commentf("%s specifies an initial %s value to which defaults will be applied", fmt.Sprintf("With%s", optionType), optionType)
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id(fmt.Sprintf("With%s", optionType)).
		Params(
			g.Id("options").Qual(workflowPkg, optionType),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("options").Op("=").Id("options"),
			g.Return(g.Id("o")),
		)

	if !local {
		f.Comment("WithHeartbeatTimeout sets the HeartbeatTimeout value")
		f.Func().
			Params(g.Id("o").Op("*").Id(typeName)).
			Id("WithHeartbeatTimeout").
			Params(
				g.Id("d").Qual("time", "Duration"),
			).
			Op("*").Id(typeName).
			Block(
				g.Id("o").Dot("heartbeatTimeout").Op("=").Op("&").Id("d"),
				g.Return(g.Id("o")),
			)
	}

	f.Comment("WithRetryPolicy sets the RetryPolicy value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithRetryPolicy").
		Params(
			g.Id("policy").Op("*").Qual(temporalPkg, "RetryPolicy"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("retryPolicy").Op("=").Id("policy"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithScheduleToCloseTimeout sets the ScheduleToCloseTimeout value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithScheduleToCloseTimeout").
		Params(
			g.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("scheduleToCloseTimeout").Op("=").Op("&").Id("d"),
			g.Return(g.Id("o")),
		)

	if !local {
		f.Comment("WithScheduleToStartTimeout sets the ScheduleToStartTimeout value")
		f.Func().
			Params(g.Id("o").Op("*").Id(typeName)).
			Id("WithScheduleToStartTimeout").
			Params(
				g.Id("d").Qual("time", "Duration"),
			).
			Op("*").Id(typeName).
			Block(
				g.Id("o").Dot("scheduleToStartTimeout").Op("=").Op("&").Id("d"),
				g.Return(g.Id("o")),
			)
	}

	f.Comment("WithStartToCloseTimeout sets the StartToCloseTimeout value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithStartToCloseTimeout").
		Params(
			g.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("startToCloseTimeout").Op("=").Op("&").Id("d"),
			g.Return(g.Id("o")),
		)

	if !local {
		f.Comment("WithTaskQueue sets the TaskQueue value")
		f.Func().
			Params(g.Id("o").Op("*").Id(typeName)).
			Id("WithTaskQueue").
			Params(
				g.Id("tq").String(),
			).
			Op("*").Id(typeName).
			Block(
				g.Id("o").Dot("taskQueue").Op("=").Op("&").Id("tq"),
				g.Return(g.Id("o")),
			)
	}

	if !local {
		f.Comment("WithWaitForCancellation sets the WaitForCancellation value")
		f.Func().
			Params(g.Id("o").Op("*").Id(typeName)).
			Id("WithWaitForCancellation").
			Params(
				g.Id("wait").Bool(),
			).
			Op("*").Id(typeName).
			Block(
				g.Id("o").Dot("waitForCancellation").Op("=").Op("&").Id("wait"),
				g.Return(g.Id("o")),
			)
	}
}
