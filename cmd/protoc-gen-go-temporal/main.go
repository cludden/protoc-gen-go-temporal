package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/cludden/protoc-gen-go-temporal/internal/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

var Version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-temporal: %s\n", Version)
		fmt.Printf("go: %s\n", runtime.Version())
		return
	}

	p := plugin.Plugin{}

	opts := protogen.Options{
		ParamFunc: p.Param,
	}

	opts.Run(p.Run)
}
