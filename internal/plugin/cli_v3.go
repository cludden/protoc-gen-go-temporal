package plugin

import (
	"cmp"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/proto"
)

// renderCLIV3 generates cli v3 resources
func (m *Manifest) renderCLIV3(f *j.File) {
	opts := proto.GetExtension(m.Service.Desc.Options(), temporalv1.E_Cli).(*temporalv1.CLIOptions)
	if opts != nil && opts.GetIgnore() {
		return
	}

	m.genCliV3OptionsImpl(f)
	m.genCliV3New(f)
	m.genCliV3NewCommands(f)

	// initialize cli flag unmarshaller index for go package
	if _, ok := m.cliFlagUnmarshallers[string(m.GoImportPath)]; !ok {
		m.cliFlagUnmarshallers[string(m.GoImportPath)] = make(map[string]struct{})
	}
	unmarshallers := m.cliFlagUnmarshallers[string(m.GoImportPath)]

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

	m.cliFlagUnmarshallers[string(m.GoImportPath)] = unmarshallers
}

// genCliNew generates a New<Service>Cli constructor function for CLI v3
func (m *Manifest) genCliV3New(f *j.File) {
	functionName := m.Names().cliV3Ctor()
	optionsName := m.Names().cliV3Options()
	commandsCtor := m.Names().cliV3CommandsCtor()
	cmdOpts := proto.GetExtension(m.Service.Desc.Options(), temporalv1.E_Cli).(*temporalv1.CLIOptions)
	name := cmp.Or(cmdOpts.GetName(), m.caser.ToKebab(m.Service.GoName))

	f.Commentf("%s initializes a cli app for a(n) %s service", functionName, m.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			j.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			j.Op("*").Qual(cliV3Pkg, "Command"),
			j.Error(),
		).
		Block(
			j.List(j.Id("commands"), j.Err()).Op(":=").Id(commandsCtor).Call(j.Id("options").Op("...")),
			j.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing subcommands: %w"), j.Err())),
			),
			j.Return(
				j.Op("&").Qual(cliV3Pkg, "Command").CustomFunc(multiLineValues, func(g *j.Group) {
					g.Id("Name").Op(":").Lit(name)
					if usage := cmdOpts.GetUsage(); usage != "" {
						g.Id("Usage").Op(":").Lit(usage)
					} else {
						g.Id("Usage").Op(":").Lit(string(m.Service.Desc.FullName()) + " operations")
					}
					g.Id("Commands").Op(":").Id("commands")
					g.Id("DisableSliceFlagSeparator").Op(":").True()
				}),
				j.Nil(),
			),
		)
}

// genCliNewCommands generates a new<Service>CommandsV3 constructor function
func (m *Manifest) genCliV3NewCommands(f *j.File) {
	functionName := m.Names().cliV3CommandsCtor()
	optionsName := m.Names().cliV3Options()

	f.Commentf("%s initializes (sub)commands for a %s cli or command", functionName, m.Service.Desc.FullName())
	f.Func().Id(functionName).
		Params(
			j.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			j.Index().Op("*").Qual(cliV3Pkg, "Command"),
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
					Params(
						j.Id("ctx").Qual("context", "Context"),
						j.Id("cmd").Op("*").Qual(cliV3Pkg, "Command"),
					).
					Params(j.Qual(clientPkg, "Client"), j.Error()).
					Block(
						j.Return(j.Qual(clientPkg, "DialContext").Call(
							j.Id("ctx"),
							j.Qual(clientPkg, "Options").Values()),
						),
					),
			),

			// initialize commands
			j.Id("commands").Op(":=").Index().Op("*").Qual(cliV3Pkg, "Command").CustomFunc(j.Options{
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
				j.Id("commands").Op("=").Append(j.Id("commands"), j.Index().Op("*").Qual(cliV3Pkg, "Command").CustomFunc(multiLineValues, func(g *j.Group) {
					m.genCliWorkerCommand(g)
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

// genCliV3OptionsImpl generates a CLIV3Options struct
func (m *Manifest) genCliV3OptionsImpl(f *j.File) {
	typeName := m.Names().cliV3Options()

	// generate type definition
	f.Commentf("%s describes runtime configuration for %s cli v3", typeName, m.Service.Desc.FullName())
	f.Type().Id(typeName).Struct(
		j.Id("after").Func().
			Params(
				j.Qual("context", "Context"),
				j.Op("*").Qual(cliV3Pkg, "Command"),
			).
			Error(),
		j.Id("before").Func().
			Params(
				j.Qual("context", "Context"),
				j.Op("*").Qual(cliV3Pkg, "Command"),
			).
			Params(
				j.Qual("context", "Context"),
				j.Error(),
			),
		j.Id("clientForCommand").Func().
			Params(
				j.Qual("context", "Context"),
				j.Op("*").Qual(cliV3Pkg, "Command"),
			).
			Params(j.Qual(clientPkg, "Client"), j.Error()),
		j.Id("worker").Func().
			Params(
				j.Qual("context", "Context"),
				j.Op("*").Qual(cliV3Pkg, "Command"),
				j.Qual(clientPkg, "Client"),
			).
			Params(j.Qual(workerPkg, "Worker"), j.Error()),
	)

	// generate New<Service>CliOptions
	functionName := m.Names().cliV3OptionsCtor()
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
				Params(
					j.Qual("context", "Context"),
					j.Op("*").Qual(cliV3Pkg, "Command"),
				).
				Params(
					j.Error(),
				),
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
				Params(
					j.Qual("context", "Context"),
					j.Op("*").Qual(cliV3Pkg, "Command"),
				).
				Params(
					j.Qual("context", "Context"),
					j.Error(),
				),
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
				Params(
					j.Qual("context", "Context"),
					j.Op("*").Qual(cliV3Pkg, "Command"),
				).
				Params(
					j.Qual(clientPkg, "Client"),
					j.Error(),
				),
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
				Params(
					j.Qual("context", "Context"),
					j.Op("*").Qual(cliV3Pkg, "Command"),
					j.Qual(clientPkg, "Client"),
				).
				Params(j.Qual(workerPkg, "Worker"), j.Error()),
		).
		Op("*").Id(typeName).
		Block(
			j.Id("opts").Dot("worker").Op("=").Id("fn"),
			j.Return(j.Id("opts")),
		)
}

func (n *names) cliV3CommandsCtor() string {
	return n.toLowerCamel("new%sCommands", n.Service.GoName)
}

func (n *names) cliV3Ctor() string {
	return n.toCamel("New%sCli", n.Service.GoName)
}

func (n *names) cliV3Options() string {
	return n.toCamel("%sCliOptions", n.Service.GoName)
}

func (n *names) cliV3OptionsCtor() string {
	return n.toCamel("New%sCliOptions", n.Service.GoName)
}
