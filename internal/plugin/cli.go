package plugin

import (
	"fmt"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	j "github.com/dave/jennifer/jen"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/compiler/protogen"
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

func (n *names) unmarshalCliFlagsTo(message *protogen.Message) string {
	return n.toCamel("UnmarshalCliFlagsTo%s", n.getMessageName(message))
}

type Cli struct{}

// renderCLI generates cli resources
func (m *Manifest) renderCLI(f *j.File) {
	opts := proto.GetExtension(m.Service.Desc.Options(), temporalv1.E_Cli).(*temporalv1.CLIOptions)
	if opts != nil && opts.GetIgnore() {
		return
	}

	m.genCliOptionsImpl(f)
	m.genCliNew(f)
	m.genCliNewCommand(f)
	m.genCliNewCommands(f)

	// cache unmarshal functions to void duplicate declarations
	unmarshallers := map[string]struct{}{}

	// generate query request unmarshallers
	for _, query := range m.queriesOrdered {
		if m.methods[query].Desc.Parent() != m.Service.Desc {
			continue
		}
		if opts, ok := m.commands[query]; ok && opts.GetIgnore() {
			continue
		}
		if isEmpty(m.methods[query].Input) {
			continue
		}
		if _, ok := unmarshallers[m.methods[query].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[m.methods[query].Input.GoIdent.GoName] = struct{}{}
		m.genCliUnmarshalMessage(f, m.methods[query].Input)
	}

	// generate signal request unmarshallers
	for _, signal := range m.signalsOrdered {
		if m.methods[signal].Desc.Parent() != m.Service.Desc {
			continue
		}
		if opts, ok := m.commands[signal]; ok && opts.GetIgnore() {
			continue
		}
		if isEmpty(m.methods[signal].Input) {
			continue
		}
		if _, ok := unmarshallers[m.methods[signal].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[m.methods[signal].Input.GoIdent.GoName] = struct{}{}
		m.genCliUnmarshalMessage(f, m.methods[signal].Input)
	}

	// generate update request unmarshallers
	for _, update := range m.updatesOrdered {
		if m.methods[update].Desc.Parent() != m.Service.Desc {
			continue
		}
		if opts, ok := m.commands[update]; ok && opts.GetIgnore() {
			continue
		}
		if isEmpty(m.methods[update].Input) {
			continue
		}
		if _, ok := unmarshallers[m.methods[update].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[m.methods[update].Input.GoIdent.GoName] = struct{}{}
		m.genCliUnmarshalMessage(f, m.methods[update].Input)
	}

	// generate workflow request unmarshallers
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		if opts, ok := m.commands[workflow]; ok && opts.GetIgnore() {
			continue
		}
		if isEmpty(m.methods[workflow].Input) {
			continue
		}
		if _, ok := unmarshallers[m.methods[workflow].Input.GoIdent.GoName]; ok {
			continue
		}
		unmarshallers[m.methods[workflow].Input.GoIdent.GoName] = struct{}{}
		m.genCliUnmarshalMessage(f, m.methods[workflow].Input)
	}
}

// genCliNew generates a New<Service>Cli constructor function
func (m *Manifest) genCliNew(f *j.File) {
	functionName := m.toCamel("New%sCli", m.Service.GoName)
	optionsName := m.toCamel("%sCliOptions", m.Service.GoName)

	f.Commentf("%s initializes a cli for a(n) %s service", functionName, m.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			j.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			j.Op("*").Qual(cliPkg, "App"),
			j.Error(),
		).
		Block(
			j.List(j.Id("commands"), j.Err()).Op(":=").Id(m.toLowerCamel("new%sCommands", m.Service.GoName)).Call(j.Id("options").Op("...")),
			j.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing subcommands: %w"), j.Err())),
			),
			j.Return(
				j.Op("&").Qual(cliPkg, "App").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit(m.caser.ToKebab(m.Service.GoName))
					fields.Id("Commands").Op(":").Id("commands")
					fields.Id("DisableSliceFlagSeparator").Op(":").True()
				}),
				j.Nil(),
			),
		)
}

// genCliNewCommand generates a New<Service>CliCommand constructor function
func (m *Manifest) genCliNewCommand(f *j.File) {
	functionName := m.toCamel("New%sCliCommand", m.Service.GoName)
	optionsName := m.toCamel("%sCliOptions", m.Service.GoName)

	f.Commentf("%s initializes a cli command for a %s service with subcommands for each query, signal, update, and workflow", functionName, m.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			j.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			j.Op("*").Qual(cliPkg, "Command"),
			j.Error(),
		).
		Block(
			j.List(j.Id("subcommands"), j.Err()).Op(":=").Id(m.toLowerCamel("new%sCommands", m.Service.GoName)).Call(j.Id("options").Op("...")),
			j.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing subcommands: %w"), j.Err())),
			),
			j.Return(
				j.Op("&").Qual(cliPkg, "Command").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit(m.caser.ToKebab(m.Service.GoName))
					fields.Id("Subcommands").Op(":").Id("subcommands")
				}),
				j.Nil(),
			),
		)
}

