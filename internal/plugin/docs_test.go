package plugin

import (
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/cludden/protoc-gen-go-temporal/internal/plugin/docs"
	"github.com/stretchr/testify/require"
)

func newDocsTestData() *docs.Data {
	return &docs.Data{
		Activities: map[string]docs.Activity{
			"a.v1.Svc.Act": {Descriptor: docs.Descriptor{FullName: "a.v1.Svc.Act", Package: "a.v1"}},
		},
		Enums: map[string]docs.Enum{},
		Messages: map[string]docs.Message{
			"a.v1.Foo": {Descriptor: docs.Descriptor{FullName: "a.v1.Foo", Package: "a.v1"}},
			"b.v1.Bar": {Descriptor: docs.Descriptor{FullName: "b.v1.Bar", Package: "b.v1"}},
			"c.v1.Baz": {Descriptor: docs.Descriptor{FullName: "c.v1.Baz", Package: "c.v1"}},
		},
		Packages: map[string]docs.Package{
			"a.v1": {Dir: "proto/a/v1", Descriptor: docs.Descriptor{Name: "a.v1"}},
			"b.v1": {Dir: "proto/b/v1", Descriptor: docs.Descriptor{Name: "b.v1"}},
			"c.v1": {Dir: "", Descriptor: docs.Descriptor{Name: "c.v1"}},
		},
		Queries:  map[string]docs.Query{},
		Services: map[string]docs.Service{},
		Signals:  map[string]docs.Signal{},
		Updates:  map[string]docs.Update{},
		Workflows: map[string]docs.Workflow{
			"a.v1.Svc.Run": {Descriptor: docs.Descriptor{FullName: "a.v1.Svc.Run", Package: "a.v1"}},
		},
	}
}

func TestIsDocsDirMode(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "out.md")
	require.NoError(t, os.WriteFile(tmpFile, []byte("x"), 0o644))

	cases := []struct {
		name string
		path string
		want bool
	}{
		{"empty", "", false},
		{"trailing slash", "./docs/", true},
		{"file with extension", "./README.md", false},
		{"existing directory", tmpDir, true},
		{"existing file", tmpFile, false},
		{"extensionless non-existent", filepath.Join(tmpDir, "does-not-exist"), false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, isDocsDirMode(tc.path))
		})
	}
}

func TestResolveDocsPackage(t *testing.T) {
	data := newDocsTestData()

	cases := []struct {
		name string
		fqdn string
		want string
	}{
		{"message", "a.v1.Foo", "a.v1"},
		{"workflow", "a.v1.Svc.Run", "a.v1"},
		{"activity", "a.v1.Svc.Act", "a.v1"},
		{"bare package fqdn", "a.v1", "a.v1"},
		{"unknown", "nonexistent.X", ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, resolveDocsPackage(tc.fqdn, data))
		})
	}
}

func TestDocslinkFor(t *testing.T) {
	data := newDocsTestData()

	cases := []struct {
		name           string
		fqdn           string
		suffix         string
		currentPkg     string
		data           *docs.Data
		perPkgFilename string
		want           string
	}{
		{"single-file mode (no currentPkg)", "a.v1.Foo", "", "", data, "README.md", "#a-v1-foo"},
		{"nil data", "a.v1.Foo", "", "a.v1", nil, "README.md", "#a-v1-foo"},
		{"intra-package message", "a.v1.Foo", "", "a.v1", data, "README.md", "#a-v1-foo"},
		{"cross-package peer dirs (default filename)", "b.v1.Bar", "", "a.v1", data, "README.md", "../../b/v1/README.md#b-v1-bar"},
		{"cross-package peer dirs (mdx filename)", "b.v1.Bar", "", "a.v1", data, "README.mdx", "../../b/v1/README.mdx#b-v1-bar"},
		{"target with empty Dir falls back to anchor", "c.v1.Baz", "", "a.v1", data, "README.md", "#c-v1-baz"},
		{"workflow with suffix", "a.v1.Svc.Run", "-workflow", "b.v1", data, "README.md", "../../a/v1/README.md#a-v1-svc-run-workflow"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, docslinkFor(tc.fqdn, tc.suffix, tc.currentPkg, tc.data, tc.perPkgFilename))
		})
	}
}

func TestResolvePkgOutPath(t *testing.T) {
	cases := []struct {
		name     string
		outDir   string
		pkgDir   string
		filename string
		want     string
		wantErr  bool
	}{
		{"simple relative + default filename", "out", "a/v1", "README.md", filepath.Join("out", "a", "v1", "README.md"), false},
		{"simple relative + mdx filename", "out", "a/v1", "README.mdx", filepath.Join("out", "a", "v1", "README.mdx"), false},
		{"dot dir", "out", ".", "README.md", filepath.Join("out", "README.md"), false},
		{"trailing slash is normalized", "out", "a/v1/", "README.md", filepath.Join("out", "a", "v1", "README.md"), false},
		{"traversal rejected", "out", "../escape", "README.md", "", true},
		{"absolute rejected", "out", "/etc", "README.md", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := resolvePkgOutPath(tc.outDir, tc.pkgDir, tc.filename)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestValidateDocsFilename(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"plain README.md", "README.md", false},
		{"docusaurus mdx", "index.mdx", false},
		{"hugo underscore", "_index.md", false},
		{"empty", "", true},
		{"contains forward slash", "docs/README.md", true},
		{"contains backslash", "docs\\README.md", true},
		{"just dot-dot", "..", true},
		{"traversal", "../README.md", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateDocsFilename(tc.input, "docs-filename")
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestRequireDirTemplates(t *testing.T) {
	cases := []struct {
		name    string
		body    string
		wantErr bool
		errFrag string
	}{
		{
			name:    "both defined",
			body:    `{{- define "package" -}}p{{- end -}}{{- define "index" -}}i{{- end -}}`,
			wantErr: false,
		},
		{
			name:    "missing package",
			body:    `{{- define "index" -}}i{{- end -}}`,
			wantErr: true,
			errFrag: "package",
		},
		{
			name:    "missing index",
			body:    `{{- define "package" -}}p{{- end -}}`,
			wantErr: true,
			errFrag: "index",
		},
		{
			name:    "missing both",
			body:    `{{- define "message" -}}m{{- end -}}`,
			wantErr: true,
			errFrag: "package",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tpl, err := template.New("docs").Parse(tc.body)
			require.NoError(t, err)
			err = requireDirTemplates(tpl)
			if tc.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errFrag)
				return
			}
			require.NoError(t, err)
		})
	}
}
