package plugin

import (
	"strconv"

	"github.com/dave/jennifer/jen"
	"github.com/hako/durafmt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (svc *Manifest) renderNexus(f *jen.File) {
	svc.genSectionHeader(f, "Nexus")
	svc.genNexusService(f)

	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc || svc.workflows[workflow].GetNexus().GetDisabled() {
			continue
		}
		svc.genNexusWorkflowOperationNameConst(f, workflow)
		svc.genNexusWorkflowOperation(f, workflow)
		svc.genNexusWorkflowOperationFutureIface(f, workflow)
		svc.genNexusWorkflowOperationFutureImpl(f, workflow)
		svc.genNexusWorkflowOperationOptions(f, workflow)
	}

	svc.genNexusClientIface(f)
	svc.genNexusClientImpl(f)
}

func (svc *Manifest) genNexusClientIface(f *jen.File) {
	ifaceName := svc.toCamel("%sNexusClient", svc.Service.GoName)
	f.Commentf("%s describes a(n) %s nexus client", ifaceName, svc.Service.Desc.FullName())
	f.Type().
		Id(ifaceName).
		InterfaceFunc(func(g *jen.Group) {
			for _, workflow := range svc.workflowsOrdered {
				if svc.methods[workflow].Desc.Parent() != svc.Service.Desc || svc.workflows[workflow].GetNexus().GetDisabled() {
					continue
				}

				method := svc.methods[workflow]
				hasInput := !isEmpty(method.Input)
				hasOutput := !isEmpty(method.Output)
				optionsType := svc.toCamel("%sWorkflowOperationOptions", workflow)

				methodName := svc.methods[workflow].GoName
				commentWithDefaultf(g, methodSet(method), "%s executes a(n) %s nexus workflow operation and blocks until complete", methodName, workflow)
				g.Id(methodName).
					ParamsFunc(func(g *jen.Group) {
						g.Id("ctx").Qual(workflowPkg, "Context")
						if hasInput {
							g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), method.Input.GoIdent.GoName)
						}
						g.Id("options").Op("...").Op("*").Id(optionsType)
					}).
					ParamsFunc(func(g *jen.Group) {
						if hasOutput {
							g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), method.Output.GoIdent.GoName)
						}
						g.Error()
					})

				methodName = svc.toCamel("%sAsync", svc.methods[workflow].GoName)
				commentWithDefaultf(g, methodSet(method), "%s starts a(n) %s nexus workflow operation and returns a handle to the operation", methodName, workflow)
				g.Id(methodName).
					ParamsFunc(func(g *jen.Group) {
						g.Id("ctx").Qual(workflowPkg, "Context")
						if hasInput {
							g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), method.Input.GoIdent.GoName)
						}
						g.Id("options").Op("...").Op("*").Id(optionsType)
					}).
					ParamsFunc(func(g *jen.Group) {
						g.Id(svc.toCamel("%sWorkflowOperationFuture", workflow))
					})
			}
		})
}