// genCliNewCommands generates a new<Service>Commands constructor function
func (m *Manifest) genCliNewCommands(f *j.File) {
	functionName := m.toLowerCamel("new%sCommands", m.Service.GoName)
	optionsName := m.toCamel("%sCliOptions", m.Service.GoName)

	f.Commentf("%s initializes (sub)commands for a %s cli or command", functionName, m.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			j.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			j.Index().Op("*").Qual(cliPkg, "Command"),
			j.Error(),
		).
		Block(
			// initialize options
			j.Id("opts").Op(":=").Op("&").Id(optionsName).Values(),
			j.If(j.Len(j.Id("options")).Op(">").Lit(0)).Block(
				j.Id("opts").Op("=").Id("options").Index(j.Lit(0)),
			),

			// set default client for command
			j.If(j.Id("opts").Dot("clientForCommand").Op("==").Nil()).Block(
				j.Id("opts").Dot("clientForCommand").Op("=").Func().
					Params(j.Op("*").Qual(cliPkg, "Context")).
					Params(j.Qual(clientPkg, "Client"), j.Error()).
					Block(
						j.Return(j.Qual(clientPkg, "Dial").Call(j.Qual(clientPkg, "Options").Values())),
					),
			),

			// initialize commands
			j.Id("commands").Op(":=").Index().Op("*").Qual(cliPkg, "Command").CustomFunc(j.Options{
				Close:     "}",
				Multi:     true,
				Open:      "{",
				Separator: ",",
			}, func(g *j.Group) {
				// generate client query methods
				for _, query := range m.queriesOrdered {
					if m.methods[query].Desc.Parent() != m.Service.Desc {
						continue
					}
					if opts, ok := m.commands[query]; ok && opts.GetIgnore() {
						continue
					}
					if m.queries[query].GetCli().GetIgnore() {
						continue
					}
					m.genCliQueryCommand(g, query)
				}

				// generate client signal methods
				for _, signal := range m.signalsOrdered {
					if m.methods[signal].Desc.Parent() != m.Service.Desc {
						continue
					}
					if opts, ok := m.commands[signal]; ok && opts.GetIgnore() {
						continue
					}
					if m.signals[signal].GetCli().GetIgnore() {
						continue
					}
					m.genCliSignalCommand(g, signal)
				}

				// generate client update methods
				for _, update := range m.updatesOrdered {
					if m.methods[update].Desc.Parent() != m.Service.Desc {
						continue
					}
					if opts, ok := m.commands[update]; ok && opts.GetIgnore() {
						continue
					}
					if m.updates[update].GetCli().GetIgnore() {
						continue
					}
					m.genCliUpdateCommand(g, update)
				}

				// generate client workflow methods
				for _, workflow := range m.workflowsOrdered {
					if m.methods[workflow].Desc.Parent() != m.Service.Desc {
						continue
					}
					if opts, ok := m.commands[workflow]; ok && opts.GetIgnore() {
						continue
					}
					if m.workflows[workflow].GetCli().GetIgnore() {
						continue
					}
					m.genCliWorkflowCommand(g, workflow)

					for _, signal := range m.workflows[workflow].GetSignal() {
						if !signal.GetStart() {
							continue
						}
						if signal.GetCli().GetIgnore() {
							continue
						}
						m.genCliWorkflowWithSignalCommand(g, workflow, getFullyQualifiedRef(workflow, signal.GetRef()), signal)
					}

					for _, update := range m.workflows[workflow].GetUpdate() {
						if !update.GetStart() {
							continue
						}
						if update.GetCli().GetIgnore() {
							continue
						}
						m.genCliWorkflowWithUpdateCommand(g, workflow, getFullyQualifiedRef(workflow, update.GetRef()), update)
					}
				}
			}),

			// append worker command if initializer provided
			j.If(j.Id("opts").Dot("worker").Op("!=").Nil()).Block(
				j.Id("commands").Op("=").Append(j.Id("commands"), j.Index().Op("*").Qual(cliPkg, "Command").CustomFunc(multiLineValues, func(cmds *j.Group) {
					m.genCliWorkerCommand(cmds)
				}).Op("...")),
			),

			j.Qual("sort", "Slice").Call(
				j.Id("commands"),
				j.Func().Params(j.Id("i"), j.Id("j").Int()).Bool().Block(
					j.Return(j.Id("commands").Index(j.Id("i")).Dot("Name").Op("<").Id("commands").Index(j.Id("j")).Dot("Name")),
				),
			),
			j.Return(j.Id("commands"), j.Nil()),
		)
}

