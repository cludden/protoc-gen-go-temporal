package plugin

import (
	"bytes"
	"fmt"
	"strings"

	g "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// define cli-specific import constants
const (
	base64Pkg    = "encoding/base64"
	cliPkg       = "github.com/urfave/cli/v2"
	protojsonPkg = "google.golang.org/protobuf/encoding/protojson"
	strcasePkg   = "github.com/iancoleman/strcase"
)

var (
	multiLineValues = g.Options{
		Close:     "}",
		Multi:     true,
		Open:      "{",
		Separator: ",",
	}
)

type Cli struct {
	services        map[string]*Service
	servicesOrdered []string
}

// renderCLI generates cli resources
func (svc *Service) renderCLI(f *g.File) {
	svc.genCliOptionsImpl(f)
	svc.genCliNew(f)
	svc.genCliNewCommand(f)
	svc.genCliNewCommands(f)

	// cache unmarshal functions to void duplicate declarations
	unmarshallers := map[string]struct{}{}

	// generate query request unmarshallers
	for _, query := range svc.queriesOrdered {
		if isEmpty(svc.methods[query].Input) {
			continue
		}
		if _, ok := unmarshallers[svc.methods[query].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[svc.methods[query].Input.GoIdent.GoName] = struct{}{}
		svc.genCliUnmarshalMessage(f, svc.methods[query].Input)
	}

	// generate signal request unmarshallers
	for _, signal := range svc.signalsOrdered {
		if isEmpty(svc.methods[signal].Input) {
			continue
		}
		if _, ok := unmarshallers[svc.methods[signal].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[svc.methods[signal].Input.GoIdent.GoName] = struct{}{}
		svc.genCliUnmarshalMessage(f, svc.methods[signal].Input)
	}

	// generate update request unmarshallers
	for _, update := range svc.updatesOrdered {
		if isEmpty(svc.methods[update].Input) {
			continue
		}
		if _, ok := unmarshallers[svc.methods[update].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[svc.methods[update].Input.GoIdent.GoName] = struct{}{}
		svc.genCliUnmarshalMessage(f, svc.methods[update].Input)
	}

	// generate workflow request unmarshallers
	for _, workflow := range svc.workflowsOrdered {
		if isEmpty(svc.methods[workflow].Input) {
			continue
		}
		if _, ok := unmarshallers[svc.methods[workflow].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[svc.methods[workflow].Input.GoIdent.GoName] = struct{}{}
		svc.genCliUnmarshalMessage(f, svc.methods[workflow].Input)
	}
}

// genCliFlagForField generates a cli flag for a message field
func (svc *Service) genCliFlagForField(flags *g.Group, field *protogen.Field, category string) {
	name := field.GoName
	flagName := strcase.ToKebab(name)
	usage := strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(field.Comments.Leading.String(), "//"), "\n//", ""))
	if usage == "" {
		usage = fmt.Sprintf("set the value of the operation's %q parameter", name)
	}

	// determine cli flag type
	flagType := "String"
	switch {
	case field.Desc.IsMap():
		field.Desc.MapKey()
	default:
		switch field.Desc.Kind() {
		case protoreflect.BytesKind:
			usage += " (base64-encoded)"
		case protoreflect.BoolKind:
			flagType = "Bool"
		case protoreflect.DoubleKind, protoreflect.FloatKind:
			flagType = "Float64"
		case protoreflect.EnumKind:
			var values []string
			for _, v := range field.Enum.Values {
				values = append(values, string(v.Desc.Name()))
			}
			usage += fmt.Sprintf(" (%s)", strings.Join(values, ", "))
		case protoreflect.Fixed32Kind, protoreflect.Fixed64Kind, protoreflect.Uint32Kind, protoreflect.Uint64Kind:
			flagType = "Uint64"
		case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind:
			flagType = "Int64"
		case protoreflect.MessageKind:
			additionalUsage := &bytes.Buffer{}
			fmt.Fprint(additionalUsage, usage)
			switch field.Message.Desc.FullName() {
			case "google.protobuf.Timestamp":
				fmt.Fprint(additionalUsage, " (e.g. \"2017-01-15T01:30:15.01Z\")")
			case "google.protobuf.Duration":
				fmt.Fprint(additionalUsage, " (e.g. \"3.000000001s\")")
			default:
				fmt.Fprint(additionalUsage, " (json-encoded: {")
				var fieldDocs []string
				for _, f := range field.Message.Fields {
					kind := f.Desc.Kind().String()
					if f.Message != nil {
						kind = string(f.Message.Desc.FullName())
					} else if f.Enum != nil {
						kind = string(f.Enum.Desc.FullName())
					}
					fieldDocs = append(fieldDocs, fmt.Sprintf("%s: <%s>", f.Desc.JSONName(), kind))
				}
				fmt.Fprint(additionalUsage, strings.Join(fieldDocs, ", "))
				fmt.Fprint(additionalUsage, "})")
			}
			usage = additionalUsage.String()
		case protoreflect.StringKind:
		default:
			svc.Plugin.Error(fmt.Errorf("unable to generate cli flag for field with type %q", field.Desc.Kind().String()))
			return
		}
	}
	if field.Desc.IsList() {
		flagType += "Slice"
	}
	flagType += "Flag"

	// generate flag
	flags.Op("&").Qual(cliPkg, flagType).CustomFunc(multiLineValues, func(fields *g.Group) {
		fields.Id("Name").Op(":").Lit(flagName)
		fields.Id("Usage").Op(":").Lit(strings.TrimSpace(usage))
		if svc.opts.GetFeatures().GetCli().GetCategories() && category != "" {
			fields.Id("Category").Op(":").Lit(category)
		}
	})
}

// genCliNew generates a New<Service>Cli constructor function
func (svc *Service) genCliNew(f *g.File) {
	functionName := toCamel("New%sCli", svc.Service.GoName)
	optionsName := toCamel("%sCliOptions", svc.Service.GoName)

	f.Commentf("%s initializes a cli for a(n) %s service", functionName, svc.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			g.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			g.Op("*").Qual(cliPkg, "App"),
			g.Error(),
		).
		Block(
			g.List(g.Id("commands"), g.Err()).Op(":=").Id(toLowerCamel("new%sCommands", svc.Service.GoName)).Call(g.Id("options").Op("...")),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing subcommands: %w"), g.Err())),
			),
			g.Return(
				g.Op("&").Qual(cliPkg, "App").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit(strcase.ToKebab(svc.Service.GoName))
					fields.Id("Commands").Op(":").Id("commands")
				}),
				g.Nil(),
			),
		)
}

