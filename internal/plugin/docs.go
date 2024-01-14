package plugin

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/cludden/protoc-gen-go-temporal/internal/plugin/docs"
	"github.com/cludden/protoc-gen-go-temporal/internal/templates"
	"github.com/gosimple/slug"
	"github.com/hako/durafmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func renderDocs(p *protogen.Plugin, cfg *Config) error {
	data, err := docs.Parse(p)
	if err != nil {
		return fmt.Errorf("error extracting template data: %v", err)
	}

	var raw []byte
	switch cfg.DocsTemplate {
	case "basic":
		raw, err = templates.Templates.ReadFile("basic.tpl")
	default:
		raw, err = os.ReadFile(cfg.DocsTemplate)
	}
	if err != nil {
		return fmt.Errorf("error reading template: %v", err)
	}

	tpl, err := template.New("docs").
		Funcs(template.FuncMap{
			"fmtduration": func(d time.Duration) string {
				return durafmt.Parse(d).String()
			},
			"slug": func(input string) string {
				return slug.Make(input)
			},
		}).
		Funcs(sprig.FuncMap()).
		Parse(string(raw))
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	docs, err := os.Create(cfg.DocsOut)
	if err != nil {
		return fmt.Errorf("error creating doc file: %w", err)
	}

	if err := tpl.Execute(docs, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	return nil
}