// genCliOptionsImpl generates a CLIOptions struct
func (m *Manifest) genCliOptionsImpl(f *j.File) {
	typeName := m.toCamel("%sCliOptions", m.Service.GoName)

	// generate type definition
	f.Commentf("%s describes runtime configuration for %s cli", typeName, m.Service.Desc.FullName())
	f.Type().Id(typeName).Struct(
		j.Id("after").Func().
			Params(j.Op("*").Qual(cliPkg, "Context")).
			Error(),
		j.Id("before").Func().
			Params(j.Op("*").Qual(cliPkg, "Context")).
			Error(),
		j.Id("clientForCommand").Func().
			Params(j.Op("*").Qual(cliPkg, "Context")).
			Params(j.Qual(clientPkg, "Client"), j.Error()),
		j.Id("worker").Func().
			Params(j.Op("*").Qual(cliPkg, "Context"), j.Qual(clientPkg, "Client")).
			Params(j.Qual(workerPkg, "Worker"), j.Error()),
	)

	// generate New<Service>CliOptions
	functionName := m.toCamel("New%s", typeName)
	f.Commentf("%s initializes a new %s value", functionName, typeName)
	f.Func().Id(functionName).Params().Op("*").Id(typeName).Block(
		j.Return(j.Op("&").Id(typeName).Values()),
	)

	// generate WithAfter method
	f.Commentf("WithAfter injects a custom After hook to be run after any command invocation")
	f.Func().
		Params(j.Id("opts").Op("*").Id(typeName)).
		Id("WithAfter").
		Params(
			j.Id("fn").Func().
				Params(j.Op("*").Qual(cliPkg, "Context")).
				Error(),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("after").Op("=").Id("fn"),
			j.Return(j.Id("opts")),
		)

	// generate WithBefore method
	f.Commentf("WithBefore injects a custom Before hook to be run prior to any command invocation")
	f.Func().
		Params(j.Id("opts").Op("*").Id(typeName)).
		Id("WithBefore").
		Params(
			j.Id("fn").Func().
				Params(j.Op("*").Qual(cliPkg, "Context")).
				Error(),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("before").Op("=").Id("fn"),
			j.Return(j.Id("opts")),
		)

	// generate WithClient method
	f.Comment("WithClient provides a Temporal client factory for use by commands")
	f.Func().
		Params(j.Id("opts").Op("*").Id(typeName)).
		Id("WithClient").
		Params(
			j.Id("fn").Func().
				Params(j.Op("*").Qual(cliPkg, "Context")).
				Params(j.Qual(clientPkg, "Client"), j.Error()),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("clientForCommand").Op("=").Id("fn"),
			j.Return(j.Id("opts")),
		)

	// generate WithWorker method
	f.Comment("WithWorker provides an method for initializing a worker")
	f.Func().
		Params(j.Id("opts").Op("*").Id(typeName)).
		Id("WithWorker").
		Params(
			j.Id("fn").Func().
				Params(j.Op("*").Qual(cliPkg, "Context"), j.Qual(clientPkg, "Client")).
				Params(j.Qual(workerPkg, "Worker"), j.Error()),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("worker").Op("=").Id("fn"),
			j.Return(j.Id("opts")),
		)
}

// genCliPrintMessage serializes a proto message as json and pretty prints it
func genCliPrintMessage(b *j.Group, varName string) {
	b.List(j.Id("b"), j.Err()).Op(":=").Qual(protojsonPkg, "Marshal").Call(j.Id(varName))
	b.If(j.Err().Op("!=").Nil()).Block(
		j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error serializing response json: %w"), j.Err())),
	)
	b.Var().Id("out").Qual("bytes", "Buffer")
	b.If(
		j.Err().Op(":=").Qual("encoding/json", "Indent").Call(j.Op("&").Id("out"), j.Id("b"), j.Lit(""), j.Lit("  ")),
		j.Err().Op("!=").Nil(),
	).Block(
		j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error formatting json: %w"), j.Err())),
	)
	b.Qual("fmt", "Println").Call(j.Id("out").Dot("String").Call())
}

// genCliQueryCommand generates a <Query> command
func (m *Manifest) genCliQueryCommand(cmds *j.Group, query protoreflect.FullName) {
	method := m.methods[query]
	opts := m.queries[query]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	name := m.caser.ToKebab(m.methods[query].GoName)
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
		usage = fmt.Sprintf("executes a %s query and blocks until error or response received", m.fqnForQuery(query))
	}

	cmds.CustomFunc(multiLineValues, func(cmd *j.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *j.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if m.cfg.CliCategories {
			cmd.Id("Category").Op(":").Lit("QUERIES")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *j.Group) {
			// add workflow-id required flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("workflow-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("workflow id"))
				fields.Id("Required").Op(":").True()
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("w"))
			})
			// add run-id optional flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("run-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run id"))
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("r"))
			})
			if hasInput {
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("f"))
					fields.Id("Category").Op(":").Lit("INPUT")
				})
				// add request flags
				for _, field := range method.Input.Fields {
					m.genCliFlagForField(flags, field, "INPUT", "")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(j.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *j.Group) {
			// initialize client
			fn.List(j.Id("c"), j.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(j.Id("cmd"))
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client for command: %w"), j.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(m.toCamel("New%sClient", m.Service.GoName)).Call(j.Id("c"))

			// unmarshal input
			if hasInput {
				inputName := m.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", m.toCamel("%s", inputName))
				fn.List(j.Id("req"), j.Err()).Op(":=").Id(unmarshaller).Call(j.Id("cmd"), j.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.Dict{
					j.Id("FromFile"): j.Lit("input-file"),
				}))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling request: %w"), j.Err())),
				)
			}

			// execute query
			fn.
				If(
					j.ListFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("client").Dot(m.methods[query].GoName).CallFunc(func(args *j.Group) {
						args.Id("cmd").Dot("Context")
						args.Id("cmd").Dot("String").Call(j.Lit("workflow-id"))
						args.Id("cmd").Dot("String").Call(j.Lit("run-id"))
						if hasInput {
							args.Id("req")
						}
					}),
					j.Err().Op("!=").Nil(),
				).
				Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error executing %q query: %w"), j.Id(m.toCamel("%sQueryName", query)), j.Err())),
				).
				Else().
				BlockFunc(func(b *j.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					} else {
						fn.Qual("fmt", "Println").Call(j.Lit("success"))
					}
					b.Return(j.Nil())
				})
		})
	})
}