func (svc *Manifest) genNexusClientImpl(f *jen.File) {
	ifaceName := svc.toCamel("%sNexusClient", svc.Service.GoName)
	implName := svc.toLowerCamel("%sNexusClient", svc.Service.GoName)

	name := svc.opts.GetNexus().GetName()
	if name == "" {
		name = string(svc.Service.Desc.FullName())
	}

	f.Commentf("%s provides an internal %s implementation", implName, ifaceName)
	f.Type().
		Id(implName).
		StructFunc(func(g *jen.Group) {
			g.Id("client").Qual(workflowPkg, "NexusClient")
		})

	ctorName := svc.toCamel("New%sNexusClient", svc.Service.GoName)
	f.Commentf("%s initializes a new %s nexus client", ctorName, name)
	f.Func().
		Id(ctorName).
		ParamsFunc(func(g *jen.Group) {
			g.Id("endpoint").String()
		}).
		Id(ifaceName).
		BlockFunc(func(g *jen.Group) {
			g.ReturnFunc(func(g *jen.Group) {
				g.Op("&").Id(implName).CustomFunc(multiLineValues, func(g *jen.Group) {
					g.Id("client").Op(":").Qual(workflowPkg, "NewNexusClient").CallFunc(func(g *jen.Group) {
						g.Id("endpoint")
						g.Id(svc.toCamel("%sServiceName", svc.Service.GoName))
					})
				})
			})
		})

	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc || svc.workflows[workflow].GetNexus().GetDisabled() {
			continue
		}
		method := svc.methods[workflow]
		hasInput := !isEmpty(method.Input)
		hasOutput := !isEmpty(method.Output)
		optionsType := svc.toCamel("%sWorkflowOperationOptions", workflow)

		methodName := method.GoName
		commentWithDefaultf(f, methodSet(method), "%s executes a(n) %s nexus workflow operation and blocks until complete", methodName, workflow)
		f.Func().
			ParamsFunc(func(g *jen.Group) {
				g.Id("c").Op("*").Id(implName)
			}).
			Id(methodName).
			ParamsFunc(func(g *jen.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), method.Input.GoIdent.GoName)
				}
				g.Id("options").Op("...").Op("*").Id(optionsType)
			}).
			ParamsFunc(func(g *jen.Group) {
				if hasOutput {
					g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), method.Output.GoIdent.GoName)
				}
				g.Error()
			}).
			BlockFunc(func(g *jen.Group) {
				g.ReturnFunc(func(g *jen.Group) {
					g.Id("c").Dot(svc.toCamel("%sAsync", workflow)).
						CallFunc(func(g *jen.Group) {
							g.Id("ctx")
							if hasInput {
								g.Id("input")
							}
							g.Id("options").Op("...")
						}).
						Dot("Get").
						CallFunc(func(g *jen.Group) {
							g.Id("ctx")
						})
				})
			})

		methodName = svc.toCamel("%sAsync", workflow)
		commentWithDefaultf(f, methodSet(method), "%s starts a(n) %s nexus workflow operation, returning a handle to the operation")
		f.Func().
			ParamsFunc(func(g *jen.Group) {
				g.Id("c").Op("*").Id(implName)
			}).
			Id(methodName).
			ParamsFunc(func(g *jen.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), method.Input.GoIdent.GoName)
				}
				g.Id("options").Op("...").Op("*").Id(optionsType)
			}).
			Id(svc.toCamel("%sWorkflowOperationFuture", workflow)).
			BlockFunc(func(g *jen.Group) {
				g.Var().Id("o").Op("*").Id(optionsType)
				g.IfFunc(func(g *jen.Group) {
					g.Len(jen.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(jen.Lit(0)).Op("!=").Nil()
				}).BlockFunc(func(g *jen.Group) {
					g.Id("o").Op("=").Id("options").Index(jen.Lit(0))
				}).Line()

				g.List(jen.Id("opts"), jen.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *jen.Group) {
					g.Id("ctx")
				})
				g.IfFunc(func(g *jen.Group) {
					g.Err().Op("!=").Nil()
				}).BlockFunc(func(g *jen.Group) {
					g.ReturnFunc(func(g *jen.Group) {
						g.Op("&").Id(svc.toLowerCamel("%sWorkflowOperationFuture", workflow)).CustomFunc(multiLineValues, func(g *jen.Group) {
							g.Id("f").Op(":").Qual(errsPkg, "NewNexusOperationFutureError").Call(jen.Err())
						})
					})
				}).Line()

				g.ReturnFunc(func(g *jen.Group) {
					g.Op("&").Id(svc.toLowerCamel("%sWorkflowOperationFuture", workflow)).CustomFunc(multiLineValues, func(g *jen.Group) {
						g.Id("f").Op(":").Id("c").Dot("client").Dot("ExecuteOperation").CallFunc(func(g *jen.Group) {
							g.Id("ctx")
							g.Id(svc.toCamel("%sWorkflowOperationName", workflow))
							if hasInput {
								g.Id("input")
							} else {
								g.Nil()
							}
							g.Id("opts")
						})
					})
				})
			})
	}
}