// genCliNewCommand generates a New<Service>CliCommand constructor function
func (svc *Service) genCliNewCommand(f *g.File) {
	functionName := toCamel("New%sCliCommand", svc.Service.GoName)
	optionsName := toCamel("%sCliOptions", svc.Service.GoName)

	f.Commentf("%s initializes a cli command for a %s service with subcommands for each query, signal, update, and workflow", functionName, svc.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			g.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			g.Op("*").Qual(cliPkg, "Command"),
			g.Error(),
		).
		Block(
			g.List(g.Id("subcommands"), g.Err()).Op(":=").Id(toLowerCamel("new%sCommands", svc.Service.GoName)).Call(g.Id("options").Op("...")),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing subcommands: %w"), g.Err())),
			),
			g.Return(
				g.Op("&").Qual(cliPkg, "Command").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit(strcase.ToKebab(svc.Service.GoName))
					fields.Id("Subcommands").Op(":").Id("subcommands")
				}),
				g.Nil(),
			),
		)
}

// genCliNewCommands generates a new<Service>Commands contructor function
func (svc *Service) genCliNewCommands(f *g.File) {
	functionName := toLowerCamel("new%sCommands", svc.Service.GoName)
	optionsName := toCamel("%sCliOptions", svc.Service.GoName)

	f.Commentf("%s initializes (sub)commands for a %s cli or command", functionName, svc.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			g.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			g.Index().Op("*").Qual(cliPkg, "Command"),
			g.Error(),
		).
		Block(
			// initialize options
			g.Id("opts").Op(":=").Op("&").Id(optionsName).Values(),
			g.If(g.Len(g.Id("options")).Op(">").Lit(0)).Block(
				g.Id("opts").Op("=").Id("options").Index(g.Lit(0)),
			),

			// set default client for command
			g.If(g.Id("opts").Dot("clientForCommand").Op("==").Nil()).Block(
				g.Id("opts").Dot("clientForCommand").Op("=").Func().
					Params(g.Op("*").Qual(cliPkg, "Context")).
					Params(g.Qual(clientPkg, "Client"), g.Error()).
					Block(
						g.Return(g.Qual(clientPkg, "Dial").Call(g.Qual(clientPkg, "Options").Values())),
					),
			),

			// initialize commands
			g.Id("commands").Op(":=").Index().Op("*").Qual(cliPkg, "Command").CustomFunc(g.Options{
				Close:     "}",
				Multi:     true,
				Open:      "{",
				Separator: ",",
			}, func(cmds *g.Group) {
				// generate client query methods
				for _, query := range svc.queriesOrdered {
					svc.genCliQueryCommand(cmds, query)
				}

				// generate client signal methods
				for _, signal := range svc.signalsOrdered {
					svc.genCliSignalCommand(cmds, signal)
				}

				// generate client update methods
				for _, update := range svc.updatesOrdered {
					svc.genCliUpdateCommand(cmds, update)
				}

				// generate client workflow methods
				for _, workflow := range svc.workflowsOrdered {
					svc.genCliWorkflowCommand(cmds, workflow)
					for _, signal := range svc.workflows[workflow].GetSignal() {
						if !signal.GetStart() {
							continue
						}
						svc.genCliWorkflowWithSignalCommand(cmds, workflow, signal.GetRef())
					}
				}
			}),

			// append worker command if initializer provided
			g.If(g.Id("opts").Dot("worker").Op("!=").Nil()).Block(
				g.Id("commands").Op("=").Append(g.Id("commands"), g.Index().Op("*").Qual(cliPkg, "Command").CustomFunc(multiLineValues, func(cmds *g.Group) {
					svc.genCliWorkerCommand(cmds)
				}).Op("...")),
			),

			g.Qual("sort", "Slice").Call(
				g.Id("commands"),
				g.Func().Params(g.Id("i"), g.Id("j").Int()).Bool().Block(
					g.Return(g.Id("commands").Index(g.Id("i")).Dot("Name").Op("<").Id("commands").Index(g.Id("j")).Dot("Name")),
				),
			),
			g.Return(g.Id("commands"), g.Nil()),
		)
}