// genCliSignalCommand generates a <Signal> command
func (m *Manifest) genCliSignalCommand(cmds *j.Group, signal protoreflect.FullName) {
	method := m.methods[signal]
	opts := m.signals[signal]
	hasInput := !isEmpty(method.Input)

	name := m.caser.ToKebab(m.methods[signal].GoName)
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
		usage = fmt.Sprintf("executes a %s signal", m.fqnForSignal(signal))
	}

	cmds.CustomFunc(multiLineValues, func(cmd *j.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *j.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if m.cfg.CliCategories {
			cmd.Id("Category").Op(":").Lit("SIGNALS")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *j.Group) {
			// add workflow-id required flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("workflow-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("workflow id"))
				fields.Id("Required").Op(":").True()
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("w"))
			})
			// add run-id optional flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("run-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run id"))
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("r"))
			})
			if hasInput {
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("f"))
					fields.Id("Category").Op(":").Lit("INPUT")
				})
				// add request flags
				for _, field := range method.Input.Fields {
					m.genCliFlagForField(flags, field, "INPUT", "")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(j.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *j.Group) {
			// initialize client
			fn.List(j.Id("c"), j.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(j.Id("cmd"))
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client for command: %w"), j.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(m.toCamel("New%sClient", m.Service.GoName)).Call(j.Id("c"))

			// unmarshal input
			if hasInput {
				inputName := m.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", m.toCamel("%s", inputName))
				fn.List(j.Id("req"), j.Err()).Op(":=").Id(unmarshaller).Call(j.Id("cmd"), j.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.Dict{
					j.Id("FromFile"): j.Lit("input-file"),
				}))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling request: %w"), j.Err())),
				)
			}

			fn.If(
				j.Err().Op(":=").Id("client").Dot(m.methods[signal].GoName).CallFunc(func(args *j.Group) {
					args.Id("cmd").Dot("Context")
					args.Id("cmd").Dot("String").Call(j.Lit("workflow-id"))
					args.Id("cmd").Dot("String").Call(j.Lit("run-id"))
					if hasInput {
						args.Id("req")
					}
				}),
				j.Err().Op("!=").Nil(),
			).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error sending %q signal: %w"), j.Id(m.toCamel("%sSignalName", signal)), j.Err())),
			)

			// print response
			fn.Qual("fmt", "Println").Call(j.Lit("success"))
			fn.Return(j.Nil())
		})
	})
}

// genCliUpdateCommand generates an <Update> command
func (m *Manifest) genCliUpdateCommand(f *j.Group, update protoreflect.FullName) {
	method := m.methods[update]
	opts := m.updates[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	name := m.caser.ToKebab(m.methods[update].GoName)
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
		usage = fmt.Sprintf("executes a(n) %s update", m.fqnForUpdate(update))
	}

	f.CustomFunc(multiLineValues, func(cmd *j.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *j.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if m.cfg.CliCategories {
			cmd.Id("Category").Op(":").Lit("UPDATES")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *j.Group) {
			// add async flag
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run workflow update in the background and print workflow, execution, and udpate id"))
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("d"))
			})
			// add workflow-id required flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("workflow-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("workflow id"))
				fields.Id("Required").Op(":").True()
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("w"))
			})
			// add run-id optional flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("run-id")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run id"))
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("r"))
			})
			if hasInput {
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("f"))
					fields.Id("Category").Op(":").Lit("INPUT")
				})
				// add request flags
				for _, field := range method.Input.Fields {
					m.genCliFlagForField(flags, field, "INPUT", "")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(j.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *j.Group) {
			// initialize client
			fn.List(j.Id("c"), j.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(j.Id("cmd"))
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client for command: %w"), j.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(m.toCamel("New%sClient", m.Service.GoName)).Call(j.Id("c"))

			// unmarshal input
			if hasInput {
				inputName := m.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", m.toCamel("%s", inputName))
				fn.List(j.Id("req"), j.Err()).Op(":=").Id(unmarshaller).Call(j.Id("cmd"), j.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.Dict{
					j.Id("FromFile"): j.Lit("input-file"),
				}))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling request: %w"), j.Err())),
				)
			}

			// execute update operation
			fn.List(j.Id("handle"), j.Err()).Op(":=").Id("client").Dot(m.toCamel("%sAsync", update)).CallFunc(func(args *j.Group) {
				args.Id("cmd").Dot("Context")
				args.Id("cmd").Dot("String").Call(j.Lit("workflow-id"))
				args.Id("cmd").Dot("String").Call(j.Lit("run-id"))
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error executing %s update: %w"), j.Id(m.toCamel("%sUpdateName", update)), j.Err())),
			)

			// handle async invocation
			fn.If(j.Id("cmd").Dot("Bool").Call(j.Lit("detach"))).Block(
				j.Qual("fmt", "Println").Call(j.Lit("success")),
				j.Qual("fmt", "Printf").Call(j.Lit("workflow id: %s\n"), j.Id("handle").Dot("WorkflowID").Call()),
				j.Qual("fmt", "Printf").Call(j.Lit("run id: %s\n"), j.Id("handle").Dot("RunID").Call()),
				j.Qual("fmt", "Printf").Call(j.Lit("update id: %s\n"), j.Id("handle").Dot("UpdateID").Call()),
				j.Return(j.Nil()),
			)

			// handle synchronous invocation
			fn.
				If(
					j.ListFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("handle").Dot("Get").Call(j.Id("cmd").Dot("Context")),
					j.Err().Op("!=").Nil(),
				).
				Block(
					j.Return(j.Err()),
				).
				Else().
				BlockFunc(func(b *j.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					}
					b.Return(j.Nil())
				})
		})
	})
}

