package main

import (
	"context"
	"log"
	"sync"

	"github.com/cludden/protoc-gen-go-temporal/example/mutexv1"
	"go.temporal.io/sdk/client"
)

func main() {
	// initialize temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("error initializing client: %v", err)
	}
	defer c.Close()

	// initialize mutex client
	mutex := mutexv1.NewClient(c)
	resource := "foo"

	// start two SampleWorkflowWithMutex for the same resource concurrently
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := mutex.SampleWorkflowWithMutex(context.Background(), nil, &mutexv1.SampleWorkflowWithMutexRequest{
				Resource: resource,
			}); err != nil {
				log.Printf("error executing sample workflow: %v\n", err)
			}
		}()
	}
	wg.Wait()
}