func (svc *Manifest) genNexusService(f *jen.File) {
	name := svc.opts.GetNexus().GetName()
	if name == "" {
		name = string(svc.Service.Desc.FullName())
	}

	serviceName := svc.toCamel("%sServiceName", svc.Service.GoName)
	f.Commentf("%s defines the fully-qualified %s nexus service name", serviceName, svc.Service.GoName)
	f.Const().Id(serviceName).Op("=").Lit(name)

	registerOperations := svc.toCamel("Register%sOperations", svc.Service.GoName)
	f.Commentf("%s registers all %s nexus operations with the given service", registerOperations, svc.Service.Desc.FullName())
	f.Func().
		Id(registerOperations).
		ParamsFunc(func(g *jen.Group) {
			g.Id("svc").Op("*").Qual(nexusPkg, "Service")
		}).
		Error().
		BlockFunc(func(g *jen.Group) {
			g.ReturnFunc(func(g *jen.Group) {
				g.Id("svc").Dot("Register").CustomFunc(multiLineArgs, func(g *jen.Group) {
					for _, workflow := range svc.workflowsOrdered {
						if svc.methods[workflow].Desc.Parent() != svc.Service.Desc || svc.workflows[workflow].GetNexus().GetDisabled() {
							continue
						}
						g.Id(svc.toCamel("%sWorkflowOperation", svc.methods[workflow].GoName))
					}
				})
			})
		})

	registerService := svc.toCamel("Register%sService", svc.Service.GoName)
	f.Commentf("%s registers a %s nexus service with a given worker", registerService, svc.Service.Desc.FullName())
	f.Func().
		Id(registerService).
		ParamsFunc(func(g *jen.Group) {
			g.Id("r").Qual(workerPkg, "NexusServiceRegistry")
		}).
		Error().
		BlockFunc(func(g *jen.Group) {
			g.Id("svc").Op(":=").Qual(nexusPkg, "NewService").CallFunc(func(g *jen.Group) {
				g.Id(serviceName)
			})

			g.
				IfFunc(func(g *jen.Group) {
					g.Err().Op(":=").Id(registerOperations).CallFunc(func(g *jen.Group) {
						g.Id("svc")
					})
					g.Err().Op("!=").Nil()
				}).
				BlockFunc(func(g *jen.Group) {
					g.ReturnFunc(func(g *jen.Group) {
						g.Err()
					})
				})

			g.Id("r").Dot("RegisterNexusService").CallFunc(func(g *jen.Group) {
				g.Id("svc")
			})
			g.ReturnFunc(func(g *jen.Group) {
				g.Nil()
			})
		})
}

func (svc *Manifest) genNexusWorkflowOperationNameConst(f *jen.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	operationName := svc.toCamel("%sWorkflowOperationName", workflow)

	name := svc.workflows[workflow].GetNexus().GetName()
	if name == "" {
		name = string(method.Desc.Name())
	}

	f.Commentf("%s defines the fully-qualified name for a %s nexus workflow operation", operationName, workflow)
	f.Const().Id(operationName).Op("=").Lit(name)
}

func (svc *Manifest) genNexusWorkflowOperation(f *jen.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	operationName := svc.toCamel("%sWorkflowOperationName", workflow)
	operation := svc.toCamel("%sWorkflowOperation", workflow)

	f.Commentf("%s defines a(n) %s nexus workflow operation", operation, workflow)
	f.Var().Id(operation).Op("=").Qual(temporalnexusPkg, "MustNewWorkflowRunOperationWithOptions").CustomFunc(multiLineArgs, func(g *jen.Group) {
		g.Qual(temporalnexusPkg, "WorkflowRunOperationOptions").
			CustomFunc(typeParams, func(g *jen.Group) {
				if hasInput {
					g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
				} else {
					g.Id("input").Op("*").Qual(emptypbPkg, "Empty")
				}
				if hasOutput {
					g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				} else {
					g.Op("*").Qual(emptypbPkg, "Empty")
				}
			}).
			CustomFunc(multiLineValues, func(g *jen.Group) {
				g.Id("Name").Op(":").Id(operationName)
				g.Id("Handler").Op(":").Func().
					ParamsFunc(func(g *jen.Group) {
						g.Id("ctx").Qual("context", "Context")
						if hasInput {
							g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
						} else {
							g.Id("input").Op("*").Qual(emptypbPkg, "Empty")
						}
						g.Id("opts").Qual(nexusPkg, "StartOperationOptions")
					}).
					ParamsFunc(func(g *jen.Group) {
						g.Qual(temporalnexusPkg, "WorkflowHandle").CustomFunc(typeParams, func(g *jen.Group) {
							if hasOutput {
								g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
							} else {
								g.Op("*").Qual(emptypbPkg, "Empty")
							}
						})
						g.Error()
					}).
					BlockFunc(func(g *jen.Group) {
						g.List(jen.Id("o"), jen.Err()).Op(":=").Id(svc.toCamel("New%sOptions", workflow)).Call().Dot("Build").CallFunc(func(g *jen.Group) {
							g.Id("input").Dot("ProtoReflect").Call()
						})
						g.If(jen.Err().Op("!=").Nil()).BlockFunc(func(g *jen.Group) {
							g.Return(jen.Nil(), jen.Err())
						})
						g.Return(
							jen.Qual(temporalnexusPkg, "ExecuteUntypedWorkflow").CustomFunc(typeParams, func(g *jen.Group) {
								if hasOutput {
									g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
								} else {
									g.Op("*").Qual(emptypbPkg, "Empty")
								}
							}).CallFunc(func(g *jen.Group) {
								g.Id("ctx")
								g.Id("opts")
								g.Id("o")
								g.Id(svc.toCamel("%sWorkflowName", workflow))
								if hasInput {
									g.Id("input")
								}
							}),
						)
					})
			})
	})
}

