package plugin

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	g "github.com/dave/jennifer/jen"
)

const (
	anypbPkg = "google.golang.org/protobuf/types/known/anypb"
	xnsv1Pkg = "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	xnsPkg   = "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
)

func (svc *Service) renderXNS(f *g.File) {
	for _, workflow := range svc.workflowsOrdered {
		svc.genXNSWorkflowOptions(f, workflow)
		svc.genXNSWorkflowRunInterface(f, workflow)
		svc.genXNSWorkflowRunImpl(f, workflow)
		svc.genXNSWorkflowFunction(f, workflow)
		svc.genXNSWorkflowFunctionAsync(f, workflow)

		for _, signal := range svc.workflows[workflow].GetSignal() {
			if !signal.GetStart() {
				continue
			}
			svc.genXNSWorkflowWithStartFunction(f, workflow, signal.GetRef())
			svc.genXNSWorkflowWithStartFunctionAsync(f, workflow, signal.GetRef())
		}
	}
	for _, query := range svc.queriesOrdered {
		svc.genXNSQueryOptions(f, query)
		svc.genXNSQueryHandleInterface(f, query)
		svc.genXNSQueryHandleImpl(f, query)
		svc.genXNSQueryFunction(f, query)
		svc.genXNSQueryFunctionAsync(f, query)
	}
	for _, signal := range svc.signalsOrdered {
		svc.genXNSSignalOptions(f, signal)
		svc.genXNSSignalHandleInterface(f, signal)
		svc.genXNSSignalHandleImpl(f, signal)
		svc.genXNSSignalFunction(f, signal)
		svc.genXNSSignalFunctionAsync(f, signal)
	}
	for _, update := range svc.updatesOrdered {
		svc.genXNSUpdateOptions(f, update)
		svc.genXNSUpdateHandleInterface(f, update)
		svc.genXNSUpdateHandleImpl(f, update)
		svc.genXNSUpdateFunction(f, update)
		svc.genXNSUpdateFunctionAsync(f, update)
	}
	svc.genXNSCancelWorkflowFunction(f)

	svc.genXNSActivities(f)
	svc.genXNSRegisterActivities(f)
	for _, workflow := range svc.workflowsOrdered {
		svc.genXNSActivitiesWorkflowMethod(f, workflow)
		for _, signal := range svc.workflows[workflow].GetSignal() {
			if !signal.GetStart() {
				continue
			}
			svc.genXNSActivitiesWorkflowWithStartMethod(f, workflow, signal.GetRef())
		}
	}
	for _, query := range svc.queriesOrdered {
		svc.genXNSActivitiesQueryMethod(f, query)
	}
	for _, signal := range svc.signalsOrdered {
		svc.genXNSActivitiesSignalMethod(f, signal)
	}
	for _, update := range svc.updatesOrdered {
		svc.genXNSActivitiesUpdateMethod(f, update)
	}
}

func (svc *Service) genXNSActivities(f *g.File) {
	typeName := toLowerCamel("%sActivities", svc.GoName)

	f.Commentf("%s provides activities that can be used to interact with a(n) %s service's workflow, queries, signals, and updates across namespaces", typeName, svc.GoName)
	f.Type().Id(typeName).Struct(
		g.Id("client").Qual(string(svc.File.GoImportPath), toCamel("%sClient", svc.GoName)),
	)

	f.Comment("CancelWorkflow cancels an existing workflow execution")
	f.Func().
		Params(
			g.Id("a").Op("*").Id(typeName),
		).
		Id("CancelWorkflow").
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Error().
		Block(
			g.Return(
				g.Id("a").Dot("client").Dot("CancelWorkflow").Call(g.Id("ctx"), g.Id("workflowID"), g.Id("runID")),
			),
		)
}

func (svc *Service) genXNSActivitiesQueryMethod(f *g.File, query string) {
	methodName := toCamel(query)
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a(n) %s query via an activity", methodName, svc.fqnForQuery(query))
	f.Func().
		Params(
			g.Id("a").Op("*").Id(toLowerCamel("%sActivities", svc.GoName)),
		).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("input").Op("*").Qual(xnsv1Pkg, "QueryRequest"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasInput {
				fn.Comment("unmarshal query request")
				fn.Var().Id("req").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
				fn.If(g.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(g.Op("&").Id("req")), g.Err().Op("!=").Nil()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				)
			}

			fn.Return(
				g.Id("a").Dot("client").Dot(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("input").Dot("GetWorkflowId").Call()
					args.Id("input").Dot("GetRunId").Call()
					if hasInput {
						args.Op("&").Id("req")
					}
				}),
			)
		})
}

func (svc *Service) genXNSActivitiesSignalMethod(f *g.File, signal string) {
	methodName := toCamel(signal)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s executes a(n) %s signal via an activity", methodName, svc.fqnForSignal(signal))
	f.Func().
		Params(
			g.Id("a").Op("*").Id(toLowerCamel("%sActivities", svc.GoName)),
		).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("input").Op("*").Qual(xnsv1Pkg, "SignalRequest"),
		).
		Error().
		BlockFunc(func(fn *g.Group) {
			if hasInput {
				fn.Comment("unmarshal signal request")
				fn.Var().Id("req").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
				fn.If(g.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(g.Op("&").Id("req")), g.Err().Op("!=").Nil()).Block(
					g.Return(g.Err()),
				)
			}

			fn.Return(
				g.Id("a").Dot("client").Dot(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("input").Dot("GetWorkflowId").Call()
					args.Id("input").Dot("GetRunId").Call()
					if hasInput {
						args.Op("&").Id("req")
					}
				}),
			)
		})
}

