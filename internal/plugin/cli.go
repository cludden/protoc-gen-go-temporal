package plugin

import (
	"fmt"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// define cli-specific import constants
const (
	base64Pkg    = "encoding/base64"
	cliPkg       = "github.com/urfave/cli/v2"
	homedirPkg   = "github.com/mitchellh/go-homedir"
	protojsonPkg = "google.golang.org/protobuf/encoding/protojson"
)

var (
	multiLineArgs = g.Options{
		Close:     ")",
		Multi:     true,
		Open:      "(",
		Separator: ",",
	}

	multiLineValues = g.Options{
		Close:     "}",
		Multi:     true,
		Open:      "{",
		Separator: ",",
	}
)

type Cli struct{}

// renderCLI generates cli resources
func (svc *Manifest) renderCLI(f *g.File) {
	opts := proto.GetExtension(svc.Service.Desc.Options(), temporalv1.E_Cli).(*temporalv1.CLIOptions)
	if opts != nil && opts.GetIgnore() {
		return
	}

	svc.genCliOptionsImpl(f)
	svc.genCliNew(f)
	svc.genCliNewCommand(f)
	svc.genCliNewCommands(f)

	// cache unmarshal functions to void duplicate declarations
	unmarshallers := map[string]struct{}{}

	// generate query request unmarshallers
	for _, query := range svc.queriesOrdered {
		if svc.methods[query].Desc.Parent() != svc.Service.Desc {
			continue
		}
		if opts, ok := svc.commands[query]; ok && opts.GetIgnore() {
			continue
		}
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
		if svc.methods[signal].Desc.Parent() != svc.Service.Desc {
			continue
		}
		if opts, ok := svc.commands[signal]; ok && opts.GetIgnore() {
			continue
		}
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
		if svc.methods[update].Desc.Parent() != svc.Service.Desc {
			continue
		}
		if opts, ok := svc.commands[update]; ok && opts.GetIgnore() {
			continue
		}
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
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		if opts, ok := svc.commands[workflow]; ok && opts.GetIgnore() {
			continue
		}
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

// genCliNew generates a New<Service>Cli constructor function
func (svc *Manifest) genCliNew(f *g.File) {
	functionName := svc.toCamel("New%sCli", svc.Service.GoName)
	optionsName := svc.toCamel("%sCliOptions", svc.Service.GoName)

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
			g.List(g.Id("commands"), g.Err()).Op(":=").Id(svc.toLowerCamel("new%sCommands", svc.Service.GoName)).Call(g.Id("options").Op("...")),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing subcommands: %w"), g.Err())),
			),
			g.Return(
				g.Op("&").Qual(cliPkg, "App").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit(svc.caser.ToKebab(svc.Service.GoName))
					fields.Id("Commands").Op(":").Id("commands")
					fields.Id("DisableSliceFlagSeparator").Op(":").True()
				}),
				g.Nil(),
			),
		)
}

// genCliNewCommand generates a New<Service>CliCommand constructor function
func (svc *Manifest) genCliNewCommand(f *g.File) {
	functionName := svc.toCamel("New%sCliCommand", svc.Service.GoName)
	optionsName := svc.toCamel("%sCliOptions", svc.Service.GoName)

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
			g.List(g.Id("subcommands"), g.Err()).Op(":=").Id(svc.toLowerCamel("new%sCommands", svc.Service.GoName)).Call(g.Id("options").Op("...")),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing subcommands: %w"), g.Err())),
			),
			g.Return(
				g.Op("&").Qual(cliPkg, "Command").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit(svc.caser.ToKebab(svc.Service.GoName))
					fields.Id("Subcommands").Op(":").Id("subcommands")
				}),
				g.Nil(),
			),
		)
}