// genCliOptionsImpl generates a CLIOptions struct
func (svc *Service) genCliOptionsImpl(f *g.File) {
	typeName := toCamel("%sCliOptions", svc.Service.GoName)

	// generate type definition
	f.Commentf("%s describes runtime configuration for %s cli", typeName, svc.Service.Desc.FullName())
	f.Type().Id(typeName).Struct(
		g.Id("after").Func().
			Params(g.Op("*").Qual(cliPkg, "Context")).
			Error(),
		g.Id("before").Func().
			Params(g.Op("*").Qual(cliPkg, "Context")).
			Error(),
		g.Id("clientForCommand").Func().
			Params(g.Op("*").Qual(cliPkg, "Context")).
			Params(g.Qual(clientPkg, "Client"), g.Error()),
		g.Id("worker").Func().
			Params(g.Op("*").Qual(cliPkg, "Context"), g.Qual(clientPkg, "Client")).
			Params(g.Qual(workerPkg, "Worker"), g.Error()),
	)

	// generate New<Service>CliOptions
	functionName := toCamel("New%s", typeName)
	f.Commentf("%s initializes a new %s value", functionName, typeName)
	f.Func().Id(functionName).Params().Op("*").Id(typeName).Block(
		g.Return(g.Op("&").Id(typeName).Values()),
	)

	// generate WithAfter method
	f.Commentf("WithAfter injects a custom After hook to be run after any command invocation")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithAfter").
		Params(
			g.Id("fn").Func().
				Params(g.Op("*").Qual(cliPkg, "Context")).
				Error(),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("after").Op("=").Id("fn"),
			g.Return(g.Id("opts")),
		)

	// generate WithBefore method
	f.Commentf("WithBefore injects a custom Before hook to be run prior to any command invocation")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithBefore").
		Params(
			g.Id("fn").Func().
				Params(g.Op("*").Qual(cliPkg, "Context")).
				Error(),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("before").Op("=").Id("fn"),
			g.Return(g.Id("opts")),
		)

	// generate WithClient method
	f.Comment("WithClient provides a Temporal client factory for use by commands")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithClient").
		Params(
			g.Id("fn").Func().
				Params(g.Op("*").Qual(cliPkg, "Context")).
				Params(g.Qual(clientPkg, "Client"), g.Error()),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("clientForCommand").Op("=").Id("fn"),
			g.Return(g.Id("opts")),
		)

	// generate WithWorker method
	f.Comment("WithWorker provides an method for initializing a worker")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithWorker").
		Params(
			g.Id("fn").Func().
				Params(g.Op("*").Qual(cliPkg, "Context"), g.Qual(clientPkg, "Client")).
				Params(g.Qual(workerPkg, "Worker"), g.Error()),
		).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("worker").Op("=").Id("fn"),
			g.Return(g.Id("opts")),
		)
}

