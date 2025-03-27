package plugin

import j "github.com/dave/jennifer/jen"

func (m *Manifest) debugActivity(g *j.Group, msg string, fields ...j.Code) {
	if !m.cfg.EnableDebugLogging {
		return
	}
	fields = append([]j.Code{j.Lit(msg)}, fields...)
	g.Qual(activityPkg, "GetLogger").Call(j.Id("ctx")).Dot("Debug").Call(fields...)
}