// genCliNewCommands generates a new<Service>Commands constructor function
func (svc *Manifest) genCliNewCommands(f *g.File) {
	functionName := svc.toLowerCamel("new%sCommands", svc.Service.GoName)
	optionsName := svc.toCamel("%sCliOptions", svc.Service.GoName)

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
					if svc.methods[query].Desc.Parent() != svc.Service.Desc {
						continue
					}
					if opts, ok := svc.commands[query]; ok && opts.GetIgnore() {
						continue
					}
					if svc.queries[query].GetCli().GetIgnore() {
						continue
					}
					svc.genCliQueryCommand(cmds, query)
				}

				// generate client signal methods
				for _, signal := range svc.signalsOrdered {
					if svc.methods[signal].Desc.Parent() != svc.Service.Desc {
						continue
					}
					if opts, ok := svc.commands[signal]; ok && opts.GetIgnore() {
						continue
					}
					if svc.signals[signal].GetCli().GetIgnore() {
						continue
					}
					svc.genCliSignalCommand(cmds, signal)
				}

				// generate client update methods
				for _, update := range svc.updatesOrdered {
					if svc.methods[update].Desc.Parent() != svc.Service.Desc {
						continue
					}
					if opts, ok := svc.commands[update]; ok && opts.GetIgnore() {
						continue
					}
					if svc.updates[update].GetCli().GetIgnore() {
						continue
					}
					svc.genCliUpdateCommand(cmds, update)
				}

				// generate client workflow methods
				for _, workflow := range svc.workflowsOrdered {
					if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
						continue
					}
					if opts, ok := svc.commands[workflow]; ok && opts.GetIgnore() {
						continue
					}
					if svc.workflows[workflow].GetCli().GetIgnore() {
						continue
					}
					svc.genCliWorkflowCommand(cmds, workflow)
					for _, signal := range svc.workflows[workflow].GetSignal() {
						if !signal.GetStart() {
							continue
						}
						if signal.GetCli().GetIgnore() {
							continue
						}
						svc.genCliWorkflowWithSignalCommand(cmds, workflow, getFullyQualifiedRef(workflow, signal.GetRef()), signal)
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
func (svc *Manifest) genCliOptionsImpl(f *g.File) {
	typeName := svc.toCamel("%sCliOptions", svc.Service.GoName)

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
	functionName := svc.toCamel("New%s", typeName)
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
func (svc *Manifest) genCliQueryCommand(cmds *g.Group, query protoreflect.FullName) {
	method := svc.methods[query]
	opts := svc.queries[query]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	name := svc.caser.ToKebab(svc.methods[query].GoName)
	if v := opts.GetCli().GetName(); v != "" {
		name = v
	}

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = method.Comments.Leading.String()
	}

	if usage != "" {
		usage = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(usage, "//"), "\n//", ""))
	} else {
		usage = fmt.Sprintf("executes a %s query and blocks until error or response received", svc.fqnForQuery(query))
	}

	cmds.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *g.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if svc.cfg.CliCategories {
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
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("f"))
				})
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT", "")
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
			fn.Id("client").Op(":=").Id(svc.toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal input
			if hasInput {
				inputName := svc.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", svc.toCamel("%s", inputName))
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
					}).Op(":=").Id("client").Dot(svc.methods[query].GoName).CallFunc(func(args *g.Group) {
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
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error executing %q query: %w"), g.Id(svc.toCamel("%sQueryName", query)), g.Err())),
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
func (svc *Manifest) genCliSignalCommand(cmds *g.Group, signal protoreflect.FullName) {
	method := svc.methods[signal]
	opts := svc.signals[signal]
	hasInput := !isEmpty(method.Input)

	name := svc.caser.ToKebab(svc.methods[signal].GoName)
	if v := opts.GetCli().GetName(); v != "" {
		name = v
	}

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = method.Comments.Leading.String()
	}

	if usage != "" {
		usage = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(usage, "//"), "\n//", ""))
	} else {
		usage = fmt.Sprintf("executes a %s signal", svc.fqnForSignal(signal))
	}

	cmds.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *g.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if svc.cfg.CliCategories {
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
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("f"))
				})
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT", "")
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
			fn.Id("client").Op(":=").Id(svc.toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal input
			if hasInput {
				inputName := svc.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", svc.toCamel("%s", inputName))
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			fn.If(
				g.Err().Op(":=").Id("client").Dot(svc.methods[signal].GoName).CallFunc(func(args *g.Group) {
					args.Id("cmd").Dot("Context")
					args.Id("cmd").Dot("String").Call(g.Lit("workflow-id"))
					args.Id("cmd").Dot("String").Call(g.Lit("run-id"))
					if hasInput {
						args.Id("req")
					}
				}),
				g.Err().Op("!=").Nil(),
			).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error sending %q signal: %w"), g.Id(svc.toCamel("%sSignalName", signal)), g.Err())),
			)

			// print response
			fn.Qual("fmt", "Println").Call(g.Lit("success"))
			fn.Return(g.Nil())
		})
	})
}