func (svc *Service) genXNSActivitiesUpdateMethod(f *g.File, update string) {
	methodName := toCamel(update)
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a(n) %s update via an activity", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(
			g.Id("a").Op("*").Id(toLowerCamel("%sActivities", svc.GoName)),
		).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("input").Op("*").Qual(xnsv1Pkg, "UpdateRequest"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Id("resp").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Err().Error()
		}).
		BlockFunc(func(fn *g.Group) {
			fn.Var().Id("update").Qual(string(svc.File.GoImportPath), toCamel("%sHandle", update))
			fn.If(g.Qual(activityPkg, "HasHeartbeatDetails").Call(g.Id("ctx"))).Block(
				g.Comment("extract update id from heartbeat details"),
				g.Var().Id("updateID").String(),
				g.If(
					g.Err().Op(":=").Qual(activityPkg, "GetHeartbeatDetails").Call(g.Id("ctx"), g.Op("&").Id("updateID")),
					g.Err().Op("!=").Nil(),
				).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				),
				g.Line(),
				g.Comment("retrieve handle for existing update"),
				g.List(g.Id("update"), g.Err()).Op("=").Id("a").Dot("client").Dot(toCamel("Get%s", update)).Call(
					g.Id("ctx"),
					g.Qual(clientPkg, "GetWorkflowUpdateHandleOptions").Custom(
						multiLineValues,
						g.Id("WorkflowID").Op(":").Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetWorkflowId").Call(),
						g.Id("RunID").Op(":").Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetRunId").Call(),
						g.Id("UpdateID").Op(":").Id("updateID"),
					),
				),
				g.If(g.Err().Op("!=").Nil()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				),
			).Else().BlockFunc(func(bl *g.Group) {
				if hasInput {
					bl.Comment("unmarshal update request")
					bl.Var().Id("req").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
					bl.If(g.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(g.Op("&").Id("req")), g.Err().Op("!=").Nil()).Block(
						g.ReturnFunc(func(returnVals *g.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Err()
						}),
					)
					bl.Line()
				}

				bl.Comment("initialize update execution")
				bl.List(g.Id("update"), g.Err()).Op("=").Id("a").Dot("client").Dot(toCamel("%sAsync", methodName)).CustomFunc(multiLineArgs, func(args *g.Group) {
					args.Id("ctx")
					args.Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetWorkflowId").Call()
					args.Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetRunId").Call()
					if hasInput {
						args.Op("&").Id("req")
					}
					args.Qual(string(svc.File.GoImportPath), toCamel("New%sOptions", update)).
						Call().
						Dot("WithUpdateWorkflowOptions").
						Custom(multiLineArgs, g.Qual(xnsPkg, "UnmarshalUpdateWorkflowOptions").Call(
							g.Id("input").Dot("GetUpdateWorkflowOptions").Call(),
						))
				})
				bl.If(g.Err().Op("!=").Nil()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				)
				bl.Qual(activityPkg, "RecordHeartbeat").Call(g.Id("ctx"), g.Id("update").Dot("UpdateID").Call())
			})
			fn.Line()

			fn.Comment("wait for update to complete in child goroutine")
			fn.Id("doneCh").Op(":=").Make(g.Chan().Struct())
			fn.Go().Func().Params().Block(
				g.ListFunc(func(ls *g.Group) {
					if hasOutput {
						ls.Id("resp")
					}
					ls.Err()
				}).Op("=").Id("update").Dot("Get").Call(g.Id("ctx")),
				g.Close(g.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			fn.If(g.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("heartbeatInterval").Op("=").Qual("time", "Minute"),
			)
			fn.Line()

			fn.Comment("heartbeat activity while waiting for workflow update to complete")
			fn.For().Block(
				g.Select().Block(
					g.Case(g.Op("<-").Qual("time", "After").Call(g.Id("heartbeatInterval"))).Block(
						g.Qual(activityPkg, "RecordHeartbeat").Call(g.Id("ctx"), g.Id("update").Dot("UpdateID").Call()),
					),
					g.Case(g.Op("<-").Id("ctx").Dot("Done").Call()).Block(
						g.ReturnFunc(func(returnVals *g.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Qual(workflowPkg, "ErrCanceled")
						}),
					),
					g.Case(g.Op("<-").Id("doneCh")).Block(
						g.ReturnFunc(func(returnVals *g.Group) {
							if hasOutput {
								returnVals.Id("resp")
							}
							returnVals.Err()
						}),
					),
				),
			)
		})
}

func (svc *Service) genXNSActivitiesWorkflowMethod(f *g.File, workflow string) {
	methodName := toCamel(workflow)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a(n) %s workflow via an activity", methodName, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(
			g.Id("a").Op("*").Id(toLowerCamel("%sActivities", svc.GoName)),
		).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("input").Op("*").Qual(xnsv1Pkg, "WorkflowRequest"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Id("resp").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Err().Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasInput {
				fn.Comment("unmarshal workflow request")
				fn.Var().Id("req").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
				fn.If(g.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(g.Op("&").Id("req")), g.Err().Op("!=").Nil()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				)
				fn.Line()
			}

			fn.Comment("initialize workflow execution")
			fn.Var().Id("run").Qual(string(svc.File.GoImportPath), toCamel("%sRun", workflow))
			fn.List(g.Id("run"), g.Err()).Op("=").Id("a").Dot("client").Dot(toCamel("%sAsync", methodName)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Op("&").Id("req")
				}
				args.Qual(string(svc.File.GoImportPath), toCamel("New%sOptions", workflow)).
					Call().
					Dot("WithStartWorkflowOptions").
					Custom(multiLineArgs, g.Qual(xnsPkg, "UnmarshalStartWorkflowOptions").Call(g.Id("input").Dot("GetStartWorkflowOptions").Call()))
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Line()

			fn.Comment("exit early if detached enabled")
			fn.If(g.Id("input").Dot("GetDetached").Call()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Nil()
				}),
			)
			fn.Line()

			fn.Comment("otherwise, wait for execution to complete in child goroutine")
			fn.Id("doneCh").Op(":=").Make(g.Chan().Struct())
			fn.Go().Func().Params().Block(
				g.ListFunc(func(ls *g.Group) {
					if hasOutput {
						ls.Id("resp")
					}
					ls.Err()
				}).Op("=").Id("run").Dot("Get").Call(g.Id("ctx")),
				g.Close(g.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			fn.If(g.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("heartbeatInterval").Op("=").Qual("time", "Minute"),
			)
			fn.Line()

			fn.Comment("heartbeat activity while waiting for workflow execution to complete")
			fn.For().Block(
				g.Select().Block(
					g.Case(g.Op("<-").Qual("time", "After").Call(g.Id("heartbeatInterval"))).Block(
						g.Qual(activityPkg, "RecordHeartbeat").Call(g.Id("ctx"), g.Id("run").Dot("ID").Call()),
					),
					g.Case(g.Op("<-").Id("ctx").Dot("Done").Call()).Block(
						g.If(
							g.Err().Op(":=").Id("run").Dot("Cancel").Call(g.Id("ctx")),
							g.Err().Op("!=").Nil(),
						).Block(
							g.ReturnFunc(func(returnVals *g.Group) {
								if hasOutput {
									returnVals.Nil()
								}
								returnVals.Err()
							}),
						),
						g.ReturnFunc(func(returnVals *g.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Qual(workflowPkg, "ErrCanceled")
						}),
					),
					g.Case(g.Op("<-").Id("doneCh")).Block(
						g.ReturnFunc(func(returnVals *g.Group) {
							if hasOutput {
								returnVals.Id("resp")
							}
							returnVals.Err()
						}),
					),
				),
			)
		})
}

func (svc *Service) genXNSActivitiesWorkflowWithStartMethod(f *g.File, workflow, signal string) {
	methodName := toCamel("%sWith%s", workflow, signal)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := svc.methods[signal]
	handlerInput := !isEmpty(handler.Input)

	f.Commentf("%s sends a(n) %s signal to a(n) %s workflow via an activity", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(
			g.Id("a").Op("*").Id(toLowerCamel("%sActivities", svc.GoName)),
		).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("input").Op("*").Qual(xnsv1Pkg, "WorkflowRequest"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Id("resp").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Err().Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasInput {
				fn.Comment("unmarshal workflow request")
				fn.Var().Id("req").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
				fn.If(g.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(g.Op("&").Id("req")), g.Err().Op("!=").Nil()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				)
				fn.Line()
			}
			if handlerInput {
				fn.Comment("unmarshal signal request")
				fn.Var().Id("signal").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
				fn.If(g.Err().Op(":=").Id("input").Dot("Signal").Dot("UnmarshalTo").Call(g.Op("&").Id("signal")), g.Err().Op("!=").Nil()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				)
				fn.Line()
			}

			fn.Comment("initialize workflow execution")
			fn.Var().Id("run").Qual(string(svc.File.GoImportPath), toCamel("%sRun", workflow))
			fn.List(g.Id("run"), g.Err()).Op("=").Id("a").Dot("client").Dot(toCamel("%sWith%sAsync", workflow, signal)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Op("&").Id("req")
				}
				if handlerInput {
					args.Op("&").Id("signal")
				}
				args.Qual(string(svc.File.GoImportPath), toCamel("New%sOptions", workflow)).
					Call().
					Dot("WithStartWorkflowOptions").
					Custom(multiLineArgs, g.Qual(xnsPkg, "UnmarshalStartWorkflowOptions").Call(g.Id("input").Dot("GetStartWorkflowOptions").Call()))
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Line()

			fn.Comment("exit early if detached enabled")
			fn.If(g.Id("input").Dot("GetDetached").Call()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Nil()
				}),
			)
			fn.Line()

			fn.Comment("otherwise, wait for execution to complete in child goroutine")
			fn.Id("doneCh").Op(":=").Make(g.Chan().Struct())
			fn.Go().Func().Params().Block(
				g.ListFunc(func(ls *g.Group) {
					if hasOutput {
						ls.Id("resp")
					}
					ls.Err()
				}).Op("=").Id("run").Dot("Get").Call(g.Id("ctx")),
				g.Close(g.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			fn.If(g.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("heartbeatInterval").Op("=").Qual("time", "Minute"),
			)
			fn.Line()

			fn.Comment("heartbeat activity while waiting for workflow execution to complete")
			fn.For().Block(
				g.Select().Block(
					g.Case(g.Op("<-").Qual("time", "After").Call(g.Id("heartbeatInterval"))).Block(
						g.Qual(activityPkg, "RecordHeartbeat").Call(g.Id("ctx"), g.Id("run").Dot("ID").Call()),
					),
					g.Case(g.Op("<-").Id("ctx").Dot("Done").Call()).Block(
						g.If(
							g.Err().Op(":=").Id("run").Dot("Cancel").Call(g.Id("ctx")),
							g.Err().Op("!=").Nil(),
						).Block(
							g.ReturnFunc(func(returnVals *g.Group) {
								if hasOutput {
									returnVals.Nil()
								}
								returnVals.Err()
							}),
						),
						g.ReturnFunc(func(returnVals *g.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Qual(workflowPkg, "ErrCanceled")
						}),
					),
					g.Case(g.Op("<-").Id("doneCh")).Block(
						g.ReturnFunc(func(returnVals *g.Group) {
							if hasOutput {
								returnVals.Id("resp")
							}
							returnVals.Err()
						}),
					),
				),
			)
		})
}

