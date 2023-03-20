package plugin

import (
	"fmt"
	"strconv"
	"strings"

	g "github.com/dave/jennifer/jen"
)

// genActivitiesInterface generates an Activities interface
func (svc *Service) genActivitiesInterface(f *g.File) {
	f.Comment("Activities describes available worker activites")
	f.Type().Id("Activities").InterfaceFunc(func(methods *g.Group) {
		for _, activity := range svc.activitiesOrdered {
			method := svc.methods[activity]
			methods.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			methods.Id(activity).
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
				})
		}
	})
}

// genActivitiesInterface generates a RegisterActivities public function
func (svc *Service) genRegisterActivities(f *g.File) {
	f.Comment("RegisterActivities registers activities with a worker")
	f.Func().Id("RegisterActivities").
		Params(
			g.Id("r").Qual(workerPkg, "Registry"),
			g.Id("activities").Id("Activities"),
		).
		BlockFunc(func(fn *g.Group) {
			for _, activity := range svc.activitiesOrdered {
				fn.Id(fmt.Sprintf("Register%sActivity", activity)).Call(
					g.Id("r"), g.Id("activities").Dot(activity),
				)
			}
		})
}

// genRegisterActivity generates a Register<Activity> public function
func (svc *Service) genRegisterActivity(f *g.File, activity string) {
	method := svc.methods[activity]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	f.Commentf("Register%sActivity registers a %s activity", activity, activity)
	f.Func().Id(fmt.Sprintf("Register%sActivity", activity)).
		Params(
			g.Id("r").Qual(workerPkg, "Registry"),
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
					g.Id("Name").Op(":").Id(fmt.Sprintf("%sActivityName", activity)).Op(","),
				),
			),
		)
}

// genActivityFuture generates a <Activity>Future struct
func (svc *Service) genActivityFuture(f *g.File, activity string) {
	future := fmt.Sprintf("%sFuture", activity)

	f.Commentf("%s describes a %s activity execution", future, activity)
	f.Type().Id(future).Struct(
		g.Id("Future").Qual(workflowPkg, "Future"),
	)
}

// genActivityFutureGetMethod generates a <Workflow>Future's Get method
func (svc *Service) genActivityFutureGetMethod(f *g.File, activity string) {
	method := svc.methods[activity]
	hasOutput := !isEmpty(method.Output)
	future := fmt.Sprintf("%sFuture", activity)

	f.Commentf("Get blocks on a %s execution, returning the response", activity)
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
func (svc *Service) genActivityFutureSelectMethod(f *g.File, activity string) {
	future := fmt.Sprintf("%sFuture", activity)

	f.Commentf("Select adds the %s completion to the selector, callback can be nil", activity)
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

// genActivityFunction generates a public <Activity>[Local] function
func (svc *Service) genActivityFunction(f *g.File, activity string, local bool) {
	method := svc.methods[activity]
	methodName := method.GoName
	if local {
		methodName = fmt.Sprintf("%sLocal", methodName)
	}
	opts := svc.activities[activity].GetDefaultOptions()
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if local {
				args.Id("opts").Op("*").Qual(workflowPkg, "LocalActivityOptions")
			} else {
				args.Id("opts").Op("*").Qual(workflowPkg, "ActivityOptions")
			}
			if local {
				args.Id("fn").
					Func().
					ParamsFunc(func(fnargs *g.Group) {
						fnargs.Qual("context", "Context")
						if hasInput {
							fnargs.Op("*").Id(method.Input.GoIdent.GoName)
						}
					}).
					ParamsFunc(func(fnreturn *g.Group) {
						if hasOutput {
							fnreturn.Op("*").Id(method.Output.GoIdent.GoName)
						}
						fnreturn.Error()
					})
			}
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		Params(
			g.Op("*").Id(fmt.Sprintf("%sFuture", method.GoName)),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize activity options if nil
			fn.If(g.Id("opts").Op("==").Nil()).BlockFunc(func(bl *g.Group) {
				optionsFn := "GetActivityOptions"
				if local {
					optionsFn = "GetLocalActivityOptions"
				}
				bl.Id("activityOpts").Op(":=").Qual(workflowPkg, optionsFn).Call(
					g.Id("ctx"),
				)
				bl.Id("opts").Op("=").Op("&").Id("activityOpts")
			})

			// set default retry policy
			if policy := opts.GetRetryPolicy(); policy != nil {
				fn.If(g.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
					g.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").BlockFunc(func(fields *g.Group) {
						if d := policy.GetInitialInterval(); d.IsValid() {
							fields.Id("InitialInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Op(",").Comment(d.AsDuration().String())
						}
						if d := policy.GetMaxInterval(); d.IsValid() {
							fields.Id("MaximumInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Op(",").Comment(d.AsDuration().String())
						}
						if n := policy.GetBackoffCoefficient(); n != 0 {
							fields.Id("BackoffCoefficient").Op(":").Lit(n).Op(",")
						}
						if n := policy.GetMaxAttempts(); n != 0 {
							fields.Id("MaximumAttempts").Op(":").Lit(n).Op(",")
						}
						if errs := policy.GetNonRetryableErrorTypes(); len(errs) > 0 {
							fields.Id("NonRetryableErrorTypes").Op(":").Lit(errs).Op(",")
						}
					}),
				)
			}

			// set default heartbeat timeout
			if timeout := opts.GetHeartbeatTimeout(); !local && timeout.IsValid() {
				fn.If(g.Id("opts").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("HeartbeatTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// set default schedule to close timeout
			if timeout := opts.GetScheduleToCloseTimeout(); timeout.IsValid() {
				fn.If(g.Id("opts").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// set default schedule to start timeout
			if timeout := opts.GetScheduleToStartTimeout(); !local && timeout.IsValid() {
				fn.If(g.Id("opts").Dot("ScheduleToStartTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("ScheduleToStartTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// set default start to close timeout
			if timeout := opts.GetStartToCloseTimeout(); timeout.IsValid() {
				fn.If(g.Id("opts").Dot("StartToCloseTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("StartToCloseTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
				)
			}

			// inject ctx with activity options
			if local {
				fn.Id("ctx").Op("=").Qual(workflowPkg, "WithLocalActivityOptions").Call(
					g.Id("ctx"), g.Op("*").Id("opts"),
				)

			} else {
				fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(
					g.Id("ctx"), g.Op("*").Id("opts"),
				)
			}

			// if activity function nil for local activity, replace with activity name
			if local {
				fn.Var().Id("activity").Any()
				fn.If(g.Id("fn").Op("==").Nil()).
					Block(
						g.Id("activity").Op("=").Id(fmt.Sprintf("%sActivityName", activity)),
					).
					Else().
					Block(
						g.Id("activity").Op("=").Id("fn"),
					)
			}

			fn.Return(
				g.Op("&").Id(fmt.Sprintf("%sFuture", method.GoName)).ValuesFunc(func(bl *g.Group) {
					future := bl.Id("Future").Op(":")
					if local {
						future.Qual(workflowPkg, "ExecuteLocalActivity").CallFunc(func(returnVals *g.Group) {
							returnVals.Id("ctx")
							returnVals.Id("activity")
							if hasInput {
								returnVals.Id("req")
							}
						}).Op(",")
					} else {
						future.Qual(workflowPkg, "ExecuteActivity").CallFunc(func(returnVals *g.Group) {
							returnVals.Id("ctx")
							returnVals.Id(fmt.Sprintf("%sActivityName", method.GoName))
							if hasInput {
								returnVals.Id("req")
							}
						}).Op(",")
					}
				}),
			)
		})
}