// genCliPrintMessage serializes a proto message as json and pretty prints it
func genCliPrintMessage(b *g.Group, varName string) {
	b.List(g.Id("b"), g.Err()).Op(":=").Qual(protojsonPkg, "Marshal").Call(g.Id(varName))
	b.If(g.Err().Op("!=").Nil()).Block(
		g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error serializing response json: %w"), g.Err())),
	)
	b.Var().Id("out").Qual("bytes", "Buffer")
	b.If(
		g.Err().Op(":=").Qual("encoding/json", "Indent").Call(g.Op("&").Id("out"), g.Id("b"), g.Lit(""), g.Lit("  ")),
		g.Err().Op("!=").Nil(),
	).Block(
		g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error formatting json: %w"), g.Err())),
	)
	b.Qual("fmt", "Println").Call(g.Id("out").Dot("String").Call())
}

// genCliQueryCommand generates a <Query> command
func (svc *Service) genCliQueryCommand(cmds *g.Group, query string) {
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	desc := method.Comments.Leading.String()
	if desc != "" {
		desc = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
	} else {
		desc = fmt.Sprintf("executes a %s query and blocks until error or response received", query)
	}
	cmds.Comment(desc)
	cmds.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(strcase.ToKebab(query))
		cmd.Id("Usage").Op(":").Lit(desc)
		if svc.opts.GetFeatures().GetCli().GetCategories() {
			cmd.Id("Category").Op(":").Lit("QUERIES")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *g.Group) {
			// add workflow-id required flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("workflow-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("workflow id"))
				fields.Id("Required").Op(":").True()
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("w"))
			})
			// add run-id optional flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("run-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run id"))
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("r"))
			})
			if hasInput {
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *g.Group) {
			// initialize client
			fn.List(g.Id("c"), g.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(g.Id("cmd"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client for command: %w"), g.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal input
			if hasInput {
				unmarshaller := fmt.Sprintf("unmarshalCliFlagsTo%s", method.Input.GoIdent.GoName)
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			// execute query
			fn.
				If(
					g.ListFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("client").Dot(query).CallFunc(func(args *g.Group) {
						args.Id("cmd").Dot("Context")
						args.Id("cmd").Dot("String").Call(g.Lit("workflow-id"))
						args.Id("cmd").Dot("String").Call(g.Lit("run-id"))
						if hasInput {
							args.Id("req")
						}
					}),
					g.Err().Op("!=").Nil(),
				).
				Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error executing %q query: %w"), g.Id(fmt.Sprintf("%sQueryName", query)), g.Err())),
				).
				Else().
				BlockFunc(func(b *g.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					} else {
						fn.Qual("fmt", "Println").Call(g.Lit("success"))
					}
					b.Return(g.Nil())
				})
		})
	})
}

