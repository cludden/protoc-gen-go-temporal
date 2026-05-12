{{- /*
basic.tpl renders Temporal documentation in one of two modes:

  - Single-file mode: the unnamed root body (below the defined templates)
    iterates every package and emits one README. Used when --docs-out
    resolves to a file path.
  - Directory mode: the plugin invokes "package" once per proto package
    and "index" once for the top-level table of contents. Used when
    --docs-out resolves to a directory.

The "message", "enum", and "package" templates are shared by both modes.
CurrentPackage is set to the rendering package in directory mode so
docslink can produce cross-package relative file links; in single-file
mode it is "" and docslink falls back to bare anchors.
*/ -}}
{{- define "message" -}}
{{- $msg := .Msg -}}
{{- $currentPkg := .CurrentPackage -}}
{{- $data := .Data -}}
{{ if gt (len $msg.Fields) 0 -}}
<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
{{ range $f := $msg.Fields -}}
<tr>
<td>{{ $f.Descriptor.Name }}</td>
<td>{{ if contains "." $f.Type }}<a href="{{ docslink $f.Type "" $currentPkg $data }}">{{ $f.Type }}</a>{{ else }}{{ $f.Type }}{{ end }}</td>
<td>
{{- if or $f.Comments.Leading $f.GoName $f.GoTags $f.JSONName }}<pre>
{{- if $f.Comments.Leading }}
{{ $f.Comments.Leading | trim }}<br>
{{ end }}
{{- if $f.JSONName }}
json_name: {{ $f.JSONName }}{{- end }}
{{- if $f.GoName }}
go_name: {{ $f.GoName }}{{- end }}
{{- if $f.GoTags }}
go_tags: {{ $f.GoTags }}{{- end }}</pre>{{- end }}</td>
</tr>
{{- end }}
</table>
{{- end }}
{{- end -}}




{{- define "enum" -}}
{{ if gt (len .Values) 0 -}}
<table>
<tr><th>Value</th><th>Description</th></tr>
{{ range $v := .Values -}}
<tr>
<td>{{ $v.Name }}</td>
<td>{{ if $v.Comments.Leading }}<pre>
{{ $v.Comments.Leading | trim }}
</pre>{{ end }}</td>
</tr>
{{- end }}
</table>
{{- end }}
{{- end -}}

{{- define "package" }}
{{- $data := .Data }}
{{- $pkgName := .Package }}
{{- $currentPkg := .CurrentPackage }}
{{- $pkg := index $data.Packages $pkgName }}

<a name="{{ $pkgName | slug }}"></a>
# {{ $pkgName }}

{{- if gt (len $pkg.Services) 0 }}

<a name="{{ $pkgName | slug }}-services"></a>
## Services
{{- range $svcI, $svcName := $pkg.Services -}}
{{- $svc := (index $data.Services $svcName) }}
{{- if $svc.HasTemporalResources }}

<a name="{{ $svc.FullName | slug }}"></a>
## {{ $svc.FullName }}

{{- if $svc.Comments.Leading }}

<pre>
{{ $svc.Comments.Leading | trim }}
</pre>
{{- end }}

{{- if (gt (len $svc.Workflows) 0) }}

<a name="{{ $svc.FullName | slug }}-workflows"></a>
### Workflows

{{- range $wI, $wName := $svc.Workflows }}
{{- $w := (index $data.Workflows $wName )}}

---
<a name="{{ $w.Name | slug }}-workflow"></a>
### {{ $w.Name }}

{{- if $w.Comments.Leading }}

<pre>
{{ $w.Comments.Leading | trim }}
</pre>
{{- end }}

{{- if $w.Input }}