func (svc *Manifest) genNexusWorkflowOperationFutureIface(f *jen.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)
	ifaceName := svc.toCamel("%sWorkflowOperationFuture", workflow)
	f.Commentf("%s describes a handle to an asynchronous %s nexus workflow operation", ifaceName, workflow)
	f.Type().
		Id(ifaceName).
		InterfaceFunc(func(g *jen.Group) {
			g.Comment("Future returns the underlying NexusOperationFuture")
			g.Id("Future").Params().Qual(workflowPkg, "NexusOperationFuture")

			g.Comment("Get blocks until the nexus operation is complete, returning the result or error")
			g.Id("Get").
				ParamsFunc(func(g *jen.Group) {
					g.Id("ctx").Qual(workflowPkg, "Context")
				}).
				ParamsFunc(func(g *jen.Group) {
					if hasOutput {
						g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
					}
					g.Error()
				})
		})
}

func (svc *Manifest) genNexusWorkflowOperationFutureImpl(f *jen.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	ifaceName := svc.toCamel("%sWorkflowOperationFuture", workflow)
	implName := svc.toLowerCamel("%sWorkflowOperationFuture", method.GoName)
	f.Commentf("%s provides an internal %s implementation", implName, ifaceName)
	f.Type().
		Id(implName).
		StructFunc(func(g *jen.Group) {
			g.Id("f").Qual(workflowPkg, "NexusOperationFuture")
		})

	f.Comment("Future returns the underlying NexusOperationFuture")
	f.Func().
		ParamsFunc(func(g *jen.Group) {
			g.Id("f").Op("*").Id(implName)
		}).
		Id("Future").
		Params().
		Qual(workflowPkg, "NexusOperationFuture").
		BlockFunc(func(g *jen.Group) {
			g.ReturnFunc(func(g *jen.Group) {
				g.Id("f").Dot("f")
			})
		})

	f.Comment("Get blocks until the nexus operation is complete, returning the result or error")
	f.Func().
		ParamsFunc(func(g *jen.Group) {
			g.Id("f").Op("*").Id(implName)
		}).
		Id("Get").
		ParamsFunc(func(g *jen.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
		}).
		ParamsFunc(func(g *jen.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *jen.Group) {
			if hasOutput {
				g.Var().Id("out").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				g.ReturnFunc(func(g *jen.Group) {
					g.Op("&").Id("out")
					g.Id("f").Dot("f").Dot("Get").CallFunc(func(g *jen.Group) {
						g.Id("ctx")
						g.Op("&").Id("out")
					})
				})
			} else {
				g.ReturnFunc(func(g *jen.Group) {
					g.Id("f").Dot("f").Dot("Get").CallFunc(func(g *jen.Group) {
						g.Id("ctx")
						g.Nil()
					})
				})
			}
		})
}

