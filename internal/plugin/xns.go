package plugin

import (
	"fmt"
	"strconv"
	"time"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	j "github.com/dave/jennifer/jen"
	"github.com/hako/durafmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	anypbPkg = "google.golang.org/protobuf/types/known/anypb"
	xnsv1Pkg = "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	xnsPkg   = "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
)

func (n *names) xnsActivities() string {
	return n.toLowerCamel("%sActivities", n.Service.GoName)
}

func (n *names) xnsOptionsType() string {
	return n.toCamel("%sOptions", n.Service.GoName)
}

func (n *names) xnsOptionsVar() string {
	return n.toLowerCamel("%sOptions", n.Service.GoName)
}

func (n *names) xnsUpdateHandleIface(update protoreflect.FullName) string {
	return n.toCamel("%sHandle", update)
}

func (n *names) xnsUpdateHandleImpl(update protoreflect.FullName) string {
	return n.toLowerCamel("%sHandle", update)
}

func (n *names) xnsUpdateOptions(update protoreflect.FullName) string {
	return n.toCamel("%sUpdateOptions", update)
}

func (n *names) xnsUpdateOptionsCtor(update protoreflect.FullName) string {
	return n.toCamel("New%sUpdateOptions", update)
}

func (n *names) xnsUpdateWithStartFunction(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%s", workflow, update)
}
func (n *names) xnsUpdateWithStartFunctionAsync(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%sAsync", workflow, update)
}

func (n *names) xnsUpdateWithStartOptions(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%sOptions", workflow, update)
}

func (n *names) xnsUpdateWithStartOptionsCtor(workflow, update protoreflect.FullName) string {
	return n.toCamel("New%sWith%sOptions", workflow, update)
}

func (n *names) xnsWorkflowFunction(workflow protoreflect.FullName) string {
	return n.toCamel("%s", workflow)
}

func (n *names) xnsWorkflowFunctionAsync(workflow protoreflect.FullName) string {
	return n.toCamel("%sAsync", workflow)
}

func (n *names) xnsWorkflowGetFunction(workflow protoreflect.FullName) string {
	return n.toCamel("Get%s", workflow)
}

func (n *names) xnsWorkflowGetFunctionAsync(workflow protoreflect.FullName) string {
	return n.toCamel("Get%sAsync", workflow)
}

func (n *names) xnsWorkflowGetOptions(workflow protoreflect.FullName) string {
	return n.toCamel("Get%sOptions", workflow)
}

func (n *names) xnsWorkflowGetOptionsCtor(workflow protoreflect.FullName) string {
	return n.toCamel("NewGet%sOptions", workflow)
}

func (n *names) xnsWorkflowOptions(workflow protoreflect.FullName) string {
	return n.toCamel("%sWorkflowOptions", workflow)
}

func (n *names) xnsWorkflowOptionsCtor(workflow protoreflect.FullName) string {
	return n.toCamel("New%sWorkflowOptions", workflow)
}

func (n *names) xnsWorkflowRunIface(workflow protoreflect.FullName) string {
	return n.toCamel("%sRun", workflow)
}

func (n *names) xnsWorkflowRunImpl(workflow protoreflect.FullName) string {
	return n.toLowerCamel("%sRun", workflow)
}

func (m *Manifest) renderXNS(f *j.File) {
	m.genXNSRegisterActivities(f)
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSWorkflowOptions(f, workflow)
		m.genXNSWorkflowRunInterface(f, workflow)
		m.genXNSWorkflowRunImpl(f, workflow)
		m.genXNSWorkflowFunction(f, workflow, "")
		m.genXNSWorkflowFunctionAsync(f, workflow, "")
		m.genXNSWorkflowGetFunction(f, workflow)
		m.genXNSWorkflowGetFunctionAsync(f, workflow)

		for _, signal := range m.workflows[workflow].GetSignal() {
			if !signal.GetStart() {
				continue
			}
			m.genXNSWorkflowFunction(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
			m.genXNSWorkflowFunctionAsync(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
		}

		for _, update := range m.workflows[workflow].GetUpdate() {
			if !update.GetStart() {
				continue
			}
			m.genXNSUpdateWithStartOptions(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
			m.genXNSUpdateWithStartFunction(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
			m.genXNSUpdateWithStartFunctionAsync(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
		}
	}
	for _, query := range m.queriesOrdered {
		if m.methods[query].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSQueryOptions(f, query)
		m.genXNSQueryHandleInterface(f, query)
		m.genXNSQueryHandleImpl(f, query)
		m.genXNSQueryFunction(f, query)
		m.genXNSQueryFunctionAsync(f, query)
	}
	for _, signal := range m.signalsOrdered {
		if m.methods[signal].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSSignalOptions(f, signal)
		m.genXNSSignalHandleInterface(f, signal)
		m.genXNSSignalHandleImpl(f, signal)
		m.genXNSSignalFunction(f, signal)
		m.genXNSSignalFunctionAsync(f, signal)
	}
	for _, update := range m.updatesOrdered {
		if m.methods[update].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSUpdateOptions(f, update)
		m.genXNSUpdateHandleInterface(f, update)
		m.genXNSUpdateHandleImpl(f, update)
		m.genXNSUpdateFunction(f, update)
		m.genXNSUpdateFunctionAsync(f, update)
	}
	m.genXNSCancelWorkflowFunction(f)

	m.genXNSActivities(f)
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSActivitiesWorkflowGetMethod(f, workflow)
		m.genXNSActivitiesWorkflowMethod(f, workflow, "")
		for _, signal := range m.workflows[workflow].GetSignal() {
			if !signal.GetStart() {
				continue
			}
			m.genXNSActivitiesWorkflowMethod(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
		}
		for _, update := range m.workflows[workflow].GetUpdate() {
			if !update.GetStart() {
				continue
			}
			m.genXNSActivitiesUpdateWithStartMethod(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
		}
	}
	for _, query := range m.queriesOrdered {
		if m.methods[query].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSActivitiesQueryMethod(f, query)
	}
	for _, signal := range m.signalsOrdered {
		if m.methods[signal].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSActivitiesSignalMethod(f, signal)
	}
	for _, update := range m.updatesOrdered {
		if m.methods[update].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genXNSActivitiesUpdateMethod(f, update)
	}
}

func (m *Manifest) genXNSActivities(f *j.File) {
	typeName := m.toLowerCamel("%sActivities", m.GoName)

	f.Commentf("%s provides activities that can be used to interact with a(n) %s service's workflow, queries, signals, and updates across namespaces", typeName, m.GoName)
	f.Type().Id(typeName).Struct(
		j.Id("client").Qual(string(m.File.GoImportPath), m.toCamel("%sClient", m.GoName)),
	)

	f.Comment("CancelWorkflow cancels an existing workflow execution")
	f.Func().
		Params(
			j.Id("a").Op("*").Id(typeName),
		).
		Id("CancelWorkflow").
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("workflowID").String(),
			j.Id("runID").String(),
		).
		Error().
		Block(
			j.Return(
				j.Id("a").Dot("client").Dot("CancelWorkflow").Call(j.Id("ctx"), j.Id("workflowID"), j.Id("runID")),
			),
		)
}

func (m *Manifest) genXNSActivitiesQueryMethod(f *j.File, query protoreflect.FullName) {
	methodName := m.methods[query].GoName
	method := m.methods[query]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentf(f, methodSet(method), "%s executes a(n) %s query via an activity", methodName, m.fqnForQuery(query))
	f.Func().
		Params(
			j.Id("a").Op("*").Id(m.toLowerCamel("%sActivities", m.GoName)),
		).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("input").Op("*").Qual(xnsv1Pkg, "QueryRequest"),
		).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Id("resp").Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Err().Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated query detected"), j.Lit("query"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sQueryName", query))).Line()
			}
			if hasInput {
				fn.Comment("unmarshal query request")
				fn.Var().Id("req").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Input))
				fn.If(j.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(j.Op("&").Id("req")), j.Err().Op("!=").Nil()).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
							j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
								multiLineArgs,
								j.Qual("fmt", "Sprintf").Call(
									j.Lit(fmt.Sprintf("error unmarshalling query request of type %%s as %s.%s", string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))),
									j.Id("input").Dot("Request").Dot("GetTypeUrl").Call(),
								),
								j.Lit("InvalidArgument"),
								j.Err(),
							),
						)
					}),
				)
			}

			fn.Comment("execute signal in child goroutine")
			fn.Id("doneCh").Op(":=").Make(j.Chan().Struct())
			fn.Go().Func().Params().Block(
				j.ListFunc(func(ls *j.Group) {
					if hasOutput {
						ls.Id("resp")
					}
					ls.Err()
				}).Op("=").Id("a").Dot("client").Dot(methodName).CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("input").Dot("GetWorkflowId").Call()
					args.Id("input").Dot("GetRunId").Call()
					if hasInput {
						args.Op("&").Id("req")
					}
				}),
				j.Close(j.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			fn.If(j.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				j.Id("heartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(10),
			)
			fn.Line()

			fn.Comment("heartbeat activity while waiting for signal to complete")
			fn.For().Block(
				j.Select().Block(
					j.Case(j.Op("<-").Qual("time", "After").Call(j.Id("heartbeatInterval"))).Block(
						j.Qual(activityPkg, "RecordHeartbeat").Call(j.Id("ctx")),
					),
					j.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).Block(
						j.ReturnFunc(func(returnVals *j.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
								j.Id("ctx").Dot("Err").Call(),
							)
						}),
					),
					j.Case(j.Op("<-").Id("doneCh")).Block(
						j.ReturnFunc(func(returnVals *j.Group) {
							if hasOutput {
								returnVals.Id("resp")
							}
							returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
						}),
					),
				),
			)
		})
}

func (m *Manifest) genXNSActivitiesSignalMethod(f *j.File, signal protoreflect.FullName) {
	methodName := m.methods[signal].GoName
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentf(f, methodSet(method), "%s executes a(n) %s signal via an activity", methodName, m.fqnForSignal(signal))
	f.Func().
		Params(
			j.Id("a").Op("*").Id(m.toLowerCamel("%sActivities", m.GoName)),
		).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("input").Op("*").Qual(xnsv1Pkg, "SignalRequest"),
		).
		Params(j.Err().Error()).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated signal detected"), j.Lit("signal"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sSignalName", signal))).Line()
			}
			if hasInput {
				fn.Comment("unmarshal signal request")
				fn.Var().Id("req").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				fn.If(j.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(j.Op("&").Id("req")), j.Err().Op("!=").Nil()).Block(
					j.Return(j.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
						j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
							multiLineArgs,
							j.Qual("fmt", "Sprintf").Call(
								j.Lit(fmt.Sprintf("error unmarshalling signal request of type %%s as %s.%s", string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))),
								j.Id("input").Dot("Request").Dot("GetTypeUrl").Call(),
							),
							j.Lit("InvalidArgument"),
							j.Err(),
						),
					)),
				)
			}

			fn.Comment("execute signal in child goroutine")
			fn.Id("doneCh").Op(":=").Make(j.Chan().Struct())
			fn.Go().Func().Params().Block(
				j.Err().Op("=").Id("a").Dot("client").Dot(methodName).CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("input").Dot("GetWorkflowId").Call()
					args.Id("input").Dot("GetRunId").Call()
					if hasInput {
						args.Op("&").Id("req")
					}
				}),
				j.Close(j.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			fn.If(j.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				j.Id("heartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(10),
			)
			fn.Line()

			fn.Comment("heartbeat activity while waiting for signal to complete")
			fn.For().Block(
				j.Select().Block(
					j.Case(j.Op("<-").Qual("time", "After").Call(j.Id("heartbeatInterval"))).Block(
						j.Qual(activityPkg, "RecordHeartbeat").Call(j.Id("ctx")),
					),
					j.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).Block(
						j.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Id("ctx").Dot("Err").Call()),
					),
					j.Case(j.Op("<-").Id("doneCh")).Block(
						j.Return(j.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())),
					),
				),
			)
		})
}

func (m *Manifest) genXNSActivitiesUpdateMethod(f *j.File, update protoreflect.FullName) {
	methodName := m.methods[update].GoName
	method := m.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentf(f, methodSet(method), "%s executes a(n) %s update via an activity", methodName, m.fqnForUpdate(update))
	f.Func().
		Params(
			j.Id("a").Op("*").Id(m.toLowerCamel("%sActivities", m.GoName)),
		).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("input").Op("*").Qual(xnsv1Pkg, "UpdateRequest"),
		).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Id("resp").Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Err().Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated update detected"), j.Lit("update"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sUpdateName", update))).Line()
			}
			fn.Var().Id("handle").Qual(string(m.File.GoImportPath), m.toCamel("%sHandle", update))
			fn.If(j.Qual(activityPkg, "HasHeartbeatDetails").Call(j.Id("ctx"))).Block(
				j.Comment("extract update id from heartbeat details"),
				j.Var().Id("updateID").String(),
				j.If(
					j.Err().Op(":=").Qual(activityPkg, "GetHeartbeatDetails").Call(j.Id("ctx"), j.Op("&").Id("updateID")),
					j.Err().Op("!=").Nil(),
				).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
					}),
				),
				j.Line(),
				j.Comment("retrieve handle for existing update"),
				j.List(j.Id("handle"), j.Err()).Op("=").Id("a").Dot("client").Dot(m.toCamel("Get%s", update)).Call(
					j.Id("ctx"),
					j.Qual(clientPkg, "GetWorkflowUpdateHandleOptions").Custom(
						multiLineValues,
						j.Id("WorkflowID").Op(":").Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetWorkflowId").Call(),
						j.Id("RunID").Op(":").Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetRunId").Call(),
						j.Id("UpdateID").Op(":").Id("updateID"),
					),
				),
				j.If(j.Err().Op("!=").Nil()).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
					}),
				),
			).Else().BlockFunc(func(bl *j.Group) {
				if hasInput {
					bl.Comment("unmarshal update request")
					bl.Var().Id("req").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					bl.If(j.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(j.Op("&").Id("req")), j.Err().Op("!=").Nil()).Block(
						j.ReturnFunc(func(returnVals *j.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
								j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
									multiLineArgs,
									j.Qual("fmt", "Sprintf").Call(
										j.Lit(fmt.Sprintf("error unmarshalling update request of type %%s as %s.%s", string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))),
										j.Id("input").Dot("Request").Dot("GetTypeUrl").Call(),
									),
									j.Lit("InvalidArgument"),
									j.Err(),
								),
							)
						}),
					)
					bl.Line()
				}

				bl.Id("uo").Op(":=").Qual(xnsPkg, "UnmarshalUpdateWorkflowOptions").Call(
					j.Id("input").Dot("GetUpdateWorkflowOptions").Call(),
				)
				bl.Id("uo").Dot("WaitForStage").Op("=").Qual(clientPkg, "WorkflowUpdateStageAccepted").Line()

				bl.Comment("initialize update execution")
				bl.List(j.Id("handle"), j.Err()).Op("=").Id("a").Dot("client").Dot(m.toCamel("%sAsync", methodName)).CustomFunc(multiLineArgs, func(args *j.Group) {
					args.Id("ctx")
					args.Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetWorkflowId").Call()
					args.Id("input").Dot("GetUpdateWorkflowOptions").Call().Dot("GetRunId").Call()
					if hasInput {
						args.Op("&").Id("req")
					}
					args.Qual(string(m.File.GoImportPath), m.toCamel("New%sOptions", update)).
						Call().
						Dot("WithUpdateWorkflowOptions").
						Call(j.Id("uo"))
				})
				bl.If(j.Err().Op("!=").Nil()).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
					}),
				)
				bl.Qual(activityPkg, "RecordHeartbeat").Call(j.Id("ctx"), j.Id("handle").Dot("UpdateID").Call())
			})
			fn.Line()

			fn.Comment("wait for update to complete in child goroutine")
			fn.Id("doneCh").Op(":=").Make(j.Chan().Struct())
			fn.Go().Func().Params().Block(
				j.ListFunc(func(ls *j.Group) {
					if hasOutput {
						ls.Id("resp")
					}
					ls.Err()
				}).Op("=").Id("handle").Dot("Get").Call(j.Id("ctx")),
				j.Close(j.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			fn.If(j.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				j.Id("heartbeatInterval").Op("=").Qual("time", "Minute"),
			)
			fn.Line()

			fn.Comment("heartbeat activity while waiting for workflow update to complete")
			fn.For().Block(
				j.Select().Block(
					j.Case(j.Op("<-").Qual("time", "After").Call(j.Id("heartbeatInterval"))).Block(
						j.Qual(activityPkg, "RecordHeartbeat").Call(j.Id("ctx"), j.Id("handle").Dot("UpdateID").Call()),
					),
					j.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).Block(
						j.ReturnFunc(func(returnVals *j.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Id("ctx").Dot("Err").Call())
						}),
					),
					j.Case(j.Op("<-").Id("doneCh")).Block(
						j.ReturnFunc(func(returnVals *j.Group) {
							if hasOutput {
								returnVals.Id("resp")
							}
							returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
						}),
					),
				),
			)
		})
}