**Input:** [{{ $w.Input }}]({{ docslink $w.Input "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $w.Input) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}

{{- if $w.Output }}

**Output:** [{{ $w.Output }}]({{ docslink $w.Output "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $w.Output) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
{{- if gt $w.ExecutionTimeout 0 }}
<tr><td>execution_timeout</td><td>{{ $w.ExecutionTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if $w.ID }}
<tr><td>id</td><td><pre><code>{{ $w.ID }}</code></pre></td></tr>
{{- end }}
<tr><td>id_reuse_policy</td><td><pre><code>{{ $w.IDReusePolicy }}</code></pre></td></tr>
{{- if $w.ParentClosePolicy }}
<tr><td>parent_close_policy</td><td><pre><code>{{ $w.ParentClosePolicy }}</code></pre></td></tr>
{{- end }}
{{- if $w.RetryPolicy }}
{{- if $w.RetryPolicy.BackoffCoefficient }}
<tr><td>retry_policy.backoff_coefficient</td><td>{{ $w.RetryPolicy.BackoffCoefficient }}</td></tr>
{{- end }}
{{- if gt $w.RetryPolicy.InitialInterval 0 }}
<tr><td>retry_policy.initial_interval</td><td>{{ $w.RetryPolicy.InitialInterval | fmtduration }}</td></tr>
{{- end }}
{{- if $w.RetryPolicy.MaxAttempts }}
<tr><td>retry_policy.max_attempts</td><td>{{ $w.RetryPolicy.MaxAttempts }}</td></tr>
{{- end }}
{{- if gt $w.RetryPolicy.MaxInterval 0 }}
<tr><td>retry_policy.max_interval</td><td>{{ $w.RetryPolicy.MaxInterval | fmtduration }}</td></tr>
{{- end }}
{{- if gt (len $w.RetryPolicy.NonRetryableErrorTypes) 0 }}
<tr><td>retry_policy.non_retryable_error_types</td><td>{{ $w.RetryPolicy.NonRetryableErrorTypes | join "," }}</td></tr>
{{- end }}
{{- end }}
{{- if gt $w.RunTimeout 0 }}
<tr><td>run_timeout</td><td>{{ $w.RunTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if $w.SearchAttributes }}
<tr><td>search_attributes</td><td><pre><code>{{ $w.SearchAttributes }}</code></pre></td></tr>
{{- end }}
{{- if $w.TaskQueue }}
<tr><td>task_queue</td><td><pre><code>{{ $w.TaskQueue }}</code></pre></td></tr>
{{- end }}
{{- if gt $w.TaskTimeout 0 }}
<tr><td>task_timeout</td><td>{{ $w.TaskTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if $w.XNS }}
{{- if gt $w.XNS.HeartbeatInterval 0 }}
<tr><td>xns.heartbeat_interval</td><td>{{ $w.XNS.HeartbeatInterval | fmtduration }}</td></tr>
{{- end }}
{{- if gt $w.XNS.HeartbeatTimeout 0 }}
<tr><td>xns.heartbeat_timeout</td><td>{{ $w.XNS.HeartbeatTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if $w.XNS.Name }}
<tr><td>xns.name</td><td>{{ $w.XNS.Name }}</td></tr>
{{- end }}
{{- if $w.XNS.RetryPolicy }}
{{- if $w.XNS.RetryPolicy.BackoffCoefficient }}
<tr><td>xns.retry_policy.backoff_coefficient</td><td>{{ $w.XNS.RetryPolicy.BackoffCoefficient }}</td></tr>
{{- end }}
{{- if gt $w.XNS.RetryPolicy.InitialInterval 0 }}
<tr><td>xns.retry_policy.initial_interval</td><td>{{ $w.XNS.RetryPolicy.InitialInterval | fmtduration }}</td></tr>
{{- end }}
{{- if $w.XNS.RetryPolicy.MaxAttempts }}
<tr><td>xns.retry_policy.max_attempts</td><td>{{ $w.XNS.RetryPolicy.MaxAttempts }}</td></tr>
{{- end }}
{{- if gt $w.XNS.RetryPolicy.MaxInterval 0 }}
<tr><td>xns.retry_policy.max_interval</td><td>{{ $w.XNS.RetryPolicy.MaxInterval | fmtduration }}</td></tr>
{{- end }}
{{- if gt (len $w.XNS.RetryPolicy.NonRetryableErrorTypes) 0 }}
<tr><td>xns.retry_policy.non_retryable_error_types</td><td>{{ $w.XNS.RetryPolicy.NonRetryableErrorTypes | join "," }}</td></tr>
{{- end }}
{{- end }}
{{- if gt $w.XNS.ScheduleToCloseTimeout 0 }}
<tr><td>xns.schedule_to_close_timeout</td><td>{{ $w.XNS.ScheduleToCloseTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if gt $w.XNS.ScheduleToStartTimeout 0 }}
<tr><td>xns.schedule_to_start_timeout</td><td>{{ $w.XNS.ScheduleToStartTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if gt $w.XNS.StartToCloseTimeout 0 }}
<tr><td>xns.start_to_close_timeout</td><td>{{ $w.XNS.StartToCloseTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if $w.XNS.TaskQueue }}
<tr><td>xns.task_queue</td><td>{{ $w.XNS.TaskQueue }}</td></tr>
{{- end }}
{{- end }}
</table>

{{- if gt (len $w.Queries) 0 }}

**Queries:**

<table>
<tr><th>Query</th></tr>
{{- range $q, $opt := $w.Queries }}
<tr><td><a href="{{ docslink $q "-query" $currentPkg $data }}">{{ $q }}</a></td></tr>
{{- end }}
</table>
{{- end }}

{{- if gt (len $w.Signals) 0 }}

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
{{- range $s, $opt := $w.Signals }}
<tr><td><a href="{{ docslink $s "-signal" $currentPkg $data }}">{{ $s }}</a></td><td>{{ $opt.Start }}</td></tr>
{{- end }}
</table>
{{- end }}

{{- if gt (len $w.Updates) 0 }}

**Updates:**

<table>
<tr><th>Update</th></tr>
{{- range $u, $opt := $w.Updates }}
<tr><td><a href="{{ docslink $u "-update" $currentPkg $data }}">{{ $u }}</a></td></tr>
{{- end }}
</table>
{{- end }}
{{- end }} {{/* range workflows */}}
{{- end }} {{/* if workflows > 0 */}}




{{- if (gt (len $svc.Queries) 0) }}

<a name="{{ $svc.FullName | slug }}-queries"></a>
### Queries
{{- range $qI, $qName := $svc.Queries -}}
{{- $q := (index $data.Queries $qName ) }}

---
<a name="{{ $q.Name | slug }}-query"></a>
### {{ $q.Name }}

{{ if $q.Comments.Leading -}}
<pre>
{{ $q.Comments.Leading | trim }}
</pre>
{{- end }}

{{- if $q.Input }}

**Input:** [{{ $q.Input }}]({{ docslink $q.Input "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $q.Input) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}

{{- if $q.Output }}

**Output:** [{{ $q.Output }}]({{ docslink $q.Output "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $q.Output) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}
{{- end }} {{/* range queries */}}
{{- end }} {{/* if queries > 0 */}}

{{- if (gt (len $svc.Signals) 0) }}

<a name="{{ $svc.FullName | slug }}-signals"></a>
### Signals
{{- range $sI, $sName := $svc.Signals -}}
{{- $s := (index $data.Signals $sName )}}

---
<a name="{{ $s.Name | slug }}-signal"></a>
### {{ $s.Name }}

{{ if $s.Comments.Leading -}}
<pre>
{{ $s.Comments.Leading | trim }}
</pre>
{{- end }}

{{- if $s.Input }}

**Input:** [{{ $s.Input }}]({{ docslink $s.Input "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $s.Input) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}
{{- end }} {{/* range signals */}}
{{- end }} {{/* if signals > 0 */}}



{{- if (gt (len $svc.Updates) 0) }}

<a name="{{ $svc.FullName | slug }}-updates"></a>
### Updates
{{- range $uI, $uName := $svc.Updates -}}
{{- $u := (index $data.Updates $uName ) }}

---
<a name="{{ $u.Name | slug }}-update"></a>
### {{ $u.Name }}

{{ if $u.Comments.Leading -}}
<pre>
{{ $u.Comments.Leading | trim }}
</pre>
{{- end }}

{{- if $u.Input }}

**Input:** [{{ $u.Input }}]({{ docslink $u.Input "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $u.Input) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}

{{- if $u.Output }}

**Output:** [{{ $u.Output }}]({{ docslink $u.Output "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $u.Output) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}

{{- if $u.HasNonDefaultOptions }}

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
{{- if $u.Validate }}
<tr><td>validate</td><td>{{ $u.Validate }}</td></tr>
{{- end }}
{{- if $u.WaitPolicy }}
<tr><td>wait_policy</td><td><pre>{{ $u.WaitPolicy }}</pre></td></tr>
{{- end }}
</table>
{{- end }}
{{- end }}
{{- end }}



{{- if (gt (len $svc.Activities) 0) }}

<a name="{{ $svc.FullName | slug }}-activities"></a>
### Activities
{{- range $ai, $aName := $svc.Activities -}}
{{- $a := (index $data.Activities $aName ) }}

---
<a name="{{ $a.Name | slug }}-activity"></a>
### {{ $a.Name }}

{{ if $a.Comments.Leading -}}
<pre>
{{ $a.Comments.Leading | trim }}
</pre>
{{- end }}

{{- if $a.Input }}

**Input:** [{{ $a.Input }}]({{ docslink $a.Input "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $a.Input) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}

{{- if $a.Output }}

**Output:** [{{ $a.Output }}]({{ docslink $a.Output "" $currentPkg $data }})

{{ template "message" (dict "Msg" (index $data.Messages $a.Output) "CurrentPackage" $currentPkg "Data" $data) }}
{{- end }}

{{- if $a.HasNonDefaultOptions }}

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
{{- if gt $a.HeartbeatTimeout 0 }}
<tr><td>heartbeat_timeout</td><td>{{ $a.HeartbeatTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if $a.RetryPolicy }}
{{- if $a.RetryPolicy.BackoffCoefficient }}
<tr><td>retry_policy.backoff_coefficient</td><td>{{ $a.RetryPolicy.BackoffCoefficient }}</td></tr>
{{- end }}
{{- if gt $a.RetryPolicy.InitialInterval 0 }}
<tr><td>retry_policy.initial_interval</td><td>{{ $a.RetryPolicy.InitialInterval | fmtduration }}</td></tr>
{{- end }}
{{- if $a.RetryPolicy.MaxAttempts }}
<tr><td>retry_policy.max_attempts</td><td>{{ $a.RetryPolicy.MaxAttempts }}</td></tr>
{{- end }}
{{- if gt $a.RetryPolicy.MaxInterval 0 }}
<tr><td>retry_policy.max_interval</td><td>{{ $a.RetryPolicy.MaxInterval | fmtduration }}</td></tr>
{{- end }}
{{- if gt (len $a.RetryPolicy.NonRetryableErrorTypes) 0 }}
<tr><td>retry_policy.non_retryable_error_types</td><td>{{ $a.RetryPolicy.NonRetryableErrorTypes | join "," }}</td></tr>
{{- end }}
{{- end }}
{{- if gt $a.ScheduleToCloseTimeout 0 }}
<tr><td>schedule_to_close_timeout</td><td>{{ $a.ScheduleToCloseTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if gt $a.ScheduleToStartTimeout 0 }}
<tr><td>schedule_to_start_timeout</td><td>{{ $a.ScheduleToStartTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if gt $a.StartToCloseTimeout 0 }}
<tr><td>start_to_close_timeout</td><td>{{ $a.StartToCloseTimeout | fmtduration }}</td></tr>
{{- end }}
{{- if $a.WaitForCancellation }}
<tr><td>wait_for_cancellation</td><td>{{ $a.WaitForCancellation }}</td></tr>
{{- end }}
{{- if $a.TaskQueue }}
<tr><td>task_queue</td><td>{{ $a.TaskQueue }}</td></tr>
{{- end }}
</table>
{{- end }} {{/* if $a.HasNonDefaultOptions */}}
{{- end }} {{/* range $pkg.Activities */}}
{{- end }} {{/* if gt len($pkg.Activities) 0 */}}
{{- end }}
{{- end }}
{{- end -}}

{{- if gt (len $pkg.ReferencedMessages) 0 }}

<a name="{{ $pkgName | slug }}-messages"></a>
## Messages
{{- range $msgI, $msgName := $pkg.ReferencedMessages }}

<a name="{{ $msgName | slug }}"></a>
### {{ $msgName }}
{{- if (index $data.Messages $msgName) }}
{{- $msg := index $data.Messages $msgName }}
{{- if $msg.Comments.Leading }}

<pre>
{{ $msg.Comments.Leading | trim }}
</pre>
{{- end }}
{{- if gt (len $msg.Fields) 0 }}

{{ template "message" (dict "Msg" $msg "CurrentPackage" $currentPkg "Data" $data) }}
{{- end -}}
{{- end -}}
{{- if (index $data.Enums $msgName) }}
{{- $enum := index $data.Enums $msgName -}}
{{- if $enum.Comments.Leading }}

<pre>
{{ $enum.Comments.Leading | trim }}
</pre>
{{- end }}

{{ template "enum" $enum }}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- $data := . -}}


# Table of Contents
{{ range $pkgName, $pkg := $data.Packages -}}
{{- if $pkg.HasTemporalResources }}
- [{{ $pkgName }}](#{{ $pkgName | slug }})
  - Services
    {{- range $svcI, $svcName := $pkg.Services }}
    {{- $svc := index $data.Services $svcName }}
    {{- if $svc.HasTemporalResources }}
    - [{{ $svcName }}](#{{ $svcName | slug }})
      {{- if (gt (len $svc.Workflows) 0) }}
      - [Workflows](#{{ $svcName | slug }}-workflows)
        {{- range $wI, $wName := $svc.Workflows }}
        {{- $w := (index $data.Workflows $wName )}}
        - [{{ $w.Name }}](#{{ $w.Name | slug }}-workflow)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Queries) 0) }}
      - [Queries](#{{ $svcName | slug }}-queries)
        {{- range $qI, $qName := $svc.Queries }}
        {{- $q := (index $data.Queries $qName )}}
        - [{{ $q.Name }}](#{{ $q.Name | slug }}-query)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Signals) 0) }}
      - [Signals](#{{ $svcName | slug }}-signals)
        {{- range $sI, $sName := $svc.Signals }}
        {{- $s := (index $data.Signals $sName )}}
        - [{{ $s.Name }}](#{{ $s.Name | slug }}-signal)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Updates) 0) }}
      - [Updates](#{{ $svcName | slug }}-updates)
        {{- range $uI, $uName := $svc.Updates }}
        {{- $u := (index $data.Updates $uName )}}
        - [{{ $u.Name }}](#{{ $u.Name | slug }}-update)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Activities) 0) }}
      - [Activities](#{{ $svcName | slug }}-activities)
        {{- range $aI, $aName := $svc.Activities }}
        {{- $a := (index $data.Activities $aName )}}
        - [{{ $a.Name }}](#{{ $a.Name | slug }}-activity)
        {{- end }}
      {{- end }}
    {{- end }}
    {{- end }}
  {{- if gt (len $pkg.ReferencedMessages) 0 }}
  - Messages
    {{- range $msgI, $msgName := $pkg.ReferencedMessages }}
    {{- $msg := index $data.Messages $msgName }}
    - [{{ $msgName }}](#{{ $msgName | slug }})
    {{- end }}
  {{- end }}
{{- end }} 
{{- end }}
{{- range $pkgName, $pkg := $data.Packages }}
{{- if and (not $pkg.HasTemporalResources) (gt (len $pkg.ReferencedMessages) 0) }}
- [{{ $pkgName }}](#{{ $pkgName | slug }})
  - Messages
    {{- range $msgI, $msgName := $pkg.ReferencedMessages }}
    {{- $msg := index $data.Messages $msgName }}
    - [{{ $msgName }}](#{{ $msgName | slug }})
    {{- end }}
{{- end }}
{{- end }}

{{- range $pkgName, $pkg := $data.Packages -}}
{{ if $pkg.HasTemporalResources -}}
{{ template "package" (dict "Data" $data "Package" $pkgName "CurrentPackage" "") }}
{{- end }}
{{- end -}}

{{- range $pkgName, $pkg := $data.Packages -}}
{{- if and (not $pkg.HasTemporalResources) (gt (len $pkg.ReferencedMessages) 0) }}
{{ template "package" (dict "Data" $data "Package" $pkgName "CurrentPackage" "") }}
{{- end }}
{{- end -}}

{{- define "index" -}}
{{- $data := .Data }}
{{- $filename := .Filename }}
# Documentation
{{ range $pkgName, $pkg := $data.Packages -}}
{{- if $pkg.HasTemporalResources }}
- [{{ $pkgName }}]({{ $pkg.Dir }}/{{ $filename }})
  - Services
    {{- range $svcI, $svcName := $pkg.Services }}
    {{- $svc := index $data.Services $svcName }}
    {{- if $svc.HasTemporalResources }}
    - [{{ $svcName }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $svcName | slug }})
      {{- if (gt (len $svc.Workflows) 0) }}
      - [Workflows]({{ $pkg.Dir }}/{{ $filename }}#{{ $svcName | slug }}-workflows)
        {{- range $wI, $wName := $svc.Workflows }}
        {{- $w := (index $data.Workflows $wName )}}
        - [{{ $w.Name }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $w.Name | slug }}-workflow)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Queries) 0) }}
      - [Queries]({{ $pkg.Dir }}/{{ $filename }}#{{ $svcName | slug }}-queries)
        {{- range $qI, $qName := $svc.Queries }}
        {{- $q := (index $data.Queries $qName )}}
        - [{{ $q.Name }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $q.Name | slug }}-query)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Signals) 0) }}
      - [Signals]({{ $pkg.Dir }}/{{ $filename }}#{{ $svcName | slug }}-signals)
        {{- range $sI, $sName := $svc.Signals }}
        {{- $s := (index $data.Signals $sName )}}
        - [{{ $s.Name }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $s.Name | slug }}-signal)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Updates) 0) }}
      - [Updates]({{ $pkg.Dir }}/{{ $filename }}#{{ $svcName | slug }}-updates)
        {{- range $uI, $uName := $svc.Updates }}
        {{- $u := (index $data.Updates $uName )}}
        - [{{ $u.Name }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $u.Name | slug }}-update)
        {{- end }}
      {{- end }}
      {{- if (gt (len $svc.Activities) 0) }}
      - [Activities]({{ $pkg.Dir }}/{{ $filename }}#{{ $svcName | slug }}-activities)
        {{- range $aI, $aName := $svc.Activities }}
        {{- $a := (index $data.Activities $aName )}}
        - [{{ $a.Name }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $a.Name | slug }}-activity)
        {{- end }}
      {{- end }}
    {{- end }}
    {{- end }}
  {{- if gt (len $pkg.ReferencedMessages) 0 }}
  - Messages
    {{- range $msgI, $msgName := $pkg.ReferencedMessages }}
    - [{{ $msgName }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $msgName | slug }})
    {{- end }}
  {{- end }}
{{- end }}
{{- end }}
{{- range $pkgName, $pkg := $data.Packages }}
{{- if and (not $pkg.HasTemporalResources) (gt (len $pkg.ReferencedMessages) 0) }}
- [{{ $pkgName }}]({{ $pkg.Dir }}/{{ $filename }})
  - Messages
    {{- range $msgI, $msgName := $pkg.ReferencedMessages }}
    - [{{ $msgName }}]({{ $pkg.Dir }}/{{ $filename }}#{{ $msgName | slug }})
    {{- end }}
{{- end }}
{{- end -}}
{{- end -}}