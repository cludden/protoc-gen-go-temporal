package plugin

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	g "github.com/dave/jennifer/jen"
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
			commentf(methods, methodSet(method), "%s implements a(n) %s activity definition", activity, svc.fqnForActivity(activity))
			methods.Id(svc.methods[activity].GoName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
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
	opts := svc.activities[activity]
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

	commentf(f, methodSet(method), desc)
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
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
					returnVals.Op("*").Id(method.Output.GoIdent.GoName)
				}
				returnVals.Error()
			}
		}).
		BlockFunc(func(fn *g.Group) {
			if isDeprecated(method) {
				fn.Qual(workflowPkg, "GetLogger").Call(g.Id("ctx")).Dot("Warn").Call(g.Lit("use of deprecated activity detected"), g.Lit("activity"), g.Id(svc.toCamel("%sActivityName", activity))).Line()
			}
			// initialize activity options if nil
			if local {
				fn.Var().Id("opts").Op("*").Id(svc.toCamel("%sLocalActivityOptions", activity))
			} else {
				fn.Var().Id("opts").Op("*").Id(svc.toCamel("%sActivityOptions", activity))
			}

			// initialize opts
			fn.
				If(
					g.Len(g.Id("options")).Op(">").Lit(0).
						Op("&&").Id("options").Index(g.Lit(0)).Op("!=").Nil(),
				).
				Block(
					g.Id("opts").Op("=").Id("options").Index(g.Lit(0)),
				).
				Else().
				BlockFunc(func(bl *g.Group) {
					if local {
						bl.Id("opts").Op("=").Id(svc.toCamel("New%sLocalActivityOptions", activity)).Call()
					} else {
						bl.Id("opts").Op("=").Id(svc.toCamel("New%sActivityOptions", activity)).Call()
					}
				})

			// initialize activity options
			fn.If(g.Id("opts").Dot("opts").Op("==").Nil()).BlockFunc(func(bl *g.Group) {
				if local {
					bl.Id("activityOpts").Op(":=").Qual(workflowPkg, "GetLocalActivityOptions").Call(g.Id("ctx"))
				} else {
					bl.Id("activityOpts").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(g.Id("ctx"))
				}
				bl.Id("opts").Dot("opts").Op("=").Op("&").Id("activityOpts")
			})

			// set default task queue
			if tq := opts.GetTaskQueue(); !local && tq != "" {
				fn.If(g.Id("opts").Dot("opts").Dot("TaskQueue").Op("==").Lit("")).Block(
					g.Id("opts").Dot("opts").Dot("TaskQueue").Op("=").Lit(tq),
				)
			}

			// set default retry policy
			if policy := opts.GetRetryPolicy(); policy != nil {
				fn.If(g.Id("opts").Dot("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
					g.Id("opts").Dot("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").ValuesFunc(func(fields *g.Group) {
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
							fields.Id("NonRetryableErrorTypes").Op(":").Lit(errs)
						}
					}),
				)
			}

			// set default heartbeat timeout
			if timeout := opts.GetHeartbeatTimeout(); !local && timeout.IsValid() {
				fn.If(g.Id("opts").Dot("opts").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("opts").Dot("HeartbeatTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// set default schedule to close timeout
			if timeout := opts.GetScheduleToCloseTimeout(); timeout.IsValid() {
				fn.If(g.Id("opts").Dot("opts").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("opts").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// set default schedule to start timeout
			if timeout := opts.GetScheduleToStartTimeout(); !local && timeout.IsValid() {
				fn.If(g.Id("opts").Dot("opts").Dot("ScheduleToStartTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("opts").Dot("ScheduleToStartTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// set default start to close timeout
			if timeout := opts.GetStartToCloseTimeout(); timeout.IsValid() {
				fn.If(g.Id("opts").Dot("opts").Dot("StartToCloseTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("opts").Dot("StartToCloseTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// inject ctx with activity options
			if local {
				fn.Id("ctx").Op("=").Qual(workflowPkg, "WithLocalActivityOptions").Call(
					g.Id("ctx"), g.Op("*").Id("opts").Dot("opts"),
				)

			} else {
				fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(
					g.Id("ctx"), g.Op("*").Id("opts").Dot("opts"),
				)
			}

			// initialize activity reference
			fn.Var().Id("activity").Any()
			if local {
				fn.If(g.Id("opts").Dot("fn").Op("!=").Nil()).
					Block(
						g.Id("activity").Op("=").Id("opts").Dot("fn"),
					).
					Else().
					Block(
						g.Id("activity").Op("=").Id(svc.toCamel("%sActivityName", activity)),
					)
			} else {
				fn.Id("activity").Op("=").Id(svc.toCamel("%sActivityName", activity))
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
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Id(method.Output.GoIdent.GoName)
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
						args.Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
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

// genActivityLocalOptions generates an <Activity>LocalActivityOptions struct
func (svc *Manifest) genActivityLocalOptions(f *g.File, activity protoreflect.FullName) {
	typeName := svc.toCamel("%sLocalActivityOptions", activity)
	method := svc.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	// generate type definition
	f.Commentf("%s provides configuration for a local %s activity", typeName, svc.fqnForActivity(activity))
	f.Type().Id(typeName).Struct(
		g.Id("fn").Func().
			ParamsFunc(func(args *g.Group) {
				args.Qual("context", "Context")
				if hasInput {
					args.Op("*").Id(method.Input.GoIdent.GoName)
				}
			}).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Id(method.Output.GoIdent.GoName)
				}
				returnVals.Error()
			}),
		g.Id("opts").Op("*").Qual(workflowPkg, "LocalActivityOptions"),
	)

	// generate New<Activity>LocalActivityOptions method
	ctorName := "New" + typeName
	f.Commentf("%s sets default LocalActivityOptions", ctorName)
	f.Func().
		Id(ctorName).
		Params().
		Op("*").Id(typeName).
		Block(
			g.Return(g.Op("&").Id(typeName).Values()),
		)

	// generate Local method
	f.Commentf("Local provides a local %s activity implementation", svc.fqnForActivity(activity))
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("Local").
		Params(
			g.Id("fn").Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual("context", "Context")
					if hasInput {
						args.Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
					}
					returnVals.Error()
				}),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("fn").Op("=").Id("fn"),
			g.Return(g.Id("opts")),
		)

	// generate WithLocalActivityOptions method
	f.Comment("WithLocalActivityOptions sets default LocalActivityOptions")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithLocalActivityOptions").
		Params(
			g.Id("options").Qual(workflowPkg, "LocalActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("opts").Op("=").Op("&").Id("options"),
			g.Return(g.Id("opts")),
		)
}

// genActivityOptions generates an <Activity>ActivityOptions struct
func (svc *Manifest) genActivityOptions(f *g.File, activity protoreflect.FullName) {
	typeName := svc.toCamel("%sActivityOptions", activity)

	// generate type definition
	f.Commentf("%s provides configuration for a(n) %s activity", typeName, svc.fqnForActivity(activity))
	f.Type().Id(typeName).Struct(
		g.Id("opts").Op("*").Qual(workflowPkg, "ActivityOptions"),
	)

	// generate New<Activity>ActivityOptions method
	ctorName := "New" + typeName
	f.Commentf("%s sets default ActivityOptions", ctorName)
	f.Func().
		Id(ctorName).
		Params().
		Op("*").Id(typeName).
		Block(
			g.Return(g.Op("&").Id(typeName).Values()),
		)

	// generate WithActivityOptions method
	f.Comment("WithActivityOptions sets default ActivityOptions")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithActivityOptions").
		Params(
			g.Id("options").Qual(workflowPkg, "ActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("opts").Op("=").Op("&").Id("options"),
			g.Return(g.Id("opts")),
		)
}
