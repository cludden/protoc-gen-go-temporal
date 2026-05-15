package docs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeOrdering(t *testing.T) {
	data := &Data{
		Activities: map[string]Activity{
			"a.v1.Svc.Zeta":  {Descriptor: Descriptor{FullName: "a.v1.Svc.Zeta"}, Name: "zeta"},
			"a.v1.Svc.Alpha": {Descriptor: Descriptor{FullName: "a.v1.Svc.Alpha"}, Name: "alpha"},
			"a.v1.Svc.Mid":   {Descriptor: Descriptor{FullName: "a.v1.Svc.Mid"}, Name: "mid"},
		},
		Workflows: map[string]Workflow{
			"a.v1.Svc.Run":    {Descriptor: Descriptor{FullName: "a.v1.Svc.Run"}, Name: "zRun"},
			"a.v1.Svc.Create": {Descriptor: Descriptor{FullName: "a.v1.Svc.Create"}, Name: "aCreate"},
		},
		Queries: map[string]Query{
			"a.v1.Svc.GetThing": {Descriptor: Descriptor{FullName: "a.v1.Svc.GetThing"}, Name: "getThing"},
			"a.v1.Svc.GetOther": {Descriptor: Descriptor{FullName: "a.v1.Svc.GetOther"}, Name: "abcOther"},
		},
		Signals: map[string]Signal{
			"a.v1.Svc.Notify": {Descriptor: Descriptor{FullName: "a.v1.Svc.Notify"}, Name: "yNotify"},
			"a.v1.Svc.Ack":    {Descriptor: Descriptor{FullName: "a.v1.Svc.Ack"}, Name: "aAck"},
		},
		Updates: map[string]Update{
			"a.v1.Svc.UpdateB": {Descriptor: Descriptor{FullName: "a.v1.Svc.UpdateB"}, Name: "bUpdate"},
			"a.v1.Svc.UpdateA": {Descriptor: Descriptor{FullName: "a.v1.Svc.UpdateA"}, Name: "aUpdate"},
		},
		Services: map[string]Service{
			"a.v1.Svc": {
				Descriptor: Descriptor{FullName: "a.v1.Svc"},
				Workflows:  []string{"a.v1.Svc.Run", "a.v1.Svc.Create"},
				Queries:    []string{"a.v1.Svc.GetThing", "a.v1.Svc.GetOther"},
				Signals:    []string{"a.v1.Svc.Notify", "a.v1.Svc.Ack"},
				Updates:    []string{"a.v1.Svc.UpdateB", "a.v1.Svc.UpdateA"},
				Activities: []string{"a.v1.Svc.Zeta", "a.v1.Svc.Alpha", "a.v1.Svc.Mid"},
			},
		},
		Packages: map[string]Package{
			"a.v1": {
				Descriptor: Descriptor{Name: "a.v1"},
				Services:   []string{"a.v1.Zulu", "a.v1.Alpha", "a.v1.Mike"},
				Messages:   []string{"a.v1.Z", "a.v1.A", "a.v1.M"},
				Enums:      []string{"a.v1.Z", "a.v1.A"},
			},
		},
	}

	normalizeOrdering(data)

	svc := data.Services["a.v1.Svc"]
	require.Equal(t, []string{"a.v1.Svc.Create", "a.v1.Svc.Run"}, svc.Workflows, "workflows by display Name")
	require.Equal(t, []string{"a.v1.Svc.GetOther", "a.v1.Svc.GetThing"}, svc.Queries, "queries by display Name")
	require.Equal(t, []string{"a.v1.Svc.Ack", "a.v1.Svc.Notify"}, svc.Signals, "signals by display Name")
	require.Equal(t, []string{"a.v1.Svc.UpdateA", "a.v1.Svc.UpdateB"}, svc.Updates, "updates by display Name")
	require.Equal(t, []string{"a.v1.Svc.Alpha", "a.v1.Svc.Mid", "a.v1.Svc.Zeta"}, svc.Activities, "activities by display Name")

	pkg := data.Packages["a.v1"]
	require.Equal(t, []string{"a.v1.Alpha", "a.v1.Mike", "a.v1.Zulu"}, pkg.Services, "services by FullName")
	require.Equal(t, []string{"a.v1.A", "a.v1.M", "a.v1.Z"}, pkg.Messages, "messages by FullName")
	require.Equal(t, []string{"a.v1.A", "a.v1.Z"}, pkg.Enums, "enums by FullName")
}
