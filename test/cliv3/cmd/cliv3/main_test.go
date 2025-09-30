package main

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/cludden/protoc-gen-go-temporal/gen/test/cliv3"
	"github.com/cludden/protoc-gen-go-temporal/pkg/convert"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCliV3Timestamp(t *testing.T) {
	cases := []struct {
		args     map[string]string
		expected *cliv3.CreateFooInput
		errors   []string
	}{
		{
			args: map[string]string{
				"name":        "test",
				"description": "this is a test foo",
				"expires-at":  "2025-12-01T01:05:03Z",
			},
			expected: cliv3.CreateFooInput_builder{
				Name:        proto.String("test"),
				Description: proto.String("this is a test foo"),
				ExpiresAt: timestamppb.New(
					convert.Must(time.Parse(time.RFC3339, "2025-12-01T01:05:03Z")),
				),
			}.Build(),
		},
	}

	for _, c := range cases {
		var name strings.Builder
		for _, k := range workflow.DeterministicKeys(c.args) {
			v := c.args[k]
			name.WriteString(fmt.Sprintf("--%s %s", k, v))
		}
		t.Run(name.String(), func(t *testing.T) {
			root, err := cliv3.NewExampleServiceCli(cliv3.NewExampleServiceCliOptions())
			require.NoError(t, err)
			var cmd *cli.Command
			for _, subcmd := range root.Commands {
				if subcmd.Name == "create-foo-with-signal-foo" {
					cmd = subcmd
					for k, v := range c.args {
						cmd.Set(k, v)
					}
					break
				}
			}
			if cmd == nil {
				t.Fatalf("create-foo-with-signal-foo command not found")
			}

			actual, err := cliv3.UnmarshalCliFlagsToCreateFooInput(cmd)
			if len(c.errors) > 0 {
				require.Error(t, err)
				for _, e := range c.errors {
					require.ErrorContains(t, err, e)
				}
			} else {
				require.NoError(t, err)
				require.Empty(t, cmp.Diff(c.expected, actual, protocmp.Transform()))
			}
		})
	}
}