func (m *Manifest) genXNSActivitiesUpdateWithStartMethod(f *j.File, workflow, update protoreflect.FullName) {
	activities := m.Names().xnsActivities()
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)
	// hasWorkflowOutput := !isEmpty(method.Output)
	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)
	hasUpdateOutput := !isEmpty(handler.Output)

	methodName := m.Names().xnsUpdateWithStartFunction(workflow, update)
	methodOptionsCtor := m.Names().clientUpdateWithStartOptionsCtor(workflow, update)
	asyncName := m.Names().xnsUpdateWithStartFunctionAsync(workflow, update)
	xnsOptions := m.Names().xnsOptionsVar()

	commentf(f, methodSet(method, handler), "%s executes a(n) %s workflow with a(n) %s update via an activity", methodName, m.fqnForWorkflow(workflow), m.fqnForUpdate(update))
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("a").Op("*").Id(activities)
		}).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("input").Op("*").Qual(xnsv1Pkg, "UpdateWithStartRequest")
		}).
		ParamsFunc(func(g *j.Group) {
			if hasUpdateOutput {
				g.Id("out").Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			g.Err().Error()
		}).
		BlockFunc(func(g *j.Group) {
			// log deprecation warnings
			methodDeprecated, updateDeprecated := isDeprecated(method), isDeprecated(handler)
			if methodDeprecated || updateDeprecated {
				if methodDeprecated {
					g.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated workflow detected"), j.Lit("workflow"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sWorkflowName", workflow))).Line()
				}
				if updateDeprecated {
					g.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated update detected"), j.Lit("update"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sUpdateName", update))).Line()
				}
				g.Line()
			}

			// unmarshal workflow and update request
			if hasWorkflowInput {
				g.Comment("unmarshal workflow request")
				g.Var().Id("req").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				g.IfFunc(func(g *j.Group) {
					g.Err().Op(":=").Id("input").Dot("GetInput").Call().Dot("UnmarshalTo").CallFunc(func(g *j.Group) {
						g.Op("&").Id("req")
					})
					g.Err().Op("!=").Nil()
				}).BlockFunc(func(g *j.Group) {
					g.ReturnFunc(func(g *j.Group) {
						if hasUpdateOutput {
							g.Nil()
						}
						g.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
							j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
								multiLineArgs,
								j.Qual("fmt", "Sprintf").Call(
									j.Lit(fmt.Sprintf("error unmarshalling workflow request of type %%s as %s.%s", string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))),
									j.Id("input").Dot("GetInput").Call().Dot("GetTypeUrl").Call(),
								),
								j.Lit("InvalidArgument"),
								j.Err(),
							),
						)
					})
				})
				g.Line()
			}

			if hasUpdateInput {
				g.Comment("unmarshal update request")
				g.Var().Id("update").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				g.IfFunc(func(g *j.Group) {
					g.Err().Op(":=").Id("input").Dot("GetUpdate").Call().Dot("UnmarshalTo").CallFunc(func(g *j.Group) {
						g.Op("&").Id("update")
					})
					g.Err().Op("!=").Nil()
				}).BlockFunc(func(g *j.Group) {
					g.ReturnFunc(func(g *j.Group) {
						if hasUpdateOutput {
							g.Nil()
						}
						g.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
							j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
								multiLineArgs,
								j.Qual("fmt", "Sprintf").Call(
									j.Lit(fmt.Sprintf("error unmarshalling update request of type %%s as %s.%s", string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))),
									j.Id("input").Dot("GetUpdate").Call().Dot("GetTypeUrl").Call(),
								),
								j.Lit("InvalidArgument"),
								j.Err(),
							),
						)
					})
				})
				g.Line()
			}

			// unmarshal workflow and update options
			g.Comment("unmarshal workflow and update options")
			g.Id("swo").Op(":=").Qual(xnsPkg, "UnmarshalStartWorkflowOptions").CallFunc(func(g *j.Group) {
				g.Id("input").Dot("GetStartWorkflowOptions").Call()
			})
			g.Id("uwo").Op(":=").Qual(xnsPkg, "UnmarshalUpdateWorkflowOptions").CallFunc(func(g *j.Group) {
				g.Id("input").Dot("GetUpdateWorkflowOptions").Call()
			})
			g.Line()
			withWorkflowOptions := m.toCamel("With%sOptions", workflow)
			workflowOptionsCtor := m.Names().clientWorkflowOptionsCtor(workflow)
			withUpdateOptions := m.toCamel("With%sOptions", update)
			updateOptionsCtor := m.Names().clientUpdateOptionsCtor(update)

			// execute update with start asyncronously
			g.Comment("execute update with start asynchronously")
			g.List(j.Id("handle"), j.Id("run"), j.Err()).Op(":=").Id("a").Dot("client").Dot(asyncName).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasWorkflowInput {
					g.Op("&").Id("req")
				}
				if hasUpdateInput {
					g.Op("&").Id("update")
				}
				g.Qual(string(m.File.GoImportPath), methodOptionsCtor).Call().
					Dot(withWorkflowOptions).
					CallFunc(func(g *j.Group) {
						g.Qual(string(m.File.GoImportPath), workflowOptionsCtor).Call().
							Dot("WithStartWorkflowOptions").Call(j.Id("swo"))
					}).
					Dot(withUpdateOptions).
					CallFunc(func(g *j.Group) {
						g.Qual(string(m.File.GoImportPath), updateOptionsCtor).Call().
							Dot("WithUpdateWorkflowOptions").Call(j.Id("uwo"))
					})
			})
			g.IfFunc(func(g *j.Group) {
				g.Err().Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Nil()
					}
					g.Id(xnsOptions).Dot("convertError").Call(j.Err())
				})
			})
			g.Line()

			// return early if detached
			g.Comment("return early if detached")
			g.IfFunc(func(g *j.Group) {
				g.Id("input").Dot("GetDetached").Call()
			}).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Nil()
					}
					g.Nil()
				})
			})
			g.Line()

			// initialize heartbeat interval duration
			g.Comment("initialize heartbeat interval duration")
			g.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			g.IfFunc(func(g *j.Group) {
				g.Id("heartbeatInterval").Op("==").Lit(0)
			}).BlockFunc(func(g *j.Group) {
				g.Id("heartbeatTimeout").Op(":=").Qual(activityPkg, "GetInfo").CallFunc(func(g *j.Group) {
					g.Id("ctx")
				}).Dot("HeartbeatTimeout")
				g.IfFunc(func(g *j.Group) {
					g.Id("heartbeatTimeout").Op(">").Lit(0)
				}).BlockFunc(func(g *j.Group) {
					g.Id("heartbeatInterval").Op("=").Id("heartbeatTimeout").Op("/").Lit(2)
				}).Else().BlockFunc(func(g *j.Group) {
					g.Id("heartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30)
				})
			})
			g.Line()

			// wait on update in child goroutine
			g.Comment("wait for update to complete in child goroutine")
			g.Id("doneCh").Op(":=").Make(j.Chan().Struct())
			g.Go().Func().Params().BlockFunc(func(g *j.Group) {
				g.Defer().Close(j.Id("doneCh"))
				g.ListFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Id("out")
					}
					g.Err()
				}).Op("=").Id("handle").Dot("Get").CallFunc(func(g *j.Group) {
					g.Id("ctx")
				})
			}).Call()
			g.Line()

			// heartbeat activity while waiting for update to complete
			g.Comment("heartbeat activity while waiting for update to complete")
			g.For().BlockFunc(func(g *j.Group) {
				g.Select().BlockFunc(func(g *j.Group) {
					// record heartbeat every heartbeatInterval
					g.Case(j.Op("<-").Qual("time", "After").Call(j.Id("heartbeatInterval"))).BlockFunc(func(g *j.Group) {
						g.Qual(activityPkg, "RecordHeartbeat").Call(j.Id("ctx"), j.Id("handle").Dot("UpdateID").Call())
					})
					g.Line()

					// return retryable error if the worker is stopping
					g.Case(j.Op("<-").Qual(activityPkg, "GetWorkerStopChannel").Call(j.Id("ctx"))).BlockFunc(func(g *j.Group) {
						g.ReturnFunc(func(g *j.Group) {
							if hasUpdateOutput {
								g.Nil()
							}
							g.Id(xnsOptions).Dot("convertError").Call(
								j.Qual(temporalPkg, "NewApplicationError").Call(
									j.Lit("worker is stopping"),
									j.Lit("WorkerStopping"),
								),
							)
						})
					})
					g.Line()

					g.Comment("catch parent activity context cancellation. in most cases, this should indicate a")
					g.Comment("server-sent cancellation, but there's a non-zero possibility that this cancellation")
					g.Comment("is received due to the worker stopping, prior to detecting the closing of the worker")
					g.Comment("stop channel. to give us an opportunity to detect a cancellation stemming from the")
					g.Comment("worker closing, we again check to see if the worker stop channel is closed before")
					g.Comment("propagating the cancellation")
					g.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).BlockFunc(func(g *j.Group) {
						g.Select().BlockFunc(func(g *j.Group) {
							g.Case(j.Op("<-").Qual(activityPkg, "GetWorkerStopChannel").Call(j.Id("ctx"))).BlockFunc(func(g *j.Group) {
								g.ReturnFunc(func(g *j.Group) {
									if hasUpdateOutput {
										g.Nil()
									}
									g.Id(xnsOptions).Dot("convertError").Call(
										j.Qual(temporalPkg, "NewApplicationError").Call(
											j.Lit("worker is stopping"),
											j.Lit("WorkerStopping"),
										),
									)
								})
							})

							g.Default().BlockFunc(func(g *j.Group) {
								g.Id("parentClosePolicy").Op(":=").Id("input").Dot("GetParentClosePolicy").Call()
								g.IfFunc(func(g *j.Group) {
									g.Id("parentClosePolicy").Op("==").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL").Op("||").
										Id("parentClosePolicy").Op("==").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE")
								}).BlockFunc(func(g *j.Group) {
									g.List(j.Id("disconnectedCtx"), j.Id("cancel")).Op(":=").Qual("context", "WithTimeout").CallFunc(func(g *j.Group) {
										g.Id("ctx")
										g.Qual("time", "Minute")
									})
									g.Defer().Id("cancel").Call()

									g.IfFunc(func(g *j.Group) {
										g.Id("parentClosePolicy").Op("==").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL")
									}).BlockFunc(func(g *j.Group) {
										g.Err().Op("=").Id("run").Dot("Cancel").Call(j.Id("disconnectedCtx"))
									}).Else().BlockFunc(func(g *j.Group) {
										g.Err().Op("=").Id("run").Dot("Terminate").Call(j.Id("disconnectedCtx"), j.Lit("xns activity cancellation received"), j.Lit("error"), j.Id("ctx").Dot("Err").Call())
									})
									g.IfFunc(func(g *j.Group) {
										g.Err().Op("!=").Nil()
									}).BlockFunc(func(g *j.Group) {
										g.ReturnFunc(func(g *j.Group) {
											if hasUpdateOutput {
												g.Nil()
											}
											g.Id(xnsOptions).Dot("convertError").Call(j.Err())
										})
									})
								})
								g.ReturnFunc(func(g *j.Group) {
									if hasUpdateOutput {
										g.Nil()
									}
									g.Id(xnsOptions).Dot("convertError").CallFunc(func(g *j.Group) {
										g.Qual(temporalPkg, "NewCanceledError").CallFunc(func(g *j.Group) {
											g.Id("ctx").Dot("Err").Call().Dot("Error").Call()
										})
									})
								})
							})
						})
					})
					g.Line()

					// handle doneCh
					g.Case(j.Op("<-").Id("doneCh")).BlockFunc(func(g *j.Group) {
						g.ReturnFunc(func(g *j.Group) {
							if hasUpdateOutput {
								g.Id("out")
							}
							g.Id(xnsOptions).Dot("convertError").Call(j.Err())
						})
					})
				})
			})
		})
}

func (m *Manifest) genXNSActivitiesWorkflowGetMethod(f *j.File, workflow protoreflect.FullName) {
	activities := m.Names().xnsActivities()
	method := m.methods[workflow]
	hasWorkflowOutput := !isEmpty(method.Output)
	get := m.Names().xnsWorkflowGetFunction(workflow)
	clientGet := m.Names().clientWorkflowGet(workflow)
	options := m.Names().xnsOptionsVar()

	commentf(f, methodSet(method), "%s retrieves a(n) %s workflow via an activity", get, m.fqnForWorkflow(workflow))
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("a").Op("*").Id(activities)
		}).
		Id(get).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("input").Op("*").Qual(xnsv1Pkg, "GetWorkflowRequest")
		}).
		ParamsFunc(func(g *j.Group) {
			if hasWorkflowOutput {
				g.Id("out").Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Err().Error()
		}).
		BlockFunc(func(g *j.Group) {
			if isDeprecated(method) {
				g.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated workflow detected"), j.Lit("workflow"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sWorkflowName", workflow))).Line()
			}

			// initialize heartbeat interval duration
			g.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			g.If(j.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				j.Id("heartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30),
			)

			// call client's GetWorkflow method in a goroutine
			g.Id("done").Op(":=").Make(j.Chan().Struct())
			g.Go().Func().Params().BlockFunc(func(g *j.Group) {
				g.Defer().Close(j.Id("done"))
				g.ListFunc(func(g *j.Group) {
					if hasWorkflowOutput {
						g.Id("out")
					}
					g.Err()
				}).Op("=").Id("a").Dot("client").Dot(clientGet).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("input").Dot("GetWorkflowId").Call()
					g.Id("input").Dot("GetRunId").Call()
				}).Dot("Get").CallFunc(func(g *j.Group) {
					g.Id("ctx")
				})
			}).Call()

			// wait for done
			g.For().BlockFunc(func(g *j.Group) {
				g.Select().BlockFunc(func(g *j.Group) {
					// record heartbeat every heartbeatInterval
					g.Case(j.Op("<-").Qual("time", "After").Call(j.Id("heartbeatInterval"))).BlockFunc(func(g *j.Group) {
						g.Qual(activityPkg, "RecordHeartbeat").Call(j.Id("ctx"))
					})

					// return retryable error if the worker is stopping
					g.Case(j.Op("<-").Qual(activityPkg, "GetWorkerStopChannel").Call(j.Id("ctx"))).BlockFunc(func(g *j.Group) {
						g.ReturnFunc(func(g *j.Group) {
							if hasWorkflowOutput {
								g.Nil()
							}
							g.Id(options).Dot("convertError").Call(j.Qual(temporalPkg, "NewApplicationError").Call(
								j.Lit("worker is stopping"),
								j.Lit("WorkerStopping"),
							))
						})
					})

					// handle done
					g.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).BlockFunc(func(g *j.Group) {
						g.ReturnFunc(func(g *j.Group) {
							if hasWorkflowOutput {
								g.Nil()
							}
							g.Id(options).Dot("convertError").Call(j.Err())
						})
					})
				})
			})
		})
}

