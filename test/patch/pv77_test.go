package main

import (
	"context"
	"fmt"
	"testing"

	patchv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/patch/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/patch"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func TestPv77_WithoutContextPropagator(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	ctx := context.Background()

	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, srv.Stop()) })

	c := srv.Client()
	t.Cleanup(c.Close)

	var wf Pv77Workflows
	var a Pv77Activities
	customTaskQueue := "custom"
	for _, w := range []worker.Worker{
		worker.New(c, patchv1.Pv77FooServiceTaskQueue, worker.Options{}),
		worker.New(c, patchv1.Pv77QuxServiceTaskQueue, worker.Options{}),
		worker.New(c, patchv1.Pv77BarServiceTaskQueue, worker.Options{}),
		worker.New(c, patchv1.Pv77BazServiceTaskQueue, worker.Options{}),
		worker.New(c, customTaskQueue, worker.Options{}),
	} {
		patchv1.RegisterPv77FooServiceWorkflows(w, &wf)
		patchv1.RegisterPv77FooServiceActivities(w, &a)
		patchv1.RegisterPv77QuxServiceWorkflows(w, &wf)
		patchv1.RegisterPv77QuxServiceActivities(w, &a)
		patchv1.RegisterPv77BarServiceWorkflows(w, &wf)
		patchv1.RegisterPv77BarServiceActivities(w, &a)
		patchv1.RegisterPv77BazServiceWorkflows(w, &wf)
		patchv1.RegisterPv77BazServiceActivities(w, &a)
		require.NoError(t, w.Start())
		t.Cleanup(w.Stop)
	}

	cases := []struct {
		taskQueue string
		next      []string
		expected  []string
	}{
		{
			taskQueue: patchv1.Pv77FooServiceTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77BazWorkflowName,
				patchv1.Pv77QuxWorkflowName,
			},
			expected: []string{
				patchv1.Pv77FooServiceTaskQueue,
				patchv1.Pv77BarServiceTaskQueue,
				patchv1.Pv77BazServiceTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77BazWorkflowName,
				patchv1.Pv77QuxWorkflowName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77BazServiceTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77BazWorkflowName,
				patchv1.Pv77FooWorkflowName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77BazServiceTaskQueue,
				patchv1.Pv77FooServiceTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77FooActivityName,
				patchv1.Pv77QuxWorkflowName,
				patchv1.Pv77FooActivityName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
				patchv1.Pv77FooServiceTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77FooActivityName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77QuxActivityName,
				patchv1.Pv77FooActivityName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
				customTaskQueue,
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var out Pv77Output
			var err error
			switch tc.next[0] {
			case patchv1.Pv77FooWorkflowName:
				out, err = patchv1.NewPv77FooServiceClient(c).Pv77Foo(ctx, &patchv1.Pv77FooInput{Next: tc.next[1:]}, patchv1.NewPv77FooOptions().WithTaskQueue(tc.taskQueue))
			case patchv1.Pv77BarWorkflowName:
				out, err = patchv1.NewPv77BarServiceClient(c).Pv77Bar(ctx, &patchv1.Pv77BarInput{Next: tc.next[1:]}, patchv1.NewPv77BarOptions().WithTaskQueue(tc.taskQueue))
			case patchv1.Pv77BazWorkflowName:
				out, err = patchv1.NewPv77BazServiceClient(c).Pv77Baz(ctx, &patchv1.Pv77BazInput{Next: tc.next[1:]}, patchv1.NewPv77BazOptions().WithTaskQueue(tc.taskQueue))
			case patchv1.Pv77QuxWorkflowName:
				out, err = patchv1.NewPv77QuxServiceClient(c).Pv77Qux(ctx, &patchv1.Pv77QuxInput{Next: tc.next[1:]}, patchv1.NewPv77QuxOptions().WithTaskQueue(tc.taskQueue))
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, out.GetTaskQueues())
		})
	}
}