func (svc *Service) genXNSCancelWorkflowFunction(f *g.File) {
	funcName := toCamel("Cancel%sWorkflow", svc.GoName)
	f.Commentf("%s cancels an existing workflow", funcName)
	f.Func().
		Id(funcName).
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Error().
		Block(
			g.Return(
				g.Id(toCamel("%sAsync", funcName)).
					Call(
						g.Id("ctx"),
						g.Id("workflowID"),
						g.Id("runID"),
					).
					Dot("Get").
					Call(
						g.Id("ctx"),
						g.Nil(),
					),
			),
		)

	funcName = toCamel("%sAsync", funcName)
	f.Commentf("%s cancels an existing workflow", funcName)
	f.Func().
		Id(funcName).
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Qual(workflowPkg, "Future").
		Block(
			g.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(g.Id("ctx")),
			g.If(g.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0).Op("&&").Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("StartToCloseTimeout").Op("=").Qual("time", "Minute"),
			),
			g.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(g.Id("ctx"), g.Id("ao")),
			g.Return(
				g.Qual(workflowPkg, "ExecuteActivity").Call(
					g.Id("ctx"),
					g.Lit(fmt.Sprintf("%s.CancelWorkflow", svc.Service.Desc.FullName())),
					g.Id("workflowID"),
					g.Id("runID"),
				),
			),
		)
}

func (svc *Service) genXNSRegisterActivities(f *g.File) {
	funcName := toCamel("Register%sActivities", svc.GoName)
	f.Commentf("%s registers %s cross-namespace activities", funcName, string(svc.Service.Desc.FullName()))
	f.Func().
		Id(funcName).
		Params(
			g.Id("r").Qual(workerPkg, "ActivityRegistry"),
			g.Id("c").Qual(string(svc.File.GoImportPath), toCamel("%sClient", svc.GoName)),
		).
		BlockFunc(func(fn *g.Group) {
			fn.Id("a").Op(":=").Op("&").Id(toLowerCamel("%sActivities", svc.GoName)).Values(
				g.Id("c"),
			)
			fn.Id("r").Dot("RegisterActivityWithOptions").Call(
				g.Id("a").Dot("CancelWorkflow"),
				g.Qual(activityPkg, "RegisterOptions").Values(
					g.Id("Name").Op(":").Lit(fmt.Sprintf("%s.CancelWorkflow", svc.Service.Desc.FullName())),
				),
			)
			for _, workflow := range svc.workflowsOrdered {
				fn.Id("r").Dot("RegisterActivityWithOptions").Call(
					g.Id("a").Dot(toCamel(workflow)),
					g.Qual(activityPkg, "RegisterOptions").Values(
						g.Id("Name").Op(":").Qual(string(svc.File.GoImportPath), toCamel("%sWorkflowName", workflow)),
					),
				)
				for _, signal := range svc.workflows[workflow].GetSignal() {
					if !signal.GetStart() {
						continue
					}
					fn.Id("r").Dot("RegisterActivityWithOptions").Call(
						g.Id("a").Dot(toCamel("%sWith%s", workflow, signal.GetRef())),
						g.Qual(activityPkg, "RegisterOptions").Values(
							g.Id("Name").Op(":").Lit(fmt.Sprintf("%s.%s", string(svc.Service.Desc.FullName()), toCamel("%sWith%s", workflow, signal.GetRef()))),
						),
					)
				}
			}
			for _, query := range svc.queriesOrdered {
				fn.Id("r").Dot("RegisterActivityWithOptions").Call(
					g.Id("a").Dot(toCamel(query)),
					g.Qual(activityPkg, "RegisterOptions").Values(
						g.Id("Name").Op(":").Qual(string(svc.File.GoImportPath), toCamel("%sQueryName", query)),
					),
				)
			}
			for _, signal := range svc.signalsOrdered {
				fn.Id("r").Dot("RegisterActivityWithOptions").Call(
					g.Id("a").Dot(toCamel(signal)),
					g.Qual(activityPkg, "RegisterOptions").Values(
						g.Id("Name").Op(":").Qual(string(svc.File.GoImportPath), toCamel("%sSignalName", signal)),
					),
				)
			}
			for _, update := range svc.updatesOrdered {
				fn.Id("r").Dot("RegisterActivityWithOptions").Call(
					g.Id("a").Dot(toCamel(update)),
					g.Qual(activityPkg, "RegisterOptions").Values(
						g.Id("Name").Op(":").Qual(string(svc.File.GoImportPath), toCamel("%sUpdateName", update)),
					),
				)
			}
		})
}

func (svc *Service) genXNSQueryFunction(f *g.File, query string) {
	methodName := toCamel(query)
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s query and blocks until error or response received", methodName, svc.fqnForQuery(query))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sQueryOptions", query))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.List(g.Id("handle"), g.Err()).Op(":=").Id(toCamel("%sAsync", query)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			}),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			),
			g.Return(g.Id("handle").Dot("Get").Call(g.Id("ctx"))),
		)
}