func (m *Manifest) genXNSActivitiesWorkflowMethod(f *j.File, workflow, signal protoreflect.FullName) {
	methodName := m.methods[workflow].GoName
	clientMethodName := m.toCamel("%sAsync", methodName)
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	var handler *protogen.Method
	var handlerInput bool
	if signal.IsValid() {
		methodName = m.toCamel("%sWith%s", workflow, signal)
		clientMethodName = m.toCamel("%sWith%sAsync", workflow, signal)
		handler = m.methods[signal]
		handlerInput = !isEmpty(handler.Input)
	}

	if signal.IsValid() {
		commentf(f, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow via an activity", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))

	} else {
		commentf(f, methodSet(method), "%s executes a(n) %s workflow via an activity", methodName, m.fqnForWorkflow(workflow))
	}
	f.Func().
		Params(
			j.Id("a").Op("*").Id(m.toLowerCamel("%sActivities", m.GoName)),
		).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("input").Op("*").Qual(xnsv1Pkg, "WorkflowRequest"),
		).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Id("resp").Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Err().Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) || (signal.IsValid() && isDeprecated(handler)) {
				if isDeprecated(method) {
					fn.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated workflow detected"), j.Lit("workflow"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sWorkflowName", workflow)))
				}
				if signal.IsValid() && isDeprecated(handler) {
					fn.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated signal detected"), j.Lit("signal"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sSignalName", signal)))
				}
				fn.Line()
			}

			if hasInput {
				fn.Comment("unmarshal workflow request")
				fn.Var().Id("req").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				fn.If(j.Err().Op(":=").Id("input").Dot("Request").Dot("UnmarshalTo").Call(j.Op("&").Id("req")), j.Err().Op("!=").Nil()).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
							j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
								multiLineArgs,
								j.Qual("fmt", "Sprintf").Call(
									j.Lit(fmt.Sprintf("error unmarshalling workflow request of type %%s as %s.%s", string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))),
									j.Id("input").Dot("Request").Dot("GetTypeUrl").Call(),
								),
								j.Lit("InvalidArgument"),
								j.Err(),
							),
						)
					}),
				).Line()
			}

			if handlerInput {
				fn.Comment("unmarshal signal request")
				fn.Var().Id("signal").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				fn.If(j.Err().Op(":=").Id("input").Dot("Signal").Dot("UnmarshalTo").Call(j.Op("&").Id("signal")), j.Err().Op("!=").Nil()).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
							j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
								multiLineArgs,
								j.Qual("fmt", "Sprintf").Call(
									j.Lit(fmt.Sprintf("error unmarshalling signal request of type %%s as %s.%s", string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))),
									j.Id("input").Dot("Signal").Dot("GetTypeUrl").Call(),
								),
								j.Lit("InvalidArgument"),
								j.Err(),
							),
						)
					}),
				).Line()
			}

			fn.Comment("initialize workflow execution")
			fn.Var().Id("run").Qual(string(m.File.GoImportPath), m.toCamel("%sRun", workflow))
			fn.List(j.Id("run"), j.Err()).Op("=").Id("a").Dot("client").Dot(clientMethodName).CallFunc(func(args *j.Group) {
				args.Id("ctx")
				if hasInput {
					args.Op("&").Id("req")
				}
				if handlerInput {
					args.Op("&").Id("signal")
				}
				args.Qual(string(m.File.GoImportPath), m.toCamel("New%sOptions", workflow)).
					Call().
					Dot("WithStartWorkflowOptions").
					Custom(multiLineArgs, j.Qual(xnsPkg, "UnmarshalStartWorkflowOptions").Call(j.Id("input").Dot("GetStartWorkflowOptions").Call()))
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
				}),
			).Line()

			fn.Comment("exit early if detached enabled")
			fn.If(j.Id("input").Dot("GetDetached").Call()).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Nil()
				}),
			).Line()

			fn.Comment("otherwise, wait for execution to complete in child goroutine")
			fn.Id("doneCh").Op(":=").Make(j.Chan().Struct())
			fn.Go().Func().Params().Block(
				j.ListFunc(func(ls *j.Group) {
					if hasOutput {
						ls.Id("resp")
					}
					ls.Err()
				}).Op("=").Id("run").Dot("Get").Call(j.Id("ctx")),
				j.Close(j.Id("doneCh")),
			).Call().Line()

			fn.Id("heartbeatInterval").Op(":=").Id("input").Dot("GetHeartbeatInterval").Call().Dot("AsDuration").Call()
			fn.If(j.Id("heartbeatInterval").Op("==").Lit(0)).Block(
				j.Id("heartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30),
			).Line()

			fn.Comment("heartbeat activity while waiting for workflow execution to complete")
			fn.For().Block(
				j.Select().Block(
					j.Comment("send heartbeats periodically"),
					j.Case(j.Op("<-").Qual("time", "After").Call(j.Id("heartbeatInterval"))).Block(
						j.Qual(activityPkg, "RecordHeartbeat").Call(j.Id("ctx"), j.Id("run").Dot("ID").Call()),
					).Line(),

					j.Comment("return retryable error on worker close"),
					j.Case(j.Op("<-").Qual(activityPkg, "GetWorkerStopChannel").Call(j.Id("ctx"))).Block(
						j.ReturnFunc(func(returnVals *j.Group) {
							if hasOutput {
								returnVals.Nil()
							}
							returnVals.Qual(temporalPkg, "NewApplicationError").Call(j.Lit("worker is stopping"), j.Lit("WorkerStopped"))
						}),
					).Line(),

					j.Comment("catch parent activity context cancellation. in most cases, this should indicate a"),
					j.Comment("server-sent cancellation, but there's a non-zero possibility that this cancellation"),
					j.Comment("is received due to the worker stopping, prior to detecting the closing of the worker"),
					j.Comment("stop channel. to give us an opportunity to detect a cancellation stemming from the"),
					j.Comment("worker closing, we again check to see if the worker stop channel is closed before"),
					j.Comment("propagating the cancellation"),
					j.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).Block(
						j.Select().Block(
							j.Case(j.Op("<-").Qual(activityPkg, "GetWorkerStopChannel").Call(j.Id("ctx"))).Block(
								j.ReturnFunc(func(returnVals *j.Group) {
									if hasOutput {
										returnVals.Nil()
									}
									returnVals.Qual(temporalPkg, "NewApplicationError").Call(j.Lit("worker is stopping"), j.Lit("WorkerStopped"))
								}),
							),
							j.Default().BlockFunc(func(b *j.Group) {
								b.Id("parentClosePolicy").Op(":=").Id("input").Dot("GetParentClosePolicy").Call()
								b.If(
									j.Id("parentClosePolicy").Op("==").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL").Op("||").
										Id("parentClosePolicy").Op("==").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE"),
								).BlockFunc(func(b *j.Group) {
									// initialize cancellation context
									b.List(j.Id("disconnectedCtx"), j.Id("cancel")).Op(":=").Qual("context", "WithTimeout").Call(j.Qual("context", "Background").Call(), j.Qual("time", "Minute"))
									b.Defer().Id("cancel").Call()

									// cancel or terminate workflow depending on desired parent close policy
									b.If(j.Id("parentClosePolicy").Op("==").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL")).Block(
										j.Err().Op("=").Id("run").Dot("Cancel").Call(j.Id("disconnectedCtx")),
									).Else().Block(
										j.Err().Op("=").Id("run").Dot("Terminate").Call(j.Id("disconnectedCtx"), j.Lit("xns activity cancellation received"), j.Lit("error"), j.Id("ctx").Dot("Err").Call()),
									)
									b.If(j.Err().Op("!=").Nil()).Block(
										j.ReturnFunc(func(returnVals *j.Group) {
											if hasOutput {
												returnVals.Nil()
											}
											returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
										}),
									)
								})
								b.ReturnFunc(func(returnVals *j.Group) {
									if hasOutput {
										returnVals.Nil()
									}
									returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(
										j.Qual(temporalPkg, "NewCanceledError").Call(j.Id("ctx").Dot("Err").Call().Dot("Error").Call()),
									)
								})
							}),
						),
					).Line(),

					j.Comment("handle workflow completion"),
					j.Case(j.Op("<-").Id("doneCh")).Block(
						j.ReturnFunc(func(returnVals *j.Group) {
							if hasOutput {
								returnVals.Id("resp")
							}
							returnVals.Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("convertError").Call(j.Err())
						}),
					),
				),
			)
		})
}

func (m *Manifest) genXNSCancelWorkflowFunction(f *j.File) {
	funcName := m.toCamel("Cancel%sWorkflow", m.GoName)
	f.Commentf("%s cancels an existing workflow", funcName)
	f.Func().
		Id(funcName).
		Params(
			j.Id("ctx").Qual(workflowPkg, "Context"),
			j.Id("workflowID").String(),
			j.Id("runID").String(),
		).
		Error().
		Block(
			j.Return(
				j.Id(m.toCamel("%sAsync", funcName)).
					Call(
						j.Id("ctx"),
						j.Id("workflowID"),
						j.Id("runID"),
					).
					Dot("Get").
					Call(
						j.Id("ctx"),
						j.Nil(),
					),
			),
		)

	funcName = m.toCamel("%sAsync", funcName)
	f.Commentf("%s cancels an existing workflow", funcName)
	f.Func().
		Id(funcName).
		Params(
			j.Id("ctx").Qual(workflowPkg, "Context"),
			j.Id("workflowID").String(),
			j.Id("runID").String(),
		).
		Qual(workflowPkg, "Future").
		Block(
			j.Id("activityName").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(
				j.Lit(fmt.Sprintf("%s.CancelWorkflow", m.Service.Desc.FullName())),
			),
			j.If(j.Id("activityName").Op("==").Lit("")).Block(
				j.List(j.Id("f"), j.Id("s")).Op(":=").Qual(workflowPkg, "NewFuture").Call(j.Id("ctx")),
				j.Id("s").Dot("SetError").Call(j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
					multiLineArgs,
					j.Lit(fmt.Sprintf("no activity registered for %s.CancelWorkflow", m.Service.Desc.FullName())),
					j.Lit("Unimplemented"),
					j.Nil(),
				)),
				j.Return(j.Id("f")),
			),
			j.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(j.Id("ctx")),
			j.If(j.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0).Op("&&").Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
				j.Id("ao").Dot("StartToCloseTimeout").Op("=").Qual("time", "Minute"),
			),
			j.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(j.Id("ctx"), j.Id("ao")),
			j.Return(
				j.Qual(workflowPkg, "ExecuteActivity").Call(
					j.Id("ctx"),
					j.Id("activityName"),
					j.Id("workflowID"),
					j.Id("runID"),
				),
			),
		)
}

func (m *Manifest) genXNSRegisterActivities(f *j.File) {
	optionsTypeName := m.toCamel("%sOptions", m.GoName)
	optionsName := m.toLowerCamel("%sOptions", m.GoName)
	f.Commentf("%s is used to configure %s xns activity registration", optionsTypeName, string(m.Service.Desc.FullName()))
	f.Type().Id(optionsTypeName).Struct(
		j.Comment("errorConverter is used to customize error"),
		j.Id("errorConverter").Func().Params(j.Error()).Error(),
		j.Comment("filter is used to filter xns activity registrations. It receives as"),
		j.Comment("input the original activity name, and should return one of the following:"),
		j.Comment("1. the original activity name, for no changes"),
		j.Comment("2. a modified activity name, to override the original activity name"),
		j.Comment("3. an empty string, to skip registration"),
		j.Id("filter").Func().Params(j.String()).String(),
	)

	optionsConstructor := m.toCamel("New%sOptions", m.GoName)
	f.Commentf("%s initializes a new %s value", optionsConstructor, optionsTypeName)
	f.Func().
		Id(optionsConstructor).
		Params().
		Op("*").Id(optionsTypeName).
		Block(
			j.Return(j.Op("&").Id(optionsTypeName).Values()),
		)

	f.Comment("WithErrorConverter overrides the default error converter applied to xns activity errors")
	f.Func().
		Params(j.Id("opts").Op("*").Id(optionsTypeName)).
		Id("WithErrorConverter").
		Params(
			j.Id("errorConverter").Func().Params(j.Error()).Error(),
		).
		Op("*").Id(optionsTypeName).
		Block(
			j.Id("opts").Dot("errorConverter").Op("=").Id("errorConverter"),
			j.Return(j.Id("opts")),
		)

	f.Comment("Filter is used to filter registered xns activities or customize their name")
	f.Func().
		Params(j.Id("opts").Op("*").Id(optionsTypeName)).
		Id("WithFilter").
		Params(
			j.Id("filter").Func().Params(j.String()).String(),
		).
		Op("*").Id(optionsTypeName).
		Block(
			j.Id("opts").Dot("filter").Op("=").Id("filter"),
			j.Return(j.Id("opts")),
		)

	f.Comment("convertError is applied to all xns activity errors")
	f.Func().
		Params(j.Id("opts").Op("*").Id(optionsTypeName)).
		Id("convertError").
		Params(j.Err().Error()).
		Error().
		BlockFunc(func(fn *j.Group) {
			fn.If(j.Err().Op("==").Nil()).Block(
				j.Return(j.Nil()),
			)
			fn.If(j.Id("opts").Op("!=").Nil().Op("&&").Id("opts").Dot("errorConverter").Op("!=").Nil()).Block(
				j.Return(j.Id("opts").Dot("errorConverter").Call(j.Err())),
			)
			fn.Return(j.Qual(xnsPkg, "ErrorToApplicationError").Call(j.Err()))
		})

	f.Comment("filterActivity is used to filter xns activity registrations")
	f.Func().
		Params(j.Id("opts").Op("*").Id(optionsTypeName)).
		Id("filterActivity").
		Params(j.Id("name").String()).
		String().
		Block(
			j.If(j.Id("opts").Op("==").Nil().Op("||").Id("opts").Dot("filter").Op("==").Nil()).Block(
				j.Return(j.Id("name")),
			),
			j.Return(j.Id("opts").Dot("filter").Call(j.Id("name"))),
		)

	f.Commentf("%s is a reference to the %s initialized at registration", optionsName, optionsTypeName)
	f.Var().Id(optionsName).Op("*").Id(optionsTypeName)

	funcName := m.toCamel("Register%sActivities", m.GoName)
	f.Commentf("%s registers %s cross-namespace activities", funcName, string(m.Service.Desc.FullName()))
	f.Func().
		Id(funcName).
		Params(
			j.Id("r").Qual(workerPkg, "ActivityRegistry"),
			j.Id("c").Qual(string(m.File.GoImportPath), m.toCamel("%sClient", m.GoName)),
			j.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", m.GoName)),
		).
		BlockFunc(func(fn *j.Group) {
			fn.If(j.Id(optionsName).Op("==").Nil().Op("&&").Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id(optionsName).Op("=").Id("options").Index(j.Lit(0)),
			)

			fn.Id("a").Op(":=").Op("&").Id(m.toLowerCamel("%sActivities", m.GoName)).Values(
				j.Id("c"),
			)

			// register CancelWorkflow
			fn.If(
				j.Id("name").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(j.Lit(fmt.Sprintf("%s.CancelWorkflow", m.Service.Desc.FullName()))),
				j.Id("name").Op("!=").Lit(""),
			).Block(
				j.Id("r").Dot("RegisterActivityWithOptions").Call(
					j.Id("a").Dot("CancelWorkflow"),
					j.Qual(activityPkg, "RegisterOptions").Values(
						j.Id("Name").Op(":").Id("name"),
					),
				),
			)

			for _, workflow := range m.workflowsOrdered {
				if m.methods[workflow].Desc.Parent() != m.Service.Desc {
					continue
				}
				fn.If(
					j.Id("name").Op(":=").Id(optionsName).Dot("filterActivity").Call(j.Qual(string(m.File.GoImportPath), m.toCamel("%sWorkflowName", workflow))),
					j.Id("name").Op("!=").Lit(""),
				).Block(
					j.Id("r").Dot("RegisterActivityWithOptions").Call(
						j.Id("a").Dot(m.methods[workflow].GoName),
						j.Qual(activityPkg, "RegisterOptions").Values(
							j.Id("Name").Op(":").Id("name"),
						),
					),
				)
				for _, signal := range m.workflows[workflow].GetSignal() {
					if !signal.GetStart() {
						continue
					}
					fn.If(
						j.Id("name").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(j.Lit(fmt.Sprintf("%s.%s", string(m.Service.Desc.FullName()), m.toCamel("%sWith%s", workflow, getFullyQualifiedRef(workflow, signal.GetRef()))))),
						j.Id("name").Op("!=").Lit(""),
					).Block(
						j.Id("r").Dot("RegisterActivityWithOptions").Call(
							j.Id("a").Dot(m.toCamel("%sWith%s", workflow, getFullyQualifiedRef(workflow, signal.GetRef()))),
							j.Qual(activityPkg, "RegisterOptions").Values(
								j.Id("Name").Op(":").Id("name"),
							),
						),
					)
				}
				for _, update := range m.workflows[workflow].GetUpdate() {
					if !update.GetStart() {
						continue
					}
					serviceName := string(m.Service.Desc.FullName())
					methodName := m.Names().xnsUpdateWithStartFunction(workflow, getFullyQualifiedRef(workflow, update.GetRef()))
					activityName := j.Lit(fmt.Sprintf("%s.%s", serviceName, methodName))
					fn.If(
						j.Id("name").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(activityName),
						j.Id("name").Op("!=").Lit(""),
					).Block(
						j.Id("r").Dot("RegisterActivityWithOptions").Call(
							j.Id("a").Dot(methodName),
							j.Qual(activityPkg, "RegisterOptions").Values(
								j.Id("Name").Op(":").Id("name"),
							),
						),
					)
				}
			}
			for _, query := range m.queriesOrdered {
				if m.methods[query].Desc.Parent() != m.Service.Desc {
					continue
				}
				fn.If(
					j.Id("name").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(j.Qual(string(m.File.GoImportPath), m.toCamel("%sQueryName", query))),
					j.Id("name").Op("!=").Lit(""),
				).Block(
					j.Id("r").Dot("RegisterActivityWithOptions").Call(
						j.Id("a").Dot(m.methods[query].GoName),
						j.Qual(activityPkg, "RegisterOptions").Values(
							j.Id("Name").Op(":").Id("name"),
						),
					),
				)
			}
			for _, signal := range m.signalsOrdered {
				if m.methods[signal].Desc.Parent() != m.Service.Desc {
					continue
				}
				fn.If(
					j.Id("name").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(j.Qual(string(m.File.GoImportPath), m.toCamel("%sSignalName", signal))),
					j.Id("name").Op("!=").Lit(""),
				).Block(
					j.Id("r").Dot("RegisterActivityWithOptions").Call(
						j.Id("a").Dot(m.methods[signal].GoName),
						j.Qual(activityPkg, "RegisterOptions").Values(
							j.Id("Name").Op(":").Id("name"),
						),
					),
				)
			}
			for _, update := range m.updatesOrdered {
				if m.methods[update].Desc.Parent() != m.Service.Desc {
					continue
				}
				fn.If(
					j.Id("name").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(j.Qual(string(m.File.GoImportPath), m.toCamel("%sUpdateName", update))),
					j.Id("name").Op("!=").Lit(""),
				).Block(
					j.Id("r").Dot("RegisterActivityWithOptions").Call(
						j.Id("a").Dot(m.methods[update].GoName),
						j.Qual(activityPkg, "RegisterOptions").Values(
							j.Id("Name").Op(":").Id("name"),
						),
					),
				)
			}
		})
}

