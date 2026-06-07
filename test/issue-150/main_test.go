package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
	require.Contains(t, explicitZeroBuild, "opts.Priority = temporal.Priority{PriorityKey: 4}")

	unsetBuild := getBuildFunctionBody(t, src, "UnsetPriorityActivityOptions")
	require.NotContains(t, unsetBuild, "opts.Priority = temporal.Priority{PriorityKey: 4}")
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