func (svc *Service) genXNSQueryFunctionAsync(f *g.File, query string) {
	methodName := toCamel("%sAsync", query)
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s query and blocks until error or response received", methodName, svc.fqnForQuery(query))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sQueryOptions", query))
		}).
		Params(
			g.Id(toCamel("%sQueryHandle", query)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(toCamel("%sQueryOptions", query)).Values()
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("opt").Op("=").Id("opts").Index(g.Lit(0)),
			)
			fn.If(g.Id("opt").Dot("HeartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("opt").Dot("HeartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30),
			)
			fn.Line()

			fn.Comment("build and set activity options")
			fn.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(g.Id("ctx"))
			fn.If(g.Id("opt").Dot("ActivityOptions").Op("!=").Nil()).Block(
				g.Id("ao").Op("=").Op("*").Id("opt").Dot("ActivityOptions"),
			)
			fn.If(g.Id("ao").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("HeartbeatTimeout").Op("=").Id("opt").Dot("HeartbeatInterval").Op("*").Lit(2),
			)
			fn.If(g.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0).Op("&&").Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("StartToCloseTimeout").Op("=").Qual("time", "Minute").Op("*").Lit(2),
			)
			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(g.Id("ctx"), g.Id("ao"))
			fn.Line()

			if hasInput {
				fn.Comment("marshal workflow request")
				fn.List(g.Id("wreq"), g.Err()).Op(":=").Qual(anypbPkg, "New").Call(g.Id("req"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling workflow request: %w"), g.Err())),
				)
				fn.Line()
			}

			fn.Comment("create cancellable context")
			fn.List(g.Id("ctx"), g.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(g.Id("ctx"))
			fn.Line()

			// return run with execute activity future
			fn.Return(
				g.Op("&").Id(toLowerCamel("%sQueryHandle", query)).Custom(multiLineValues,
					g.Id("cancel").Op(":").Id("cancel"),
					g.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						g.Id("ctx"),
						g.Qual(string(svc.File.GoImportPath), toCamel("%sQueryName", query)),
						g.Op("&").Qual(xnsv1Pkg, "QueryRequest").CustomFunc(multiLineValues, func(fields *g.Group) {
							fields.Id("WorkflowId").Op(":").Id("workflowID")
							fields.Id("RunId").Op(":").Id("runID")
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
						}),
					),
				),
				g.Nil(),
			)
		})
}

func (svc *Service) genXNSQueryHandleImpl(f *g.File, query string) {
	typeName := toLowerCamel("%sQueryHandle", query)
	method := svc.methods[query]
	// opts := svc.workflows[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a(n) %s implementation", typeName, toCamel("%sQueryHandle", query))
	f.Type().Id(typeName).Struct(
		g.Id("cancel").Func().Params(),
		g.Id("future").Qual(workflowPkg, "Future"),
	)

	f.Comment("Cancel the underlying query activity")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			g.Id("r").Dot("cancel").Call(),
			g.If(
				g.ListFunc(func(ls *g.Group) {
					if hasOutput {
						ls.Id("_")
					}
					ls.Err()
				}).Op(":=").Id("r").Dot("Get").Call(g.Id("ctx")),
				g.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(g.Err(), g.Qual(workflowPkg, "ErrCanceled")),
			).Block(
				g.Return(g.Err()),
			),
			g.Return(g.Nil()),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		Block(
			g.Return(g.Id("r").Dot("future")),
		)

	f.Comment("Get blocks on activity completion and returns the underlying query result")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			fn.If(
				g.Err().Op(":=").Id("r").Dot("future").Dot("Get").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasOutput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				g.Err().Op("!=").Nil(),
			).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.ReturnFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("&").Id("resp")
				}
				returnVals.Nil()
			})
		})
}

func (svc *Service) genXNSQueryHandleInterface(f *g.File, query string) {
	typeName := toCamel("%sQueryHandle", query)
	method := svc.methods[query]
	// opts := svc.workflows[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a handle for a %s query activity", typeName, query)
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		methods.Comment("Cancel cancels the workflow")
		methods.Id("Cancel").
			Params(g.Qual(workflowPkg, "Context")).
			Error()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future")

		methods.Comment("Get returns the inner workflow.Future")
		methods.Id("Get").
			Params(g.Qual(workflowPkg, "Context")).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			})
	})
}

func (svc *Service) genXNSQueryOptions(f *g.File, query string) {
	typeName := toCamel("%sQueryOptions", query)

	f.Commentf("%s are used to configure a(n) %s query execution", typeName, svc.fqnForQuery(query))
	f.Type().Id(typeName).Struct(
		g.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		g.Id("HeartbeatInterval").Qual("time", "Duration"),
	)

	f.Commentf("New%s initializes a new %s value", typeName, typeName)
	f.Func().
		Id(toCamel("New%s", typeName)).
		Params().
		Op("*").Id(typeName).
		Block(
			g.Return(
				g.Op("&").Id(typeName).Values(),
			),
		)

	f.Comment("WithActivityOptions can be used to customize the activity options")
	f.Func().
		Params(
			g.Id("opts").Op("*").Id(typeName),
		).
		Id("WithActivityOptions").
		Params(
			g.Id("ao").Qual(workflowPkg, "ActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("ActivityOptions").Op("=").Op("&").Id("ao"),
			g.Return(g.Id("opts")),
		)

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		Params(
			g.Id("opts").Op("*").Id(typeName),
		).
		Id("WithHeartbeatInterval").
		Params(
			g.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("HeartbeatInterval").Op("=").Id("d"),
			g.Return(g.Id("opts")),
		)
}

func (svc *Service) genXNSSignalFunction(f *g.File, signal string) {
	methodName := toCamel(signal)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s signal", methodName, svc.fqnForSignal(signal))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sSignalOptions", signal))
		}).
		Error().
		Block(
			g.List(g.Id("handle"), g.Err()).Op(":=").Id(toCamel("%sAsync", signal)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			}),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Err()),
			),
			g.Return(g.Id("handle").Dot("Get").Call(g.Id("ctx"))),
		)
}

func (svc *Service) genXNSSignalFunctionAsync(f *g.File, signal string) {
	methodName := toCamel("%sAsync", signal)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s query", methodName, svc.fqnForSignal(signal))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sSignalOptions", signal))
		}).
		Params(
			g.Id(toCamel("%sSignalHandle", signal)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(toCamel("%sSignalOptions", signal)).Values()
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("opt").Op("=").Id("opts").Index(g.Lit(0)),
			)
			fn.If(g.Id("opt").Dot("HeartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("opt").Dot("HeartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30),
			)
			fn.Line()

			fn.Comment("build and set activity options")
			fn.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(g.Id("ctx"))
			fn.If(g.Id("opt").Dot("ActivityOptions").Op("!=").Nil()).Block(
				g.Id("ao").Op("=").Op("*").Id("opt").Dot("ActivityOptions"),
			)
			fn.If(g.Id("ao").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("HeartbeatTimeout").Op("=").Id("opt").Dot("HeartbeatInterval").Op("*").Lit(2),
			)
			fn.If(g.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0).Op("&&").Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("StartToCloseTimeout").Op("=").Qual("time", "Minute").Op("*").Lit(2),
			)
			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(g.Id("ctx"), g.Id("ao"))
			fn.Line()

			if hasInput {
				fn.Comment("marshal workflow request")
				fn.List(g.Id("wreq"), g.Err()).Op(":=").Qual(anypbPkg, "New").Call(g.Id("req"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling workflow request: %w"), g.Err())),
				)
				fn.Line()
			}

			fn.Comment("create cancellable context")
			fn.List(g.Id("ctx"), g.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(g.Id("ctx"))
			fn.Line()

			// return run with execute activity future
			fn.Return(
				g.Op("&").Id(toLowerCamel("%sSignalHandle", signal)).Custom(multiLineValues,
					g.Id("cancel").Op(":").Id("cancel"),
					g.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						g.Id("ctx"),
						g.Qual(string(svc.File.GoImportPath), toCamel("%sSignalName", signal)),
						g.Op("&").Qual(xnsv1Pkg, "SignalRequest").CustomFunc(multiLineValues, func(fields *g.Group) {
							fields.Id("WorkflowId").Op(":").Id("workflowID")
							fields.Id("RunId").Op(":").Id("runID")
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
						}),
					),
				),
				g.Nil(),
			)
		})
}

