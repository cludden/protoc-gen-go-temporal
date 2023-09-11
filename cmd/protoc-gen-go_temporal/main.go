package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/cludden/protoc-gen-go-temporal/internal/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	version = "dev"
	commit  = "latest"
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-temporal: %s\n", version)
		fmt.Printf("go: %s\n", runtime.Version())
		return
	}

	p := plugin.New(commit, version)

	opts := protogen.Options{
		ParamFunc: p.Param,
	}

	opts.Run(p.Run)
}
