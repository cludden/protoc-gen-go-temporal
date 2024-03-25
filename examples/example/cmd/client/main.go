package main

import (
	"context"
	"log"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	// initialize service client with sdk client
	c, _ := client.Dial(client.Options{})
	client, ctx := examplev1.NewExampleClient(c), context.Background()

	// execute a workflow asynchronously
	run, _ := client.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "test"})
	log.Printf("started workflow: workflow_id=%s, run_id=%s\n", run.ID(), run.RunID())

	// send a signal to the workflow
	log.Println("signalling progress")
	_ = run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7})

	// query the workflow
	progress, _ := run.GetFooProgress(ctx)
	log.Printf("queried progress: %s\n", progress.String())

	// update the workflow
	update, _ := run.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
	log.Printf("updated progress: %s\n", update.String())

	// block on workflow completion
	resp, _ := run.Get(ctx)
	log.Printf("workflow completed: %s\n", resp.String())
}