// genCliWorkerCommand generates a <Workflow> command
func (m *Manifest) genCliWorkerCommand(f *j.Group) {
	f.CustomFunc(multiLineValues, func(cmd *j.Group) {
		cmd.Id("Name").Op(":").Lit("worker")
		cmd.Id("Usage").Op(":").Lit(fmt.Sprintf("runs a %s worker process", m.Service.Desc.FullName()))
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Action").Op(":").Func().Params(j.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *j.Group) {
			// initialize client
			fn.List(j.Id("c"), j.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(j.Id("cmd"))
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client for command: %w"), j.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()

			// initialize worker
			fn.List(j.Id("w"), j.Err()).Op(":=").Id("opts").Dot("worker").Call(j.Id("cmd"), j.Id("c"))
			fn.If(j.Id("opts").Dot("worker").Op("!=").Nil()).Block(
				j.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing worker: %w"), j.Err())),
				),
			)

			// run worker
			fn.If(j.Err().Op(":=").Id("w").Dot("Start").Call(), j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error starting worker: %w"), j.Err())),
			)
			fn.Defer().Id("w").Dot("Stop").Call()
			fn.Op("<-").Id("cmd").Dot("Context").Dot("Done").Call()
			fn.Return(j.Nil())
		})
	})
}

// genCliWorkflowCommand generates a <Workflow> command
func (m *Manifest) genCliWorkflowCommand(f *j.Group, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	name := m.caser.ToKebab(m.methods[workflow].GoName)
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
		usage = fmt.Sprintf("executes a(n) %s workflow", m.fqnForWorkflow(workflow))
	}

	f.CustomFunc(multiLineValues, func(cmd *j.Group) {
		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *j.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if m.cfg.CliCategories {
			cmd.Id("Category").Op(":").Lit("WORKFLOWS")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *j.Group) {
			// add async flag
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit("run workflow in the background and print workflow and execution id")
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("d"))
			})
			// add task-queue flag
			flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("task-queue")
				fields.Id("Usage").Op(":").Lit("task queue name")
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("t"))
				fields.Id("EnvVars").Op(":").Index().String().Values(j.Lit("TEMPORAL_TASK_QUEUE_NAME"), j.Lit("TEMPORAL_TASK_QUEUE"), j.Lit("TASK_QUEUE_NAME"), j.Lit("TASK_QUEUE"))
				tq := m.opts.GetTaskQueue()
				if tq == "" {
					fields.Id("Required").Op(":").True()
				} else {
					fields.Id("Value").Op(":").Lit(tq)
				}
			})
			if hasInput {
				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("f"))
					fields.Id("Category").Op(":").Lit("INPUT")
				})
				// add request flags
				for _, field := range method.Input.Fields {
					m.genCliFlagForField(flags, field, "INPUT", "")
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(j.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *j.Group) {
			// initialize client
			fn.List(j.Id("tc"), j.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(j.Id("cmd"))
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client for command: %w"), j.Err())),
			)
			fn.Defer().Id("tc").Dot("Close").Call()
			fn.Id("c").Op(":=").Id(m.toCamel("New%sClient", m.Service.GoName)).Call(j.Id("tc"))

			// unmarshal input
			if hasInput {
				inputName := m.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", m.toCamel("%s", inputName))
				fn.List(j.Id("req"), j.Err()).Op(":=").Id(unmarshaller).Call(j.Id("cmd"), j.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.Dict{
					j.Id("FromFile"): j.Lit("input-file"),
				}))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling request: %w"), j.Err())),
				)
			}

			// build workflow options
			fn.Id("opts").Op(":=").Qual(clientPkg, "StartWorkflowOptions").Values()
			fn.If(j.Id("tq").Op(":=").Id("cmd").Dot("String").Call(j.Lit("task-queue")), j.Id("tq").Op("!=").Lit("")).Block(
				j.Id("opts").Dot("TaskQueue").Op("=").Id("tq"),
			)

			// execute operation
			fn.List(j.Id("run"), j.Err()).Op(":=").Id("c").Dot(m.toCamel("%sAsync", workflow)).CallFunc(func(args *j.Group) {
				args.Id("cmd").Dot("Context")
				if hasInput {
					args.Id("req")
				}
				args.Id(m.toCamel("New%sOptions", workflow)).Call().Dot("WithStartWorkflowOptions").Call(j.Id("opts"))
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error starting %s workflow: %w"), j.Id(m.toCamel("%sWorkflowName", workflow)), j.Err())),
			)

			// handle async invocation
			fn.If(j.Id("cmd").Dot("Bool").Call(j.Lit("detach"))).Block(
				j.Qual("fmt", "Println").Call(j.Lit("success")),
				j.Qual("fmt", "Printf").Call(j.Lit("workflow id: %s\n"), j.Id("run").Dot("ID").Call()),
				j.Qual("fmt", "Printf").Call(j.Lit("run id: %s\n"), j.Id("run").Dot("RunID").Call()),
				j.Return(j.Nil()),
			)

			// handle synchronous invocation
			fn.
				If(
					j.ListFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("run").Dot("Get").Call(j.Id("cmd").Dot("Context")),
					j.Err().Op("!=").Nil(),
				).
				Block(
					j.Return(j.Err()),
				).
				Else().
				BlockFunc(func(b *j.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					}
					b.Return(j.Nil())
				})
		})
	})
}

