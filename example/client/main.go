package main

import (
	"context"
	"log"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	c, _ := client.Dial(client.Options{})
	client, ctx := examplev1.NewClient(c), context.Background()

	run, _ := client.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "test"})
	log.Printf("started workflow: workflow_id=%s, run_id=%s\n", run.ID(), run.RunID())

	log.Println("signalling progress")
	_ = run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7})

	progress, _ := run.GetFooProgress(ctx)
	log.Printf("queried progress: %s\n", progress.String())

	update, _ := run.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
	log.Printf("updated progress: %s\n", update.String())

	resp, _ := run.Get(ctx)
	log.Printf("workflow completed: %s\n", resp.String())
}
