package main

import (
	"flag"
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "0.1.0"

const (
	contextPackage = protogen.GoImportPath("context")
	fmtPackage     = protogen.GoImportPath("fmt")
	reflectPackage = protogen.GoImportPath("reflect")

	activityPackage = protogen.GoImportPath("go.temporal.io/sdk/activity")
	clientPackage   = protogen.GoImportPath("go.temporal.io/sdk/client")
	workerPackage   = protogen.GoImportPath("go.temporal.io/sdk/worker")
	workflowPackage = protogen.GoImportPath("go.temporal.io/sdk/workflow")
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-temporal %v\n", version)
		return
	}

	var flags flag.FlagSet
	// TODO(cretz): Flags?

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}
			if err := (&gen{Plugin: plugin, File: f}).gen(); err != nil {
				return fmt.Errorf("file %v failed: %w", f.Desc.Name(), err)
			}
		}
		return nil
	})
}

func isEmpty(m *protogen.Message) bool {
	return m.Desc.FullName() == "google.protobuf.Empty"
}

func private(s string) string {
	// Just lowercase first char
	// TODO(cretz): More sophisticated if starts with more than one upper char
	return strings.ToLower(s[:1]) + s[1:]
}