// genCliSignalCommand generates a <Signal> command
func (svc *Service) genCliSignalCommand(cmds *g.Group, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)
	desc := method.Comments.Leading.String()
	if desc != "" {
		desc = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
	} else {
		desc = fmt.Sprintf("executes a %s signal", signal)
	}
	cmds.Comment(desc)
	cmds.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(strcase.ToKebab(signal))
		cmd.Id("Usage").Op(":").Lit(desc)
		if svc.opts.GetFeatures().GetCli().GetCategories() {
			cmd.Id("Category").Op(":").Lit("SIGNALS")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *g.Group) {
			// add workflow-id required flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("workflow-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("workflow id"))
				fields.Id("Required").Op(":").True()
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("w"))
			})
			// add run-id optional flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("run-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run id"))
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("r"))
			})
			if hasInput {
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *g.Group) {
			// initialize client
			fn.List(g.Id("c"), g.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(g.Id("cmd"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client for command: %w"), g.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal input
			if hasInput {
				unmarshaller := fmt.Sprintf("unmarshalCliFlagsTo%s", method.Input.GoIdent.GoName)
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			fn.If(
				g.Err().Op(":=").Id("client").Dot(signal).CallFunc(func(args *g.Group) {
					args.Id("cmd").Dot("Context")
					args.Id("cmd").Dot("String").Call(g.Lit("workflow-id"))
					args.Id("cmd").Dot("String").Call(g.Lit("run-id"))
					if hasInput {
						args.Id("req")
					}
				}),
				g.Err().Op("!=").Nil(),
			).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error sending %q signal: %w"), g.Id(fmt.Sprintf("%sSignalName", signal)), g.Err())),
			)

			// print response
			fn.Qual("fmt", "Println").Call(g.Lit("success"))
			fn.Return(g.Nil())
		})
	})
}

// genCliUpdateCommand generates an <Update> command
func (svc *Service) genCliUpdateCommand(f *g.Group, update string) {
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	desc := method.Comments.Leading.String()
	if desc != "" {
		desc = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
	} else {
		desc = fmt.Sprintf("%s executes a(n) %s update", update, update)
	}
	f.Comment(desc)
	f.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(strcase.ToKebab(update))
		cmd.Id("Usage").Op(":").Lit(desc)
		if svc.opts.GetFeatures().GetCli().GetCategories() {
			cmd.Id("Category").Op(":").Lit("UPDATES")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *g.Group) {
			// add async flag
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run workflow in the background and print workflow and execution id"))
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("d"))
			})
			// add workflow-id required flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("workflow-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("workflow id"))
				fields.Id("Required").Op(":").True()
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("w"))
			})
			// add run-id optional flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("run-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run id"))
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("r"))
			})
			if hasInput {
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *g.Group) {
			// initialize client
			fn.List(g.Id("c"), g.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(g.Id("cmd"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client for command: %w"), g.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal input
			if hasInput {
				unmarshaller := fmt.Sprintf("unmarshalCliFlagsTo%s", method.Input.GoIdent.GoName)
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			// execute update operation
			fn.List(g.Id("handle"), g.Err()).Op(":=").Id("client").Dot(fmt.Sprintf("%sAsync", update)).CallFunc(func(args *g.Group) {
				args.Id("cmd").Dot("Context")
				args.Id("cmd").Dot("String").Call(g.Lit("workflow-id"))
				args.Id("cmd").Dot("String").Call(g.Lit("run-id"))
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error executing %s update: %w"), g.Id(fmt.Sprintf("%sUpdateName", update)), g.Err())),
			)

			// handle async invocation
			fn.If(g.Id("cmd").Dot("Bool").Call(g.Lit("detach"))).Block(
				g.Qual("fmt", "Println").Call(g.Lit("success")),
				g.Qual("fmt", "Printf").Call(g.Lit("workflow id: %s\n"), g.Id("handle").Dot("WorkflowID").Call()),
				g.Qual("fmt", "Printf").Call(g.Lit("run id: %s\n"), g.Id("handle").Dot("RunID").Call()),
				g.Qual("fmt", "Printf").Call(g.Lit("update id: %s\n"), g.Id("handle").Dot("UpdateID").Call()),
				g.Return(g.Nil()),
			)

			// handle synchronous invocation
			fn.
				If(
					g.ListFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("handle").Dot("Get").Call(g.Id("cmd").Dot("Context")),
					g.Err().Op("!=").Nil(),
				).
				Block(
					g.Return(g.Err()),
				).
				Else().
				BlockFunc(func(b *g.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					}
					b.Return(g.Nil())
				})
		})
	})
}

