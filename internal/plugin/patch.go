package plugin

import (
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/patch"
	g "github.com/dave/jennifer/jen"
)

var (
	minVersion = g.Qual(workflowPkg, "DefaultVersion")
	maxVersion = g.Lit(1)
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
	temporalv1.Patch_PV_77: {
		id:      patch.PV_77_UseParentTaskQueue,
		comment: "use parent workflow task queue for child workflows and activities",
		link:    "https://cludden.github.io/protoc-gen-go-temporal/docs/guides/patches#pv_77-use-parent-task-queue",
	},
}

func patchComment(b *g.Group, pv temporalv1.Patch_Version) {
	b.Comment(patches[pv].comment)
	b.Commentf("more info: %s", patches[pv].link)
}

func patchVersion(pv temporalv1.Patch_Version, pvm temporalv1.Patch_Mode) g.Code {
	min := minVersion
	if pvm == temporalv1.Patch_PVM_MARKER {
		min = maxVersion
	}
	return g.Qual(workflowPkg, "GetVersion").Call(g.Id("ctx"), g.Lit(patches[pv].id), min, maxVersion)
}
