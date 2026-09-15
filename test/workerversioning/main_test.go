package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	workerversioningv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/workerversioning/v1"
	workermocks "github.com/cludden/protoc-gen-go-temporal/mocks/go.temporal.io/sdk/worker"
	"github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func TestWorkerVersioning(t *testing.T) {
	r := workermocks.NewMockRegistry(t)
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name: workerversioningv1.FooWorkflowName,
	})
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name: workerversioningv1.BarWorkflowName,
	})
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name:               workerversioningv1.BazWorkflowName,
		VersioningBehavior: workflow.VersioningBehaviorPinned,
	})
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name:               workerversioningv1.QuxWorkflowName,
		VersioningBehavior: workflow.VersioningBehaviorAutoUpgrade,
	})
	Register(r)
}

func TestVersioningOverride(t *testing.T) {
	pinned := &client.PinnedVersioningOverride{
		Version: worker.WorkerDeploymentVersion{DeploymentName: "example", BuildID: "abc123"},
	}
	autoUpgrade := &client.AutoUpgradeVersioningOverride{}

	cases := map[string]struct {
		options  *workerversioningv1.FooOptions
		expected client.VersioningOverride
	}{
		"unset": {
			options:  workerversioningv1.NewFooOptions().WithTaskQueue("tq"),
			expected: nil,
		},
		"with versioning override": {
			options:  workerversioningv1.NewFooOptions().WithTaskQueue("tq").WithVersioningOverride(pinned),
			expected: pinned,
		},
		"with auto upgrade override": {
			options:  workerversioningv1.NewFooOptions().WithTaskQueue("tq").WithVersioningOverride(autoUpgrade),
			expected: autoUpgrade,
		},
		"from start workflow options": {
			options: workerversioningv1.NewFooOptions().
				WithStartWorkflowOptions(client.StartWorkflowOptions{
					TaskQueue:          "tq",
					VersioningOverride: autoUpgrade,
				}),
			expected: autoUpgrade,
		},
		// the runtime helper wins over a value supplied via WithStartWorkflowOptions
		"with versioning override and start workflow options": {
			options: workerversioningv1.NewFooOptions().
				WithStartWorkflowOptions(client.StartWorkflowOptions{
					TaskQueue:          "tq",
					VersioningOverride: autoUpgrade,
				}).
				WithVersioningOverride(pinned),
			expected: pinned,
		},
	}

	for _, name := range workflow.DeterministicKeys(cases) {
		c := cases[name]
		t.Run(name, func(t *testing.T) {
			opts, err := c.options.Build((&workerversioningv1.FooInput{}).ProtoReflect())
			require.NoError(t, err)
			require.Equal(t, c.expected, opts.VersioningOverride)
		})
	}
}

// TestVersioningOverrideAcrossXNS covers the composition the generated xns
// activity performs: a caller's start workflow options are serialized as
// protobuf, sent to the remote namespace, and rebuilt there.
func TestVersioningOverrideAcrossXNS(t *testing.T) {
	cases := map[string]client.VersioningOverride{
		"pinned": &client.PinnedVersioningOverride{
			Version: worker.WorkerDeploymentVersion{DeploymentName: "example", BuildID: "abc123"},
		},
		"auto upgrade": &client.AutoUpgradeVersioningOverride{},
		"unset":        nil,
	}

	for _, name := range workflow.DeterministicKeys(cases) {
		expected := cases[name]
		t.Run(name, func(t *testing.T) {
			pb, err := xns.MarshalStartWorkflowOptions(client.StartWorkflowOptions{
				TaskQueue:          "tq",
				VersioningOverride: expected,
			})
			require.NoError(t, err)

			opts, err := workerversioningv1.NewFooOptions().
				WithStartWorkflowOptions(xns.UnmarshalStartWorkflowOptions(pb)).
				Build((&workerversioningv1.FooInput{}).ProtoReflect())
			require.NoError(t, err)
			require.Equal(t, expected, opts.VersioningOverride)
			require.Equal(t, "tq", opts.TaskQueue)
		})
	}
}

// TestVersioningOverrideNotOnChildOptions guards the scope of the helper:
// workflow.ChildWorkflowOptions carries no versioning override, so child
// options must not offer one.
func TestVersioningOverrideNotOnChildOptions(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	require.True(t, ok)

	generated := filepath.Join(
		filepath.Dir(thisFile),
		"..", "..",
		"gen", "test", "workerversioning", "v1", "example_temporal.pb.go",
	)
	content, err := os.ReadFile(generated)
	require.NoError(t, err)
	src := string(content)

	require.Contains(t, src, "func (o *FooOptions) WithVersioningOverride(")
	require.NotContains(t, src, "ChildOptions) WithVersioningOverride(")

	// nor may the field reach the child options struct
	require.Contains(t, structBody(t, src, "FooOptions"), "versioningOverride")
	require.NotContains(t, structBody(t, src, "FooChildOptions"), "versioningOverride")
}

// structBody returns the field block of the named generated struct.
func structBody(t *testing.T, src, name string) string {
	t.Helper()

	prefix := "type " + name + " struct {"
	start := strings.Index(src, prefix)
	require.NotEqual(t, -1, start, "missing struct %s", name)

	bodyStart := start + len(prefix)
	end := strings.Index(src[bodyStart:], "\n}")
	require.NotEqual(t, -1, end, "missing struct terminator for %s", name)

	return src[bodyStart : bodyStart+end]
}