func (svc *Service) genXNSSignalHandleImpl(f *g.File, signal string) {
	typeName := toLowerCamel("%sSignalHandle", signal)
	f.Commentf("%s provides a(n) %s implementation", typeName, toCamel("%sQueryHandle", signal))
	f.Type().Id(typeName).Struct(
		g.Id("cancel").Func().Params(),
		g.Id("future").Qual(workflowPkg, "Future"),
	)

	f.Comment("Cancel the underlying signal activity")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			g.Id("r").Dot("cancel").Call(),
			g.If(
				g.Err().Op(":=").Id("r").Dot("Get").Call(g.Id("ctx")),
				g.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(g.Err(), g.Qual(workflowPkg, "ErrCanceled")),
			).Block(
				g.Return(g.Err()),
			),
			g.Return(g.Nil()),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		Block(
			g.Return(g.Id("r").Dot("future")),
		)

	f.Comment("Get blocks on activity completion")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			g.Return(g.Id("r").Dot("future").Dot("Get").Call(g.Id("ctx"), g.Nil())),
		)
}

func (svc *Service) genXNSSignalHandleInterface(f *g.File, signal string) {
	typeName := toCamel("%sSignalHandle", signal)

	f.Commentf("%s provides a handle for a %s signal activity", typeName, signal)
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		methods.Comment("Cancel cancels the workflow")
		methods.Id("Cancel").
			Params(g.Qual(workflowPkg, "Context")).
			Error()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future")

		methods.Comment("Get returns the inner workflow.Future")
		methods.Id("Get").
			Params(g.Qual(workflowPkg, "Context")).
			Error()
	})
}

func (svc *Service) genXNSSignalOptions(f *g.File, signal string) {
	typeName := toCamel("%sSignalOptions", signal)

	f.Commentf("%s are used to configure a(n) %s signal execution", typeName, svc.fqnForSignal(signal))
	f.Type().Id(typeName).Struct(
		g.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		g.Id("HeartbeatInterval").Qual("time", "Duration"),
	)

	f.Commentf("New%s initializes a new %s value", typeName, typeName)
	f.Func().
		Id(toCamel("New%s", typeName)).
		Params().
		Op("*").Id(typeName).
		Block(
			g.Return(
				g.Op("&").Id(typeName).Values(),
			),
		)

	f.Comment("WithActivityOptions can be used to customize the activity options")
	f.Func().
		Params(
			g.Id("opts").Op("*").Id(typeName),
		).
		Id("WithActivityOptions").
		Params(
			g.Id("ao").Qual(workflowPkg, "ActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("ActivityOptions").Op("=").Op("&").Id("ao"),
			g.Return(g.Id("opts")),
		)

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		Params(
			g.Id("opts").Op("*").Id(typeName),
		).
		Id("WithHeartbeatInterval").
		Params(
			g.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("HeartbeatInterval").Op("=").Id("d"),
			g.Return(g.Id("opts")),
		)
}

func (svc *Service) genXNSUpdateFunction(f *g.File, update string) {
	methodName := update
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s update and blocks until error or response received", methodName, svc.fqnForUpdate(update))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sUpdateOptions", update))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.List(g.Id("run"), g.Err()).Op(":=").Id(toCamel("%sAsync", update)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			}),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			),
			g.Return(g.Id("run").Dot("Get").Call(g.Id("ctx"))),
		)
}

func (svc *Service) genXNSUpdateFunctionAsync(f *g.File, update string) {
	methodName := toCamel("%sAsync", update)
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	opts := svc.updates[update]

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s update and blocks until error or response received", methodName, svc.fqnForUpdate(update))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sUpdateOptions", update))
		}).
		Params(
			g.Id(toCamel("%sHandle", update)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(toCamel("%sUpdateOptions", update)).Values()
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("opt").Op("=").Id("opts").Index(g.Lit(0)),
			)
			fn.If(g.Id("opt").Dot("HeartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("opt").Dot("HeartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30),
			)
			fn.Line()

			// build and set activity options
			fn.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(g.Id("ctx"))
			fn.If(g.Id("opt").Dot("ActivityOptions").Op("!=").Nil()).Block(
				g.Id("ao").Op("=").Op("*").Id("opt").Dot("ActivityOptions"),
			)
			fn.If(g.Id("ao").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("HeartbeatTimeout").Op("=").Id("opt").Dot("HeartbeatInterval").Op("*").Lit(2),
			)
			fn.If(g.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0).Op("&&").Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("StartToCloseTimeout").Op("=").Qual("time", "Hour").Op("*").Lit(24),
			)
			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(g.Id("ctx"), g.Id("ao"))
			fn.Line()

			// build update options
			fn.Id("uo").Op(":=").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest").Values()
			fn.If(g.Id("opt").Dot("UpdateWorkflowOptions").Op("!=").Nil()).Block(
				g.Id("uo").Op("=").Op("*").Id("opt").Dot("UpdateWorkflowOptions"),
			)
			fn.Id("uo").Dot("WorkflowID").Op("=").Id("workflowID")
			fn.Id("uo").Dot("RunID").Op("=").Id("runID")
			// set workflow id if unset and  id field and/or prefix defined
			if idExpr := opts.GetId(); idExpr != "" {
				fn.If(g.Id("uo").Dot("UpdateID").Op("==").Lit("")).Block(
					g.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
						args.Qual(string(svc.File.GoImportPath), toCamel("%sIDExpression", update))
						if hasInput {
							args.Id("req").Dot("ProtoReflect").Call()
						} else {
							args.Nil()
						}
					}),
					g.If(g.Err().Op("!=").Nil()).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error evaluating id expression for %q update: %%w", update)), g.Err())),
					),
					g.Id("uo").Dot("UpdateID").Op("=").Id("id"),
				)
			}
			fn.If(g.Id("uo").Dot("UpdateID").Op("==").Lit("")).Block(
				g.List(g.Id("id"), g.Err()).Op(":=").Qual(uuidPkg, "NewRandom").Call(),
				g.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error generating update id: %w"), g.Err())),
				),
				g.Id("uo").Dot("UpdateID").Op("=").Id("id").Dot("String").Call(),
			)
			fn.Line()

			// marshal update options
			fn.List(g.Id("uopb"), g.Err()).Op(":=").Qual(xnsPkg, "MarshalUpdateWorkflowOptions").Call(g.Id("uo"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling update workflow options: %w"), g.Err())),
			)
			fn.Line()

			// marshal workflow request
			if hasInput {
				fn.List(g.Id("wreq"), g.Err()).Op(":=").Qual(anypbPkg, "New").Call(g.Id("req"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling update request: %w"), g.Err())),
				)
				fn.Line()
			}

			// create cancellable context
			fn.List(g.Id("ctx"), g.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(g.Id("ctx"))
			fn.Line()

			// return run with execute activity future
			fn.Return(
				g.Op("&").Id(toLowerCamel("%sHandle", update)).Custom(multiLineValues,
					g.Id("cancel").Op(":").Id("cancel"),
					g.Id("id").Op(":").Id("uo").Dot("UpdateID"),
					g.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						g.Id("ctx"),
						g.Qual(string(svc.File.GoImportPath), toCamel("%sUpdateName", update)),
						g.Op("&").Qual(xnsv1Pkg, "UpdateRequest").CustomFunc(multiLineValues, func(fields *g.Group) {
							fields.Id("HeartbeatInterval").Op(":").Qual(durationpbPkg, "New").Call(g.Id("opt").Dot("HeartbeatInterval"))
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
							fields.Id("UpdateWorkflowOptions").Op(":").Id("uopb")
						}),
					),
				),
				g.Nil(),
			)
		})
}