// genCliUnmarshalMessage generates an unmarshalCliFlagsTo<Message> function
func (svc *Service) genCliUnmarshalMessage(f *g.File, msg *protogen.Message) {
	name := msg.GoIdent.GoName
	fnName := fmt.Sprintf("unmarshalCliFlagsTo%s", name)
	f.Commentf("%s unmarshals a %s from command line flags", fnName, name)
	f.Func().Id(fnName).
		Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).
		Params(
			g.Op("*").Id(name),
			g.Error(),
		).BlockFunc(func(fn *g.Group) {
		fn.Var().Id("result").Id(name)
		fn.Var().Id("hasValues").Bool()
		for _, field := range msg.Fields {
			flag := strcase.ToKebab(field.GoName)
			fn.If(g.Id("cmd").Dot("IsSet").Call(g.Lit(flag))).BlockFunc(func(b *g.Group) {
				// indicate presence of value
				b.Id("hasValues").Op("=").True()

				switch {
				case field.Desc.IsList():
					fallthrough
				case field.Desc.IsMap():
					b.Var().Id("tmp").Id(name)
					b.If(
						g.Err().Op(":=").Qual(protojsonPkg, "Unmarshal").Call(
							g.Index().Byte().Call(g.Qual("fmt", "Sprintf").Call(g.Lit(fmt.Sprintf(`{"%s":%%s}`, field.Desc.JSONName())), g.Id("cmd").Dot("String").Call(g.Lit(flag)))),
							g.Op("&").Id("tmp"),
						),
						g.Err().Op("!=").Nil(),
					).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error unmarshalling %q map flag: %%w", flag)), g.Err())),
					)
					b.Id("result").Dot(field.GoName).Op("=").Id("tmp").Dot(field.GoName)
					return
				}

				switch field.Desc.Kind() {
				case protoreflect.BoolKind:
					b.Id("result").Dot(field.GoName).Op("=").Id("cmd").Dot("Bool").Call(g.Lit(flag))
				case protoreflect.BytesKind:
					b.List(g.Id("v"), g.Err()).Op(":=").Qual(base64Pkg, "StdEncoding").Dot("DecodeString").Call(g.Id("cmd").Dot("String").Call(g.Lit(flag)))
					b.If(g.Err().Op("!=").Nil()).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error base64-decoding %q flag: %%w", flag)), g.Err())),
					)
					b.Id("result").Dot(field.GoName).Op("=").Id("v")
				case protoreflect.DoubleKind:
					b.Id("result").Dot(field.GoName).Op("=").Id("cmd").Dot("Float64").Call(g.Lit(flag))
				case protoreflect.EnumKind:
					b.List(g.Id("v"), g.Id("ok")).Op(":=").Id(fmt.Sprintf("%s_value", field.Enum.GoIdent.GoName)).Index(g.Id("cmd").Dot("String").Call(g.Lit(flag)))
					b.If(g.Op("!").Id("ok")).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("unsupported enum value for %q flag: %%q", flag)), g.Id("cmd").Dot("String").Call(g.Lit(flag)))),
					)
					b.Id("result").Dot(field.GoName).Op("=").Id(field.Enum.GoIdent.GoName).Call(g.Id("v"))
				case protoreflect.Fixed32Kind, protoreflect.Uint32Kind:
					b.Id("result").Dot(field.GoName).Op("=").Uint32().Call(g.Id("cmd").Dot("Uint64").Call(g.Lit(flag)))
				case protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
					b.Id("result").Dot(field.GoName).Op("=").Id("cmd").Dot("Uint64").Call(g.Lit(flag))
				case protoreflect.FloatKind:
					b.Id("result").Dot(field.GoName).Op("=").Float32().Call(g.Id("cmd").Dot("Float64").Call(g.Lit(flag)))
				case protoreflect.Int32Kind, protoreflect.Sfixed32Kind, protoreflect.Sint32Kind:
					b.Id("result").Dot(field.GoName).Op("=").Int32().Call(g.Id("cmd").Dot("Int64").Call(g.Lit(flag)))
				case protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Sint64Kind:
					b.Id("result").Dot(field.GoName).Op("=").Id("cmd").Dot("Int64").Call(g.Lit(flag))
				case protoreflect.GroupKind:
				case protoreflect.MessageKind:
					if field.Message.GoIdent.GoImportPath != svc.File.GoImportPath {
						b.Var().Id("v").Qual(string(field.Message.GoIdent.GoImportPath), field.Message.GoIdent.GoName)
					} else {
						b.Var().Id("v").Id(field.Message.GoIdent.GoName)
					}
					b.If(g.Err().Op(":=").Qual(protojsonPkg, "Unmarshal").Call(g.Index().Byte().Call(g.Id("cmd").Dot("String").Call(g.Lit(flag))), g.Op("&").Id("v")), g.Err().Op("!=").Nil()).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error unmarhsalling %q flag: %%w", flag)), g.Err())),
					)
					b.Id("result").Dot(field.GoName).Op("=").Op("&").Id("v")
				case protoreflect.StringKind:
					b.Id("result").Dot(field.GoName).Op("=").Id("cmd").Dot("String").Call(g.Lit(flag))
				}
			})
		}
		fn.If(g.Op("!").Id("hasValues")).Block(
			g.Return(g.Nil(), g.Nil()),
		)
		fn.Return(g.Op("&").Id("result"), g.Nil())
	})
}