func (m *Manifest) genXNSQueryFunction(f *j.File, query protoreflect.FullName) {
	methodName := m.methods[query].GoName
	method := m.methods[query]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentWithDefaultf(f, methodSet(method), "%s executes a(n) %s query and blocks until error or response received", methodName, m.fqnForQuery(query))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sQueryOptions", query))
		}).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			fn.List(j.Id("handle"), j.Err()).Op(":=").Id(m.toCamel("%sAsync", query)).CallFunc(func(args *j.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Return(j.Id("handle").Dot("Get").Call(j.Id("ctx")))
		})
}

func (m *Manifest) genXNSQueryFunctionAsync(f *j.File, query protoreflect.FullName) {
	methodName := m.toCamel("%sAsync", query)
	method := m.methods[query]
	opts := m.queries[query]
	hasInput := !isEmpty(method.Input)

	commentf(f, methodSet(method), "%s executes a(n) %s query and returns a handle to the activity", methodName, m.fqnForQuery(query))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sQueryOptions", query))
		}).
		Params(
			j.Id(m.toCamel("%sQueryHandle", query)),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated query detected"), j.Lit("query"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sQueryName", query))).Line()
			}
			fn.Id("activityName").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(
				j.Qual(string(m.File.GoImportPath), m.toCamel("%sQueryName", query)),
			)
			fn.If(j.Id("activityName").Op("==").Lit("")).Block(
				j.Return(
					j.Nil(),
					j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
						multiLineArgs,
						j.Qual("fmt", "Sprintf").Call(j.Lit("no activity registered for %s"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sQueryName", query))),
						j.Lit("Unimplemented"),
						j.Nil(),
					),
				),
			)
			fn.Line()

			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(m.toCamel("%sQueryOptions", query)).Values()
			fn.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
			)
			fn.Line()

			initializeXNSOptions(fn, opts.GetXns(), time.Minute)

			if hasInput {
				fn.Comment("marshal workflow request")
				fn.List(j.Id("wreq"), j.Err()).Op(":=").Qual(anypbPkg, "New").Call(j.Id("req"))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling workflow request: %w"), j.Err())),
				)
				fn.Line()
			}

			// return run with execute activity future
			fn.List(j.Id("ctx"), j.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(j.Id("ctx"))
			fn.Return(
				j.Op("&").Id(m.toLowerCamel("%sQueryHandle", query)).Custom(multiLineValues,
					j.Id("cancel").Op(":").Id("cancel"),
					j.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						j.Id("ctx"),
						j.Id("activityName"),
						j.Op("&").Qual(xnsv1Pkg, "QueryRequest").CustomFunc(multiLineValues, func(fields *j.Group) {
							fields.Id("HeartbeatInterval").Op(":").Qual(durationpbPkg, "New").Call(j.Id("opt").Dot("HeartbeatInterval"))
							fields.Id("WorkflowId").Op(":").Id("workflowID")
							fields.Id("RunId").Op(":").Id("runID")
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
						}),
					),
				),
				j.Nil(),
			)
		})
}

func (m *Manifest) genXNSQueryHandleImpl(f *j.File, query protoreflect.FullName) {
	typeName := m.toLowerCamel("%sQueryHandle", query)
	method := m.methods[query]
	// opts := m.workflows[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a(n) %s implementation", typeName, m.toCamel("%sQueryHandle", query))
	f.Type().Id(typeName).Struct(
		j.Id("cancel").Func().Params(),
		j.Id("future").Qual(workflowPkg, "Future"),
	)

	f.Comment("Cancel the underlying query activity")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			j.Id("r").Dot("cancel").Call(),
			j.If(
				j.ListFunc(func(ls *j.Group) {
					if hasOutput {
						ls.Id("_")
					}
					ls.Err()
				}).Op(":=").Id("r").Dot("Get").Call(j.Id("ctx")),
				j.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(j.Err(), j.Qual(workflowPkg, "ErrCanceled")),
			).Block(
				j.Return(j.Err()),
			),
			j.Return(j.Nil()),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		Block(
			j.Return(j.Id("r").Dot("future")),
		)

	f.Comment("Get blocks on activity completion and returns the underlying query result")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			fn.If(
				j.Err().Op(":=").Id("r").Dot("future").Dot("Get").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					if hasOutput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				j.Err().Op("!=").Nil(),
			).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.ReturnFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("&").Id("resp")
				}
				returnVals.Nil()
			})
		})
}

func (m *Manifest) genXNSQueryHandleInterface(f *j.File, query protoreflect.FullName) {
	typeName := m.toCamel("%sQueryHandle", query)
	method := m.methods[query]
	// opts := m.workflows[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a handle for a %s query activity", typeName, query)
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		methods.Comment("Cancel cancels the workflow")
		methods.Id("Cancel").
			Params(j.Qual(workflowPkg, "Context")).
			Error().
			Line()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future").Line()

		methods.Comment("Get returns the inner workflow.Future")
		methods.Id("Get").
			Params(j.Qual(workflowPkg, "Context")).
			ParamsFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				returnVals.Error()
			}).
			Line()
	})
}

func (m *Manifest) genXNSQueryOptions(f *j.File, query protoreflect.FullName) {
	typeName := m.toCamel("%sQueryOptions", query)

	f.Commentf("%s are used to configure a(n) %s query execution", typeName, m.fqnForQuery(query))
	f.Type().Id(typeName).Struct(
		j.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		j.Id("HeartbeatInterval").Qual("time", "Duration"),
	)

	f.Commentf("New%s initializes a new %s value", typeName, typeName)
	f.Func().
		Id(m.toCamel("New%s", typeName)).
		Params().
		Op("*").Id(typeName).
		Block(
			j.Return(
				j.Op("&").Id(typeName).Values(),
			),
		)

	f.Comment("WithActivityOptions can be used to customize the activity options")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithActivityOptions").
		Params(
			j.Id("ao").Qual(workflowPkg, "ActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("ActivityOptions").Op("=").Op("&").Id("ao"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithHeartbeatInterval").
		Params(
			j.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("HeartbeatInterval").Op("=").Id("d"),
			j.Return(j.Id("opts")),
		)
}

func (m *Manifest) genXNSSignalFunction(f *j.File, signal protoreflect.FullName) {
	methodName := m.methods[signal].GoName
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s executes a(n) %s signal", methodName, m.fqnForSignal(signal))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sSignalOptions", signal))
		}).
		Error().
		BlockFunc(func(fn *j.Group) {
			fn.List(j.Id("handle"), j.Err()).Op(":=").Id(m.toCamel("%sAsync", signal)).CallFunc(func(args *j.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Err()),
			)
			fn.Return(j.Id("handle").Dot("Get").Call(j.Id("ctx")))
		})
}

func (m *Manifest) genXNSSignalFunctionAsync(f *j.File, signal protoreflect.FullName) {
	methodName := m.toCamel("%sAsync", signal)
	method := m.methods[signal]
	opts := m.signals[signal]
	hasInput := !isEmpty(method.Input)

	commentf(f, methodSet(method), "%s executes a(n) %s signal", methodName, m.fqnForSignal(signal))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sSignalOptions", signal))
		}).
		Params(
			j.Id(m.toCamel("%sSignalHandle", signal)),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated signal detected"), j.Lit("signal"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sSignalName", signal))).Line()
			}
			fn.Id("activityName").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(
				j.Qual(string(m.File.GoImportPath), m.toCamel("%sSignalName", signal)),
			)
			fn.If(j.Id("activityName").Op("==").Lit("")).Block(
				j.Return(
					j.Nil(),
					j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
						multiLineArgs,
						j.Qual("fmt", "Sprintf").Call(j.Lit("no activity registered for %s"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sSignalName", signal))),
						j.Lit("Unimplemented"),
						j.Nil(),
					),
				),
			)
			fn.Line()

			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(m.toCamel("%sSignalOptions", signal)).Values()
			fn.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
			)
			fn.Line()

			initializeXNSOptions(fn, opts.GetXns(), time.Minute)

			if hasInput {
				fn.Comment("marshal workflow request")
				fn.List(j.Id("wreq"), j.Err()).Op(":=").Qual(anypbPkg, "New").Call(j.Id("req"))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling workflow request: %w"), j.Err())),
				)
				fn.Line()
			}

			// return run with execute activity future
			fn.List(j.Id("ctx"), j.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(j.Id("ctx"))
			fn.Return(
				j.Op("&").Id(m.toLowerCamel("%sSignalHandle", signal)).Custom(multiLineValues,
					j.Id("cancel").Op(":").Id("cancel"),
					j.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						j.Id("ctx"),
						j.Id("activityName"),
						j.Op("&").Qual(xnsv1Pkg, "SignalRequest").CustomFunc(multiLineValues, func(fields *j.Group) {
							fields.Id("HeartbeatInterval").Op(":").Qual(durationpbPkg, "New").Call(j.Id("opt").Dot("HeartbeatInterval"))
							fields.Id("WorkflowId").Op(":").Id("workflowID")
							fields.Id("RunId").Op(":").Id("runID")
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
						}),
					),
				),
				j.Nil(),
			)
		})
}

func (m *Manifest) genXNSSignalHandleImpl(f *j.File, signal protoreflect.FullName) {
	typeName := m.toLowerCamel("%sSignalHandle", signal)
	f.Commentf("%s provides a(n) %s implementation", typeName, m.toCamel("%sQueryHandle", signal))
	f.Type().Id(typeName).Struct(
		j.Id("cancel").Func().Params(),
		j.Id("future").Qual(workflowPkg, "Future"),
	)

	f.Comment("Cancel the underlying signal activity")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			j.Id("r").Dot("cancel").Call(),
			j.If(
				j.Err().Op(":=").Id("r").Dot("Get").Call(j.Id("ctx")),
				j.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(j.Err(), j.Qual(workflowPkg, "ErrCanceled")),
			).Block(
				j.Return(j.Err()),
			),
			j.Return(j.Nil()),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		Block(
			j.Return(j.Id("r").Dot("future")),
		)

	f.Comment("Get blocks on activity completion")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			j.Return(j.Id("r").Dot("future").Dot("Get").Call(j.Id("ctx"), j.Nil())),
		)
}

func (m *Manifest) genXNSSignalHandleInterface(f *j.File, signal protoreflect.FullName) {
	typeName := m.toCamel("%sSignalHandle", signal)

	f.Commentf("%s provides a handle for a %s signal activity", typeName, signal)
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		methods.Comment("Cancel cancels the workflow")
		methods.Id("Cancel").
			Params(j.Qual(workflowPkg, "Context")).
			Error()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future")

		methods.Comment("Get returns the inner workflow.Future")
		methods.Id("Get").
			Params(j.Qual(workflowPkg, "Context")).
			Error()
	})
}

func (m *Manifest) genXNSSignalOptions(f *j.File, signal protoreflect.FullName) {
	typeName := m.toCamel("%sSignalOptions", signal)

	f.Commentf("%s are used to configure a(n) %s signal execution", typeName, m.fqnForSignal(signal))
	f.Type().Id(typeName).Struct(
		j.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		j.Id("HeartbeatInterval").Qual("time", "Duration"),
	)

	f.Commentf("New%s initializes a new %s value", typeName, typeName)
	f.Func().
		Id(m.toCamel("New%s", typeName)).
		Params().
		Op("*").Id(typeName).
		Block(
			j.Return(
				j.Op("&").Id(typeName).Values(),
			),
		)

	f.Comment("WithActivityOptions can be used to customize the activity options")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithActivityOptions").
		Params(
			j.Id("ao").Qual(workflowPkg, "ActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("ActivityOptions").Op("=").Op("&").Id("ao"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithHeartbeatInterval").
		Params(
			j.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("HeartbeatInterval").Op("=").Id("d"),
			j.Return(j.Id("opts")),
		)
}

func (m *Manifest) genXNSUpdateFunction(f *j.File, update protoreflect.FullName) {
	methodName := m.methods[update].GoName
	method := m.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentWithDefaultf(f, methodSet(method), "%s executes a(n) %s update and blocks until error or response received", methodName, m.fqnForUpdate(update))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sUpdateOptions", update))
		}).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			fn.List(j.Id("run"), j.Err()).Op(":=").Id(m.toCamel("%sAsync", update)).CallFunc(func(args *j.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Return(j.Id("run").Dot("Get").Call(j.Id("ctx")))
		})
}

