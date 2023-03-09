package plugin

import (
	"fmt"

	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
)

// Plugin provides a protoc plugin for generating temporal workers and clients in go
type Plugin struct {
	*protogen.Plugin
}

// Param provides a protogen ParamFunc handler
func (p *Plugin) Param(key, value string) error {
	return nil
}

// Run defines the plugin entrypoint
func (p *Plugin) Run(plugin *protogen.Plugin) error {
	p.Plugin = plugin
	servicesByPkg := map[string]int{}

	for _, file := range p.Files {
		if !file.Generate {
			continue
		}

		pkgName := string(file.GoPackageName)
		f := g.NewFile(string(pkgName))
		var hasContent bool

		for _, service := range file.Services {
			svc := parseService(service)
			if len(svc.activities) == 0 && len(svc.workflows) == 0 && len(svc.signals) == 0 && len(svc.queries) == 0 {
				continue
			}

			if _, ok := servicesByPkg[string(file.GoPackageName)]; !ok {
				servicesByPkg[pkgName] = 0
			}
			servicesByPkg[pkgName]++
			if n := servicesByPkg[pkgName]; n > 1 {
				return fmt.Errorf("only one temporal service per package is currently supported, observed violation for package: %s", pkgName)
			}
			svc.render(f)
			hasContent = true
			break
		}

		if !hasContent {
			continue
		}

		if err := f.Render(p.NewGeneratedFile(fmt.Sprintf("%s_temporal.pb.go", file.GeneratedFilenamePrefix), file.GoImportPath)); err != nil {
			return fmt.Errorf("error rendering file: %w", err)
		}
	}
	return nil
}

// isEmpty checks if the message is a google.protobuf.Empty message
func isEmpty(m *protogen.Message) bool {
	return m.Desc.FullName() == "google.protobuf.Empty"
}