// genCliWorkerCommand generates a <Workflow> command
func (svc *Service) genCliWorkerCommand(f *g.Group) {
	f.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit("worker")
		cmd.Id("Usage").Op(":").Lit(fmt.Sprintf("runs a %s worker process", svc.Service.Desc.FullName()))
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Action").Op(":").Func().Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *g.Group) {
			// initialize client
			fn.List(g.Id("c"), g.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(g.Id("cmd"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client for command: %w"), g.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()

			// initialize worker
			fn.List(g.Id("w"), g.Err()).Op(":=").Id("opts").Dot("worker").Call(g.Id("cmd"), g.Id("c"))
			fn.If(g.Id("opts").Dot("worker").Op("!=").Nil()).Block(
				g.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing worker: %w"), g.Err())),
				),
			)

			// run worker
			fn.If(g.Err().Op(":=").Id("w").Dot("Start").Call(), g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error starting worker: %w"), g.Err())),
			)
			fn.Defer().Id("w").Dot("Stop").Call()
			fn.Op("<-").Id("cmd").Dot("Context").Dot("Done").Call()
			fn.Return(g.Nil())
		})
	})
}

// genCliWorkflowCommand generates a <Workflow> command
func (svc *Service) genCliWorkflowCommand(f *g.Group, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	desc := method.Comments.Leading.String()
	if desc != "" {
		desc = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
	} else {
		desc = fmt.Sprintf("%s executes a(n) %s workflow", workflow, workflow)
	}
	f.Comment(desc)
	f.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(strcase.ToKebab(workflow))
		cmd.Id("Usage").Op(":").Lit(desc)
		if svc.opts.GetFeatures().GetCli().GetCategories() {
			cmd.Id("Category").Op(":").Lit("WORKFLOWS")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *g.Group) {
			// add async flag
			// generate flag
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run workflow in the background and print workflow and execution id"))
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("d"))
			})
			if hasInput {
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *g.Group) {
			// initialize client
			fn.List(g.Id("c"), g.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(g.Id("cmd"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client for command: %w"), g.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal input
			if hasInput {
				unmarshaller := fmt.Sprintf("unmarshalCliFlagsTo%s", method.Input.GoIdent.GoName)
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			// execute operation
			fn.List(g.Id("run"), g.Err()).Op(":=").Id("client").Dot(fmt.Sprintf("%sAsync", workflow)).CallFunc(func(args *g.Group) {
				args.Id("cmd").Dot("Context")
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error starting %s workflow: %w"), g.Id(fmt.Sprintf("%sWorkflowName", workflow)), g.Err())),
			)

			// handle async invocation
			fn.If(g.Id("cmd").Dot("Bool").Call(g.Lit("detach"))).Block(
				g.Qual("fmt", "Println").Call(g.Lit("success")),
				g.Qual("fmt", "Printf").Call(g.Lit("workflow id: %s\n"), g.Id("run").Dot("ID").Call()),
				g.Qual("fmt", "Printf").Call(g.Lit("run id: %s\n"), g.Id("run").Dot("RunID").Call()),
				g.Return(g.Nil()),
			)

			// handle synchronous invocation
			fn.
				If(
					g.ListFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("run").Dot("Get").Call(g.Id("cmd").Dot("Context")),
					g.Err().Op("!=").Nil(),
				).
				Block(
					g.Return(g.Err()),
				).
				Else().
				BlockFunc(func(b *g.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					}
					b.Return(g.Nil())
				})
		})
	})
}

