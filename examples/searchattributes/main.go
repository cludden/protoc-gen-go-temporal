package main

import (
	"log"
	"os"
	"strings"
	"time"

	searchattributesv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/searchattributes/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	customKeywordField     = temporal.NewSearchAttributeKeyKeyword("CustomKeywordField")
	customKeywordListField = temporal.NewSearchAttributeKeyKeywordList("CustomKeywordListField")
	customTextField        = temporal.NewSearchAttributeKeyString("CustomTextField")
	customIntField         = temporal.NewSearchAttributeKeyInt64("CustomIntField")
	customDoubleField      = temporal.NewSearchAttributeKeyFloat64("CustomDoubleField")
	customBoolField        = temporal.NewSearchAttributeKeyBool("CustomBoolField")
	customDatetimeField    = temporal.NewSearchAttributeKeyTime("CustomDatetimeField")
)

func main() {
	app, err := searchattributesv1.NewExampleCli(
		searchattributesv1.NewExampleCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
			w := worker.New(c, searchattributesv1.ExampleTaskQueue, worker.Options{})
			searchattributesv1.RegisterExampleWorkflows(w, &Workflows{})
			return w, nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type Workflows struct{}

type SearchAttributesWorkflow struct {
	*searchattributesv1.SearchAttributesWorkflowInput
	log tlog.Logger
}

func (w *Workflows) SearchAttributes(
	ctx workflow.Context,
	input *searchattributesv1.SearchAttributesWorkflowInput,
) (searchattributesv1.SearchAttributesWorkflow, error) {
	return &SearchAttributesWorkflow{input, workflow.GetLogger(ctx)}, nil
}

func (w *SearchAttributesWorkflow) Execute(ctx workflow.Context) (err error) {
	sa := workflow.GetInfo(ctx).SearchAttributes
	for _, attr := range strings.Split("CustomBoolField,CustomDatetimeField,CustomDoubleField,CustomIntField,CustomKeywordField,CustomTextField", ",") {
		if p, ok := sa.IndexedFields[attr]; ok {
			switch attr {
			case "CustomBoolField":
				var result bool
				err = converter.GetDefaultDataConverter().FromPayload(p, &result)
				w.log.Info("search attribute", "name", attr, "value", result, "error", err)
			case "CustomDatetimeField":
				var result time.Time
				err = converter.GetDefaultDataConverter().FromPayload(p, &result)
				w.log.Info("search attribute", "name", attr, "value", result, "error", err)
			case "CustomDoubleField":
				var result float64
				err = converter.GetDefaultDataConverter().FromPayload(p, &result)
				w.log.Info("search attribute", "name", attr, "value", result, "error", err)
			case "CustomIntField":
				var result int
				err = converter.GetDefaultDataConverter().FromPayload(p, &result)
				w.log.Info("search attribute", "name", attr, "value", result, "error", err)
			case "CustomKeywordField":
				var result string
				err = converter.GetDefaultDataConverter().FromPayload(p, &result)
				w.log.Info("search attribute", "name", attr, "value", result, "error", err)
			case "CustomTextField":
				var result string
				err = converter.GetDefaultDataConverter().FromPayload(p, &result)
				w.log.Info("search attribute", "name", attr, "value", result, "error", err)
			}
		}
	}

	return nil
}

type TypedSearchAttributesWorkflow struct {
	*searchattributesv1.TypedSearchAttributesWorkflowInput
	log tlog.Logger
}

func (w *Workflows) TypedSearchAttributes(
	ctx workflow.Context, input *searchattributesv1.TypedSearchAttributesWorkflowInput,
) (searchattributesv1.TypedSearchAttributesWorkflow, error) {
	return &TypedSearchAttributesWorkflow{input, workflow.GetLogger(ctx)}, nil
}

func (w *TypedSearchAttributesWorkflow) Execute(
	ctx workflow.Context,
) (*searchattributesv1.TypedSearchAttributesOutput, error) {
	out := &searchattributesv1.TypedSearchAttributesOutput{}
	sa := workflow.GetTypedSearchAttributes(ctx)
	if v, ok := sa.GetKeyword(customKeywordField); ok {
		out.CustomKeywordField = v
	}
	if v, ok := sa.GetFloat64(customDoubleField); ok {
		out.CustomDoubleField = v
	}
	if v, ok := sa.GetBool(customBoolField); ok {
		out.CustomBoolField = v
	}
	if v, ok := sa.GetTime(customDatetimeField); ok {
		out.CustomDatetimeField = timestamppb.New(v)
	}
	if v, ok := sa.GetKeywordList(customKeywordListField); ok {
		out.CustomKeywordListField = v
	}
	if v, ok := sa.GetString(customTextField); ok {
		out.CustomTextField = v
	}
	if v, ok := sa.GetInt64(customIntField); ok {
		out.CustomIntField = v
	}
	return out, nil
}
