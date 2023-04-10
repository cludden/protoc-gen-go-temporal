package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	mutex "github.com/cludden/protoc-gen-go-temporal/example"
	"github.com/cludden/protoc-gen-go-temporal/example/mutexv1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// initialize temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("error initializing client: %v", err)
	}
	defer c.Close()

	// initializ temporal worker and register workflows, activities using
	// generated helpers
	w := worker.New(c, mutexv1.MutexTaskQueue, worker.Options{})
	mutexv1.RegisterActivities(w, &mutex.Activites{
		Client: mutexv1.NewClient(c),
	})
	mutexv1.RegisterWorkflows(w, &mutex.Workflows{})

	// start worker
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	if err := w.Start(); err != nil {
		log.Fatalf("error starting worker: %v", err)
	}

	// wait for close signal before stopping worker
	<-ctx.Done()
	w.Stop()
}