// genCliWorkflowWithSignalCommand generates a <Workflow>-with-<Signal> command
func (m *Manifest) genCliWorkflowWithSignalCommand(cmds *j.Group, w, signal protoreflect.FullName, opts *temporalv1.WorkflowOptions_Signal) {
	method := m.methods[w]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := m.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	name := opts.GetCli().GetName()
	if name == "" {
		workflowName := m.workflows[w].GetCli().GetName()
		if workflowName == "" {
			workflowName = m.methods[w].GoName
		}
		signalName := m.signals[signal].GetCli().GetName()
		if signalName == "" {
			signalName = m.methods[signal].GoName
		}
		name = m.caser.ToKebab(strings.Join([]string{workflowName, "with", signalName}, "-"))
	}

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = fmt.Sprintf("sends a %s signal to a %s workflow, starting it if necessary", signal, w)
	}

	cmds.Comment(usage)
	cmds.CustomFunc(multiLineValues, func(cmd *j.Group) {
		fields := map[string]struct{}{}
		collisions := map[string]struct{}{}

		cmd.Id("Name").Op(":").Lit(name)
		cmd.Id("Usage").Op(":").Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			cmd.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *j.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if m.cfg.CliCategories {
			cmd.Id("Category").Op(":").Lit("WORKFLOWS")
		}
		cmd.Id("UseShortOptionHandling").Op(":").True()
		cmd.Id("Before").Op(":").Id("opts").Dot("before")
		cmd.Id("After").Op(":").Id("opts").Dot("after")
		cmd.Id("Flags").Op(":").Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(flags *j.Group) {
			flags.Op("&").Qual(cliPkg, "BoolFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
				fields.Id("Name").Op(":").Lit("detach")
				fields.Id("Usage").Op(":").Lit(strings.TrimSpace("run workflow in the background and print workflow and execution id"))
				fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("d"))
			})

			if hasInput {
				var category string
				if hasSignalInput {
					category = "INPUT"
				}

				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit("input-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("f"))
					fields.Id("Category").Op(":").Lit(category)
				})

				// add request flags
				for _, field := range method.Input.Fields {
					fields[field.GoName] = struct{}{}
					m.genCliFlagForField(flags, field, category, "")
				}
			}
			if hasSignalInput {
				var category string
				if hasSignalInput {
					category = "SIGNAL"
				}

				// add -f flag to read input from json file
				flags.Op("&").Qual(cliPkg, "StringFlag").CustomFunc(multiLineValues, func(fields *j.Group) {
					fields.Id("Name").Op(":").Lit("signal-file")
					fields.Id("Usage").Op(":").Lit("path to json-formatted input file")
					fields.Id("Aliases").Op(":").Index().String().Values(j.Lit("s"))
					fields.Id("Category").Op(":").Lit(category)
				})

				// add signal flags
				for _, field := range handler.Input.Fields {
					var prefix string
					if _, ok := fields[field.GoName]; ok {
						prefix = handler.GoName
						collisions[m.flagName(field, "")] = struct{}{}
					}
					m.genCliFlagForField(flags, field, category, prefix)
				}
			}
		})
		cmd.Id("Action").Op(":").Func().Params(j.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(fn *j.Group) {
			// initialize client
			fn.List(j.Id("c"), j.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(j.Id("cmd"))
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client for command: %w"), j.Err())),
			)
			fn.Defer().Id("c").Dot("Close").Call()
			fn.Id("client").Op(":=").Id(m.toCamel("New%sClient", m.Service.GoName)).Call(j.Id("c"))

			// unmarshal request
			if hasInput {
				inputName := m.getMessageName(method.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", m.toCamel("%s", inputName))
				fn.List(j.Id("req"), j.Err()).Op(":=").Qual(m.goImportPathForMethod(w), unmarshaller).Call(j.Id("cmd"), j.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.Dict{
					j.Id("FromFile"): j.Lit("input-file"),
				}))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling request: %w"), j.Err())),
				)
			}

			// unmarshal signal
			if hasSignalInput {
				inputName := m.getMessageName(handler.Input)
				unmarshaller := fmt.Sprintf("UnmarshalCliFlagsTo%s", m.toCamel("%s", inputName))
				fn.List(j.Id("signal"), j.Err()).Op(":=").Qual(m.goImportPathForMethod(signal), unmarshaller).CallFunc(func(b *j.Group) {
					b.Id("cmd")
					b.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.DictFunc(func(d j.Dict) {
						d[j.Id("FromFile")] = j.Lit("signal-file")
						if len(collisions) > 0 {
							b.Qual(helpersPkg, "UnmarshalCliFlagsOptions").CustomFunc(multiLineValues, func(b *j.Group) {
								b.Id("Prefix").Op(":").Lit(m.caser.ToKebab(handler.GoName))
								b.Id("PrefixFlags").Op(":").Map(j.String()).Struct().CustomFunc(multiLineValues, func(b *j.Group) {
									for _, field := range workflow.DeterministicKeys(collisions) {
										b.Lit(field).Op(":").Values()
									}
								})
							})
						}
					}))
				})
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling signal: %w"), j.Err())),
				)
			}

			// execute operation
			fn.List(j.Id("run"), j.Err()).Op(":=").Id("client").Dot(m.toCamel("%sWith%sAsync", w, signal)).CallFunc(func(args *j.Group) {
				args.Id("cmd").Dot("Context")
				if hasInput {
					args.Id("req")
				}
				if hasSignalInput {
					args.Id("signal")
				}
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error starting %s workflow with %s signal: %w"), j.Id(m.toCamel("%sWorkflowName", w)), j.Qual(m.goImportPathForMethod(signal), m.toCamel("%sSignalName", signal)), j.Err())),
			)

			// handle async invocation
			fn.If(j.Id("cmd").Dot("Bool").Call(j.Lit("detach"))).Block(
				j.Qual("fmt", "Println").Call(j.Lit("success")),
				j.Qual("fmt", "Printf").Call(j.Lit("workflow id: %s\n"), j.Id("run").Dot("ID").Call()),
				j.Qual("fmt", "Printf").Call(j.Lit("run id: %s\n"), j.Id("run").Dot("RunID").Call()),
				j.Return(j.Nil()),
			)

			// handle synchronous invocation
			fn.
				If(
					j.ListFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Id("resp")
						}
						returnVals.Err()
					}).Op(":=").Id("run").Dot("Get").Call(j.Id("cmd").Dot("Context")),
					j.Err().Op("!=").Nil(),
				).
				Block(
					j.Return(j.Err()),
				).
				Else().
				BlockFunc(func(b *j.Group) {
					// print response
					if hasOutput {
						genCliPrintMessage(b, "resp")
					}
					b.Return(j.Nil())
				})
		})
	})
}