func (svc *Service) genXNSUpdateOptions(f *g.File, update string) {
	typeName := toCamel("%sUpdateOptions", update)

	f.Commentf("%s are used to configure a(n) %s update execution", typeName, update)
	f.Type().Id(typeName).Struct(
		g.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		g.Id("HeartbeatInterval").Qual("time", "Duration"),
		g.Id("UpdateWorkflowOptions").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest"),
	)
}

func (svc *Service) genXNSUpdateHandleImpl(f *g.File, update string) {
	typeName := toLowerCamel("%sHandle", update)
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a(n) %s implementation", typeName, toCamel("%sHandle", update))
	f.Type().Id(typeName).Struct(
		g.Id("cancel").Func().Params(),
		g.Id("future").Qual(workflowPkg, "Future"),
		g.Id("id").String(),
	)

	f.Comment("Cancel the underlying workflow update")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			g.Id("r").Dot("cancel").Call(),
			g.If(
				g.ListFunc(func(ls *g.Group) {
					if hasOutput {
						ls.Id("_")
					}
					ls.Err()
				}).Op(":=").Id("r").Dot("Get").Call(g.Id("ctx")),
				g.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(g.Err(), g.Qual(workflowPkg, "ErrCanceled")),
			).Block(
				g.Return(g.Err()),
			),
			g.Return(g.Nil()),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		Block(
			g.Return(g.Id("r").Dot("future")),
		)

	f.Comment("Get blocks on activity completion and returns the underlying update result")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			fn.If(
				g.Err().Op(":=").Id("r").Dot("future").Dot("Get").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasOutput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				g.Err().Op("!=").Nil(),
			).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.ReturnFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("&").Id("resp")
				}
				returnVals.Nil()
			})
		})

	f.Comment("ID returns the underlying workflow id")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("ID").
		Params().
		String().
		Block(
			g.Return(g.Id("r").Dot("id")),
		)
}

func (svc *Service) genXNSUpdateHandleInterface(f *g.File, update string) {
	typeName := toCamel("%sHandle", update)
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a handle to a %s workflow update", typeName, update)
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		methods.Comment("Cancel cancels the update activity")
		methods.Id("Cancel").
			Params(g.Qual(workflowPkg, "Context")).
			Error()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future")

		methods.Comment("Get blocks on update completion and returns the result")
		methods.Id("Get").
			Params(g.Qual(workflowPkg, "Context")).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			})

		methods.Comment("ID returns the update id")
		methods.Id("ID").
			Params().
			String()
	})
}

func (svc *Service) genXNSWorkflowFunction(f *g.File, workflow string) {
	methodName := workflow
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s workflow and blocks until error or response received", methodName, svc.fqnForWorkflow(workflow))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sWorkflowOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.List(g.Id("run"), g.Err()).Op(":=").Id(toCamel("%sAsync", workflow)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			}),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			),
			g.Return(g.Id("run").Dot("Get").Call(g.Id("ctx"))),
		)
}

func (svc *Service) genXNSWorkflowFunctionAsync(f *g.File, workflow string) {
	methodName := toCamel("%sAsync", workflow)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	// hasOutput := !isEmpty(method.Output)
	opts := svc.workflows[workflow]

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a(n) %s workflow and blocks until error or response received", methodName, svc.fqnForWorkflow(workflow))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sWorkflowOptions", workflow))
		}).
		Params(
			g.Id(toCamel("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(toCamel("%sWorkflowOptions", workflow)).Values()
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("opt").Op("=").Id("opts").Index(g.Lit(0)),
			)
			fn.If(g.Id("opt").Dot("HeartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("opt").Dot("HeartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30),
			)
			fn.Line()

			// build and set activity options
			fn.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(g.Id("ctx"))
			fn.If(g.Id("opt").Dot("ActivityOptions").Op("!=").Nil()).Block(
				g.Id("ao").Op("=").Op("*").Id("opt").Dot("ActivityOptions"),
			)
			fn.If(g.Id("ao").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("HeartbeatTimeout").Op("=").Id("opt").Dot("HeartbeatInterval").Op("*").Lit(2),
			)
			fn.If(g.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0).Op("&&").Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).BlockFunc(func(bl *g.Group) {
				if d := opts.GetExecutionTimeout(); d.IsValid() {
					t := d.AsDuration() + (time.Minute * 5)
					bl.Id("ao").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(t.Nanoseconds(), 10)).Comment(t.String())
				} else {
					bl.Id("ao").Dot("StartToCloseTimeout").Op("=").Qual("time", "Hour").Op("*").Lit(24)
				}
			})
			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(g.Id("ctx"), g.Id("ao"))
			fn.Line()

			// build start workflow options
			fn.Id("wo").Op(":=").Qual(clientPkg, "StartWorkflowOptions").Values()
			fn.If(g.Id("opt").Dot("StartWorkflowOptions").Op("!=").Nil()).Block(
				g.Id("wo").Op("=").Op("*").Id("opt").Dot("StartWorkflowOptions"),
			)
			// set workflow id if unset and  id field and/or prefix defined
			if idExpr := opts.GetId(); idExpr != "" {
				fn.If(g.Id("wo").Dot("ID").Op("==").Lit("")).Block(
					g.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
						args.Qual(string(svc.File.GoImportPath), toCamel("%sIDExpression", workflow))
						if hasInput {
							args.Id("req").Dot("ProtoReflect").Call()
						} else {
							args.Nil()
						}
					}),
					g.If(g.Err().Op("!=").Nil()).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error evaluating id expression for %q workflow: %%w", workflow)), g.Err())),
					),
					g.Id("wo").Dot("ID").Op("=").Id("id"),
				)
			}
			fn.If(g.Id("wo").Dot("ID").Op("==").Lit("")).Block(
				g.List(g.Id("id"), g.Err()).Op(":=").Qual(uuidPkg, "NewRandom").Call(),
				g.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error generating workflow id: %w"), g.Err())),
				),
				g.Id("wo").Dot("ID").Op("=").Id("id").Dot("String").Call(),
			)
			fn.Line()

			// marshal start workflow options
			fn.List(g.Id("swo"), g.Err()).Op(":=").Qual(xnsPkg, "MarshalStartWorkflowOptions").Call(g.Id("wo"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling start workflow options: %w"), g.Err())),
			)
			fn.Line()

			// marshal workflow request
			if hasInput {
				fn.List(g.Id("wreq"), g.Err()).Op(":=").Qual(anypbPkg, "New").Call(g.Id("req"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling workflow request: %w"), g.Err())),
				)
				fn.Line()
			}

			// create cancellable context
			fn.List(g.Id("ctx"), g.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(g.Id("ctx"))
			fn.Line()

			// return run with execute activity future
			fn.Return(
				g.Op("&").Id(toLowerCamel("%sRun", workflow)).Custom(multiLineValues,
					g.Id("cancel").Op(":").Id("cancel"),
					g.Id("id").Op(":").Id("wo").Dot("ID"),
					g.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						g.Id("ctx"),
						g.Qual(string(svc.File.GoImportPath), toCamel("%sWorkflowName", workflow)),
						g.Op("&").Qual(xnsv1Pkg, "WorkflowRequest").CustomFunc(multiLineValues, func(fields *g.Group) {
							fields.Id("Detached").Op(":").Id("opt").Dot("Detached")
							fields.Id("HeartbeatInterval").Op(":").Qual(durationpbPkg, "New").Call(g.Id("opt").Dot("HeartbeatInterval"))
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
							fields.Id("StartWorkflowOptions").Op(":").Id("swo")
						}),
					),
				),
				g.Nil(),
			)
		})
}

