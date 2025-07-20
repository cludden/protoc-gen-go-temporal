package plugin

import (
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/patch"
	j "github.com/dave/jennifer/jen"
)

var (
	minVersion = j.Qual(workflowPkg, "DefaultVersion")
	maxVersion = j.Lit(1)
)

type patchDetail struct {
	id      string
	comment string
	link    string
}

var patches = map[temporalv1.Patch_Version]patchDetail{
	temporalv1.Patch_PV_64: {
		id:      patch.PV_64_ExpressionEvaluationLocalActivity,
		comment: "wrap expression evaluation in local activity",
		link:    "https://cludden.github.io/protoc-gen-go-temporal/docs/guides/patches#pv_64-expression-evaluation-local-activity",
	},
}

func patchComment(g *j.Group, pv temporalv1.Patch_Version) {
	g.Comment(patches[pv].comment)
	g.Commentf("more info: %s", patches[pv].link)
}

func patchVersion(pv temporalv1.Patch_Version, pvm temporalv1.Patch_Mode) j.Code {
	min := minVersion
	if pvm == temporalv1.Patch_PVM_MARKER {
		min = maxVersion
	}
	return j.Qual(workflowPkg, "GetVersion").Call(j.Id("ctx"), j.Lit(patches[pv].id), min, maxVersion)
}