func (m *Manifest) genCliWorkflowWithUpdateCommand(g *j.Group, w, update protoreflect.FullName, opts *temporalv1.WorkflowOptions_Update) {
	method := m.methods[w]
	hasWorkflowInput := !isEmpty(method.Input)
	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)
	hasUpdateOutput := !isEmpty(handler.Output)

	name := opts.GetCli().GetName()
	if name == "" {
		workflowName := m.workflows[w].GetCli().GetName()
		if workflowName == "" {
			workflowName = m.methods[w].GoName
		}
		updateName := m.updates[update].GetCli().GetName()
		if updateName == "" {
			updateName = m.methods[update].GoName
		}
		name = m.caser.ToKebab(strings.Join([]string{workflowName, "with", updateName}, "-"))
	}
	clientCtor := m.Names().clientCtor()
	updateWithStart := m.Names().clientUpdateWithStartAsync(w, update)

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = fmt.Sprintf("executes a(n) %s update on a %s workflow, starting it if necessary", update, w)
	}

	g.Comment(usage)
	g.Values(j.DictFunc(func(g j.Dict) {
		g[j.Id("Name")] = j.Lit(name)
		g[j.Id("Usage")] = j.Lit(usage)
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			g[j.Id("Aliases")] = j.Index().String().ValuesFunc(func(g *j.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if m.cfg.CliCategories {
			g[j.Id("Category")] = j.Lit("WORKFLOWS")
		}
		g[j.Id("UseShortOptionHandling")] = j.True()
		g[j.Id("Before")] = j.Id("opts").Dot("before")
		g[j.Id("After")] = j.Id("opts").Dot("after")

		collisions := map[string]struct{}{}
		g[j.Id("Flags")] = j.Index().Qual(cliPkg, "Flag").CustomFunc(multiLineValues, func(g *j.Group) {
			fields := map[string]struct{}{}

			// add async flag
			g.Op("&").Qual(cliPkg, "BoolFlag").Values(j.DictFunc(func(g j.Dict) {
				g[j.Id("Name")] = j.Lit("detach")
				g[j.Id("Usage")] = j.Lit(strings.TrimSpace("run workflow update in the background and print workflow, execution, and update id"))
				g[j.Id("Aliases")] = j.Index().String().Values(j.Lit("d"))
			}))

			// add workflow input flags
			if hasWorkflowInput {
				var category string
				if hasUpdateInput {
					category = "INPUT"
				}

				// add -f flag to read input from json file
				g.Op("&").Qual(cliPkg, "StringFlag").Values(j.DictFunc(func(g j.Dict) {
					g[j.Id("Name")] = j.Lit("input-file")
					g[j.Id("Usage")] = j.Lit("path to json-formatted input file")
					g[j.Id("Aliases")] = j.Index().String().Values(j.Lit("f"))
					g[j.Id("Category")] = j.Lit(category)
				}))

				for _, field := range method.Input.Fields {
					fields[field.GoName] = struct{}{}
					m.genCliFlagForField(g, field, category, "")
				}
			}

			// add update input flags
			if hasUpdateInput {
				var category string
				if hasUpdateInput {
					category = "UPDATE"
				}

				g.Op("&").Qual(cliPkg, "StringFlag").Values(j.DictFunc(func(g j.Dict) {
					g[j.Id("Name")] = j.Lit("update-file")
					g[j.Id("Usage")] = j.Lit("path to json-formatted update file")
					g[j.Id("Aliases")] = j.Index().String().Values(j.Lit("u"))
					g[j.Id("Category")] = j.Lit(category)
				}))

				for _, field := range handler.Input.Fields {
					var prefix string
					if _, ok := fields[field.GoName]; ok {
						prefix = handler.GoName
						collisions[m.flagName(field, "")] = struct{}{}
					}
					m.genCliFlagForField(g, field, category, prefix)
				}
			}
		})

		g[j.Id("Action")] = j.Func().Params(j.Id("cmd").Op("*").Qual(cliPkg, "Context")).Error().BlockFunc(func(g *j.Group) {
			// initialize client
			g.List(j.Id("c"), j.Err()).Op(":=").Id("opts").Dot("clientForCommand").Call(j.Id("cmd"))
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client for command: %w"), j.Err()))
			})
			g.Defer().Id("c").Dot("Close").Call()
			g.Id("client").Op(":=").Id(clientCtor).Call(j.Id("c"))

			// unmarshal workflow input
			if hasWorkflowInput {
				unmarshaller := m.Names().unmarshalCliFlagsTo(method.Input)

				g.List(j.Id("input"), j.Err()).Op(":=").Qual(m.goImportPathForMethod(w), unmarshaller).Call(j.Id("cmd"), j.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.Dict{
					j.Id("FromFile"): j.Lit("input-file"),
				}))
				g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
					g.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling input: %w"), j.Err()))
				})
			}

			// unmarshal update input
			if hasUpdateInput {
				unmarshaller := m.Names().unmarshalCliFlagsTo(handler.Input)

				g.List(j.Id("update"), j.Err()).Op(":=").Qual(m.goImportPathForMethod(update), unmarshaller).CallFunc(func(b *j.Group) {
					b.Id("cmd")
					b.Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values(j.DictFunc(func(d j.Dict) {
						d[j.Id("FromFile")] = j.Lit("update-file")
						if len(collisions) > 0 {
							b.Qual(helpersPkg, "UnmarshalCliFlagsOptions").CustomFunc(multiLineValues, func(b *j.Group) {
								b.Id("Prefix").Op(":").Lit(m.caser.ToKebab(handler.GoName))
								b.Id("PrefixFlags").Op(":").Map(j.String()).Struct().CustomFunc(multiLineValues, func(b *j.Group) {
									for _, field := range workflow.DeterministicKeys(collisions) {
										b.Lit(field).Op(":").Values()
									}
								})
							})
						}
					}))
				})
				g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
					g.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error unmarshalling update: %w"), j.Err()))
				})
			}

			// execute operation
			g.List(j.Id("handle"), j.Id("_"), j.Err()).Op(":=").Id("client").Dot(updateWithStart).CallFunc(func(g *j.Group) {
				g.Id("cmd").Dot("Context")
				if hasWorkflowInput {
					g.Id("input")
				}
				if hasUpdateInput {
					g.Id("update")
				}
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Qual("fmt", "Errorf").Call(j.Lit("error starting workflow with update: %w"), j.Err()))
			})

			// handle async invocation
			g.If(j.Id("cmd").Dot("Bool").Call(j.Lit("detach"))).Block(
				j.Qual("fmt", "Println").Call(j.Lit("success")),
				j.Qual("fmt", "Printf").Call(j.Lit("workflow id: %s\n"), j.Id("handle").Dot("WorkflowID").Call()),
				j.Qual("fmt", "Printf").Call(j.Lit("run id: %s\n"), j.Id("handle").Dot("RunID").Call()),
				j.Qual("fmt", "Printf").Call(j.Lit("update id: %s\n"), j.Id("handle").Dot("UpdateID").Call()),
				j.Return(j.Nil()),
			)

			// handle synchronous invocation
			g.IfFunc(func(g *j.Group) {
				g.ListFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Id("out")
					}
					g.Err()
				}).Op(":=").Id("handle").Dot("Get").Call(j.Id("cmd").Dot("Context"))
				g.Err().Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Return(j.Err())
			}).Else().BlockFunc(func(g *j.Group) {
				// print response
				if hasUpdateOutput {
					genCliPrintMessage(g, "out")
				}
				g.Return(j.Nil())
			})
		})
	}))
}