func (svc *Service) genXNSWorkflowOptions(f *g.File, workflow string) {
	typeName := toCamel("%sWorkflowOptions", workflow)

	f.Commentf("%s are used to configure a(n) %s workflow execution", typeName, workflow)
	f.Type().Id(typeName).Struct(
		g.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		g.Id("Detached").Bool(),
		g.Id("HeartbeatInterval").Qual("time", "Duration"),
		g.Id("StartWorkflowOptions").Op("*").Qual(clientPkg, "StartWorkflowOptions"),
	)
}

func (svc *Service) genXNSWorkflowRunImpl(f *g.File, workflow string) {
	typeName := toLowerCamel("%sRun", workflow)
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a(n) %s implementation", typeName, toCamel("%sRun", workflow))
	f.Type().Id(typeName).Struct(
		g.Id("cancel").Func().Params(),
		g.Id("future").Qual(workflowPkg, "Future"),
		g.Id("id").String(),
	)

	f.Comment("Cancel the underlying workflow execution")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			g.If(g.Id("r").Dot("cancel").Op("!=").Nil()).Block(
				g.Id("r").Dot("cancel").Call(),
				g.If(
					g.ListFunc(func(ls *g.Group) {
						if hasOutput {
							ls.Id("_")
						}
						ls.Err()
					}).Op(":=").Id("r").Dot("Get").Call(g.Id("ctx")),
					g.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(g.Err(), g.Qual(workflowPkg, "ErrCanceled")),
				).Block(
					g.Return(g.Err()),
				),
				g.Return(g.Nil()),
			),
			g.Return(g.Id(toCamel("Cancel%sWorkflow", svc.GoName)).Call(g.Id("ctx"), g.Id("r").Dot("id"), g.Lit(""))),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		Block(
			g.Return(g.Id("r").Dot("future")),
		)

	f.Comment("Get blocks on activity completion and returns the underlying workflow result")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			fn.If(
				g.Err().Op(":=").Id("r").Dot("future").Dot("Get").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasOutput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				g.Err().Op("!=").Nil(),
			).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.ReturnFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("&").Id("resp")
				}
				returnVals.Nil()
			})
		})

	f.Comment("ID returns the underlying workflow id")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("ID").
		Params().
		String().
		Block(
			g.Return(g.Id("r").Dot("id")),
		)

	for i := range opts.GetQuery() {
		query := opts.GetQuery()[i].GetRef()
		handler := svc.methods[query]
		handlerInput := !isEmpty(handler.Input)
		handlerOutput := !isEmpty(handler.Output)

		methodName := toCamel(query)
		if handler.Comments.Leading.String() != "" {
			f.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
		} else {
			f.Commentf("%s executes a(n) %s query and blocks until completion", methodName, svc.fqnForQuery(query))
		}
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(methodName).
			ParamsFunc(func(args *g.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
				}
				args.Id("opts").Op("...").Op("*").Id(toCamel("%sQueryOptions", query))
			}).
			ParamsFunc(func(returnVals *g.Group) {
				if handlerOutput {
					returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Output))
				}
				returnVals.Error()
			}).
			Block(
				g.Return(g.Id(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if handlerInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				})),
			)

		methodName = toCamel("%sAsync", query)
		if handler.Comments.Leading.String() != "" {
			f.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
		} else {
			f.Commentf("%s executes a(n) %s query and returns a handle to the underlying activity", methodName, svc.fqnForQuery(query))
		}
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(methodName).
			ParamsFunc(func(args *g.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
				}
				args.Id("opts").Op("...").Op("*").Id(toCamel("%sQueryOptions", query))
			}).
			Params(
				g.Id(toCamel("%sQueryHandle", query)),
				g.Error(),
			).
			Block(
				g.Return(g.Id(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if handlerInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				})),
			)
	}

	for i := range opts.GetSignal() {
		signal := opts.GetSignal()[i].GetRef()
		handler := svc.methods[signal]
		handlerInput := !isEmpty(handler.Input)

		methodName := toCamel(signal)
		if handler.Comments.Leading.String() != "" {
			f.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
		} else {
			f.Commentf("%s executes a(n) %s signal and blocks until the underlying activity completes", methodName, svc.fqnForSignal(signal))
		}
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(methodName).
			ParamsFunc(func(args *g.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
				}
				args.Id("opts").Op("...").Op("*").Id(toCamel("%sSignalOptions", signal))
			}).
			Error().
			Block(
				g.Return(g.Id(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if handlerInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				})),
			)

		methodName = toCamel("%sAsync", signal)
		if handler.Comments.Leading.String() != "" {
			f.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
		} else {
			f.Commentf("%s executes a(n) %s signal and returns a handle to the underlying activity", methodName, svc.fqnForSignal(signal))
		}
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(methodName).
			ParamsFunc(func(args *g.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
				}
				args.Id("opts").Op("...").Op("*").Id(toCamel("%sSignalOptions", signal))
			}).
			Params(
				g.Id(toCamel("%sSignalHandle", signal)),
				g.Error(),
			).
			Block(
				g.Return(g.Id(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if handlerInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				})),
			)
	}

	for i := range opts.GetUpdate() {
		update := opts.GetUpdate()[i].GetRef()
		handler := svc.methods[update]
		handlerInput := !isEmpty(handler.Input)
		handlerOutput := !isEmpty(handler.Output)

		methodName := toCamel(update)
		if handler.Comments.Leading.String() != "" {
			f.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
		} else {
			f.Commentf("%s executes a(n) %s update and blocks until completion", methodName, svc.fqnForUpdate(update))
		}
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(methodName).
			ParamsFunc(func(args *g.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
				}
				args.Id("opts").Op("...").Op("*").Id(toCamel("%sUpdateOptions", update))
			}).
			ParamsFunc(func(returnVals *g.Group) {
				if handlerOutput {
					returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Output))
				}
				returnVals.Error()
			}).
			Block(
				g.Return(g.Id(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if handlerInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				})),
			)

		methodName = toCamel("%sAsync", update)
		if handler.Comments.Leading.String() != "" {
			f.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
		} else {
			f.Commentf("%s executes a(n) %s update and returns a handle to the underlying activity", methodName, svc.fqnForUpdate(update))
		}
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(methodName).
			ParamsFunc(func(args *g.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
				}
				args.Id("opts").Op("...").Op("*").Id(toCamel("%sUpdateOptions", update))
			}).
			Params(
				g.Id(toCamel("%sHandle", update)),
				g.Error(),
			).
			Block(
				g.Return(g.Id(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if handlerInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				})),
			)
	}
}