func (m *Manifest) genXNSUpdateFunctionAsync(f *j.File, update protoreflect.FullName) {
	methodName := m.toCamel("%sAsync", update)
	method := m.methods[update]
	hasInput := !isEmpty(method.Input)
	opts := m.updates[update]

	commentf(f, methodSet(method), "%s executes a(n) %s update and blocks until error or response received", methodName, m.fqnForUpdate(update))
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sUpdateOptions", update))
		}).
		Params(
			j.Id(m.toCamel("%sHandle", update)),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated update detected"), j.Lit("update"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sUpdateName", update))).Line()
			}
			fn.Id("activityName").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(
				j.Qual(string(m.File.GoImportPath), m.toCamel("%sUpdateName", update)),
			)
			fn.If(j.Id("activityName").Op("==").Lit("")).Block(
				j.Return(
					j.Nil(),
					j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
						multiLineArgs,
						j.Qual("fmt", "Sprintf").Call(j.Lit("no activity registered for %s"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sUpdateName", update))),
						j.Lit("Unimplemented"),
						j.Nil(),
					),
				),
			)
			fn.Line()

			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(m.toCamel("%sUpdateOptions", update)).Values()
			fn.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
			)
			fn.Line()

			initializeXNSOptions(fn, opts.GetXns(), time.Minute*5)

			// build update options
			fn.Id("uo").Op(":=").Qual(clientPkg, "UpdateWorkflowOptions").Values()
			fn.If(j.Id("opt").Dot("UpdateWorkflowOptions").Op("!=").Nil()).Block(
				j.Id("uo").Op("=").Op("*").Id("opt").Dot("UpdateWorkflowOptions"),
			)
			fn.Id("uo").Dot("WorkflowID").Op("=").Id("workflowID")
			fn.Id("uo").Dot("RunID").Op("=").Id("runID")

			// set update id if unset and  id field and/or prefix defined
			if idExpr := opts.GetId(); idExpr != "" {
				fn.If(j.Id("uo").Dot("UpdateID").Op("==").Lit("")).Block(
					j.If(
						j.Err().Op(":=").Qual(workflowPkg, "SideEffect").Call(j.Id("ctx"), j.Func().Params(j.Id("ctx").Qual(workflowPkg, "Context")).Any().Block(
							j.List(j.Id("id"), j.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *j.Group) {
								args.Qual(string(m.File.GoImportPath), m.toCamel("%sIDExpression", update))
								if hasInput {
									args.Id("req").Dot("ProtoReflect").Call()
								} else {
									args.Nil()
								}
							}),
							j.If(j.Err().Op("!=").Nil()).Block(
								j.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Error").Call(
									j.Lit(fmt.Sprintf("error evaluating id expression for %q update", update)),
									j.Lit("error"),
									j.Err(),
								),
								j.Return(j.Nil()),
							),
							j.Return(j.Id("id")),
						)).Dot("Get").Call(j.Op("&").Id("uo").Dot("UpdateID")),
						j.Err().Op("!=").Nil(),
					).Block(
						j.Return(j.Nil(), j.Err()),
					),
				)
			}
			fn.If(j.Id("uo").Dot("UpdateID").Op("==").Lit("")).Block(
				j.If(
					j.Err().Op(":=").Qual(workflowPkg, "SideEffect").Call(j.Id("ctx"), j.Func().Params(j.Id("ctx").Qual(workflowPkg, "Context")).Any().Block(
						j.List(j.Id("id"), j.Err()).Op(":=").Qual(uuidPkg, "NewRandom").Call(),
						j.If(j.Err().Op("!=").Nil()).Block(
							j.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Error").Call(
								j.Lit("error generating update id"),
								j.Lit("error"),
								j.Err(),
							),
							j.Return(j.Nil()),
						),
						j.Return(j.Id("id")),
					)).Dot("Get").Call(j.Op("&").Id("uo").Dot("UpdateID")),
					j.Err().Op("!=").Nil(),
				).Block(
					j.Return(j.Nil(), j.Err()),
				),
			)
			fn.If(j.Id("uo").Dot("UpdateID").Op("==").Lit("")).Block(
				j.Return(
					j.Nil(),
					j.Qual(temporalPkg, "NewNonRetryableApplicationError").Call(
						j.Lit("update id is required"),
						j.Lit("InvalidArgument"),
						j.Nil(),
					),
				),
			)
			fn.Line()

			// marshal update options
			fn.List(j.Id("uopb"), j.Err()).Op(":=").Qual(xnsPkg, "MarshalUpdateWorkflowOptions").Call(j.Id("uo"))
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling update workflow options: %w"), j.Err())),
			)
			fn.Line()

			// marshal workflow request
			if hasInput {
				fn.List(j.Id("wreq"), j.Err()).Op(":=").Qual(anypbPkg, "New").Call(j.Id("req"))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling update request: %w"), j.Err())),
				)
				fn.Line()
			}

			// return run with execute activity future
			fn.List(j.Id("ctx"), j.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(j.Id("ctx"))
			fn.Return(
				j.Op("&").Id(m.toLowerCamel("%sHandle", update)).Custom(multiLineValues,
					j.Id("cancel").Op(":").Id("cancel"),
					j.Id("id").Op(":").Id("uo").Dot("UpdateID"),
					j.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						j.Id("ctx"),
						j.Id("activityName"),
						j.Op("&").Qual(xnsv1Pkg, "UpdateRequest").CustomFunc(multiLineValues, func(fields *j.Group) {
							fields.Id("HeartbeatInterval").Op(":").Qual(durationpbPkg, "New").Call(j.Id("opt").Dot("HeartbeatInterval"))
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
							fields.Id("UpdateWorkflowOptions").Op(":").Id("uopb")
						}),
					),
				),
				j.Nil(),
			)
		})
}

func (m *Manifest) genXNSUpdateOptions(f *j.File, update protoreflect.FullName) {
	typeName := m.toCamel("%sUpdateOptions", update)

	f.Commentf("%s are used to configure a(n) %s update execution", typeName, m.fqnForUpdate(update))
	f.Type().Id(typeName).Struct(
		j.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		j.Id("HeartbeatInterval").Qual("time", "Duration"),
		j.Id("UpdateWorkflowOptions").Op("*").Qual(clientPkg, "UpdateWorkflowOptions"),
	)

	f.Commentf("New%s initializes a new %s value", typeName, typeName)
	f.Func().
		Id(m.toCamel("New%s", typeName)).
		Params().
		Op("*").Id(typeName).
		Block(
			j.Return(
				j.Op("&").Id(typeName).Values(),
			),
		)

	f.Comment("WithActivityOptions can be used to customize the activity options")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithActivityOptions").
		Params(
			j.Id("ao").Qual(workflowPkg, "ActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("ActivityOptions").Op("=").Op("&").Id("ao"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithHeartbeatInterval").
		Params(
			j.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("HeartbeatInterval").Op("=").Id("d"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithUpdateWorkflowOptions can be used to customize the update workflow options")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithUpdateWorkflowOptions").
		Params(
			j.Id("uwo").Qual(clientPkg, "UpdateWorkflowOptions"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("UpdateWorkflowOptions").Op("=").Op("&").Id("uwo"),
			j.Return(j.Id("opts")),
		)
}

func (m *Manifest) genXNSUpdateHandleImpl(f *j.File, update protoreflect.FullName) {
	typeName := m.toLowerCamel("%sHandle", update)
	method := m.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a(n) %s implementation", typeName, m.toCamel("%sHandle", update))
	f.Type().Id(typeName).Struct(
		j.Id("cancel").Func().Params(),
		j.Id("future").Qual(workflowPkg, "Future"),
		j.Id("id").String(),
	)

	f.Comment("Cancel the underlying workflow update")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			j.Id("r").Dot("cancel").Call(),
			j.If(
				j.ListFunc(func(ls *j.Group) {
					if hasOutput {
						ls.Id("_")
					}
					ls.Err()
				}).Op(":=").Id("r").Dot("Get").Call(j.Id("ctx")),
				j.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(j.Err(), j.Qual(workflowPkg, "ErrCanceled")),
			).Block(
				j.Return(j.Err()),
			),
			j.Return(j.Nil()),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		Block(
			j.Return(j.Id("r").Dot("future")),
		)

	f.Comment("Get blocks on activity completion and returns the underlying update result")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			fn.If(
				j.Err().Op(":=").Id("r").Dot("future").Dot("Get").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					if hasOutput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				j.Err().Op("!=").Nil(),
			).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.ReturnFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("&").Id("resp")
				}
				returnVals.Nil()
			})
		})

	f.Comment("ID returns the underlying workflow id")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("ID").
		Params().
		String().
		Block(
			j.Return(j.Id("r").Dot("id")),
		)
}

func (m *Manifest) genXNSUpdateHandleInterface(f *j.File, update protoreflect.FullName) {
	typeName := m.toCamel("%sHandle", update)
	method := m.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a handle to a %s workflow update", typeName, update)
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		methods.Comment("Cancel cancels the update activity")
		methods.Id("Cancel").
			Params(j.Qual(workflowPkg, "Context")).
			Error().
			Line()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future").Line()

		methods.Comment("Get blocks on update completion and returns the result")
		methods.Id("Get").
			Params(j.Qual(workflowPkg, "Context")).
			ParamsFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				returnVals.Error()
			}).
			Line()

		methods.Comment("ID returns the update id")
		methods.Id("ID").
			Params().
			String().
			Line()
	})
}

func (m *Manifest) genXNSUpdateWithStartFunction(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)
	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)
	hasUpdateOutput := !isEmpty(handler.Output)

	async := m.Names().clientUpdateWithStartAsync(workflow, update)
	function := m.Names().clientUpdateWithStart(workflow, update)
	options := m.Names().clientUpdateWithStartOptions(workflow, update)
	runIface := m.Names().xnsWorkflowRunIface(workflow)

	commentf(f, methodSet(method, handler), "%s executes a(n) %s update for a(n) %s workflow, starting it if necessary, and blocks until error or update is complete", function, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
	f.Func().
		Id(function).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			if hasWorkflowInput {
				g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(options)
		}).
		ParamsFunc(func(g *j.Group) {
			if hasUpdateOutput {
				g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			g.Id(runIface)
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// invoke async function
			g.List(j.Id("handle"), j.Id("run"), j.Err()).Op(":=").Id(async).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasWorkflowInput {
					g.Id("input")
				}
				if hasUpdateInput {
					g.Id("update")
				}
				g.Id("options").Op("...")
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Nil()
					}
					g.Id("run")
					g.Err()
				})
			})

			// wait for the update to complete
			g.IfFunc(func(g *j.Group) {
				g.ListFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Id("out")
					}
					g.Id("err")
				}).Op(":=").Id("handle").Dot("Get").Call(j.Id("ctx"))
				g.Err().Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Nil()
					}
					g.Id("run")
					g.Err()
				})
			}).Else().BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Id("out")
					}
					g.Id("run")
					g.Nil()
				})
			})
		})
}

func (m *Manifest) genXNSUpdateWithStartFunctionAsync(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)
	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)

	async := m.Names().clientUpdateWithStartAsync(workflow, update)
	function := m.Names().clientUpdateWithStart(workflow, update)
	options := m.Names().xnsUpdateWithStartOptions(workflow, update)
	optionsCtor := m.Names().xnsUpdateWithStartOptionsCtor(workflow, update)
	handleIface := m.Names().xnsUpdateHandleIface(update)
	handleImpl := m.Names().xnsUpdateHandleImpl(update)
	runIface := m.Names().xnsWorkflowRunIface(workflow)
	runImpl := m.Names().xnsWorkflowRunImpl(workflow)
	xnsOptions := m.Names().xnsOptionsVar()
	workflowName := j.Qual(string(m.File.GoImportPath), m.Names().workflowName(workflow))
	updateName := j.Qual(string(m.File.GoImportPath), m.Names().updateName(update))

	commentf(f, methodSet(method, handler), "%s executes a(n) %s update for a(n) %s workflow, starting it if necessary, and returns a handle to the update and workflow execution", async, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
	f.Func().
		Id(async).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			if hasWorkflowInput {
				g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(options)
		}).
		ParamsFunc(func(g *j.Group) {
			g.Id(handleIface)
			g.Id(runIface)
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			if isDeprecated(method) {
				g.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated workflow detected"), j.Lit("workflow"), workflowName)
			}
			if isDeprecated(handler) {
				g.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated update detected"), j.Lit("update"), updateName)
			}

			// lookup activity name
			activityName := j.Lit(fmt.Sprintf("%s.%s", string(m.Service.Desc.FullName()), function))
			g.Id("activityName").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(
				activityName,
			)
			g.If(j.Id("activityName").Op("==").Lit("")).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					g.Nil()
					g.Nil()
					g.Id(xnsOptions).Dot("convertError").CallFunc(func(g *j.Group) {
						g.Qual(temporalPkg, "NewNonRetryableApplicationError").CallFunc(func(g *j.Group) {
							g.Qual("fmt", "Sprintf").CallFunc(func(g *j.Group) {
								g.Lit("no activity registered for %s")
								g.Id("activityName")
							})
							g.Lit("Unimplemented")
							g.Nil()
						})
					})
				})
			})

			// initialize method options
			g.Var().Id("o").Op("*").Id(options)
			g.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id("options").Index(j.Lit(0))
			}).Else().BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id(optionsCtor).Call()
			})

			// build activity context and input using method options
			g.ListFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Id("req")
				g.Err()
			}).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasWorkflowInput {
					g.Id("input")
				}
				if hasUpdateInput {
					g.Id("update")
				}
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					g.Nil()
					g.Nil()
					g.Id(xnsOptions).Dot("convertError").Call(j.Err())
				})
			})

			// execute update with start activity and initialize return values
			g.List(j.Id("ctx"), j.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(j.Id("ctx"))
			g.Id("handle").Op(":=").Op("&").Id(handleImpl).Values(j.DictFunc(func(d j.Dict) {
				d[j.Id("cancel")] = j.Id("cancel")
				d[j.Id("id")] = j.Id("req").Dot("GetUpdateWorkflowOptions").Call().Dot("GetUpdateId").Call()
				d[j.Id("future")] = j.Qual(workflowPkg, "ExecuteActivity").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("activityName")
					g.Id("req")
				})
			}))
			g.Id("run").Op(":=").Op("&").Id(runImpl).Values(j.DictFunc(func(d j.Dict) {
				d[j.Id("ctx")] = j.Id("ctx")
				d[j.Id("id")] = j.Id("req").Dot("GetStartWorkflowOptions").Call().Dot("GetId").Call()
			}))
			g.ReturnFunc(func(g *j.Group) {
				g.Id("handle")
				g.Id("run")
				g.Nil()
			})
		})
}