func (svc *Manifest) genNexusWorkflowOperationOptions(f *jen.File, workflow protoreflect.FullName) {
	opts := svc.workflows[workflow]

	optionsName := svc.toCamel("%sWorkflowOperationOptions", workflow)
	f.Commentf("%s provides methods for configuration a(n) %s nexus workflow operation", optionsName, workflow)
	f.Type().
		Id(optionsName).
		StructFunc(func(g *jen.Group) {
			g.Id("opts").Op("*").Qual(workflowPkg, "NexusOperationOptions")
			g.Id("scheduleToCloseTimeout").Op("*").Qual(durationpbPkg, "Duration")
		})

	ctorName := svc.toCamel("New%sWorkflowOperationOptions", workflow)
	f.Commentf("%s initializes a new %s value", ctorName, optionsName)
	f.Func().
		Id(ctorName).
		Params().
		Op("*").Id(optionsName).
		BlockFunc(func(g *jen.Group) {
			g.ReturnFunc(func(g *jen.Group) {
				g.Op("&").Id(optionsName).Values()
			})
		})

	f.Commentf("Build converts a(n) %s value into a workflow.NexusOperationOptions value", optionsName)
	f.Func().
		ParamsFunc(func(g *jen.Group) {
			g.Id("o").Op("*").Id(optionsName)
		}).
		Id("Build").
		ParamsFunc(func(g *jen.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
		}).
		ParamsFunc(func(g *jen.Group) {
			g.Qual(workflowPkg, "NexusOperationOptions")
			g.Error()
		}).
		BlockFunc(func(g *jen.Group) {
			g.Var().Id("opts").Qual(workflowPkg, "NexusOperationOptions")

			g.IfFunc(func(g *jen.Group) {
				g.Id("o").Op("==").Nil()
			}).BlockFunc(func(g *jen.Group) {
				g.ReturnFunc(func(g *jen.Group) {
					g.Id("opts")
					g.Nil()
				})
			}).Line()

			g.IfFunc(func(g *jen.Group) {
				g.Id("v").Op(":=").Id("o").Dot("opts")
				g.Id("v").Op("!=").Nil()
			}).BlockFunc(func(g *jen.Group) {
				g.Id("opts").Op("=").Op("*").Id("o").Dot("opts")
			}).Line()

			scheduleToCloseTimeout := g.IfFunc(func(g *jen.Group) {
				g.Id("v").Op(":=").Id("o").Dot("scheduleToCloseTimeout")
				g.Id("v").Dot("IsValid").Call()
			}).BlockFunc(func(g *jen.Group) {
				g.Id("opts").Dot("ScheduleToCloseTimeout").Op("=").Id("v").Dot("AsDuration").Call()
			})
			if d := opts.GetExecutionTimeout(); d.IsValid() {
				scheduleToCloseTimeout.Else().IfFunc(func(g *jen.Group) {
					g.Id("opts").Dot("ScheduleToCloseTimeout").Op("==").Lit(0)
				}).BlockFunc(func(g *jen.Group) {
					g.Id("opts").Dot("ScheduleToCloseTimeout").Op("=").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10)).Comment(durafmt.Parse(d.AsDuration()).String())
				})
			}
			scheduleToCloseTimeout.Line()

			g.ReturnFunc(func(g *jen.Group) {
				g.Id("opts")
				g.Nil()
			})
		})

	f.Comment("WithOptions overrides the initial NexusOperationOptions to which defaults and overrides are then applied")
	f.Func().
		ParamsFunc(func(g *jen.Group) {
			g.Id("o").Op("*").Id(optionsName)
		}).
		Id("WithOptions").
		ParamsFunc(func(g *jen.Group) {
			g.Id("opts").Qual(workflowPkg, "NexusOperationOptions")
		}).
		ParamsFunc(func(g *jen.Group) {
			g.Op("*").Id(optionsName)
		}).
		BlockFunc(func(g *jen.Group) {
			g.Id("o").Dot("opts").Op("=").Op("&").Id("opts")
			g.ReturnFunc(func(g *jen.Group) {
				g.Id("o")
			})
		})

	f.Comment("WithScheduleToCloseTimeout overrides the default ScheduleToCloseTimeout")
	f.Func().
		ParamsFunc(func(g *jen.Group) {
			g.Id("o").Op("*").Id(optionsName)
		}).
		Id("WithScheduleToCloseTimeout").
		ParamsFunc(func(g *jen.Group) {
			g.Id("d").Qual("time", "Duration")
		}).
		ParamsFunc(func(g *jen.Group) {
			g.Op("*").Id(optionsName)
		}).
		BlockFunc(func(g *jen.Group) {
			g.Id("o").Dot("scheduleToCloseTimeout").Op("=").Qual(durationpbPkg, "New").Call(jen.Id("d"))
			g.ReturnFunc(func(g *jen.Group) {
				g.Id("o")
			})
		})
}
