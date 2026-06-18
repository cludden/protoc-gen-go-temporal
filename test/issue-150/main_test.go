package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	issue_150v1 "github.com/cludden/protoc-gen-go-temporal/gen/test/issue-150/v1"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func TestIssue150(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	require.True(t, ok)

	generated := filepath.Join(
		filepath.Dir(thisFile),
		"..", "..",
		"gen", "test", "issue-150", "v1", "issue-150_temporal.pb.go",
	)
	content, err := os.ReadFile(generated)
	require.NoError(t, err)
	src := string(content)

	explicitZeroBuild := getBuildFunctionBody(t, src, "ExplicitPriorityActivityOptions")
	require.Contains(t, explicitZeroBuild, "opts.Priority.PriorityKey = 4")

	unsetBuild := getBuildFunctionBody(t, src, "UnsetPriorityActivityOptions")
	require.NotContains(t, unsetBuild, "opts.Priority.PriorityKey = 4")
}

func getBuildFunctionBody(t *testing.T, src, optionsType string) string {
	t.Helper()

	prefix := "func (o *" + optionsType + ") Build(ctx workflow.Context) (workflow.Context, error) {"
	start := strings.Index(src, prefix)
	require.NotEqual(t, -1, start, "missing Build function for %s", optionsType)

	bodyStart := start + len(prefix)
	end := strings.Index(src[bodyStart:], "\n}\n\n// WithActivityOptions")
	require.NotEqual(t, -1, end, "missing Build function terminator for %s", optionsType)

	return src[bodyStart : bodyStart+end]
}

func TestIssue150_StartWorkflowOverrides(t *testing.T) {
	cases := map[string]struct {
		options  *issue_150v1.ExplicitPriorityOptions
		expected client.StartWorkflowOptions
	}{
		"proto defaults": {
			options: issue_150v1.NewExplicitPriorityOptions(),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey:    2,
					FairnessKey:    "default",
					FairnessWeight: 1,
				},
			},
		},
		"with priority": {
			options: issue_150v1.NewExplicitPriorityOptions().
				WithPriority(temporal.Priority{
					PriorityKey: 4,
				}),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey: 4,
				},
			},
		},
		"with start workflow options": {
			options: issue_150v1.NewExplicitPriorityOptions().
				WithStartWorkflowOptions(client.StartWorkflowOptions{
					TaskQueue: "custom-queue",
					Priority: temporal.Priority{
						PriorityKey: 5,
					},
				}),
			expected: client.StartWorkflowOptions{
				TaskQueue: "custom-queue",
				Priority: temporal.Priority{
					PriorityKey:    5,
					FairnessKey:    "default",
					FairnessWeight: 1,
				},
			},
		},
		"with mixed options": {
			options: issue_150v1.NewExplicitPriorityOptions().
				WithStartWorkflowOptions(client.StartWorkflowOptions{
					Priority: temporal.Priority{
						FairnessKey:    "custom-fairness",
						FairnessWeight: 2.0,
					},
				}),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey:    2,
					FairnessKey:    "custom-fairness",
					FairnessWeight: 2.0,
				},
			},
		},
		"with priority and start workflow options": {
			options: issue_150v1.NewExplicitPriorityOptions().
				WithPriority(temporal.Priority{
					PriorityKey: 1,
				}).
				WithStartWorkflowOptions(client.StartWorkflowOptions{
					Priority: temporal.Priority{
						PriorityKey:    3,
						FairnessKey:    "custom-fairness",
						FairnessWeight: 2.0,
					},
				}),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey: 1,
				},
			},
		},
		"with priority key": {
			// WithPriorityKey overrides only PriorityKey; proto fairness defaults are retained
			options: issue_150v1.NewExplicitPriorityOptions().
				WithPriorityKey(7),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey:    7,
					FairnessKey:    "default",
					FairnessWeight: 1,
				},
			},
		},
		"with fairness key": {
			// WithFairnessKey overrides only FairnessKey; proto priority/weight defaults are retained
			options: issue_150v1.NewExplicitPriorityOptions().
				WithFairnessKey("custom"),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey:    2,
					FairnessKey:    "custom",
					FairnessWeight: 1,
				},
			},
		},
		"with fairness weight": {
			// WithFairnessWeight overrides only FairnessWeight; proto priority/key defaults are retained
			options: issue_150v1.NewExplicitPriorityOptions().
				WithFairnessWeight(3),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey:    2,
					FairnessKey:    "default",
					FairnessWeight: 3,
				},
			},
		},
		"with priority key and start workflow options": {
			// individual setter overrides its field; the base options value for the
			// untouched field is preserved and proto defaults fill the rest
			options: issue_150v1.NewExplicitPriorityOptions().
				WithPriorityKey(9).
				WithStartWorkflowOptions(client.StartWorkflowOptions{
					Priority: temporal.Priority{
						FairnessKey: "base-fairness",
					},
				}),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey:    9,
					FairnessKey:    "base-fairness",
					FairnessWeight: 1,
				},
			},
		},
		"with priority and priority key": {
			// individual setter takes precedence over WithPriority for its field
			options: issue_150v1.NewExplicitPriorityOptions().
				WithPriorityKey(8).
				WithPriority(temporal.Priority{
					PriorityKey: 5,
					FairnessKey: "whole",
				}),
			expected: client.StartWorkflowOptions{
				TaskQueue: "issue-150-v1",
				Priority: temporal.Priority{
					PriorityKey: 8,
					FairnessKey: "whole",
				},
			},
		},
	}

	for _, name := range workflow.DeterministicKeys(cases) {
		c := cases[name]
		t.Run(name, func(t *testing.T) {
			opts, err := c.options.Build((&issue_150v1.Input{}).ProtoReflect())
			require.NoError(t, err)
			require.Equal(t, c.expected, opts)
		})
	}
}