func (m *Manifest) genXNSUpdateWithStartOptions(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)
	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)
	var xnsActivityOpts *temporalv1.XNSActivityOptions
	for _, u := range m.workflows[workflow].GetUpdate() {
		if getFullyQualifiedRef(m.Service.Desc.FullName(), u.GetRef()) == update {
			xnsActivityOpts = u.GetXns()
			break
		}
	}

	options := m.Names().clientUpdateWithStartOptions(workflow, update)
	optionsCtor := m.Names().xnsUpdateWithStartOptionsCtor(workflow, update)
	updateOptions := m.Names().xnsUpdateOptions(update)
	updateOptionsCtor := m.Names().xnsUpdateOptionsCtor(update)
	workflowOptions := m.Names().xnsWorkflowOptions(workflow)
	workflowOptionsCtor := m.Names().xnsWorkflowOptionsCtor(workflow)

	commentf(f, methodSet(method, handler), "%s are used to configure a(n) %s update for a(n) %s workflow", m.Names().clientUpdateWithStartOptions(workflow, update), m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
	f.Type().Id(options).StructFunc(func(g *j.Group) {
		g.Id("activityOptions").Op("*").Qual(workflowPkg, "ActivityOptions")
		g.Id("heartbeatInterval").Op("*").Qual("time", "Duration")
		g.Id("updateOptions").Op("*").Id(updateOptions)
		g.Id("workflowOptions").Op("*").Id(workflowOptions)
	})

	f.Commentf("%s initializes a new %s value", optionsCtor, options)
	f.Func().
		Id(optionsCtor).
		Params().
		Op("*").Id(options).
		BlockFunc(func(g *j.Group) {
			g.ReturnFunc(func(g *j.Group) {
				g.Op("&").Id(options).Values()
			})
		})

	f.Comment("Build builds the activity context and input for an update with start workflow activity")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("Build").
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			if hasWorkflowInput {
				g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		ParamsFunc(func(g *j.Group) {
			g.Qual(workflowPkg, "Context")
			g.Op("*").Qual(xnsv1Pkg, "UpdateWithStartRequest")
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// initialize start workflow options
			g.Id("wo").Op(":=").Id("o").Dot("workflowOptions")
			g.IfFunc(func(g *j.Group) {
				g.Id("wo").Op("==").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("wo").Op("=").Id(workflowOptionsCtor).Call()
			})

			// initialize update options
			g.Id("uo").Op(":=").Id("o").Dot("updateOptions")
			g.IfFunc(func(g *j.Group) {
				g.Id("uo").Op("==").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("uo").Op("=").Id(updateOptionsCtor).Call()
			})

			// initialize activity options
			g.Id("ao").Op(":=").Id("o").Dot("activityOptions")
			g.IfFunc(func(g *j.Group) {
				g.Id("ao").Op("==").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("ao").Op("=").Op("&").Qual(workflowPkg, "ActivityOptions").Values()
			})

			// set heartbeat interval
			g.IfFunc(func(g *j.Group) {
				g.Id("o").Dot("heartbeatInterval").Op("==").Lit(0)
			}).BlockFunc(func(g *j.Group) {
				if d := xnsActivityOpts.GetHeartbeatInterval(); d.IsValid() {

				} else {
					g.Id("o").Dot("heartbeatInterval").Op("=").Op("&").Qual("time", "Duration").Values()
				}
			})

			// set heartbeat timeout
			if d := xnsActivityOpts.GetHeartbeatTimeout(); d.IsValid() {
				g.IfFunc(func(g *j.Group) {
					g.Id("ao").Dot("HeartbeatTimeout").Op("==").Lit(0)
				}).BlockFunc(func(g *j.Group) {
					g.Id("ao").Dot("HeartbeatTimeout").Op("=").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Comment(durafmt.Parse(d.AsDuration()).String())
				})
			}

			// initialize activity input
			g.Id("req").Op("=").Op("&").Qual(xnsv1Pkg, "UpdateWithStartRequest").Values()
		})

	f.Commentf("WithActivityOptions can be used to customize the activity options")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("WithActivityOptions").
		ParamsFunc(func(g *j.Group) {
			g.Id("ao").Qual(workflowPkg, "ActivityOptions")
		}).
		ParamsFunc(func(g *j.Group) {
			g.Op("*").Id(options)
		}).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("activityOptions").Op("=").Op("&").Id("ao")
			g.Return(j.Id("o"))
		})

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("WithHeartbeatInterval").
		ParamsFunc(func(g *j.Group) {
			g.Id("d").Qual("time", "Duration")
		}).
		ParamsFunc(func(g *j.Group) {
			g.Op("*").Id(options)
		}).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("heartbeatInterval").Op("=").Op("&").Id("d")
			g.Return(j.Id("o"))
		})

	f.Comment("WithUpdateOptions can be used to customize the update options")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("WithUpdateOptions").
		ParamsFunc(func(g *j.Group) {
			g.Id("uo").Op("*").Id(updateOptions)
		}).
		ParamsFunc(func(g *j.Group) {
			g.Op("*").Id(options)
		}).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("updateOptions").Op("=").Id("uo")
			g.Return(j.Id("o"))
		})
}

func (m *Manifest) genXNSWorkflowFunction(f *j.File, workflow, signal protoreflect.FullName) {
	methodName := m.methods[workflow].GoName
	asyncMethodName := m.toCamel("%sAsync", workflow)
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	var handler *protogen.Method
	var handlerInput bool
	if signal.IsValid() {
		methodName = m.toCamel("%sWith%s", workflow, signal)
		asyncMethodName = m.toCamel("%sWith%sAsync", workflow, signal)
		handler = m.methods[signal]
		handlerInput = !isEmpty(handler.Input)
	}

	if signal.IsValid() {
		commentf(f, methodSet(method, handler), "%s sends a(n) %s signal to a %s workflow, starting it if necessary, and blocks until the workflow completes", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
	} else {
		commentWithDefaultf(f, methodSet(method), "%s executes a(n) %s workflow and blocks until error or response is received", methodName, m.fqnForWorkflow(workflow))
	}
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if signal.IsValid() && handlerInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sWorkflowOptions", workflow))
		}).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			j.List(j.Id("run"), j.Err()).Op(":=").Id(asyncMethodName).CallFunc(func(args *j.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req")
				}
				if signal.IsValid() && handlerInput {
					args.Id("signal")
				}
				args.Id("opts").Op("...")
			}),
			j.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			),
			j.Return(j.Id("run").Dot("Get").Call(j.Id("ctx"))),
		)
}

func (m *Manifest) genXNSWorkflowFunctionAsync(f *j.File, workflow, signal protoreflect.FullName) {
	methodName := m.toCamel("%sAsync", workflow)
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)
	opts := m.workflows[workflow]

	var handler *protogen.Method
	var handlerInput bool
	if signal.IsValid() {
		methodName = m.toCamel("%sWith%sAsync", workflow, signal)
		handler = m.methods[signal]
		handlerInput = !isEmpty(handler.Input)
	}

	if signal.IsValid() {
		commentf(f, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and returns a handle to the underlying activity", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
	} else {
		commentWithDefaultf(f, methodSet(method), "%s executes a(n) %s workflow and returns a handle to the underlying activity", methodName, m.fqnForWorkflow(workflow))
	}
	f.Func().
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if handlerInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sWorkflowOptions", workflow))
		}).
		Params(
			j.Id(m.toCamel("%sRun", workflow)),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			// log deprecration warnings
			if isDeprecated(method) || (signal.IsValid() && isDeprecated(handler)) {
				if isDeprecated(method) {
					fn.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated workflow detected"), j.Lit("workflow"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sWorkflowName", workflow)))
				}
				if signal.IsValid() && isDeprecated(handler) {
					fn.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Warn").Call(j.Lit("use of deprecated signal detected"), j.Lit("signal"), j.Qual(string(m.File.GoImportPath), m.toCamel("%sSignalName", signal)))
				}
				fn.Line()
			}

			// lookup xns activity name
			defaultActivityName := j.Qual(string(m.File.GoImportPath), m.toCamel("%sWorkflowName", workflow))
			if signal.IsValid() {
				defaultActivityName = j.Lit(fmt.Sprintf("%s.%s", string(m.Service.Desc.FullName()), m.toCamel("%sWith%s", workflow, signal)))
			}
			fn.Id("activityName").Op(":=").Id(m.toLowerCamel("%sOptions", m.GoName)).Dot("filterActivity").Call(
				defaultActivityName,
			)
			fn.If(j.Id("activityName").Op("==").Lit("")).Block(
				j.Return(
					j.Nil(),
					j.Qual(temporalPkg, "NewNonRetryableApplicationError").Custom(
						multiLineArgs,
						j.Qual("fmt", "Sprintf").Call(j.Lit("no activity registered for %s"), defaultActivityName),
						j.Lit("Unimplemented"),
						j.Nil(),
					),
				),
			).Line()

			// extract workflow options
			fn.Id("opt").Op(":=").Op("&").Id(m.toCamel("%sWorkflowOptions", workflow)).Values()
			fn.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
			)

			// initialize xns activity options
			xnsOpts := opts.GetXns()
			if signal.IsValid() {
				for _, s := range opts.GetSignal() {
					if sig := getFullyQualifiedRef(workflow, s.GetRef()); sig != signal {
						continue
					}
					if s.GetXns() != nil {
						xnsOpts = s.GetXns()
					}
					break
				}
			}
			initializeXNSOptions(fn, xnsOpts, opts.GetExecutionTimeout().AsDuration())

			// initialize start workflow options
			m.initializeXNSStartWorkflowOptions(fn, workflow)

			// marshal workflow input as anypb.Any
			if hasInput {
				fn.Comment("marshal workflow request protobuf message")
				fn.List(j.Id("wreq"), j.Err()).Op(":=").Qual(anypbPkg, "New").Call(j.Id("req"))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling workflow request: %w"), j.Err())),
				).Line()
			}

			// marshal signal input as anypb.Any
			if signal.IsValid() && handlerInput {
				fn.Comment("marshal signal request protobuf message")
				fn.List(j.Id("wsignal"), j.Err()).Op(":=").Qual(anypbPkg, "New").Call(j.Id("signal"))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling signal request: %w"), j.Err())),
				).Line()
			}

			// compute parent close policy from workflow default, xns default, and option override
			defaultParentClosePolicy := opts.GetParentClosePolicy()
			if opts.GetXns().GetParentClosePolicy() != temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED {
				defaultParentClosePolicy = opts.GetXns().GetParentClosePolicy()
			}
			if defaultParentClosePolicy != temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED {
				var v string
				switch defaultParentClosePolicy {
				case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON:
					v = "ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON"
				case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL:
					v = "ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL"
				case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE:
					v = "ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE"
				}
				fn.Id("parentClosePolicy").Op(":=").Qual(temporalv1Pkg, v)
			} else {
				fn.Var().Id("parentClosePolicy").Qual(temporalv1Pkg, "ParentClosePolicy")
			}
			fn.Switch(j.Id("opt").Dot("ParentClosePolicy")).Block(
				j.Case(j.Qual(enumsPkg, "PARENT_CLOSE_POLICY_ABANDON")).
					Block(
						j.Id("parentClosePolicy").Op("=").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON"),
					),
				j.Case(j.Qual(enumsPkg, "PARENT_CLOSE_POLICY_REQUEST_CANCEL")).
					Block(
						j.Id("parentClosePolicy").Op("=").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL"),
					),
				j.Case(j.Qual(enumsPkg, "PARENT_CLOSE_POLICY_TERMINATE")).
					Block(
						j.Id("parentClosePolicy").Op("=").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE"),
					),
			).Line()

			// return run with execute activity future
			fn.List(j.Id("ctx"), j.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(j.Id("ctx"))
			fn.Return(
				j.Op("&").Id(m.toLowerCamel("%sRun", workflow)).Custom(multiLineValues,
					j.Id("cancel").Op(":").Id("cancel"),
					j.Id("id").Op(":").Id("wo").Dot("ID"),
					j.Id("future").Op(":").Qual(workflowPkg, "ExecuteActivity").Call(
						j.Id("ctx"),
						j.Id("activityName"),
						j.Op("&").Qual(xnsv1Pkg, "WorkflowRequest").CustomFunc(multiLineValues, func(fields *j.Group) {
							fields.Id("Detached").Op(":").Id("opt").Dot("Detached")
							fields.Id("HeartbeatInterval").Op(":").Qual(durationpbPkg, "New").Call(j.Id("opt").Dot("HeartbeatInterval"))
							fields.Id("ParentClosePolicy").Op(":").Id("parentClosePolicy")
							if hasInput {
								fields.Id("Request").Op(":").Id("wreq")
							}
							if signal.IsValid() && handlerInput {
								fields.Id("Signal").Op(":").Id("wsignal")
							}
							fields.Id("StartWorkflowOptions").Op(":").Id("swo")
						}),
					),
				),
				j.Nil(),
			)
		})
}

func (m *Manifest) genXNSWorkflowGetFunction(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowOutput := !isEmpty(method.Output)

	function := m.Names().xnsWorkflowGetFunction(workflow)
	async := m.Names().xnsWorkflowGetFunctionAsync(workflow)

	commentf(f, methodSet(method), "%s returns a(n) %s workflow execution", function, m.fqnForWorkflow(workflow))
	f.Func().
		Id(function).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
		}).
		ParamsFunc(func(g *j.Group) {
			if hasWorkflowOutput {
				g.Id("out").Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Err().Error()
		}).
		BlockFunc(func(g *j.Group) {
			g.ListFunc(func(g *j.Group) {
				if hasWorkflowOutput {
					g.Id("out")
				}
				g.Id("err")
			}).Op("=").Id(async).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Id("workflowID")
				g.Id("runID")
			}).Dot("Get").Call(j.Id("ctx"))

			g.IfFunc(func(g *j.Group) {
				g.Err().Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasWorkflowOutput {
						g.Nil()
					}
					g.Id("err")
				})
			})
			g.ReturnFunc(func(g *j.Group) {
				if hasWorkflowOutput {
					g.Id("out")
				}
				g.Nil()
			})
		})
}

func (m *Manifest) genXNSWorkflowGetFunctionAsync(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]

	async := m.Names().xnsWorkflowGetFunctionAsync(workflow)
	getActivity := m.Names().xnsWorkflowGetFunction(workflow)
	runIface := m.Names().xnsWorkflowRunIface(workflow)
	runImpl := m.Names().xnsWorkflowRunImpl(workflow)
	options := m.Names().xnsOptionsVar()

	commentf(f, methodSet(method), "%s returns a handle to a(n) %s workflow execution", async, m.fqnForWorkflow(workflow))
	f.Func().
		Id(async).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
		}).
		ParamsFunc(func(g *j.Group) {
			g.Id(runIface)
		}).
		BlockFunc(func(g *j.Group) {
			g.Id("activityName").Op(":=").Id(options).Dot("filterActivity").Call(
				j.Lit(fmt.Sprintf("%s.%s", string(m.Service.Desc.FullName()), getActivity)),
			)
			g.If(j.Id("activityName").Op("==").Lit("")).BlockFunc(func(g *j.Group) {
				g.ListFunc(func(g *j.Group) {
					g.Id("f")
					g.Id("set")
				}).Op(":=").Qual(workflowPkg, "NewFuture").Call(j.Id("ctx"))
				g.Id("set").Dot("SetError").CallFunc(func(g *j.Group) {
					g.Qual(temporalPkg, "NewNonRetryableApplicationError").CallFunc(func(g *j.Group) {
						g.Qual("fmt", "Sprintf").CallFunc(func(g *j.Group) {
							g.Lit("no activity registered for %s")
							g.Id("activityName")
						})
						g.Lit("Unimplemented")
						g.Nil()
					})
				})
				g.ReturnFunc(func(g *j.Group) {
					g.Op("&").Id(runImpl).Values(j.DictFunc(func(d j.Dict) {
						d[j.Id("id")] = j.Id("workflowID")
						d[j.Id("future")] = j.Id("f")
					}))
				})
			})

			g.List(j.Id("ctx"), j.Id("cancel")).Op(":=").Qual(workflowPkg, "WithCancel").Call(j.Id("ctx"))
			g.ReturnFunc(func(g *j.Group) {
				g.Op("&").Id(runImpl).Values(j.DictFunc(func(d j.Dict) {
					d[j.Id("cancel")] = j.Id("cancel")
					d[j.Id("id")] = j.Id("workflowID")
					d[j.Id("future")] = j.Qual(workflowPkg, "ExecuteActivity").Call(
						j.Id("ctx"),
						j.Id("activityName"),
						j.Op("&").Qual(xnsv1Pkg, "GetWorkflowRequest").Values(
							j.Dict{
								j.Id("WorkflowId"): j.Id("workflowID"),
								j.Id("RunId"):      j.Id("runID"),
								j.Id("HeartbeatInterval"): j.Qual(durationpbPkg, "New").CallFunc(func(g *j.Group) {
									g.Qual("time", "Second").Op("*").Lit(30)
								}),
							},
						),
					)
				}))
			})
		})
}