func TestPv77_WithContextPropagator(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	ctx := context.Background()

	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ClientOptions: &client.Options{
			ContextPropagators: []workflow.ContextPropagator{
				patch.NewContextPropagator(),
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, srv.Stop()) })

	c := srv.Client()
	t.Cleanup(c.Close)

	var wf Pv77Workflows
	var a Pv77Activities
	customTaskQueue := "custom"
	for _, w := range []worker.Worker{
		worker.New(c, patchv1.Pv77FooServiceTaskQueue, worker.Options{}),
		worker.New(c, patchv1.Pv77QuxServiceTaskQueue, worker.Options{}),
		worker.New(c, patchv1.Pv77BarServiceTaskQueue, worker.Options{}),
		worker.New(c, patchv1.Pv77BazServiceTaskQueue, worker.Options{}),
		worker.New(c, customTaskQueue, worker.Options{}),
	} {
		patchv1.RegisterPv77FooServiceWorkflows(w, &wf)
		patchv1.RegisterPv77FooServiceActivities(w, &a)
		patchv1.RegisterPv77QuxServiceWorkflows(w, &wf)
		patchv1.RegisterPv77QuxServiceActivities(w, &a)
		patchv1.RegisterPv77BarServiceWorkflows(w, &wf)
		patchv1.RegisterPv77BarServiceActivities(w, &a)
		patchv1.RegisterPv77BazServiceWorkflows(w, &wf)
		patchv1.RegisterPv77BazServiceActivities(w, &a)
		require.NoError(t, w.Start())
		t.Cleanup(w.Stop)
	}

	cases := []struct {
		taskQueue string
		next      []string
		expected  []string
	}{
		{
			taskQueue: patchv1.Pv77FooServiceTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77BazWorkflowName,
				patchv1.Pv77QuxWorkflowName,
			},
			expected: []string{
				patchv1.Pv77FooServiceTaskQueue,
				patchv1.Pv77BarServiceTaskQueue,
				patchv1.Pv77BazServiceTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77BazWorkflowName,
				patchv1.Pv77QuxWorkflowName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77BazServiceTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77BazWorkflowName,
				patchv1.Pv77FooWorkflowName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77BazServiceTaskQueue,
				customTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77FooActivityName,
				patchv1.Pv77QuxWorkflowName,
				patchv1.Pv77FooActivityName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
				customTaskQueue,
			},
		},
		{
			taskQueue: customTaskQueue,
			next: []string{
				patchv1.Pv77FooWorkflowName,
				patchv1.Pv77FooActivityName,
				patchv1.Pv77BarWorkflowName,
				patchv1.Pv77QuxActivityName,
				patchv1.Pv77FooActivityName,
			},
			expected: []string{
				customTaskQueue,
				customTaskQueue,
				customTaskQueue,
				patchv1.Pv77QuxServiceTaskQueue,
				customTaskQueue,
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var out Pv77Output
			var err error
			switch tc.next[0] {
			case patchv1.Pv77FooWorkflowName:
				out, err = patchv1.NewPv77FooServiceClient(c).Pv77Foo(ctx, &patchv1.Pv77FooInput{Next: tc.next[1:]}, patchv1.NewPv77FooOptions().WithTaskQueue(tc.taskQueue))
			case patchv1.Pv77BarWorkflowName:
				out, err = patchv1.NewPv77BarServiceClient(c).Pv77Bar(ctx, &patchv1.Pv77BarInput{Next: tc.next[1:]}, patchv1.NewPv77BarOptions().WithTaskQueue(tc.taskQueue))
			case patchv1.Pv77BazWorkflowName:
				out, err = patchv1.NewPv77BazServiceClient(c).Pv77Baz(ctx, &patchv1.Pv77BazInput{Next: tc.next[1:]}, patchv1.NewPv77BazOptions().WithTaskQueue(tc.taskQueue))
			case patchv1.Pv77QuxWorkflowName:
				out, err = patchv1.NewPv77QuxServiceClient(c).Pv77Qux(ctx, &patchv1.Pv77QuxInput{Next: tc.next[1:]}, patchv1.NewPv77QuxOptions().WithTaskQueue(tc.taskQueue))
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, out.GetTaskQueues())
		})
	}
}
