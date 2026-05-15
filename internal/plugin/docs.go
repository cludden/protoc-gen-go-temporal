package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/cludden/protoc-gen-go-temporal/internal/plugin/docs"
	"github.com/cludden/protoc-gen-go-temporal/internal/templates"
	"github.com/gosimple/slug"
	"github.com/hako/durafmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func renderDocs(p *protogen.Plugin, cfg *Config) error {
	dirMode := isDocsDirMode(cfg.DocsOut)

	perPkgFilename := cfg.DocsFilename
	indexFilename := cfg.DocsIndex
	if dirMode {
		if err := validateDocsFilename(perPkgFilename, "docs-filename"); err != nil {
			return err
		}
		if err := validateDocsFilename(indexFilename, "docs-index"); err != nil {
			return err
		}
	}
	if perPkgFilename == "" {
		perPkgFilename = "README.md"
	}
	if indexFilename == "" {
		indexFilename = "README.md"
	}

	data, err := docs.Parse(p)
	if err != nil {
		return fmt.Errorf("error extracting template data: %v", err)
	}

	var raw []byte
	switch cfg.DocsTemplate {
	case "basic":
		raw, err = templates.Templates.ReadFile("basic.tpl")
	default:
		raw, err = os.ReadFile(cfg.DocsTemplate)
	}
	if err != nil {
		return fmt.Errorf("error reading template: %v", err)
	}

	tpl, err := template.New("docs").
		Funcs(template.FuncMap{
			"fmtduration": func(d time.Duration) string {
				return durafmt.Parse(d).String()
			},
			"slug": func(input string) string {
				return slug.Make(input)
			},
			"docslink": func(fqdn, suffix, currentPkg string, d *docs.Data) string {
				return docslinkFor(fqdn, suffix, currentPkg, d, perPkgFilename)
			},
			"mapkeys": sortedMapKeys,
		}).
		Funcs(sprig.FuncMap()).
		Parse(string(raw))
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	if dirMode {
		if err := requireDirTemplates(tpl); err != nil {
			return err
		}
		return renderDocsDir(tpl, data, cfg.DocsOut, perPkgFilename, indexFilename)
	}
	return renderDocsFile(tpl, data, cfg.DocsOut)
}