// genCliUpdateCommand generates an <Update> command
func (svc *Manifest) genCliUpdateCommand(f *g.Group, update protoreflect.FullName) {
	method := svc.methods[update]
	opts := svc.updates[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	name := svc.caser.ToKebab(svc.methods[update].GoName)
	if v := opts.GetCli().GetName(); v != "" {
		name = v
	}

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = method.Comments.Leading.String()
	}

	if usage != "" {
		usage = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(usage, "//"), "\n//", ""))
	} else {
		usage = fmt.Sprintf("executes a(n) %s update", svc.fqnForUpdate(update))
	}

	f.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *g.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if svc.cfg.CliCategories {
			cmd.Id("Category").Op(":").Lit("UPDATES")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *g.Group) {
			// add async flag
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run workflow update in the background and print workflow, execution, and udpate id"))
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
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("f"))
				})
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT", "")
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
			fn.Id("client").Op(":=").Id(svc.toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal input
			if hasInput {
				inputName := svc.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", svc.toCamel("%s", inputName))
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			// execute update operation
			fn.List(g.Id("handle"), g.Err()).Op(":=").Id("client").Dot(svc.toCamel("%sAsync", update)).CallFunc(func(args *g.Group) {
				args.Id("cmd").Dot("Context")
				args.Id("cmd").Dot("String").Call(g.Lit("workflow-id"))
				args.Id("cmd").Dot("String").Call(g.Lit("run-id"))
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error executing %s update: %w"), g.Id(svc.toCamel("%sUpdateName", update)), g.Err())),
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

// genCliWorkerCommand generates a <Workflow> command
func (svc *Manifest) genCliWorkerCommand(f *g.Group) {
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
func (svc *Manifest) genCliWorkflowCommand(f *g.Group, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	name := svc.caser.ToKebab(svc.methods[workflow].GoName)
	if v := opts.GetCli().GetName(); v != "" {
		name = v
	}

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = method.Comments.Leading.String()
	}

	if usage != "" {
		usage = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(usage, "//"), "\n//", ""))
	} else {
		usage = fmt.Sprintf("executes a(n) %s workflow", svc.fqnForWorkflow(workflow))
	}

	f.CustomFunc(multiLineValues, func(cmd *g.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *g.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if svc.cfg.CliCategories {
			cmd.Id("Category").Op(":").Lit("WORKFLOWS")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *g.Group) {
			// add async flag
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit("run workflow in the background and print workflow and execution id")
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("d"))
			})
			// add task-queue flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
				fields.Id("Name").Op(":").Lit("task-queue")
				fields.Id("Usage").Op(":").Lit("task queue name")
				fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("t"))
				fields.Id("EnvVars").Op(":").Index().String().Values(g.Lit("TEMPORAL_TASK_QUEUE_NAME"), g.Lit("TEMPORAL_TASK_QUEUE"), g.Lit("TASK_QUEUE_NAME"), g.Lit("TASK_QUEUE"))
				tq := svc.opts.GetTaskQueue()
				if tq == "" {
					fields.Id("Required").Op(":").True()
				} else {
					fields.Id("Value").Op(":").Lit(tq)
				}
			})
			if hasInput {
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("f"))
				})
				// add request flags
				for _, field := range method.Input.Fields {
					svc.genCliFlagForField(flags, field, "INPUT", "")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(g.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *g.Group) {
			// initialize client
			fn.List(g.Id("tc"), g.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(g.Id("cmd"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client for command: %w"), g.Err())),
			)
			fn.Defer().Id("tc").Dot("Close").Call()
			fn.Id("c").Op(":=").Id(svc.toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("tc"))

			// unmarshal input
			if hasInput {
				inputName := svc.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", svc.toCamel("%s", inputName))
				fn.List(g.Id("req"), g.Err()).Op(":=").Id(unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			// build workflow options
			fn.Id("opts").Op(":=").Qual(clientPkg, "StartWorkflowOptions").Values()
			fn.If(g.Id("tq").Op(":=").Id("cmd").Dot("String").Call(g.Lit("task-queue")), g.Id("tq").Op("!=").Lit("")).Block(
				g.Id("opts").Dot("TaskQueue").Op("=").Id("tq"),
			)

			// execute operation
			fn.List(g.Id("run"), g.Err()).Op(":=").Id("c").Dot(svc.toCamel("%sAsync", workflow)).CallFunc(func(args *g.Group) {
				args.Id("cmd").Dot("Context")
				if hasInput {
					args.Id("req")
				}
				args.Id(svc.toCamel("New%sOptions", workflow)).Call().Dot("WithStartWorkflowOptions").Call(g.Id("opts"))
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error starting %s workflow: %w"), g.Id(svc.toCamel("%sWorkflowName", workflow)), g.Err())),
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
func (svc *Manifest) genCliWorkflowWithSignalCommand(cmds *g.Group, w, signal protoreflect.FullName, opts *temporalv1.WorkflowOptions_Signal) {
	method := svc.methods[w]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := svc.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	name := opts.GetCli().GetName()
	if name == "" {
		workflowName := svc.workflows[w].GetCli().GetName()
		if workflowName == "" {
			workflowName = svc.methods[w].GoName
		}
		signalName := svc.signals[signal].GetCli().GetName()
		if signalName == "" {
			signalName = svc.methods[signal].GoName
		}
		name = svc.caser.ToKebab(strings.Join([]string{workflowName, "with", signalName}, "-"))
	}

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = fmt.Sprintf("sends a %s signal to a %s workflow, starting it if necessary", signal, w)
	}

	cmds.Comment(usage)
	cmds.CustomFunc(multiLineValues, func(cmd *g.Group) {
		fields := map[string]struct{}{}
		collisions := map[string]struct{}{}

		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *g.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if svc.cfg.CliCategories {
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
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(g.Lit("f"))
				})
				var category string
				if hasSignalInput {
					category = "INPUT"
				}
				// add request flags
				for _, field := range method.Input.Fields {
					fields[field.GoName] = struct{}{}
					svc.genCliFlagForField(flags, field, category, "")
				}
			}
			if hasSignalInput {
				var category string
				if hasSignalInput {
					category = "SIGNAL"
				}
				// add request flags
				for _, field := range handler.Input.Fields {
					var prefix string
					if _, ok := fields[field.GoName]; ok {
						prefix = handler.GoName
						collisions[svc.flagName(field, "")] = struct{}{}
					}
					svc.genCliFlagForField(flags, field, category, prefix)
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
			fn.Id("client").Op(":=").Id(svc.toCamel("New%sClient", svc.Service.GoName)).Call(g.Id("c"))

			// unmarshal request
			if hasInput {
				inputName := svc.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", svc.toCamel("%s", inputName))
				fn.List(g.Id("req"), g.Err()).Op(":=").Qual(svc.goImportPathForMethod(w), unmarshaller).Call(g.Id("cmd"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling request: %w"), g.Err())),
				)
			}

			// unmarshal signal
			if hasSignalInput {
				inputName := svc.getMessageName(handler.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", svc.toCamel("%s", inputName))
				fn.List(g.Id("signal"), g.Err()).Op(":=").Qual(svc.goImportPathForMethod(signal), unmarshaller).CallFunc(func(b *g.Group) {
					b.Id("cmd")
					if len(collisions) > 0 {
						b.Qual(helpersPkg, "UnmarshalCliFlagsOptions").CustomFunc(multiLineValues, func(b *g.Group) {
							b.Id("Prefix").Op(":").Lit(svc.caser.ToKebab(handler.GoName))
							b.Id("PrefixFlags").Op(":").Map(g.String()).Struct().CustomFunc(multiLineValues, func(b *g.Group) {
								for _, field := range workflow.DeterministicKeys(collisions) {
									b.Lit(field).Op(":").Values()
								}
							})
						})
					}
				})
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error unmarshalling signal: %w"), g.Err())),
				)
			}

			// execute operation
			fn.List(g.Id("run"), g.Err()).Op(":=").Id("client").Dot(svc.toCamel("%sWith%sAsync", w, signal)).CallFunc(func(args *g.Group) {
				args.Id("cmd").Dot("Context")
				if hasInput {
					args.Id("req")
				}
				if hasSignalInput {
					args.Id("signal")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Qual("fmt", "Errorf").Call(g.Lit("error starting %s workflow with %s signal: %w"), g.Id(svc.toCamel("%sWorkflowName", w)), g.Qual(svc.goImportPathForMethod(signal), svc.toCamel("%sSignalName", signal)), g.Err())),
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