func (m *Manifest) genXNSWorkflowGetOptions(f *j.File, workflow protoreflect.FullName) {
	options := m.Names().xnsWorkflowGetOptions(workflow)
	optionsCtor := m.Names().xnsWorkflowGetOptionsCtor(workflow)

	commentf(f, methodSet(m.methods[workflow]), "%s are used to configure a(n) %s workflow execution getter activity", options, m.fqnForWorkflow(workflow))
	f.Type().Id(options).StructFunc(func(g *j.Group) {
		g.Id("activityOptions").Op("*").Qual(workflowPkg, "ActivityOptions")
		g.Id("heartbeatInterval").Qual("time", "Duration")
		g.Id("parentClosePolicy").Qual(enumsPkg, "ParentClosePolicy")
	})

	commentf(f, methodSet(m.methods[workflow]), "%s initializes a new %s value", optionsCtor, options)
	f.Func().
		Id(m.Names().xnsWorkflowGetOptionsCtor(workflow)).
		Params().
		Op("*").Id(options).
		BlockFunc(func(g *j.Group) {
			g.ReturnFunc(func(g *j.Group) {
				g.Op("&").Id(options).Values()
			})
		})

	f.Comment("Build initializes the activity context and input")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("Build").
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
		}).
		ParamsFunc(func(g *j.Group) {
			g.Id("ao").Qual(workflowPkg, "ActivityOptions")
			g.Id("input").Op("*").Qual(xnsv1Pkg, "GetWorkflowRequest")
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// initialize activity options
			g.IfFunc(func(g *j.Group) {
				g.Id("o").Dot("activityOptions").Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("ao").Op("=").Op("*").Id("o").Dot("activityOptions")
			})

			//
		})

	f.Comment("WithActivityOptions can be used to customize the activity options")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("WithActivityOptions").
		ParamsFunc(func(g *j.Group) {
			g.Id("ao").Qual(workflowPkg, "ActivityOptions")
		}).
		Op("*").Id(options).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("activityOptions").Op("=").Op("&").Id("ao")
			g.Return(j.Id("o"))
		})

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("WithHeartbeatInterval").
		ParamsFunc(func(g *j.Group) {
			g.Id("d").Qual("time", "Duration")
		}).
		Op("*").Id(options).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("heartbeatInterval").Op("=").Id("d")
			g.Return(j.Id("o"))
		})

	f.Comment("WithParentClosePolicy can be used to customize the cancellation propagation behavior")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("o").Op("*").Id(options)
		}).
		Id("WithParentClosePolicy").
		ParamsFunc(func(g *j.Group) {
			g.Id("policy").Qual(enumsPkg, "ParentClosePolicy")
		}).
		Op("*").Id(options).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("parentClosePolicy").Op("=").Id("policy")
			g.Return(j.Id("o"))
		})
}

func (m *Manifest) genXNSWorkflowOptions(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sWorkflowOptions", workflow)
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	hasWorkflowInput := !isEmpty(method.Input)

	f.Commentf("%s are used to configure a(n) %s workflow execution", typeName, m.fqnForWorkflow(workflow))
	f.Type().Id(typeName).Struct(
		j.Id("ActivityOptions").Op("*").Qual(workflowPkg, "ActivityOptions"),
		j.Id("Detached").Bool(),
		j.Id("HeartbeatInterval").Qual("time", "Duration"),
		j.Id("ParentClosePolicy").Qual(enumsPkg, "ParentClosePolicy"),
		j.Id("StartWorkflowOptions").Op("*").Qual(clientPkg, "StartWorkflowOptions"),
	)

	f.Commentf("New%s initializes a new %s value", typeName, typeName)
	f.Func().
		Id(m.toCamel("New%s", typeName)).
		Params().
		Op("*").Id(typeName).
		Block(
			j.Return(
				j.Op("&").Id(typeName).Values(),
			),
		)

	f.Comment("Build initializes the activity context and input")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("opts").Op("*").Id(typeName)
		}).
		Id("Build").
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			if hasWorkflowInput {
				g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		ParamsFunc(func(g *j.Group) {
			g.Qual(workflowPkg, "Context")
			g.Op("*").Qual(xnsv1Pkg, "WorkflowRequest")
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// initialize start workflow options
			g.Comment("initialize start workflow options")
			g.Id("swo").Op(":=").Qual(clientPkg, "StartWorkflowOptions").Values()
			g.If(j.Id("opts").Dot("StartWorkflowOptions").Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Id("swo").Op("=").Op("*").Id("opts").Dot("StartWorkflowOptions")
			})
			g.Line()

			// initialize workflow id if not set
			g.Comment("initialize workflow id if not set")
			g.If(j.Id("swo").Dot("ID").Op("==").Lit("")).BlockFunc(func(g *j.Group) {
				g.IfFunc(func(g *j.Group) {
					g.Id("serr").Op(":=").Qual(workflowPkg, "SideEffect").
						CallFunc(func(g *j.Group) {
							g.Id("ctx")
							g.Func().
								ParamsFunc(func(g *j.Group) {
									g.Id("ctx").Qual(workflowPkg, "Context")
								}).
								Any().
								BlockFunc(func(g *j.Group) {
									// if id expression defined, use it
									if idExpr := opts.GetId(); idExpr != "" {
										g.Var().Id("id").String()
										g.IfFunc(func(g *j.Group) {
											g.List(j.Id("id"), j.Err()).Op("=").Qual(expressionPkg, "EvalExpression").CallFunc(func(g *j.Group) {
												g.Qual(string(m.File.GoImportPath), m.Names().workflowIDExpression(workflow))
												if hasWorkflowInput {
													g.Id("input").Dot("ProtoReflect").Call()
												} else {
													g.Nil()
												}
											})
											g.Err().Op("!=").Nil()
										}).BlockFunc(func(g *j.Group) {
											g.Return(j.Err())
										})
										g.Return(j.Id("id"))
									}
									// otherwise generate uuid
									g.Return(j.Qual(uuidPkg, "NewString").Call())
								})
						}).Dot("Get").Call(j.Op("&").Id("swo").Dot("ID"))
					g.Id("serr").Op("!=").Nil()
				}).BlockFunc(func(g *j.Group) {
					g.Return(j.Id("ctx"), j.Nil(), j.Id("serr"))
				}).Else().If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
					g.Return(j.Id("ctx"), j.Nil(), j.Err())
				})
			})
			g.Line()

			// marshal workflow input as anypb.Any
			if hasWorkflowInput {
				g.Comment("marshal workflow request protobuf message")
				g.List(j.Id("inputpb"), j.Err()).Op(":=").Qual(anypbPkg, "New").Call(j.Id("input"))
				g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
					g.Return(j.Id("ctx"), j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling workflow request: %w"), j.Err()))
				})
				g.Line()
			}

			// marshal start workflow options
			g.Comment("marshal start workflow options protobuf message")
			g.List(j.Id("swopb"), j.Err()).Op(":=").Qual(anypbPkg, "New").Call(j.Id("swo"))
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Id("ctx"), j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling start workflow options: %w"), j.Err()))
			})
			g.Line()

			// marshal parent close policy
			g.Comment("marshal parent close policy protobuf message")
			g.Var().Id("parentClosePolicy").Qual(temporalv1Pkg, "ParentClosePolicy")
			g.SwitchFunc(func(g *j.Group) {
				g.Id("opts").Dot("ParentClosePolicy")
			}).BlockFunc(func(g *j.Group) {
				g.Case(j.Qual(enumsPkg, "PARENT_CLOSE_POLICY_ABANDON")).Block(
					g.Id("parentClosePolicy").Op("=").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON"),
				)
				g.Case(j.Qual(enumsPkg, "PARENT_CLOSE_POLICY_REQUEST_CANCEL")).Block(
					g.Id("parentClosePolicy").Op("=").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL"),
				)
				g.Case(j.Qual(enumsPkg, "PARENT_CLOSE_POLICY_TERMINATE")).Block(
					g.Id("parentClosePolicy").Op("=").Qual(temporalv1Pkg, "ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE"),
				)
			})

		})

	f.Comment("WithActivityOptions can be used to customize the activity options")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithActivityOptions").
		Params(
			j.Id("ao").Qual(workflowPkg, "ActivityOptions"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("ActivityOptions").Op("=").Op("&").Id("ao"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithDetached can be used to start a workflow execution and exit immediately")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithDetached").
		Params(
			j.Id("d").Bool(),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("Detached").Op("=").Id("d"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithHeartbeatInterval can be used to customize the activity heartbeat interval")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithHeartbeatInterval").
		Params(
			j.Id("d").Qual("time", "Duration"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("HeartbeatInterval").Op("=").Id("d"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithParentClosePolicy can be used to customize the cancellation propagation behavior")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithParentClosePolicy").
		Params(
			j.Id("policy").Qual(enumsPkg, "ParentClosePolicy"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("ParentClosePolicy").Op("=").Id("policy"),
			j.Return(j.Id("opts")),
		)

	f.Comment("WithStartWorkflowOptions can be used to customize the start workflow options")
	f.Func().
		Params(
			j.Id("opts").Op("*").Id(typeName),
		).
		Id("WithStartWorkflow").
		Params(
			j.Id("swo").Qual(clientPkg, "StartWorkflowOptions"),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("StartWorkflowOptions").Op("=").Op("&").Id("swo"),
			j.Return(j.Id("opts")),
		)
}

func (m *Manifest) genXNSWorkflowRunImpl(f *j.File, workflow protoreflect.FullName) {
	iface := m.Names().xnsWorkflowRunIface(workflow)
	impl := m.Names().xnsWorkflowRunImpl(workflow)
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	hasOutput := !isEmpty(method.Output)
	getAsync := m.Names().xnsWorkflowGetFunctionAsync(workflow)

	f.Commentf("%s provides a(n) %s implementation", impl, iface)
	f.Type().Id(impl).Struct(
		j.Id("cancel").Func().Params(),
		j.Id("ctx").Qual(workflowPkg, "Context"),
		j.Id("future").Qual(workflowPkg, "Future"),
		j.Id("id").String(),
	)

	f.Comment("Cancel the underlying workflow execution")
	f.Func().
		Params(j.Id("r").Op("*").Id(impl)).
		Id("Cancel").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		Error().
		Block(
			j.If(j.Id("r").Dot("cancel").Op("!=").Nil()).Block(
				j.Id("r").Dot("cancel").Call(),
				j.If(
					j.ListFunc(func(g *j.Group) {
						if hasOutput {
							g.Id("_")
						}
						g.Err()
					}).Op(":=").Id("r").Dot("Get").Call(j.Id("ctx")),
					j.Err().Op("!=").Nil().Op("&&").Op("!").Qual("errors", "Is").Call(j.Err(), j.Qual(workflowPkg, "ErrCanceled")),
				).Block(
					j.Return(j.Err()),
				),
				j.Return(j.Nil()),
			),
			j.Return(j.Id(m.toCamel("Cancel%sWorkflow", m.GoName)).Call(j.Id("ctx"), j.Id("r").Dot("id"), j.Lit(""))),
		)

	f.Comment("Future returns the underlying activity future")
	f.Func().
		Params(j.Id("r").Op("*").Id(impl)).
		Id("Future").
		Params().
		Qual(workflowPkg, "Future").
		BlockFunc(func(g *j.Group) {
			g.IfFunc(func(g *j.Group) {
				g.Id("r").Dot("future").Op("==").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("rr").Op(":=").Id(getAsync).CallFunc(func(g *j.Group) {
					g.Id("r").Dot("ctx")
					g.Id("r").Dot("id")
					g.Lit("")
				}).Op(".").Parens(j.Op("*").Id(impl))
				g.Id("r").Dot("future").Op("=").Id("rr").Dot("future")
				g.Id("r").Dot("cancel").Op("=").Id("rr").Dot("cancel")
			})
			g.Return(j.Id("r").Dot("future"))
		})

	f.Comment("Get blocks on activity completion and returns the underlying workflow result")
	f.Func().
		ParamsFunc(func(g *j.Group) {
			g.Id("r").Op("*").Id(impl)
		}).
		Id("Get").
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// initialize future to get workflow activity if not already set
			g.IfFunc(func(g *j.Group) {
				g.Id("r").Dot("future").Op("==").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("rr").Op(":=").Id(getAsync).CallFunc(func(g *j.Group) {
					g.Id("r").Dot("ctx")
					g.Id("r").Dot("id")
					g.Lit("")
				}).Op(".").Parens(j.Op("*").Id(impl))
				g.Id("r").Dot("future").Op("=").Id("rr").Dot("future")
				g.Id("r").Dot("cancel").Op("=").Id("rr").Dot("cancel")
			})

			if hasOutput {
				g.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.IfFunc(func(g *j.Group) {
				g.Err().Op(":=").Id("r").Dot("future").Dot("Get").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					if hasOutput {
						g.Op("&").Id("resp")
					} else {
						g.Nil()
					}
				})
				g.Err().Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Err()
				})
			})
			g.ReturnFunc(func(g *j.Group) {
				if hasOutput {
					g.Op("&").Id("resp")
				}
				g.Nil()
			})
		})

	f.Comment("ID returns the underlying workflow id")
	f.Func().
		Params(j.Id("r").Op("*").Id(impl)).
		Id("ID").
		Params().
		String().
		Block(
			j.Return(j.Id("r").Dot("id")),
		)

	for i := range opts.GetQuery() {
		query := getFullyQualifiedRef(workflow, opts.GetQuery()[i].GetRef())
		handler := m.methods[query]
		handlerInput := !isEmpty(handler.Input)
		handlerOutput := !isEmpty(handler.Output)

		methodName := m.toCamel("%s", query)
		commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s query and blocks until completion", methodName, m.fqnForQuery(query))
		f.Func().
			Params(j.Id("r").Op("*").Id(impl)).
			Id(methodName).
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
				g.Id("opts").Op("...").Op("*").Qual(m.methodXNSPackage(query), m.toCamel("%sQueryOptions", query))
			}).
			ParamsFunc(func(g *j.Group) {
				if handlerOutput {
					g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
				}
				g.Error()
			}).
			BlockFunc(func(g *j.Group) {
				if xns := opts.GetQuery()[i].GetXns(); xns != nil {
					g.Comment("configure activity options if unset")
					g.Id("opt").Op(":=").Op("&").Qual(m.methodXNSPackage(query), m.toCamel("%sQueryOptions", query)).Values()
					g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
						j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
					)
					g.If(j.Id("opt").Dot("ActivityOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
						initializeXNSOptions(g, xns, opts.GetExecutionTimeout().AsDuration())
						g.Id("opt").Dot("ActivityOptions").Op("=").Op("&").Id("ao")
						g.Id("opts").Index(j.Lit(0)).Op("=").Id("opt")
					})
				}
				g.Return(j.Qual(m.methodXNSPackage(query), methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Lit("")
					if handlerInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}))
			})

		methodName = m.toCamel("%sAsync", query)
		commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s query and returns a handle to the underlying activity", methodName, m.fqnForQuery(query))
		f.Func().
			Params(j.Id("r").Op("*").Id(impl)).
			Id(methodName).
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
				g.Id("opts").Op("...").Op("*").Qual(m.methodXNSPackage(query), m.toCamel("%sQueryOptions", query))
			}).
			Params(
				j.Qual(m.methodXNSPackage(query), m.toCamel("%sQueryHandle", query)),
				j.Error(),
			).
			BlockFunc(func(g *j.Group) {
				if xns := opts.GetQuery()[i].GetXns(); xns != nil {
					g.Comment("configure activity options if unset")
					g.Id("opt").Op(":=").Op("&").Qual(m.methodXNSPackage(query), m.toCamel("%sQueryOptions", query)).Values()
					g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
						j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
					)
					g.If(j.Id("opt").Dot("ActivityOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
						initializeXNSOptions(g, xns, opts.GetExecutionTimeout().AsDuration())
						g.Id("opt").Dot("ActivityOptions").Op("=").Op("&").Id("ao")
						g.Id("opts").Index(j.Lit(0)).Op("=").Id("opt")
					})
				}
				g.Return(j.Qual(m.methodXNSPackage(query), methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Lit("")
					if handlerInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}))
			})
	}

	for i := range opts.GetSignal() {
		signal := getFullyQualifiedRef(workflow, opts.GetSignal()[i].GetRef())
		handler := m.methods[signal]
		handlerInput := !isEmpty(handler.Input)

		methodName := m.toCamel("%s", signal)
		commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s signal and blocks until the underlying activity completes", methodName, m.fqnForSignal(signal))
		f.Func().
			Params(j.Id("r").Op("*").Id(impl)).
			Id(methodName).
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
				g.Id("opts").Op("...").Op("*").Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalOptions", signal))
			}).
			Error().
			BlockFunc(func(g *j.Group) {
				if xns := opts.GetSignal()[i].GetXns(); xns != nil {
					g.Comment("configure activity options if unset")
					g.Id("opt").Op(":=").Op("&").Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalOptions", signal)).Values()
					g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
						j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
					)
					g.If(j.Id("opt").Dot("ActivityOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
						initializeXNSOptions(g, xns, opts.GetExecutionTimeout().AsDuration())
						g.Id("opt").Dot("ActivityOptions").Op("=").Op("&").Id("ao")
						g.Id("opts").Index(j.Lit(0)).Op("=").Id("opt")
					})
				}
				g.Return(j.Qual(m.methodXNSPackage(signal), methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Lit("")
					if handlerInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}))
			})

		methodName = m.toCamel("%sAsync", signal)
		commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s signal and returns a handle to the underlying activity", methodName, m.fqnForSignal(signal))
		f.Func().
			Params(j.Id("r").Op("*").Id(impl)).
			Id(methodName).
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
				g.Id("opts").Op("...").Op("*").Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalOptions", signal))
			}).
			Params(
				j.Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalHandle", signal)),
				j.Error(),
			).
			BlockFunc(func(g *j.Group) {
				if xns := opts.GetSignal()[i].GetXns(); xns != nil {
					g.Comment("configure activity options if unset")
					g.Id("opt").Op(":=").Op("&").Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalOptions", signal)).Values()
					g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
						j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
					)
					g.If(j.Id("opt").Dot("ActivityOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
						initializeXNSOptions(g, xns, opts.GetExecutionTimeout().AsDuration())
						g.Id("opt").Dot("ActivityOptions").Op("=").Op("&").Id("ao")
						g.Id("opts").Index(j.Lit(0)).Op("=").Id("opt")
					})
				}
				g.Return(j.Qual(m.methodXNSPackage(signal), methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Lit("")
					if handlerInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}))
			})
	}

	for i := range opts.GetUpdate() {
		update := getFullyQualifiedRef(workflow, opts.GetUpdate()[i].GetRef())
		handler := m.methods[update]
		handlerInput := !isEmpty(handler.Input)
		handlerOutput := !isEmpty(handler.Output)

		methodName := m.toCamel("%s", update)
		commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s update and blocks until completion", methodName, m.fqnForUpdate(update))
		f.Func().
			Params(j.Id("r").Op("*").Id(impl)).
			Id(methodName).
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
				g.Id("opts").Op("...").Op("*").Qual(m.methodXNSPackage(update), m.toCamel("%sUpdateOptions", update))
			}).
			ParamsFunc(func(g *j.Group) {
				if handlerOutput {
					g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
				}
				g.Error()
			}).
			BlockFunc(func(g *j.Group) {
				if xns := opts.GetUpdate()[i].GetXns(); xns != nil {
					g.Comment("configure activity options if unset")
					g.Id("opt").Op(":=").Op("&").Qual(m.methodXNSPackage(update), m.toCamel("%sUpdateOptions", update)).Values()
					g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
						j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
					)
					g.If(j.Id("opt").Dot("ActivityOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
						initializeXNSOptions(g, xns, opts.GetExecutionTimeout().AsDuration())
						g.Id("opt").Dot("ActivityOptions").Op("=").Op("&").Id("ao")
						g.Id("opts").Index(j.Lit(0)).Op("=").Id("opt")
					})
				}
				g.Return(j.Qual(m.methodXNSPackage(update), methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Lit("")
					if handlerInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}))
			})

		methodName = m.toCamel("%sAsync", update)
		commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s update and returns a handle to the underlying activity", methodName, m.fqnForUpdate(update))
		f.Func().
			Params(j.Id("r").Op("*").Id(impl)).
			Id(methodName).
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if handlerInput {
					g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
				g.Id("opts").Op("...").Op("*").Qual(m.methodXNSPackage(update), m.toCamel("%sUpdateOptions", update))
			}).
			Params(
				j.Qual(m.methodXNSPackage(update), m.toCamel("%sHandle", update)),
				j.Error(),
			).
			BlockFunc(func(g *j.Group) {
				if xns := opts.GetUpdate()[i].GetXns(); xns != nil {
					g.Comment("configure activity options if unset")
					g.Id("opt").Op(":=").Op("&").Qual(m.methodXNSPackage(update), m.toCamel("%sUpdateOptions", update)).Values()
					g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
						j.Id("opt").Op("=").Id("opts").Index(j.Lit(0)),
					)
					g.If(j.Id("opt").Dot("ActivityOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
						initializeXNSOptions(g, xns, opts.GetExecutionTimeout().AsDuration())
						g.Id("opt").Dot("ActivityOptions").Op("=").Op("&").Id("ao")
						g.Id("opts").Index(j.Lit(0)).Op("=").Id("opt")
					})
				}
				g.Return(j.Qual(m.methodXNSPackage(update), methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Lit("")
					if handlerInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}))
			})
	}
}