// validateDocsFilename ensures a filename is a flat name (no path separators,
// no traversal, non-empty). Used for the directory-mode filename options.
func validateDocsFilename(name, opt string) error {
	if name == "" {
		return fmt.Errorf("--%s must not be empty", opt)
	}
	if strings.ContainsAny(name, `/\`) {
		return fmt.Errorf("--%s %q must not contain path separators", opt, name)
	}
	if name == ".." || strings.Contains(name, "..") {
		return fmt.Errorf("--%s %q must not contain %q", opt, name, "..")
	}
	return nil
}

// requireDirTemplates verifies that the loaded template defines both the
// "package" and "index" sub-templates needed by directory mode.
func requireDirTemplates(tpl *template.Template) error {
	missing := []string{}
	if tpl.Lookup("package") == nil {
		missing = append(missing, "package")
	}
	if tpl.Lookup("index") == nil {
		missing = append(missing, "index")
	}
	if len(missing) == 0 {
		return nil
	}
	defined := []string{}
	for _, t := range tpl.Templates() {
		if name := t.Name(); name != "" && name != "docs" {
			defined = append(defined, strconv.Quote(name))
		}
	}
	sort.Strings(defined)
	return fmt.Errorf(
		"docs template must define %s sub-template(s) for directory mode (defined: %s)",
		strings.Join(missing, " and "),
		strings.Join(defined, ", "),
	)
}

func renderDocsFile(tpl *template.Template, data *docs.Data, outPath string) (err error) {
	out, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("error creating doc file: %w", err)
	}
	defer func() {
		if cerr := out.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("error closing doc file %q: %w", outPath, cerr)
		}
	}()
	if err = tpl.Execute(out, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	return nil
}

func renderDocsDir(tpl *template.Template, data *docs.Data, outDir, perPkgFilename, indexFilename string) error {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	pkgNames := make([]string, 0, len(data.Packages))
	for name := range data.Packages {
		pkgNames = append(pkgNames, name)
	}
	sort.Strings(pkgNames)

	written := make(map[string]string, len(pkgNames))
	for _, pkgName := range pkgNames {
		pkg := data.Packages[pkgName]
		if !pkg.HasTemporalResources && len(pkg.ReferencedMessages) == 0 {
			continue
		}
		if pkg.Dir == "" {
			continue
		}
		outPath, err := resolvePkgOutPath(outDir, pkg.Dir, perPkgFilename)
		if err != nil {
			return fmt.Errorf("package %q: %w", pkgName, err)
		}
		if prev, ok := written[outPath]; ok {
			return fmt.Errorf("packages %q and %q both resolve to %q; cannot generate per-package docs", prev, pkgName, outPath)
		}
		written[outPath] = pkgName
		if err := renderPackageDoc(tpl, data, pkgName, outPath); err != nil {
			return err
		}
	}

	indexPath := filepath.Join(outDir, indexFilename)
	if prev, ok := written[indexPath]; ok {
		return fmt.Errorf("package %q collides with the index file at %q", prev, indexPath)
	}
	return renderIndexDoc(tpl, data, indexPath, perPkgFilename)
}

func renderPackageDoc(tpl *template.Template, data *docs.Data, pkgName, outPath string) (err error) {
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("error creating package output dir: %w", err)
	}
	out, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("error creating package doc file: %w", err)
	}
	defer func() {
		if cerr := out.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("error closing package doc file %q: %w", outPath, cerr)
		}
	}()
	if err = tpl.ExecuteTemplate(out, "package", map[string]any{
		"Data":           data,
		"Package":        pkgName,
		"CurrentPackage": pkgName,
	}); err != nil {
		return fmt.Errorf("error executing package template for %q: %v", pkgName, err)
	}
	return nil
}

func renderIndexDoc(tpl *template.Template, data *docs.Data, outPath, perPkgFilename string) (err error) {
	out, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("error creating index file: %w", err)
	}
	defer func() {
		if cerr := out.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("error closing index file %q: %w", outPath, cerr)
		}
	}()
	if err = tpl.ExecuteTemplate(out, "index", map[string]any{
		"Data":     data,
		"Filename": perPkgFilename,
	}); err != nil {
		return fmt.Errorf("error executing index template: %v", err)
	}
	return nil
}

// resolvePkgOutPath joins outDir, a package-relative directory, and a flat
// filename, rejecting absolute paths and any ".." element that would escape
// outDir.
func resolvePkgOutPath(outDir, pkgDir, filename string) (string, error) {
	clean := filepath.Clean(filepath.FromSlash(pkgDir))
	if filepath.IsAbs(clean) {
		return "", fmt.Errorf("package dir %q is absolute", pkgDir)
	}
	if clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("package dir %q escapes output directory", pkgDir)
	}
	return filepath.Join(outDir, clean, filename), nil
}

// isDocsDirMode returns true when the configured --docs-out path should be
// treated as a directory (emit one file per proto package) rather than a
// single output file. Directory mode requires an explicit signal: either a
// trailing path separator or an existing directory at that path.
func isDocsDirMode(p string) bool {
	if p == "" {
		return false
	}
	if strings.HasSuffix(p, "/") || strings.HasSuffix(p, string(filepath.Separator)) {
		return true
	}
	if info, err := os.Stat(p); err == nil && info.IsDir() {
		return true
	}
	return false
}

// docslinkFor resolves a markdown link href for an item identified by fqdn,
// using perPkgFilename for the target file in cross-package references. In
// single-file mode (empty currentPkg) and for intra-package refs it returns
// a bare anchor; cross-package refs in directory mode produce a relative
// path to the target package's per-package file.
func docslinkFor(fqdn, suffix, currentPkg string, data *docs.Data, perPkgFilename string) string {
	anchor := slug.Make(fqdn) + suffix
	if currentPkg == "" || data == nil {
		return "#" + anchor
	}
	targetPkg := resolveDocsPackage(fqdn, data)
	if targetPkg == "" || targetPkg == currentPkg {
		return "#" + anchor
	}
	curr, currOK := data.Packages[currentPkg]
	tgt, tgtOK := data.Packages[targetPkg]
	if !currOK || !tgtOK || curr.Dir == "" || tgt.Dir == "" {
		return "#" + anchor
	}
	rel, err := filepath.Rel(curr.Dir, tgt.Dir)
	if err != nil {
		return "#" + anchor
	}
	return filepath.ToSlash(filepath.Join(rel, perPkgFilename)) + "#" + anchor
}

// sortedMapKeys returns the keys of any string-keyed map, sorted lexicographically.
// It exists because sprig's "keys" only accepts map[string]interface{}, while the
// docs templates iterate typed maps like map[string]docs.Package.
func sortedMapKeys(m any) []string {
	v := reflect.ValueOf(m)
	if !v.IsValid() || v.Kind() != reflect.Map {
		return nil
	}
	out := make([]string, 0, v.Len())
	for _, k := range v.MapKeys() {
		if k.Kind() != reflect.String {
			return nil
		}
		out = append(out, k.String())
	}
	sort.Strings(out)
	return out
}

func resolveDocsPackage(fqdn string, data *docs.Data) string {
	if v, ok := data.Messages[fqdn]; ok {
		return v.Package
	}
	if v, ok := data.Enums[fqdn]; ok {
		return v.Package
	}
	if v, ok := data.Workflows[fqdn]; ok {
		return v.Package
	}
	if v, ok := data.Queries[fqdn]; ok {
		return v.Package
	}
	if v, ok := data.Signals[fqdn]; ok {
		return v.Package
	}
	if v, ok := data.Updates[fqdn]; ok {
		return v.Package
	}
	if v, ok := data.Activities[fqdn]; ok {
		return v.Package
	}
	if _, ok := data.Packages[fqdn]; ok {
		return fqdn
	}
	return ""
}