func (svc *Service) genXNSWorkflowRunInterface(f *g.File, workflow string) {
	typeName := toCamel("%sRun", workflow)
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a handle to a %s workflow execution", typeName, workflow)
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		methods.Comment("Cancel cancels the workflow")
		methods.Id("Cancel").
			Params(g.Qual(workflowPkg, "Context")).
			Error()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future")

		methods.Comment("Get returns the inner workflow.Future")
		methods.Id("Get").
			Params(g.Qual(workflowPkg, "Context")).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			})

		methods.Comment("ID returns the workflow id")
		methods.Id("ID").
			Params().
			String()

		for i := range opts.GetQuery() {
			query := opts.GetQuery()[i].GetRef()
			handler := svc.methods[query]
			handlerInput := !isEmpty(handler.Input)
			handlerOutput := !isEmpty(handler.Output)

			// synchronous
			methodName := toCamel(query)
			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s query and blocks until completion", methodName, svc.fqnForQuery(query))
			}
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Id(toCamel("%sQueryOptions", query))
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if handlerOutput {
						returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Output))
					}
					returnVals.Error()
				})

			// async
			methodName = toCamel("%sAsync", query)
			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s query and returns a handle to the underlying activity", methodName, svc.fqnForQuery(query))
			}
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Id(toCamel("%sQueryOptions", query))
				}).
				Params(
					g.Id(toCamel("%sQueryHandle", query)),
					g.Error(),
				)
		}

		for i := range opts.GetSignal() {
			signal := opts.GetSignal()[i].GetRef()
			handler := svc.methods[signal]
			handlerInput := !isEmpty(handler.Input)

			// synchronnous
			methodName := toCamel(signal)
			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s signal and blocks until completion", methodName, svc.fqnForSignal(signal))
			}
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Id(toCamel("%sSignalOptions", signal))
				}).
				Error()

			// async
			methodName = toCamel("%sAsync", signal)
			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s signal and returns a handle to the underlying activity", methodName, svc.fqnForSignal(signal))
			}
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Id(toCamel("%sSignalOptions", signal))
				}).
				Params(
					g.Id(toCamel("%sSignalHandle", signal)),
					g.Error(),
				)
		}

		for i := range opts.GetUpdate() {
			update := opts.GetUpdate()[i].GetRef()
			handler := svc.methods[update]
			handlerInput := !isEmpty(handler.Input)
			handlerOutput := !isEmpty(handler.Output)

			// synchronous
			methodName := toCamel(update)
			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s update and blocks until completion", methodName, svc.fqnForUpdate(update))
			}
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Id(toCamel("%sUpdateOptions", update))
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if handlerOutput {
						returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Output))
					}
					returnVals.Error()
				})

			// async
			methodName = toCamel("%sAsync", update)
			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s update and returns a handle to the underlying activity", methodName, svc.fqnForUpdate(update))
			}
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Id(toCamel("%sUpdateOptions", update))
				}).
				Params(
					g.Id(toCamel("%sHandle", update)),
					g.Error(),
				)
		}
	})
}

func (svc *Service) genXNSWorkflowWithStartFunction(f *g.File, workflow, signal string) {
	methodName := toCamel("%sWith%s", workflow, signal)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := svc.methods[signal]
	handlerInput := !isEmpty(handler.Input)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s sends a(n) %s signal to a %s workflow, starting it if necessary, and blocks until the workflow completes", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
	}

	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			if handlerInput {
				args.Id("signal").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sWorkflowOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.List(g.Id("run"), g.Err()).Op(":=").Id(toCamel("%sWith%sAsync", workflow, signal)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req")
				}
				if handlerInput {
					args.Id("signal")
				}
				args.Id("opts").Op("...")
			}),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			),
			g.Return(g.Id("run").Dot("Get").Call(g.Id("ctx"))),
		)
}

func (svc *Service) genXNSWorkflowWithStartFunctionAsync(f *g.File, workflow, signal string) {
	methodName := toCamel("%sWith%sAsync", workflow, signal)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	handler := svc.methods[signal]
	handlerInput := !isEmpty(handler.Input)
	opts := svc.workflows[workflow]

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and returns a handle to the underlying activity", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
	}
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(method.Input))
			}
			if handlerInput {
				args.Id("signal").Op("*").Qual(string(svc.File.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sWorkflowOptions", workflow))
		}).
		Params(
			g.Id(toCamel("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(toCamel("%sWorkflowOptions", workflow)).Values()
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("opt").Op("=").Id("opts").Index(g.Lit(0)),
			)
			fn.If(g.Id("opt").Dot("HeartbeatInterval").Op("==").Lit(0)).Block(
				g.Id("opt").Dot("HeartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30),
			)
			fn.Line()

			// build and set activity options
			fn.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(g.Id("ctx"))
			fn.If(g.Id("opt").Dot("ActivityOptions").Op("!=").Nil()).Block(
				g.Id("ao").Op("=").Op("*").Id("opt").Dot("ActivityOptions"),
			)
			fn.If(g.Id("ao").Dot("HeartbeatTimeout").Op("==").Lit(0)).Block(
				g.Id("ao").Dot("HeartbeatTimeout").Op("=").Id("opt").Dot("HeartbeatInterval").Op("*").Lit(2),
			)
			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(g.Id("ctx"), g.Id("ao"))
			fn.Line()

			// build start workflow options
			fn.Id("wo").Op(":=").Qual(clientPkg, "StartWorkflowOptions").Values()
			fn.If(g.Id("opt").Dot("StartWorkflowOptions").Op("!=").Nil()).Block(
				g.Id("wo").Op("=").Op("*").Id("opt").Dot("StartWorkflowOptions"),
			)
			// set workflow id if unset and  id field and/or prefix defined
			if idExpr := opts.GetId(); idExpr != "" {
				fn.If(g.Id("wo").Dot("ID").Op("==").Lit("")).Block(
					g.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
						args.Qual(string(svc.File.GoImportPath), toCamel("%sIDExpression", workflow))
						if hasInput {
							args.Id("req").Dot("ProtoReflect").Call()
						} else {
							args.Nil()
						}
					}),
					g.If(g.Err().Op("!=").Nil()).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error evaluating id expression for %q workflow: %%w", workflow)), g.Err())),
					),
					g.Id("wo").Dot("ID").Op("=").Id("id"),
				)
			}
			fn.If(g.Id("wo").Dot("ID").Op("==").Lit("")).Block(
				g.List(g.Id("id"), g.Err()).Op(":=").Qual(uuidPkg, "NewRandom").Call(),
				g.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error generating workflow id: %w"), g.Err())),
				),
				g.Id("wo").Dot("ID").Op("=").Id("id").Dot("String").Call(),
			)
			fn.Line()

			// marshal start workflow options
			fn.List(g.Id("swo"), g.Err()).Op(":=").Qual(xnsPkg, "MarshalStartWorkflowOptions").Call(g.Id("wo"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling start workflow options: %w"), g.Err())),
			)
			fn.Line()

			// marshal workflow request
			if hasInput {
				fn.List(g.Id("wreq"), g.Err()).Op(":=").Qual(anypbPkg, "New").Call(g.Id("req"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling workflow request: %w"), g.Err())),
				)
				fn.Line()
			}

			// marshal signal request
			if handlerInput {
				fn.List(g.Id("wsignal"), g.Err()).Op(":=").Qual(anypbPkg, "New").Call(g.Id("signal"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error marshalling signal request: %w"), g.Err())),
				)
				fn.Line()
			}

			// create cancellable context
			fn.List(g.Id("ctx"), g.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(g.Id("ctx"))
			fn.Line()

			// return run with execute activity future
			fn.Return(
				g.Op("&").Id(toLowerCamel("%sRun", workflow)).Custom(multiLineValues,
					g.Id("cancel").Op(":").Id("cancel"),
					g.Id("id").Op(":").Id("wo").Dot("ID"),
					g.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						g.Id("ctx"),
						g.Lit(fmt.Sprintf("%s.%s", string(svc.Service.Desc.FullName()), toCamel("%sWith%s", workflow, signal))),
						g.Op("&").Qual(xnsv1Pkg, "WorkflowRequest").CustomFunc(multiLineValues, func(fields *g.Group) {
							fields.Id("Detached").Op(":").Id("opt").Dot("Detached")
							fields.Id("HeartbeatInterval").Op(":").Qual(durationpbPkg, "New").Call(g.Id("opt").Dot("HeartbeatInterval"))
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
							if handlerInput {
								fields.Id("Signal").Op(":").Id("wsignal")
							}
							fields.Id("StartWorkflowOptions").Op(":").Id("swo")
						}),
					),
				),
				g.Nil(),
			)
		})
}