func (m *Manifest) genXNSWorkflowRunInterface(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sRun", workflow)
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s provides a handle to a %s workflow execution", typeName, m.fqnForWorkflow(workflow))
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		methods.Comment("Cancel cancels the workflow")
		methods.Id("Cancel").
			Params(j.Qual(workflowPkg, "Context")).
			Error().Line()

		methods.Comment("Future returns the inner workflow.Future")
		methods.Id("Future").Params().Qual(workflowPkg, "Future").Line()

		methods.Comment("Get returns the inner workflow.Future")
		methods.Id("Get").
			Params(j.Qual(workflowPkg, "Context")).
			ParamsFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				returnVals.Error()
			}).Line()

		methods.Comment("ID returns the workflow id")
		methods.Id("ID").
			Params().
			String().
			Line()

		for i := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, opts.GetQuery()[i].GetRef())
			handler := m.methods[query]
			handlerInput := !isEmpty(handler.Input)
			handlerOutput := !isEmpty(handler.Output)

			// synchronous
			methodName := m.toCamel("%s", query)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s query and blocks until completion", methodName, m.fqnForQuery(query))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Qual(m.methodXNSPackage(query), m.toCamel("%sQueryOptions", query))
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if handlerOutput {
						returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					}
					returnVals.Error()
				}).
				Line()

			// async
			methodName += "Async"
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s query and returns a handle to the underlying activity", methodName, m.fqnForQuery(query))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Qual(m.methodXNSPackage(query), m.toCamel("%sQueryOptions", query))
				}).
				Params(
					j.Qual(m.methodXNSPackage(query), m.toCamel("%sQueryHandle", query)),
					j.Error(),
				).
				Line()
		}

		for i := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, opts.GetSignal()[i].GetRef())
			handler := m.methods[signal]
			handlerInput := !isEmpty(handler.Input)

			// synchronnous
			methodName := m.toCamel("%s", signal)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s signal and blocks until completion", methodName, m.fqnForSignal(signal))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalOptions", signal))
				}).
				Error().
				Line()

			// async
			methodName = m.toCamel("%sAsync", signal)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s signal and returns a handle to the underlying activity", methodName, m.fqnForSignal(signal))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalOptions", signal))
				}).
				Params(
					j.Qual(m.methodXNSPackage(signal), m.toCamel("%sSignalHandle", signal)),
					j.Error(),
				).
				Line()
		}

		for i := range opts.GetUpdate() {
			update := getFullyQualifiedRef(workflow, opts.GetUpdate()[i].GetRef())
			handler := m.methods[update]
			handlerInput := !isEmpty(handler.Input)
			handlerOutput := !isEmpty(handler.Output)

			// synchronous
			methodName := m.toCamel("%s", update)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s update and blocks until completion", methodName, m.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Qual(m.methodXNSPackage(update), m.toCamel("%sUpdateOptions", update))
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if handlerOutput {
						returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					}
					returnVals.Error()
				}).
				Line()

			// async
			methodName = m.toCamel("%sAsync", update)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s update and returns a handle to the underlying activity", methodName, m.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if handlerInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Op("...").Op("*").Qual(m.methodXNSPackage(update), m.toCamel("%sUpdateOptions", update))
				}).
				Params(
					j.Qual(m.methodXNSPackage(update), m.toCamel("%sHandle", update)),
					j.Error(),
				).
				Line()
		}
	})
}

func initializeXNSOptions(fn *j.Group, opts *temporalv1.XNSActivityOptions, defaultTimeout time.Duration) {
	if defaultTimeout == 0 {
		defaultTimeout = time.Hour * 24
	}

	// set default heartbeat interval if unset
	fn.If(j.Id("opt").Dot("HeartbeatInterval").Op("==").Lit(0)).BlockFunc(func(bl *j.Group) {
		if d := opts.GetHeartbeatInterval(); d.IsValid() {
			bl.Id("opt").Dot("HeartbeatInterval").Op("=").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Comment(durafmt.Parse(d.AsDuration()).String())
		} else {
			bl.Id("opt").Dot("HeartbeatInterval").Op("=").Qual("time", "Second").Op("*").Lit(30)
		}
	})
	fn.Line()

	fn.Comment("configure activity options")
	fn.Id("ao").Op(":=").Qual(workflowPkg, "GetActivityOptions").Call(j.Id("ctx"))
	// use user-specified activity options if non-nil
	fn.If(j.Id("opt").Dot("ActivityOptions").Op("!=").Nil()).Block(
		j.Id("ao").Op("=").Op("*").Id("opt").Dot("ActivityOptions"),
	)
	// set heartbeat timeout if unset
	fn.If(j.Id("ao").Dot("HeartbeatTimeout").Op("==").Lit(0)).BlockFunc(func(bl *j.Group) {
		if d := opts.GetHeartbeatTimeout(); d.IsValid() {
			bl.Id("ao").Dot("HeartbeatTimeout").Op("=").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Comment(durafmt.Parse(d.AsDuration()).String())
		} else {
			bl.Id("ao").Dot("HeartbeatTimeout").Op("=").Id("opt").Dot("HeartbeatInterval").Op("*").Lit(2)
		}
	})

	fn.Comment("WaitForCancellation must be set otherwise the underlying workflow is not guaranteed to be canceled")
	fn.Id("ao").Dot("WaitForCancellation").Op("=").Lit(true)
	fn.Line()

	// set retry policy if defined
	if v := opts.GetRetryPolicy(); v != nil {
		fn.If(j.Id("ao").Dot("RetryPolicy").Op("==").Nil()).Block(
			j.Id("ao").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").CustomFunc(multiLineValues, func(fields *j.Group) {
				if d := v.GetBackoffCoefficient(); d != 0 {
					fields.Id("BackoffCoefficient").Op(":").Lit(d)
				}
				if d := v.GetInitialInterval(); d.IsValid() {
					fields.Id("InitialInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
				}
				if d := v.GetMaxAttempts(); d != 0 {
					fields.Id("MaximumAttempts").Op(":").Lit(d)
				}
				if d := v.GetMaxInterval(); d.IsValid() {
					fields.Id("MaximumInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
				}
				if d := v.GetNonRetryableErrorTypes(); len(d) > 0 {
					fields.Id("NonRetryableErrorTypes").Op(":").Index().String().CustomFunc(multiLineValues, func(vals *j.Group) {
						for _, errT := range d {
							vals.Lit(errT)
						}
					})
				}
			}),
		)
	}
	var hasDefaultTimeout bool
	// set schedule-to-close if schema defined and unset
	if d := opts.GetScheduleToCloseTimeout(); d.IsValid() {
		hasDefaultTimeout = true
		fn.If(j.Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
			j.Id("ao").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Comment(durafmt.Parse(d.AsDuration()).String()),
		)
	}
	if d := opts.GetScheduleToStartTimeout(); d.IsValid() {
		fn.If(j.Id("ao").Dot("ScheduleToStartTimeout").Op("==").Lit(0)).Block(
			j.Id("ao").Dot("ScheduleToStartTimeout").Op("=").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Comment(durafmt.Parse(d.AsDuration()).String()),
		)
	}
	// set start-to-close if schema defined and unset
	if d := opts.GetStartToCloseTimeout(); d.IsValid() {
		hasDefaultTimeout = true
		fn.If(j.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0)).Block(
			j.Id("ao").Dot("StartToCloseTimeout").Op("=").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Comment(durafmt.Parse(d.AsDuration()).String()),
		)
	}
	if !hasDefaultTimeout {
		// ensure atleast one of start-to-close or schedule-to-close is set
		fn.If(j.Id("ao").Dot("StartToCloseTimeout").Op("==").Lit(0).Op("&&").Id("ao").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)).Block(
			j.Id("ao").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(defaultTimeout.Nanoseconds(), 10)).Comment(durafmt.Parse(defaultTimeout).String()),
		)
	}
	// set task queue if unset
	if v := opts.GetTaskQueue(); v != "" {
		fn.If(j.Id("ao").Dot("TaskQueue").Op("==").Lit("")).Block(
			j.Id("ao").Dot("TaskQueue").Op("=").Lit(v),
		)
	}
	fn.Id("ctx").Op("=").Qual(workflowPkg, "WithActivityOptions").Call(j.Id("ctx"), j.Id("ao"))
	fn.Line()
}

// initializes a `swo` variable that contains a non-nil *temporalv1.StartWorkflowOptions value
func (m *Manifest) initializeXNSStartWorkflowOptions(fn *j.Group, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	hasInput := !isEmpty(method.Input)

	fn.Comment("configure start workflow options")
	fn.Id("wo").Op(":=").Qual(clientPkg, "StartWorkflowOptions").Values()
	fn.If(j.Id("opt").Dot("StartWorkflowOptions").Op("!=").Nil()).Block(
		j.Id("wo").Op("=").Op("*").Id("opt").Dot("StartWorkflowOptions"),
	)
	// set workflow id if unset and  id field and/or prefix defined
	if idExpr := opts.GetId(); idExpr != "" {
		fn.If(j.Id("wo").Dot("ID").Op("==").Lit("")).Block(
			j.If(
				j.Err().Op(":=").Qual(workflowPkg, "SideEffect").Call(j.Id("ctx"), j.Func().Params(j.Id("ctx").Qual(workflowPkg, "Context")).Any().Block(
					j.List(j.Id("id"), j.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *j.Group) {
						args.Qual(string(m.File.GoImportPath), m.toCamel("%sIDExpression", workflow))
						if hasInput {
							args.Id("req").Dot("ProtoReflect").Call()
						} else {
							args.Nil()
						}
					}),
					j.If(j.Err().Op("!=").Nil()).Block(
						j.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Error").Call(
							j.Lit(fmt.Sprintf("error evaluating id expression for %q workflow", workflow)),
							j.Lit("error"),
							j.Err(),
						),
						j.Return(j.Nil()),
					),
					j.Return(j.Id("id")),
				)).Dot("Get").Call(j.Op("&").Id("wo").Dot("ID")),
				j.Err().Op("!=").Nil(),
			).Block(
				j.Return(j.Nil(), j.Err()),
			),
		)
	}
	fn.If(j.Id("wo").Dot("ID").Op("==").Lit("")).Block(
		j.If(
			j.Err().Op(":=").Qual(workflowPkg, "SideEffect").Call(j.Id("ctx"), j.Func().Params(j.Id("ctx").Qual(workflowPkg, "Context")).Any().Block(
				j.List(j.Id("id"), j.Err()).Op(":=").Qual(uuidPkg, "NewRandom").Call(),
				j.If(j.Err().Op("!=").Nil()).Block(
					j.Qual(workflowPkg, "GetLogger").Call(j.Id("ctx")).Dot("Error").Call(
						j.Lit("error generating workflow id"),
						j.Lit("error"),
						j.Err(),
					),
					j.Return(j.Nil()),
				),
				j.Return(j.Id("id")),
			)).Dot("Get").Call(j.Op("&").Id("wo").Dot("ID")),
			j.Err().Op("!=").Nil(),
		).Block(
			j.Return(j.Nil(), j.Err()),
		),
	)
	fn.If(j.Id("wo").Dot("ID").Op("==").Lit("")).Block(
		j.Return(
			j.Nil(),
			j.Qual(temporalPkg, "NewNonRetryableApplicationError").Call(
				j.Lit("workflow id is required"),
				j.Lit("InvalidArgument"),
				j.Nil(),
			),
		),
	)
	fn.Line()

	fn.Comment("marshal start workflow options protobuf message")
	fn.List(j.Id("swo"), j.Err()).Op(":=").Qual(xnsPkg, "MarshalStartWorkflowOptions").Call(j.Id("wo"))
	fn.If(j.Err().Op("!=").Nil()).Block(
		j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error marshalling start workflow options: %w"), j.Err())),
	)
	fn.Line()
}