// genCliWorkflowWithSignalCommand generates a <Workflow>-with-<Signal> command
func (svc *Service) genCliWorkflowWithSignalCommand(cmds *g.Group, workflow, signal string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := svc.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	cmdName := strcase.ToKebab(strings.Join([]string{workflow, "with", signal}, "-"))
	desc := fmt.Sprintf("sends a %s signal to a %s worklow, starting it if necessary", signal, workflow)

	cmds.Comment(desc)
	cmds.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(cmdName)
		cmd.Id("Usage").Op(":").Lit(desc)
		if svc.opts.GetFeatures().GetCli().GetCategories() {
			cmd.Id("Category").Op(":").Lit("WORKFLOWS")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *g.Group) {
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run workflow in the background and print workflow and execution id"))
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("d"))
			})
			if hasInput {
				var category string
				if hasSignalInput {
					category = "INPUT"
				}
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, category)
				}
			}
			if hasSignalInput {
				var category string
				if hasSignalInput {
					category = "SIGNAL"
				}
				// add request flags
				for _, field := range handler.Input.Fields {
					svc.genCliFlagForField(flags, field, category)
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *g.Group) {
			// initialize client
			fn.List(g.Id("c"), g.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(g.Id("cmd"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client for command: %w"), g.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal request
			if hasInput {
				unmarshaller := fmt.Sprintf("unmarshalCliFlagsTo%s", method.Input.GoIdent.GoName)
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			// unmarshal signal
			if hasSignalInput {
				unmarshaller := fmt.Sprintf("unmarshalCliFlagsTo%s", handler.Input.GoIdent.GoName)
				fn.List(g.Id("signal"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling signal: %w"), g.Err())),
				)
			}

			// execute operation
			fn.List(g.Id("run"), g.Err()).Op(":=").Id("client").Dot(fmt.Sprintf("%sWith%sAsync", workflow, signal)).CallFunc(func(args *g.Group) {
				args.Id("cmd").Dot("Context")
				if hasInput {
					args.Id("req")
				}
				if hasSignalInput {
					args.Id("signal")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error starting %s workflow with %s signal: %w"), g.Id(fmt.Sprintf("%sWorkflowName", workflow)), g.Id(fmt.Sprintf("%sSignalName", signal)), g.Err())),
			)

			// handle async invocation
			fn.If(g.Id("cmd").Dot("Bool").Call(g.Lit("detach"))).Block(
				g.Qual("fmt", "Println").Call(g.Lit("success")),
				g.Qual("fmt", "Printf").Call(g.Lit("workflow id: %s\n"), g.Id("run").Dot("ID").Call()),
				g.Qual("fmt", "Printf").Call(g.Lit("run id: %s\n"), g.Id("run").Dot("RunID").Call()),
				g.Return(g.Nil()),
			)

			// handle synchronous invocation
			fn.
				If(
					g.ListFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("run").Dot("Get").Call(g.Id("cmd").Dot("Context")),
					g.Err().Op("!=").Nil(),
				).
				Block(
					g.Return(g.Err()),
				).
				Else().
				BlockFunc(func(b *g.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					}
					b.Return(g.Nil())
				})
		})
	})
}
