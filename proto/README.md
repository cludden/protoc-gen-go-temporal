# Table of Contents

- [mycompany.simple](#mycompany-simple)
  - Services
    - [mycompany.simple.Simple](#mycompany-simple-simple)
      - [Workflows](#mycompany-simple-simple-workflows)
        - [mycompany.simple.SomeWorkflow1](#mycompany-simple-someworkflow1-workflow)
        - [mycompany.simple.SomeWorkflow2](#mycompany-simple-someworkflow2-workflow)
        - [mycompany.simple.Simple.SomeWorkflow3](#mycompany-simple-simple-someworkflow3-workflow)
        - [mycompany.simple.Simple.SomeWorkflow4](#mycompany-simple-simple-someworkflow4-workflow)
      - [Queries](#mycompany-simple-simple-queries)
        - [mycompany.simple.Simple.SomeQuery1](#mycompany-simple-simple-somequery1-query)
        - [mycompany.simple.Simple.SomeQuery2](#mycompany-simple-simple-somequery2-query)
      - [Signals](#mycompany-simple-simple-signals)
        - [mycompany.simple.Simple.SomeSignal1](#mycompany-simple-simple-somesignal1-signal)
        - [mycompany.simple.Simple.SomeSignal2](#mycompany-simple-simple-somesignal2-signal)
        - [mycompany.simple.Simple.SomeSignal3](#mycompany-simple-simple-somesignal3-signal)
      - [Updates](#mycompany-simple-simple-updates)
        - [mycompany.simple.Simple.SomeUpdate1](#mycompany-simple-simple-someupdate1-update)
        - [mycompany.simple.Simple.SomeUpdate2](#mycompany-simple-simple-someupdate2-update)
      - [Activities](#mycompany-simple-simple-activities)
        - [mycompany.simple.SomeActivity1](#mycompany-simple-someactivity1-activity)
        - [mycompany.simple.Simple.SomeActivity2](#mycompany-simple-simple-someactivity2-activity)
        - [mycompany.simple.Simple.SomeActivity3](#mycompany-simple-simple-someactivity3-activity)
        - [mycompany.simple.Simple.SomeActivity4](#mycompany-simple-simple-someactivity4-activity)
        - [mycompany.simple.Simple.SomeSignal1](#mycompany-simple-simple-somesignal1-activity)
        - [mycompany.simple.Simple.SomeSignal2](#mycompany-simple-simple-somesignal2-activity)
        - [mycompany.simple.Simple.SomeSignal3](#mycompany-simple-simple-somesignal3-activity)
        - [mycompany.simple.Simple.SomeUpdate1](#mycompany-simple-simple-someupdate1-activity)
    - [mycompany.simple.Other](#mycompany-simple-other)
      - [Workflows](#mycompany-simple-other-workflows)
        - [mycompany.simple.Other.OtherWorkflow](#mycompany-simple-other-otherworkflow-workflow)
        - [mycompany.simple.Other.OtherWorkflow2](#mycompany-simple-other-otherworkflow2-workflow)
      - [Queries](#mycompany-simple-other-queries)
        - [mycompany.simple.Other.OtherQuery](#mycompany-simple-other-otherquery-query)
      - [Signals](#mycompany-simple-other-signals)
        - [mycompany.simple.Other.OtherSignal](#mycompany-simple-other-othersignal-signal)
      - [Updates](#mycompany-simple-other-updates)
        - [mycompany.simple.Other.OtherUpdate](#mycompany-simple-other-otherupdate-update)
      - [Activities](#mycompany-simple-other-activities)
        - [mycompany.simple.Other.OtherWorkflow](#mycompany-simple-other-otherworkflow-activity)
    - [mycompany.simple.Ignored](#mycompany-simple-ignored)
      - [Workflows](#mycompany-simple-ignored-workflows)
        - [mycompany.simple.Ignored.What](#mycompany-simple-ignored-what-workflow)
    - [mycompany.simple.OnlyActivities](#mycompany-simple-onlyactivities)
      - [Activities](#mycompany-simple-onlyactivities-activities)
        - [mycompany.simple.OnlyActivities.LonelyActivity1](#mycompany-simple-onlyactivities-lonelyactivity1-activity)
    - [mycompany.simple.Deprecated](#mycompany-simple-deprecated)
      - [Workflows](#mycompany-simple-deprecated-workflows)
        - [mycompany.simple.Deprecated.SomeDeprecatedWorkflow1](#mycompany-simple-deprecated-somedeprecatedworkflow1-workflow)
        - [mycompany.simple.Deprecated.SomeDeprecatedWorkflow2](#mycompany-simple-deprecated-somedeprecatedworkflow2-workflow)
      - [Queries](#mycompany-simple-deprecated-queries)
        - [mycompany.simple.Deprecated.SomeDeprecatedQuery1](#mycompany-simple-deprecated-somedeprecatedquery1-query)
        - [mycompany.simple.Deprecated.SomeDeprecatedQuery2](#mycompany-simple-deprecated-somedeprecatedquery2-query)
      - [Signals](#mycompany-simple-deprecated-signals)
        - [mycompany.simple.Deprecated.SomeDeprecatedSignal1](#mycompany-simple-deprecated-somedeprecatedsignal1-signal)
        - [mycompany.simple.Deprecated.SomeDeprecatedSignal2](#mycompany-simple-deprecated-somedeprecatedsignal2-signal)
      - [Updates](#mycompany-simple-deprecated-updates)
        - [mycompany.simple.Deprecated.SomeDeprecatedUpdate1](#mycompany-simple-deprecated-somedeprecatedupdate1-update)
        - [mycompany.simple.Deprecated.SomeDeprecatedUpdate2](#mycompany-simple-deprecated-somedeprecatedupdate2-update)
      - [Activities](#mycompany-simple-deprecated-activities)
        - [mycompany.simple.Deprecated.SomeDeprecatedActivity1](#mycompany-simple-deprecated-somedeprecatedactivity1-activity)
        - [mycompany.simple.Deprecated.SomeDeprecatedActivity2](#mycompany-simple-deprecated-somedeprecatedactivity2-activity)
  - Messages
    - [mycompany.simple.Foo](#mycompany-simple-foo)
    - [mycompany.simple.LonelyActivity1Request](#mycompany-simple-lonelyactivity1request)
    - [mycompany.simple.LonelyActivity1Response](#mycompany-simple-lonelyactivity1response)
    - [mycompany.simple.OtherEnum](#mycompany-simple-otherenum)
    - [mycompany.simple.OtherQueryResponse](#mycompany-simple-otherqueryresponse)
    - [mycompany.simple.OtherSignalRequest](#mycompany-simple-othersignalrequest)
    - [mycompany.simple.OtherUpdateRequest](#mycompany-simple-otherupdaterequest)
    - [mycompany.simple.OtherUpdateResponse](#mycompany-simple-otherupdateresponse)
    - [mycompany.simple.OtherWorkflowRequest](#mycompany-simple-otherworkflowrequest)
    - [mycompany.simple.OtherWorkflowRequest.Bar](#mycompany-simple-otherworkflowrequest-bar)
    - [mycompany.simple.OtherWorkflowRequest.Baz](#mycompany-simple-otherworkflowrequest-baz)
    - [mycompany.simple.OtherWorkflowResponse](#mycompany-simple-otherworkflowresponse)
    - [mycompany.simple.Qux](#mycompany-simple-qux)
    - [mycompany.simple.SomeActivity2Request](#mycompany-simple-someactivity2request)
    - [mycompany.simple.SomeActivity3Request](#mycompany-simple-someactivity3request)
    - [mycompany.simple.SomeActivity3Response](#mycompany-simple-someactivity3response)
    - [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)
    - [mycompany.simple.SomeQuery1Response](#mycompany-simple-somequery1response)
    - [mycompany.simple.SomeQuery2Request](#mycompany-simple-somequery2request)
    - [mycompany.simple.SomeQuery2Response](#mycompany-simple-somequery2response)
    - [mycompany.simple.SomeSignal2Request](#mycompany-simple-somesignal2request)
    - [mycompany.simple.SomeSignal3Request](#mycompany-simple-somesignal3request)
    - [mycompany.simple.SomeSignal3Response](#mycompany-simple-somesignal3response)
    - [mycompany.simple.SomeUpdate1Request](#mycompany-simple-someupdate1request)
    - [mycompany.simple.SomeUpdate1Response](#mycompany-simple-someupdate1response)
    - [mycompany.simple.SomeUpdate2Request](#mycompany-simple-someupdate2request)
    - [mycompany.simple.SomeUpdate2Response](#mycompany-simple-someupdate2response)
    - [mycompany.simple.SomeWorkflow1Request](#mycompany-simple-someworkflow1request)
    - [mycompany.simple.SomeWorkflow1Response](#mycompany-simple-someworkflow1response)
    - [mycompany.simple.SomeWorkflow3Request](#mycompany-simple-someworkflow3request)
    - [mycompany.simple.WhatRequest](#mycompany-simple-whatrequest)
- [test.acronym.v1](#test-acronym-v1)
  - Services
    - [test.acronym.v1.AWS](#test-acronym-v1-aws)
      - [Workflows](#test-acronym-v1-aws-workflows)
        - [test.acronym.v1.AWS.ManageAWS](#test-acronym-v1-aws-manageaws-workflow)
        - [test.acronym.v1.AWS.ManageAWSResource](#test-acronym-v1-aws-manageawsresource-workflow)
        - [test.acronym.v1.AWS.SomethingV1FooBar](#test-acronym-v1-aws-somethingv1foobar-workflow)
        - [test.acronym.v1.AWS.SomethingV2FooBar](#test-acronym-v1-aws-somethingv2foobar-workflow)
      - [Activities](#test-acronym-v1-aws-activities)
        - [test.acronym.v1.AWS.ManageAWSResource](#test-acronym-v1-aws-manageawsresource-activity)
        - [test.acronym.v1.AWS.ManageAWSResourceURN](#test-acronym-v1-aws-manageawsresourceurn-activity)
  - Messages
    - [test.acronym.v1.ManageAWSRequest](#test-acronym-v1-manageawsrequest)
    - [test.acronym.v1.ManageAWSResourceRequest](#test-acronym-v1-manageawsresourcerequest)
    - [test.acronym.v1.ManageAWSResourceResponse](#test-acronym-v1-manageawsresourceresponse)
    - [test.acronym.v1.ManageAWSResourceURNRequest](#test-acronym-v1-manageawsresourceurnrequest)
    - [test.acronym.v1.ManageAWSResourceURNResponse](#test-acronym-v1-manageawsresourceurnresponse)
    - [test.acronym.v1.ManageAWSResponse](#test-acronym-v1-manageawsresponse)
    - [test.acronym.v1.SomethingV1FooBarRequest](#test-acronym-v1-somethingv1foobarrequest)
    - [test.acronym.v1.SomethingV1FooBarResponse](#test-acronym-v1-somethingv1foobarresponse)
    - [test.acronym.v1.SomethingV2FooBarRequest](#test-acronym-v1-somethingv2foobarrequest)
    - [test.acronym.v1.SomethingV2FooBarResponse](#test-acronym-v1-somethingv2foobarresponse)
- [test.activity.v1](#test-activity-v1)
  - Services
    - [test.activity.v1.Example](#test-activity-v1-example)
      - [Activities](#test-activity-v1-example-activities)
        - [test.activity.v1.Example.Foo](#test-activity-v1-example-foo-activity)
        - [test.activity.v1.Example.Bar](#test-activity-v1-example-bar-activity)
        - [test.activity.v1.Example.Baz](#test-activity-v1-example-baz-activity)
        - [test.activity.v1.Example.Qux](#test-activity-v1-example-qux-activity)
  - Messages
    - [test.activity.v1.BarInput](#test-activity-v1-barinput)
    - [test.activity.v1.BazOutput](#test-activity-v1-bazoutput)
    - [test.activity.v1.FooInput](#test-activity-v1-fooinput)
    - [test.activity.v1.FooOutput](#test-activity-v1-foooutput)
- [test.editions](#test-editions)
  - Services
    - [test.editions.FooService](#test-editions-fooservice)
      - [Workflows](#test-editions-fooservice-workflows)
        - [test.editions.FooService.Foo](#test-editions-fooservice-foo-workflow)
      - [Activities](#test-editions-fooservice-activities)
        - [test.editions.FooService.Foo](#test-editions-fooservice-foo-activity)
  - Messages
    - [test.editions.FooInput](#test-editions-fooinput)
    - [test.editions.FooOutput](#test-editions-foooutput)
- [test.opaque](#test-opaque)
  - Services
    - [test.opaque.Hybrid](#test-opaque-hybrid)
      - [Workflows](#test-opaque-hybrid-workflows)
        - [test.opaque.Hybrid.PutHybridExample](#test-opaque-hybrid-puthybridexample-workflow)
      - [Signals](#test-opaque-hybrid-signals)
        - [test.opaque.Hybrid.SignalHybrid](#test-opaque-hybrid-signalhybrid-signal)
    - [test.opaque.Opaque](#test-opaque-opaque)
      - [Workflows](#test-opaque-opaque-workflows)
        - [test.opaque.Opaque.PutOpaqueExample](#test-opaque-opaque-putopaqueexample-workflow)
      - [Signals](#test-opaque-opaque-signals)
        - [test.opaque.Opaque.SignalOpaque](#test-opaque-opaque-signalopaque-signal)
    - [test.opaque.Open](#test-opaque-open)
      - [Workflows](#test-opaque-open-workflows)
        - [test.opaque.Open.PutOpenExample](#test-opaque-open-putopenexample-workflow)
      - [Signals](#test-opaque-open-signals)
        - [test.opaque.Open.SignalOpen](#test-opaque-open-signalopen-signal)
    - [test.opaque.Optional](#test-opaque-optional)
      - [Workflows](#test-opaque-optional-workflows)
        - [test.opaque.Optional.PutOptionalExample](#test-opaque-optional-putoptionalexample-workflow)
      - [Signals](#test-opaque-optional-signals)
        - [test.opaque.Optional.SignalOptional](#test-opaque-optional-signaloptional-signal)
  - Messages
    - [test.opaque.Address](#test-opaque-address)
    - [test.opaque.HybridExample](#test-opaque-hybridexample)
    - [test.opaque.HybridExample.ExtraEntry](#test-opaque-hybridexample-extraentry)
    - [test.opaque.OpaqueExample](#test-opaque-opaqueexample)
    - [test.opaque.OpaqueExample.ExtraEntry](#test-opaque-opaqueexample-extraentry)
    - [test.opaque.OpenExample](#test-opaque-openexample)
    - [test.opaque.OpenExample.ExtraEntry](#test-opaque-openexample-extraentry)
    - [test.opaque.OptionalExample](#test-opaque-optionalexample)
    - [test.opaque.OptionalExample.ExtraEntry](#test-opaque-optionalexample-extraentry)
    - [test.opaque.Status](#test-opaque-status)
- [test.option.v1](#test-option-v1)
  - Services
    - [test.option.v1.Test](#test-option-v1-test)
      - [Workflows](#test-option-v1-test-workflows)
        - [test.option.v1.Test.WorkflowWithInput](#test-option-v1-test-workflowwithinput-workflow)
      - [Updates](#test-option-v1-test-updates)
        - [test.option.v1.Test.UpdateWithInput](#test-option-v1-test-updatewithinput-update)
      - [Activities](#test-option-v1-test-activities)
        - [test.option.v1.Test.ActivityWithInput](#test-option-v1-test-activitywithinput-activity)
  - Messages
    - [test.option.v1.ActivityWithInputRequest](#test-option-v1-activitywithinputrequest)
    - [test.option.v1.ActivityWithInputResponse](#test-option-v1-activitywithinputresponse)
    - [test.option.v1.UpdateWithInputRequest](#test-option-v1-updatewithinputrequest)
    - [test.option.v1.WorkflowWithInputRequest](#test-option-v1-workflowwithinputrequest)
- [test.proto3optional](#test-proto3optional)
  - Services
    - [test.proto3optional.FooService](#test-proto3optional-fooservice)
      - [Workflows](#test-proto3optional-fooservice-workflows)
        - [test.proto3optional.FooService.Foo](#test-proto3optional-fooservice-foo-workflow)
      - [Activities](#test-proto3optional-fooservice-activities)
        - [test.proto3optional.FooService.Foo](#test-proto3optional-fooservice-foo-activity)
  - Messages
    - [test.proto3optional.FooInput](#test-proto3optional-fooinput)
    - [test.proto3optional.FooOutput](#test-proto3optional-foooutput)
- [test.xnserr.v1](#test-xnserr-v1)
  - Services
    - [test.xnserr.v1.Server](#test-xnserr-v1-server)
      - [Workflows](#test-xnserr-v1-server-workflows)
        - [test.xnserr.v1.Server.Sleep](#test-xnserr-v1-server-sleep-workflow)
    - [test.xnserr.v1.Client](#test-xnserr-v1-client)
      - [Workflows](#test-xnserr-v1-client-workflows)
        - [test.xnserr.v1.Client.CallSleep](#test-xnserr-v1-client-callsleep-workflow)
  - Messages
    - [test.xnserr.v1.CallSleepRequest](#test-xnserr-v1-callsleeprequest)
    - [test.xnserr.v1.Failure](#test-xnserr-v1-failure)
    - [test.xnserr.v1.FailureInfo](#test-xnserr-v1-failureinfo)
    - [test.xnserr.v1.SleepRequest](#test-xnserr-v1-sleeprequest)
- [google.protobuf](#google-protobuf)
  - Messages
    - [google.protobuf.Any](#google-protobuf-any)
    - [google.protobuf.Duration](#google-protobuf-duration)
    - [google.protobuf.ListValue](#google-protobuf-listvalue)
    - [google.protobuf.NullValue](#google-protobuf-nullvalue)
    - [google.protobuf.Struct](#google-protobuf-struct)
    - [google.protobuf.Struct.FieldsEntry](#google-protobuf-struct-fieldsentry)
    - [google.protobuf.Timestamp](#google-protobuf-timestamp)
    - [google.protobuf.Value](#google-protobuf-value)
- [mycompany.simple.common.v1](#mycompany-simple-common-v1)
  - Messages
    - [mycompany.simple.common.v1.Example](#mycompany-simple-common-v1-example)
    - [mycompany.simple.common.v1.PaginatedRequest](#mycompany-simple-common-v1-paginatedrequest)
    - [mycompany.simple.common.v1.PaginatedResponse](#mycompany-simple-common-v1-paginatedresponse)
- [temporal.api.enums.v1](#temporal-api-enums-v1)
  - Messages
    - [temporal.api.enums.v1.WorkflowIdConflictPolicy](#temporal-api-enums-v1-workflowidconflictpolicy)
- [temporal.xns.v1](#temporal-xns-v1)
  - Messages
    - [temporal.xns.v1.IDReusePolicy](#temporal-xns-v1-idreusepolicy)
    - [temporal.xns.v1.RetryPolicy](#temporal-xns-v1-retrypolicy)
    - [temporal.xns.v1.StartWorkflowOptions](#temporal-xns-v1-startworkflowoptions)

<a name="mycompany-simple"></a>
# mycompany.simple

<a name="mycompany-simple-services"></a>
## Services

<a name="mycompany-simple-simple"></a>
## mycompany.simple.Simple

<a name="mycompany-simple-simple-workflows"></a>
### Workflows

---
<a name="mycompany-simple-someworkflow1-workflow"></a>
### mycompany.simple.SomeWorkflow1

<pre>
SomeWorkflow1 does some workflow thing.
</pre>

**Input:** [mycompany.simple.SomeWorkflow1Request](#mycompany-simple-someworkflow1request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeWorkflow1Response](#mycompany-simple-someworkflow1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>some-workflow-1/${! id }/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#mycompany-simple-simple-somequery1-query">mycompany.simple.Simple.SomeQuery1</a></td></tr>
<tr><td><a href="#mycompany-simple-simple-somequery2-query">mycompany.simple.Simple.SomeQuery2</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal1-signal">mycompany.simple.Simple.SomeSignal1</a></td><td>false</td></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal2-signal">mycompany.simple.Simple.SomeSignal2</a></td><td>false</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#mycompany-simple-simple-someupdate1-update">mycompany.simple.Simple.SomeUpdate1</a></td></tr>
<tr><td><a href="#mycompany-simple-simple-someupdate2-update">mycompany.simple.Simple.SomeUpdate2</a></td></tr>
</table>

---
<a name="mycompany-simple-someworkflow2-workflow"></a>
### mycompany.simple.SomeWorkflow2

<pre>
SomeWorkflow2 does some workflow thing.
</pre>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>some-workflow-2/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal1-signal">mycompany.simple.Simple.SomeSignal1</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#mycompany-simple-simple-someupdate1-update">mycompany.simple.Simple.SomeUpdate1</a></td></tr>
</table>

---
<a name="mycompany-simple-simple-someworkflow3-workflow"></a>
### mycompany.simple.Simple.SomeWorkflow3

<pre>
SomeWorkflow3 does some workflow thing.
Deprecated: Use SomeWorkflow2 instead.
</pre>

**Input:** [mycompany.simple.SomeWorkflow3Request](#mycompany-simple-someworkflow3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>execution_timeout</td><td>1 hour</td></tr>
<tr><td>id</td><td><pre><code>some-workflow-3/${! id }/${! requestVal }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE</code></pre></td></tr>
<tr><td>retry_policy.max_attempts</td><td>2</td></tr>
<tr><td>task_queue</td><td><pre><code>my-task-queue-2</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal2-signal">mycompany.simple.Simple.SomeSignal2</a></td><td>true</td></tr>
</table>

---
<a name="mycompany-simple-simple-someworkflow4-workflow"></a>
### mycompany.simple.Simple.SomeWorkflow4

<pre>
SomeWorkflow4 retrieves a paginated list of items
</pre>

**Input:** [mycompany.simple.common.v1.PaginatedRequest](#mycompany-simple-common-v1-paginatedrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cursor</td>
<td>bytes</td>
<td><pre>
json_name: cursor
go_name: Cursor</pre></td>
</tr><tr>
<td>limit</td>
<td>uint32</td>
<td><pre>
json_name: limit
go_name: Limit</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.common.v1.PaginatedResponse](#mycompany-simple-common-v1-paginatedresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="#google-protobuf-any">google.protobuf.Any</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>next_cursor</td>
<td>bytes</td>
<td><pre>
json_name: nextCursor
go_name: NextCursor</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>some-workflow-4/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>  

<a name="mycompany-simple-simple-queries"></a>
### Queries

---
<a name="mycompany-simple-simple-somequery1-query"></a>
### mycompany.simple.Simple.SomeQuery1

<pre>
SomeQuery1 queries some thing.
</pre>

**Output:** [mycompany.simple.SomeQuery1Response](#mycompany-simple-somequery1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

---
<a name="mycompany-simple-simple-somequery2-query"></a>
### mycompany.simple.Simple.SomeQuery2

<pre>
SomeQuery2 queries some thing.
</pre>

**Input:** [mycompany.simple.SomeQuery2Request](#mycompany-simple-somequery2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeQuery2Response](#mycompany-simple-somequery2response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>  

<a name="mycompany-simple-simple-signals"></a>
### Signals

---
<a name="mycompany-simple-simple-somesignal1-signal"></a>
### mycompany.simple.Simple.SomeSignal1

<pre>
SomeSignal1 is a signal.
</pre>

---
<a name="mycompany-simple-simple-somesignal2-signal"></a>
### mycompany.simple.Simple.SomeSignal2

<pre>
SomeSignal2 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal2Request](#mycompany-simple-somesignal2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

---
<a name="mycompany-simple-simple-somesignal3-signal"></a>
### mycompany.simple.Simple.SomeSignal3

<pre>
SomeSignal3 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal3Request](#mycompany-simple-somesignal3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>  

<a name="mycompany-simple-simple-updates"></a>
### Updates

---
<a name="mycompany-simple-simple-someupdate1-update"></a>
### mycompany.simple.Simple.SomeUpdate1

<pre>
SomeUpdate1 updates a SomeWorkflow2
</pre>

**Input:** [mycompany.simple.SomeUpdate1Request](#mycompany-simple-someupdate1request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeUpdate1Response](#mycompany-simple-someupdate1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>validate</td><td>true</td></tr>
<tr><td>wait_policy</td><td><pre>WAIT_POLICY_COMPLETED</pre></td></tr>
</table>

---
<a name="mycompany-simple-simple-someupdate2-update"></a>
### mycompany.simple.Simple.SomeUpdate2



**Input:** [mycompany.simple.SomeUpdate2Request](#mycompany-simple-someupdate2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeUpdate2Response](#mycompany-simple-someupdate2response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

<a name="mycompany-simple-simple-activities"></a>
### Activities

---
<a name="mycompany-simple-someactivity1-activity"></a>
### mycompany.simple.SomeActivity1

<pre>
SomeActivity1 does some activity thing.
</pre> 

---
<a name="mycompany-simple-simple-someactivity2-activity"></a>
### mycompany.simple.Simple.SomeActivity2

<pre>
SomeActivity2 does some activity thing.
</pre>

**Input:** [mycompany.simple.SomeActivity2Request](#mycompany-simple-someactivity2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_interval</td><td>30 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
<tr><td>wait_for_cancellation</td><td>true</td></tr>
</table> 

---
<a name="mycompany-simple-simple-someactivity3-activity"></a>
### mycompany.simple.Simple.SomeActivity3

<pre>
SomeActivity3 does some activity thing.
</pre>

**Input:** [mycompany.simple.SomeActivity3Request](#mycompany-simple-someactivity3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeActivity3Response](#mycompany-simple-someactivity3response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
<tr><td>task_queue</td><td>some-other-task-queue</td></tr>
</table> 

---
<a name="mycompany-simple-simple-someactivity4-activity"></a>
### mycompany.simple.Simple.SomeActivity4

<pre>
SomeActivity4 does some activity thing.
</pre>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>heartbeat_timeout</td><td>30 seconds</td></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>schedule_to_close_timeout</td><td>30 seconds</td></tr>
<tr><td>schedule_to_start_timeout</td><td>5 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>1 minute</td></tr>
<tr><td>wait_for_cancellation</td><td>true</td></tr>
<tr><td>task_queue</td><td>some-other-task-queue</td></tr>
</table> 

---
<a name="mycompany-simple-simple-somesignal1-activity"></a>
### mycompany.simple.Simple.SomeSignal1

<pre>
SomeSignal1 is a signal.
</pre>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-simple-somesignal2-activity"></a>
### mycompany.simple.Simple.SomeSignal2

<pre>
SomeSignal2 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal2Request](#mycompany-simple-somesignal2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-simple-somesignal3-activity"></a>
### mycompany.simple.Simple.SomeSignal3

<pre>
SomeSignal3 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal3Request](#mycompany-simple-somesignal3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeSignal3Response](#mycompany-simple-somesignal3response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-simple-someupdate1-activity"></a>
### mycompany.simple.Simple.SomeUpdate1

<pre>
SomeUpdate1 updates a SomeWorkflow2
</pre>

**Input:** [mycompany.simple.SomeUpdate1Request](#mycompany-simple-someupdate1request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeUpdate1Response](#mycompany-simple-someupdate1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>retry_policy.non_retryable_error_types</td><td>something</td></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table>   

<a name="mycompany-simple-other"></a>
## mycompany.simple.Other

<a name="mycompany-simple-other-workflows"></a>
### Workflows

---
<a name="mycompany-simple-other-otherworkflow-workflow"></a>
### mycompany.simple.Other.OtherWorkflow

**Input:** [mycompany.simple.OtherWorkflowRequest](#mycompany-simple-otherworkflowrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td><a href="#mycompany-simple-otherworkflowrequest-bar">mycompany.simple.OtherWorkflowRequest.Bar</a></td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr><tr>
<td>baz</td>
<td><a href="#mycompany-simple-otherworkflowrequest-baz">mycompany.simple.OtherWorkflowRequest.Baz</a></td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr><tr>
<td>common_enum</td>
<td><a href="#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
<td><pre>
json_name: commonEnum
go_name: CommonEnum</pre></td>
</tr><tr>
<td>example_bool</td>
<td>bool</td>
<td><pre>
json_name: exampleBool
go_name: ExampleBool</pre></td>
</tr><tr>
<td>example_bytes</td>
<td>bytes</td>
<td><pre>
json_name: exampleBytes
go_name: ExampleBytes</pre></td>
</tr><tr>
<td>example_double</td>
<td>double</td>
<td><pre>
json_name: exampleDouble
go_name: ExampleDouble</pre></td>
</tr><tr>
<td>example_duration</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="#google-protobuf-empty">google.protobuf.Empty</a></td>
<td><pre>
json_name: exampleEmpty
go_name: ExampleEmpty</pre></td>
</tr><tr>
<td>example_enum</td>
<td><a href="#mycompany-simple-otherenum">mycompany.simple.OtherEnum</a></td>
<td><pre>
json_name: exampleEnum
go_name: ExampleEnum</pre></td>
</tr><tr>
<td>example_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: exampleFixed32
go_name: ExampleFixed32</pre></td>
</tr><tr>
<td>example_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: exampleFixed64
go_name: ExampleFixed64</pre></td>
</tr><tr>
<td>example_float</td>
<td>float</td>
<td><pre>
json_name: exampleFloat
go_name: ExampleFloat</pre></td>
</tr><tr>
<td>example_int32</td>
<td>int32</td>
<td><pre>
json_name: exampleInt32
go_name: ExampleInt32</pre></td>
</tr><tr>
<td>example_int64</td>
<td>int64</td>
<td><pre>
json_name: exampleInt64
go_name: ExampleInt64</pre></td>
</tr><tr>
<td>example_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: exampleSfixed32
go_name: ExampleSfixed32</pre></td>
</tr><tr>
<td>example_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: exampleSfixed64
go_name: ExampleSfixed64</pre></td>
</tr><tr>
<td>example_sint32</td>
<td>sint32</td>
<td><pre>
json_name: exampleSint32
go_name: ExampleSint32</pre></td>
</tr><tr>
<td>example_sint64</td>
<td>sint64</td>
<td><pre>
json_name: exampleSint64
go_name: ExampleSint64</pre></td>
</tr><tr>
<td>example_timestamp</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: exampleTimestamp
go_name: ExampleTimestamp</pre></td>
</tr><tr>
<td>example_uint32</td>
<td>uint32</td>
<td><pre>
json_name: exampleUint32
go_name: ExampleUint32</pre></td>
</tr><tr>
<td>example_uint64</td>
<td>uint64</td>
<td><pre>
json_name: exampleUint64
go_name: ExampleUint64</pre></td>
</tr><tr>
<td>foo</td>
<td><a href="#mycompany-simple-foo">mycompany.simple.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr><tr>
<td>quux</td>
<td>string</td>
<td><pre>
json_name: quux
go_name: Quux</pre></td>
</tr><tr>
<td>qux</td>
<td><a href="#mycompany-simple-qux">mycompany.simple.Qux</a></td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr><tr>
<td>some_val</td>
<td>string</td>
<td><pre>
json_name: someVal
go_name: SomeVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.OtherWorkflowResponse](#mycompany-simple-otherworkflowresponse)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>other-workflow/${!uuid_v4()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="mycompany-simple-other-otherworkflow2-workflow"></a>
### mycompany.simple.Other.OtherWorkflow2

**Input:** [mycompany.simple.common.v1.PaginatedRequest](#mycompany-simple-common-v1-paginatedrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cursor</td>
<td>bytes</td>
<td><pre>
json_name: cursor
go_name: Cursor</pre></td>
</tr><tr>
<td>limit</td>
<td>uint32</td>
<td><pre>
json_name: limit
go_name: Limit</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.common.v1.PaginatedResponse](#mycompany-simple-common-v1-paginatedresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="#google-protobuf-any">google.protobuf.Any</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>next_cursor</td>
<td>bytes</td>
<td><pre>
json_name: nextCursor
go_name: NextCursor</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>other-workflow-2/${!uuid_v4()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>  

<a name="mycompany-simple-other-queries"></a>
### Queries

---
<a name="mycompany-simple-other-otherquery-query"></a>
### mycompany.simple.Other.OtherQuery



**Output:** [mycompany.simple.OtherQueryResponse](#mycompany-simple-otherqueryresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>filter</td>
<td>string</td>
<td><pre>
json_name: filter
go_name: Filter</pre></td>
</tr>
</table>  

<a name="mycompany-simple-other-signals"></a>
### Signals

---
<a name="mycompany-simple-other-othersignal-signal"></a>
### mycompany.simple.Other.OtherSignal



**Input:** [mycompany.simple.OtherSignalRequest](#mycompany-simple-othersignalrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>type</td>
<td>string</td>
<td><pre>
json_name: type
go_name: Type</pre></td>
</tr>
</table>  

<a name="mycompany-simple-other-updates"></a>
### Updates

---
<a name="mycompany-simple-other-otherupdate-update"></a>
### mycompany.simple.Other.OtherUpdate



**Input:** [mycompany.simple.OtherUpdateRequest](#mycompany-simple-otherupdaterequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>mode</td>
<td>string</td>
<td><pre>
json_name: mode
go_name: Mode</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.OtherUpdateResponse](#mycompany-simple-otherupdateresponse)



<a name="mycompany-simple-other-activities"></a>
### Activities

---
<a name="mycompany-simple-other-otherworkflow-activity"></a>
### mycompany.simple.Other.OtherWorkflow



**Input:** [mycompany.simple.OtherWorkflowRequest](#mycompany-simple-otherworkflowrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td><a href="#mycompany-simple-otherworkflowrequest-bar">mycompany.simple.OtherWorkflowRequest.Bar</a></td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr><tr>
<td>baz</td>
<td><a href="#mycompany-simple-otherworkflowrequest-baz">mycompany.simple.OtherWorkflowRequest.Baz</a></td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr><tr>
<td>common_enum</td>
<td><a href="#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
<td><pre>
json_name: commonEnum
go_name: CommonEnum</pre></td>
</tr><tr>
<td>example_bool</td>
<td>bool</td>
<td><pre>
json_name: exampleBool
go_name: ExampleBool</pre></td>
</tr><tr>
<td>example_bytes</td>
<td>bytes</td>
<td><pre>
json_name: exampleBytes
go_name: ExampleBytes</pre></td>
</tr><tr>
<td>example_double</td>
<td>double</td>
<td><pre>
json_name: exampleDouble
go_name: ExampleDouble</pre></td>
</tr><tr>
<td>example_duration</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="#google-protobuf-empty">google.protobuf.Empty</a></td>
<td><pre>
json_name: exampleEmpty
go_name: ExampleEmpty</pre></td>
</tr><tr>
<td>example_enum</td>
<td><a href="#mycompany-simple-otherenum">mycompany.simple.OtherEnum</a></td>
<td><pre>
json_name: exampleEnum
go_name: ExampleEnum</pre></td>
</tr><tr>
<td>example_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: exampleFixed32
go_name: ExampleFixed32</pre></td>
</tr><tr>
<td>example_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: exampleFixed64
go_name: ExampleFixed64</pre></td>
</tr><tr>
<td>example_float</td>
<td>float</td>
<td><pre>
json_name: exampleFloat
go_name: ExampleFloat</pre></td>
</tr><tr>
<td>example_int32</td>
<td>int32</td>
<td><pre>
json_name: exampleInt32
go_name: ExampleInt32</pre></td>
</tr><tr>
<td>example_int64</td>
<td>int64</td>
<td><pre>
json_name: exampleInt64
go_name: ExampleInt64</pre></td>
</tr><tr>
<td>example_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: exampleSfixed32
go_name: ExampleSfixed32</pre></td>
</tr><tr>
<td>example_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: exampleSfixed64
go_name: ExampleSfixed64</pre></td>
</tr><tr>
<td>example_sint32</td>
<td>sint32</td>
<td><pre>
json_name: exampleSint32
go_name: ExampleSint32</pre></td>
</tr><tr>
<td>example_sint64</td>
<td>sint64</td>
<td><pre>
json_name: exampleSint64
go_name: ExampleSint64</pre></td>
</tr><tr>
<td>example_timestamp</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: exampleTimestamp
go_name: ExampleTimestamp</pre></td>
</tr><tr>
<td>example_uint32</td>
<td>uint32</td>
<td><pre>
json_name: exampleUint32
go_name: ExampleUint32</pre></td>
</tr><tr>
<td>example_uint64</td>
<td>uint64</td>
<td><pre>
json_name: exampleUint64
go_name: ExampleUint64</pre></td>
</tr><tr>
<td>foo</td>
<td><a href="#mycompany-simple-foo">mycompany.simple.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr><tr>
<td>quux</td>
<td>string</td>
<td><pre>
json_name: quux
go_name: Quux</pre></td>
</tr><tr>
<td>qux</td>
<td><a href="#mycompany-simple-qux">mycompany.simple.Qux</a></td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr><tr>
<td>some_val</td>
<td>string</td>
<td><pre>
json_name: someVal
go_name: SomeVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.OtherWorkflowResponse](#mycompany-simple-otherworkflowresponse)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="mycompany-simple-ignored"></a>
## mycompany.simple.Ignored

<a name="mycompany-simple-ignored-workflows"></a>
### Workflows

---
<a name="mycompany-simple-ignored-what-workflow"></a>
### mycompany.simple.Ignored.What

**Input:** [mycompany.simple.WhatRequest](#mycompany-simple-whatrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>blah</td>
<td>string</td>
<td><pre>
json_name: blah
go_name: Blah</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>what/${!ksuid()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="mycompany-simple-onlyactivities"></a>
## mycompany.simple.OnlyActivities   

<a name="mycompany-simple-onlyactivities-activities"></a>
### Activities

---
<a name="mycompany-simple-onlyactivities-lonelyactivity1-activity"></a>
### mycompany.simple.OnlyActivities.LonelyActivity1



**Input:** [mycompany.simple.LonelyActivity1Request](#mycompany-simple-lonelyactivity1request)



**Output:** [mycompany.simple.LonelyActivity1Response](#mycompany-simple-lonelyactivity1response)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table>   

<a name="mycompany-simple-deprecated"></a>
## mycompany.simple.Deprecated

<a name="mycompany-simple-deprecated-workflows"></a>
### Workflows

---
<a name="mycompany-simple-deprecated-somedeprecatedworkflow1-workflow"></a>
### mycompany.simple.Deprecated.SomeDeprecatedWorkflow1

<pre>
SomeDeprecatedWorkflow1 does something
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedquery1-query">mycompany.simple.Deprecated.SomeDeprecatedQuery1</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedsignal1-signal">mycompany.simple.Deprecated.SomeDeprecatedSignal1</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedupdate1-update">mycompany.simple.Deprecated.SomeDeprecatedUpdate1</a></td></tr>
</table>

---
<a name="mycompany-simple-deprecated-somedeprecatedworkflow2-workflow"></a>
### mycompany.simple.Deprecated.SomeDeprecatedWorkflow2

<pre>
SomeDeprecatedWorkflow2 does something else

Deprecated: a custom workflow deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedquery2-query">mycompany.simple.Deprecated.SomeDeprecatedQuery2</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedsignal2-signal">mycompany.simple.Deprecated.SomeDeprecatedSignal2</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedupdate2-update">mycompany.simple.Deprecated.SomeDeprecatedUpdate2</a></td></tr>
</table>  

<a name="mycompany-simple-deprecated-queries"></a>
### Queries

---
<a name="mycompany-simple-deprecated-somedeprecatedquery1-query"></a>
### mycompany.simple.Deprecated.SomeDeprecatedQuery1

<pre>
SomeDeprecatedQuery1 does something else
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



---
<a name="mycompany-simple-deprecated-somedeprecatedquery2-query"></a>
### mycompany.simple.Deprecated.SomeDeprecatedQuery2

<pre>
SomeDeprecatedQuery2 does something else

Deprecated: a custom query deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)

  

<a name="mycompany-simple-deprecated-signals"></a>
### Signals

---
<a name="mycompany-simple-deprecated-somedeprecatedsignal1-signal"></a>
### mycompany.simple.Deprecated.SomeDeprecatedSignal1

<pre>
SomeDeprecatedSignal1 does something else
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



---
<a name="mycompany-simple-deprecated-somedeprecatedsignal2-signal"></a>
### mycompany.simple.Deprecated.SomeDeprecatedSignal2

<pre>
SomeDeprecatedSignal2 does something else

Deprecated: a custom signal deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)

  

<a name="mycompany-simple-deprecated-updates"></a>
### Updates

---
<a name="mycompany-simple-deprecated-somedeprecatedupdate1-update"></a>
### mycompany.simple.Deprecated.SomeDeprecatedUpdate1

<pre>
SomeDeprecatedUpdate1 does something else
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



---
<a name="mycompany-simple-deprecated-somedeprecatedupdate2-update"></a>
### mycompany.simple.Deprecated.SomeDeprecatedUpdate2

<pre>
SomeDeprecatedUpdate2 does something else

Deprecated: a custom signal deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



<a name="mycompany-simple-deprecated-activities"></a>
### Activities

---
<a name="mycompany-simple-deprecated-somedeprecatedactivity1-activity"></a>
### mycompany.simple.Deprecated.SomeDeprecatedActivity1

<pre>
SomeDeprecatedActivity1 does something
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-deprecated-somedeprecatedactivity2-activity"></a>
### mycompany.simple.Deprecated.SomeDeprecatedActivity2

<pre>
SomeDeprecatedActivity2 does something else

Deprecated: a custom activity deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table>   

<a name="mycompany-simple-messages"></a>
## Messages

<a name="mycompany-simple-foo"></a>
### mycompany.simple.Foo

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td>string</td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="mycompany-simple-lonelyactivity1request"></a>
### mycompany.simple.LonelyActivity1Request



<a name="mycompany-simple-lonelyactivity1response"></a>
### mycompany.simple.LonelyActivity1Response



<a name="mycompany-simple-otherenum"></a>
### mycompany.simple.OtherEnum

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>OTHER_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>OTHER_FOO</td>
<td></td>
</tr><tr>
<td>OTHER_BAR</td>
<td></td>
</tr>
</table>

<a name="mycompany-simple-otherqueryresponse"></a>
### mycompany.simple.OtherQueryResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>filter</td>
<td>string</td>
<td><pre>
json_name: filter
go_name: Filter</pre></td>
</tr>
</table>



<a name="mycompany-simple-othersignalrequest"></a>
### mycompany.simple.OtherSignalRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>type</td>
<td>string</td>
<td><pre>
json_name: type
go_name: Type</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherupdaterequest"></a>
### mycompany.simple.OtherUpdateRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>mode</td>
<td>string</td>
<td><pre>
json_name: mode
go_name: Mode</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherupdateresponse"></a>
### mycompany.simple.OtherUpdateResponse



<a name="mycompany-simple-otherworkflowrequest"></a>
### mycompany.simple.OtherWorkflowRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td><a href="#mycompany-simple-otherworkflowrequest-bar">mycompany.simple.OtherWorkflowRequest.Bar</a></td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr><tr>
<td>baz</td>
<td><a href="#mycompany-simple-otherworkflowrequest-baz">mycompany.simple.OtherWorkflowRequest.Baz</a></td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr><tr>
<td>common_enum</td>
<td><a href="#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
<td><pre>
json_name: commonEnum
go_name: CommonEnum</pre></td>
</tr><tr>
<td>example_bool</td>
<td>bool</td>
<td><pre>
json_name: exampleBool
go_name: ExampleBool</pre></td>
</tr><tr>
<td>example_bytes</td>
<td>bytes</td>
<td><pre>
json_name: exampleBytes
go_name: ExampleBytes</pre></td>
</tr><tr>
<td>example_double</td>
<td>double</td>
<td><pre>
json_name: exampleDouble
go_name: ExampleDouble</pre></td>
</tr><tr>
<td>example_duration</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="#google-protobuf-empty">google.protobuf.Empty</a></td>
<td><pre>
json_name: exampleEmpty
go_name: ExampleEmpty</pre></td>
</tr><tr>
<td>example_enum</td>
<td><a href="#mycompany-simple-otherenum">mycompany.simple.OtherEnum</a></td>
<td><pre>
json_name: exampleEnum
go_name: ExampleEnum</pre></td>
</tr><tr>
<td>example_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: exampleFixed32
go_name: ExampleFixed32</pre></td>
</tr><tr>
<td>example_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: exampleFixed64
go_name: ExampleFixed64</pre></td>
</tr><tr>
<td>example_float</td>
<td>float</td>
<td><pre>
json_name: exampleFloat
go_name: ExampleFloat</pre></td>
</tr><tr>
<td>example_int32</td>
<td>int32</td>
<td><pre>
json_name: exampleInt32
go_name: ExampleInt32</pre></td>
</tr><tr>
<td>example_int64</td>
<td>int64</td>
<td><pre>
json_name: exampleInt64
go_name: ExampleInt64</pre></td>
</tr><tr>
<td>example_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: exampleSfixed32
go_name: ExampleSfixed32</pre></td>
</tr><tr>
<td>example_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: exampleSfixed64
go_name: ExampleSfixed64</pre></td>
</tr><tr>
<td>example_sint32</td>
<td>sint32</td>
<td><pre>
json_name: exampleSint32
go_name: ExampleSint32</pre></td>
</tr><tr>
<td>example_sint64</td>
<td>sint64</td>
<td><pre>
json_name: exampleSint64
go_name: ExampleSint64</pre></td>
</tr><tr>
<td>example_timestamp</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: exampleTimestamp
go_name: ExampleTimestamp</pre></td>
</tr><tr>
<td>example_uint32</td>
<td>uint32</td>
<td><pre>
json_name: exampleUint32
go_name: ExampleUint32</pre></td>
</tr><tr>
<td>example_uint64</td>
<td>uint64</td>
<td><pre>
json_name: exampleUint64
go_name: ExampleUint64</pre></td>
</tr><tr>
<td>foo</td>
<td><a href="#mycompany-simple-foo">mycompany.simple.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr><tr>
<td>quux</td>
<td>string</td>
<td><pre>
json_name: quux
go_name: Quux</pre></td>
</tr><tr>
<td>qux</td>
<td><a href="#mycompany-simple-qux">mycompany.simple.Qux</a></td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr><tr>
<td>some_val</td>
<td>string</td>
<td><pre>
json_name: someVal
go_name: SomeVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherworkflowrequest-bar"></a>
### mycompany.simple.OtherWorkflowRequest.Bar

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td>string</td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherworkflowrequest-baz"></a>
### mycompany.simple.OtherWorkflowRequest.Baz

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>baz</td>
<td>string</td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherworkflowresponse"></a>
### mycompany.simple.OtherWorkflowResponse



<a name="mycompany-simple-qux"></a>
### mycompany.simple.Qux

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>qux</td>
<td>string</td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr>
</table>



<a name="mycompany-simple-someactivity2request"></a>
### mycompany.simple.SomeActivity2Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someactivity3request"></a>
### mycompany.simple.SomeActivity3Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someactivity3response"></a>
### mycompany.simple.SomeActivity3Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somedeprecatedmessage"></a>
### mycompany.simple.SomeDeprecatedMessage



<a name="mycompany-simple-somequery1response"></a>
### mycompany.simple.SomeQuery1Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somequery2request"></a>
### mycompany.simple.SomeQuery2Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somequery2response"></a>
### mycompany.simple.SomeQuery2Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somesignal2request"></a>
### mycompany.simple.SomeSignal2Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somesignal3request"></a>
### mycompany.simple.SomeSignal3Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somesignal3response"></a>
### mycompany.simple.SomeSignal3Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someupdate1request"></a>
### mycompany.simple.SomeUpdate1Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someupdate1response"></a>
### mycompany.simple.SomeUpdate1Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someupdate2request"></a>
### mycompany.simple.SomeUpdate2Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someupdate2response"></a>
### mycompany.simple.SomeUpdate2Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someworkflow1request"></a>
### mycompany.simple.SomeWorkflow1Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someworkflow1response"></a>
### mycompany.simple.SomeWorkflow1Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someworkflow3request"></a>
### mycompany.simple.SomeWorkflow3Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-whatrequest"></a>
### mycompany.simple.WhatRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>blah</td>
<td>string</td>
<td><pre>
json_name: blah
go_name: Blah</pre></td>
</tr>
</table>



<a name="test-acronym-v1"></a>
# test.acronym.v1

<a name="test-acronym-v1-services"></a>
## Services

<a name="test-acronym-v1-aws"></a>
## test.acronym.v1.AWS

<a name="test-acronym-v1-aws-workflows"></a>
### Workflows

---
<a name="test-acronym-v1-aws-manageaws-workflow"></a>
### test.acronym.v1.AWS.ManageAWS

<pre>
ManageAWS does some workflow thing.
</pre>

**Input:** [test.acronym.v1.ManageAWSRequest](#test-acronym-v1-manageawsrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Output:** [test.acronym.v1.ManageAWSResponse](#test-acronym-v1-manageawsresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>manage-aws/${! urn }/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="test-acronym-v1-aws-manageawsresource-workflow"></a>
### test.acronym.v1.AWS.ManageAWSResource

<pre>
ManageAWSResource does some workflow thing.
</pre>

**Input:** [test.acronym.v1.ManageAWSResourceRequest](#test-acronym-v1-manageawsresourcerequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>k8s_namespace</td>
<td>string</td>
<td><pre>
json_name: k8sNamespace
go_name: K8SNamespace</pre></td>
</tr><tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Output:** [test.acronym.v1.ManageAWSResourceResponse](#test-acronym-v1-manageawsresourceresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>manage-aws-resource/${! urn }/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="test-acronym-v1-aws-somethingv1foobar-workflow"></a>
### test.acronym.v1.AWS.SomethingV1FooBar

<pre>
SomethingV1FooBar does some workflow thing.
</pre>

**Input:** [test.acronym.v1.SomethingV1FooBarRequest](#test-acronym-v1-somethingv1foobarrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>k8s_namespace</td>
<td>string</td>
<td><pre>
json_name: k8sNamespace
go_name: K8SNamespace</pre></td>
</tr><tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Output:** [test.acronym.v1.SomethingV1FooBarResponse](#test-acronym-v1-somethingv1foobarresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>something-v1-foo-bar/${! urn }/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="test-acronym-v1-aws-somethingv2foobar-workflow"></a>
### test.acronym.v1.AWS.SomethingV2FooBar

<pre>
SomethingV2FooBar does some workflow thing.
</pre>

**Input:** [test.acronym.v1.SomethingV2FooBarRequest](#test-acronym-v1-somethingv2foobarrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>k8s_namespace</td>
<td>string</td>
<td><pre>
json_name: k8sNamespace
go_name: K8SNamespace</pre></td>
</tr><tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Output:** [test.acronym.v1.SomethingV2FooBarResponse](#test-acronym-v1-somethingv2foobarresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>something-v2-foo-bar/${! urn }/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>    

<a name="test-acronym-v1-aws-activities"></a>
### Activities

---
<a name="test-acronym-v1-aws-manageawsresource-activity"></a>
### test.acronym.v1.AWS.ManageAWSResource

<pre>
ManageAWSResource does some workflow thing.
</pre>

**Input:** [test.acronym.v1.ManageAWSResourceRequest](#test-acronym-v1-manageawsresourcerequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>k8s_namespace</td>
<td>string</td>
<td><pre>
json_name: k8sNamespace
go_name: K8SNamespace</pre></td>
</tr><tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Output:** [test.acronym.v1.ManageAWSResourceResponse](#test-acronym-v1-manageawsresourceresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>1 minute</td></tr>
</table> 

---
<a name="test-acronym-v1-aws-manageawsresourceurn-activity"></a>
### test.acronym.v1.AWS.ManageAWSResourceURN

<pre>
ManageAWSResourceURN does some workflow thing.
</pre>

**Input:** [test.acronym.v1.ManageAWSResourceURNRequest](#test-acronym-v1-manageawsresourceurnrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Output:** [test.acronym.v1.ManageAWSResourceURNResponse](#test-acronym-v1-manageawsresourceurnresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>1 minute</td></tr>
</table>   

<a name="test-acronym-v1-messages"></a>
## Messages

<a name="test-acronym-v1-manageawsrequest"></a>
### test.acronym.v1.ManageAWSRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-manageawsresourcerequest"></a>
### test.acronym.v1.ManageAWSResourceRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>k8s_namespace</td>
<td>string</td>
<td><pre>
json_name: k8sNamespace
go_name: K8SNamespace</pre></td>
</tr><tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-manageawsresourceresponse"></a>
### test.acronym.v1.ManageAWSResourceResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-manageawsresourceurnrequest"></a>
### test.acronym.v1.ManageAWSResourceURNRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-manageawsresourceurnresponse"></a>
### test.acronym.v1.ManageAWSResourceURNResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-manageawsresponse"></a>
### test.acronym.v1.ManageAWSResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-somethingv1foobarrequest"></a>
### test.acronym.v1.SomethingV1FooBarRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>k8s_namespace</td>
<td>string</td>
<td><pre>
json_name: k8sNamespace
go_name: K8SNamespace</pre></td>
</tr><tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-somethingv1foobarresponse"></a>
### test.acronym.v1.SomethingV1FooBarResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-somethingv2foobarrequest"></a>
### test.acronym.v1.SomethingV2FooBarRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>k8s_namespace</td>
<td>string</td>
<td><pre>
json_name: k8sNamespace
go_name: K8SNamespace</pre></td>
</tr><tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-acronym-v1-somethingv2foobarresponse"></a>
### test.acronym.v1.SomethingV2FooBarResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>urn</td>
<td>string</td>
<td><pre>
json_name: urn
go_name: Urn</pre></td>
</tr>
</table>



<a name="test-activity-v1"></a>
# test.activity.v1

<a name="test-activity-v1-services"></a>
## Services

<a name="test-activity-v1-example"></a>
## test.activity.v1.Example   

<a name="test-activity-v1-example-activities"></a>
### Activities

---
<a name="test-activity-v1-example-foo-activity"></a>
### test.activity.v1.Example.Foo



**Input:** [test.activity.v1.FooInput](#test-activity-v1-fooinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>input</td>
<td>string</td>
<td><pre>
json_name: input
go_name: Input</pre></td>
</tr>
</table>

**Output:** [test.activity.v1.FooOutput](#test-activity-v1-foooutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>output</td>
<td>string</td>
<td><pre>
json_name: output
go_name: Output</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="test-activity-v1-example-bar-activity"></a>
### test.activity.v1.Example.Bar



**Input:** [test.activity.v1.BarInput](#test-activity-v1-barinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>input</td>
<td>string</td>
<td><pre>
json_name: input
go_name: Input</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="test-activity-v1-example-baz-activity"></a>
### test.activity.v1.Example.Baz



**Output:** [test.activity.v1.BazOutput](#test-activity-v1-bazoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>output</td>
<td>string</td>
<td><pre>
json_name: output
go_name: Output</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="test-activity-v1-example-qux-activity"></a>
### test.activity.v1.Example.Qux



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table>   

<a name="test-activity-v1-messages"></a>
## Messages

<a name="test-activity-v1-barinput"></a>
### test.activity.v1.BarInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>input</td>
<td>string</td>
<td><pre>
json_name: input
go_name: Input</pre></td>
</tr>
</table>



<a name="test-activity-v1-bazoutput"></a>
### test.activity.v1.BazOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>output</td>
<td>string</td>
<td><pre>
json_name: output
go_name: Output</pre></td>
</tr>
</table>



<a name="test-activity-v1-fooinput"></a>
### test.activity.v1.FooInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>input</td>
<td>string</td>
<td><pre>
json_name: input
go_name: Input</pre></td>
</tr>
</table>



<a name="test-activity-v1-foooutput"></a>
### test.activity.v1.FooOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>output</td>
<td>string</td>
<td><pre>
json_name: output
go_name: Output</pre></td>
</tr>
</table>



<a name="test-editions"></a>
# test.editions

<a name="test-editions-services"></a>
## Services

<a name="test-editions-fooservice"></a>
## test.editions.FooService

<a name="test-editions-fooservice-workflows"></a>
### Workflows

---
<a name="test-editions-fooservice-foo-workflow"></a>
### test.editions.FooService.Foo

**Input:** [test.editions.FooInput](#test-editions-fooinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>optional_bool</td>
<td>bool</td>
<td><pre>
json_name: optionalBool
go_name: OptionalBool</pre></td>
</tr><tr>
<td>optional_bytes</td>
<td>bytes</td>
<td><pre>
json_name: optionalBytes
go_name: OptionalBytes</pre></td>
</tr><tr>
<td>optional_double</td>
<td>double</td>
<td><pre>
json_name: optionalDouble
go_name: OptionalDouble</pre></td>
</tr><tr>
<td>optional_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: optionalFixed32
go_name: OptionalFixed32</pre></td>
</tr><tr>
<td>optional_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: optionalFixed64
go_name: OptionalFixed64</pre></td>
</tr><tr>
<td>optional_float</td>
<td>float</td>
<td><pre>
json_name: optionalFloat
go_name: OptionalFloat</pre></td>
</tr><tr>
<td>optional_int32</td>
<td>int32</td>
<td><pre>
json_name: optionalInt32
go_name: OptionalInt32</pre></td>
</tr><tr>
<td>optional_int64</td>
<td>int64</td>
<td><pre>
json_name: optionalInt64
go_name: OptionalInt64</pre></td>
</tr><tr>
<td>optional_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: optionalSfixed32
go_name: OptionalSfixed32</pre></td>
</tr><tr>
<td>optional_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: optionalSfixed64
go_name: OptionalSfixed64</pre></td>
</tr><tr>
<td>optional_sint32</td>
<td>sint32</td>
<td><pre>
json_name: optionalSint32
go_name: OptionalSint32</pre></td>
</tr><tr>
<td>optional_sint64</td>
<td>sint64</td>
<td><pre>
json_name: optionalSint64
go_name: OptionalSint64</pre></td>
</tr><tr>
<td>optional_string</td>
<td>string</td>
<td><pre>
json_name: optionalString
go_name: OptionalString</pre></td>
</tr><tr>
<td>optional_uint32</td>
<td>uint32</td>
<td><pre>
json_name: optionalUint32
go_name: OptionalUint32</pre></td>
</tr><tr>
<td>optional_uint64</td>
<td>uint64</td>
<td><pre>
json_name: optionalUint64
go_name: OptionalUint64</pre></td>
</tr>
</table>

**Output:** [test.editions.FooOutput](#test-editions-foooutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>    

<a name="test-editions-fooservice-activities"></a>
### Activities

---
<a name="test-editions-fooservice-foo-activity"></a>
### test.editions.FooService.Foo



**Input:** [test.editions.FooInput](#test-editions-fooinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>optional_bool</td>
<td>bool</td>
<td><pre>
json_name: optionalBool
go_name: OptionalBool</pre></td>
</tr><tr>
<td>optional_bytes</td>
<td>bytes</td>
<td><pre>
json_name: optionalBytes
go_name: OptionalBytes</pre></td>
</tr><tr>
<td>optional_double</td>
<td>double</td>
<td><pre>
json_name: optionalDouble
go_name: OptionalDouble</pre></td>
</tr><tr>
<td>optional_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: optionalFixed32
go_name: OptionalFixed32</pre></td>
</tr><tr>
<td>optional_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: optionalFixed64
go_name: OptionalFixed64</pre></td>
</tr><tr>
<td>optional_float</td>
<td>float</td>
<td><pre>
json_name: optionalFloat
go_name: OptionalFloat</pre></td>
</tr><tr>
<td>optional_int32</td>
<td>int32</td>
<td><pre>
json_name: optionalInt32
go_name: OptionalInt32</pre></td>
</tr><tr>
<td>optional_int64</td>
<td>int64</td>
<td><pre>
json_name: optionalInt64
go_name: OptionalInt64</pre></td>
</tr><tr>
<td>optional_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: optionalSfixed32
go_name: OptionalSfixed32</pre></td>
</tr><tr>
<td>optional_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: optionalSfixed64
go_name: OptionalSfixed64</pre></td>
</tr><tr>
<td>optional_sint32</td>
<td>sint32</td>
<td><pre>
json_name: optionalSint32
go_name: OptionalSint32</pre></td>
</tr><tr>
<td>optional_sint64</td>
<td>sint64</td>
<td><pre>
json_name: optionalSint64
go_name: OptionalSint64</pre></td>
</tr><tr>
<td>optional_string</td>
<td>string</td>
<td><pre>
json_name: optionalString
go_name: OptionalString</pre></td>
</tr><tr>
<td>optional_uint32</td>
<td>uint32</td>
<td><pre>
json_name: optionalUint32
go_name: OptionalUint32</pre></td>
</tr><tr>
<td>optional_uint64</td>
<td>uint64</td>
<td><pre>
json_name: optionalUint64
go_name: OptionalUint64</pre></td>
</tr>
</table>

**Output:** [test.editions.FooOutput](#test-editions-foooutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>2 seconds</td></tr>
</table>   

<a name="test-editions-messages"></a>
## Messages

<a name="test-editions-fooinput"></a>
### test.editions.FooInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>optional_bool</td>
<td>bool</td>
<td><pre>
json_name: optionalBool
go_name: OptionalBool</pre></td>
</tr><tr>
<td>optional_bytes</td>
<td>bytes</td>
<td><pre>
json_name: optionalBytes
go_name: OptionalBytes</pre></td>
</tr><tr>
<td>optional_double</td>
<td>double</td>
<td><pre>
json_name: optionalDouble
go_name: OptionalDouble</pre></td>
</tr><tr>
<td>optional_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: optionalFixed32
go_name: OptionalFixed32</pre></td>
</tr><tr>
<td>optional_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: optionalFixed64
go_name: OptionalFixed64</pre></td>
</tr><tr>
<td>optional_float</td>
<td>float</td>
<td><pre>
json_name: optionalFloat
go_name: OptionalFloat</pre></td>
</tr><tr>
<td>optional_int32</td>
<td>int32</td>
<td><pre>
json_name: optionalInt32
go_name: OptionalInt32</pre></td>
</tr><tr>
<td>optional_int64</td>
<td>int64</td>
<td><pre>
json_name: optionalInt64
go_name: OptionalInt64</pre></td>
</tr><tr>
<td>optional_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: optionalSfixed32
go_name: OptionalSfixed32</pre></td>
</tr><tr>
<td>optional_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: optionalSfixed64
go_name: OptionalSfixed64</pre></td>
</tr><tr>
<td>optional_sint32</td>
<td>sint32</td>
<td><pre>
json_name: optionalSint32
go_name: OptionalSint32</pre></td>
</tr><tr>
<td>optional_sint64</td>
<td>sint64</td>
<td><pre>
json_name: optionalSint64
go_name: OptionalSint64</pre></td>
</tr><tr>
<td>optional_string</td>
<td>string</td>
<td><pre>
json_name: optionalString
go_name: OptionalString</pre></td>
</tr><tr>
<td>optional_uint32</td>
<td>uint32</td>
<td><pre>
json_name: optionalUint32
go_name: OptionalUint32</pre></td>
</tr><tr>
<td>optional_uint64</td>
<td>uint64</td>
<td><pre>
json_name: optionalUint64
go_name: OptionalUint64</pre></td>
</tr>
</table>



<a name="test-editions-foooutput"></a>
### test.editions.FooOutput



<a name="test-opaque"></a>
# test.opaque

<a name="test-opaque-services"></a>
## Services

<a name="test-opaque-hybrid"></a>
## test.opaque.Hybrid

<a name="test-opaque-hybrid-workflows"></a>
### Workflows

---
<a name="test-opaque-hybrid-puthybridexample-workflow"></a>
### test.opaque.Hybrid.PutHybridExample

**Input:** [test.opaque.HybridExample](#test-opaque-hybridexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.HybridExample](#test-opaque-hybridexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-hybrid-signalhybrid-signal">test.opaque.Hybrid.SignalHybrid</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-hybrid-signals"></a>
### Signals

---
<a name="test-opaque-hybrid-signalhybrid-signal"></a>
### test.opaque.Hybrid.SignalHybrid



**Input:** [test.opaque.HybridExample](#test-opaque-hybridexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-opaque"></a>
## test.opaque.Opaque

<a name="test-opaque-opaque-workflows"></a>
### Workflows

---
<a name="test-opaque-opaque-putopaqueexample-workflow"></a>
### test.opaque.Opaque.PutOpaqueExample

**Input:** [test.opaque.OpaqueExample](#test-opaque-opaqueexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.OpaqueExample](#test-opaque-opaqueexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-opaque-signalopaque-signal">test.opaque.Opaque.SignalOpaque</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-opaque-signals"></a>
### Signals

---
<a name="test-opaque-opaque-signalopaque-signal"></a>
### test.opaque.Opaque.SignalOpaque



**Input:** [test.opaque.OpaqueExample](#test-opaque-opaqueexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-open"></a>
## test.opaque.Open

<a name="test-opaque-open-workflows"></a>
### Workflows

---
<a name="test-opaque-open-putopenexample-workflow"></a>
### test.opaque.Open.PutOpenExample

**Input:** [test.opaque.OpenExample](#test-opaque-openexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.OpenExample](#test-opaque-openexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-open-signalopen-signal">test.opaque.Open.SignalOpen</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-open-signals"></a>
### Signals

---
<a name="test-opaque-open-signalopen-signal"></a>
### test.opaque.Open.SignalOpen



**Input:** [test.opaque.OpenExample](#test-opaque-openexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-optional"></a>
## test.opaque.Optional

<a name="test-opaque-optional-workflows"></a>
### Workflows

---
<a name="test-opaque-optional-putoptionalexample-workflow"></a>
### test.opaque.Optional.PutOptionalExample

**Input:** [test.opaque.OptionalExample](#test-opaque-optionalexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.OptionalExample](#test-opaque-optionalexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-optional-signaloptional-signal">test.opaque.Optional.SignalOptional</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-optional-signals"></a>
### Signals

---
<a name="test-opaque-optional-signaloptional-signal"></a>
### test.opaque.Optional.SignalOptional



**Input:** [test.opaque.OptionalExample](#test-opaque-optionalexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-messages"></a>
## Messages

<a name="test-opaque-address"></a>
### test.opaque.Address

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>city</td>
<td>string</td>
<td><pre>
json_name: city
go_name: City</pre></td>
</tr><tr>
<td>state</td>
<td>string</td>
<td><pre>
json_name: state
go_name: State</pre></td>
</tr><tr>
<td>street</td>
<td>string</td>
<td><pre>
json_name: street
go_name: Street</pre></td>
</tr><tr>
<td>zip</td>
<td>string</td>
<td><pre>
json_name: zip
go_name: Zip</pre></td>
</tr>
</table>



<a name="test-opaque-hybridexample"></a>
### test.opaque.HybridExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-hybridexample-extraentry"></a>
### test.opaque.HybridExample.ExtraEntry

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>key</td>
<td>string</td>
<td><pre>
json_name: key
go_name: Key</pre></td>
</tr><tr>
<td>value</td>
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-opaqueexample"></a>
### test.opaque.OpaqueExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-opaqueexample-extraentry"></a>
### test.opaque.OpaqueExample.ExtraEntry

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>key</td>
<td>string</td>
<td><pre>
json_name: key
go_name: Key</pre></td>
</tr><tr>
<td>value</td>
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-openexample"></a>
### test.opaque.OpenExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-openexample-extraentry"></a>
### test.opaque.OpenExample.ExtraEntry

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>key</td>
<td>string</td>
<td><pre>
json_name: key
go_name: Key</pre></td>
</tr><tr>
<td>value</td>
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-optionalexample"></a>
### test.opaque.OptionalExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-optionalexample-extraentry"></a>
### test.opaque.OptionalExample.ExtraEntry

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>key</td>
<td>string</td>
<td><pre>
json_name: key
go_name: Key</pre></td>
</tr><tr>
<td>value</td>
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-status"></a>
### test.opaque.Status

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>STATUS_OK</td>
<td></td>
</tr><tr>
<td>STATUS_ERROR</td>
<td></td>
</tr>
</table>

<a name="test-option-v1"></a>
# test.option.v1

<a name="test-option-v1-services"></a>
## Services

<a name="test-option-v1-test"></a>
## test.option.v1.Test

<a name="test-option-v1-test-workflows"></a>
### Workflows

---
<a name="test-option-v1-test-workflowwithinput-workflow"></a>
### test.option.v1.Test.WorkflowWithInput

**Input:** [test.option.v1.WorkflowWithInputRequest](#test-option-v1-workflowwithinputrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>execution_timeout</td><td>10 minutes</td></tr>
<tr><td>id</td><td><pre><code>workflow-with-input:${! name.or(throw("name is required")) }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY</code></pre></td></tr>
<tr><td>parent_close_policy</td><td><pre><code>PARENT_CLOSE_POLICY_REQUEST_CANCEL</code></pre></td></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>run_timeout</td><td>5 minutes</td></tr>
<tr><td>search_attributes</td><td><pre><code>name = name</code></pre></td></tr>
<tr><td>task_queue</td><td><pre><code>option-v2</code></pre></td></tr>
<tr><td>task_timeout</td><td>10 seconds</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#test-option-v1-test-updatewithinput-update">test.option.v1.Test.UpdateWithInput</a></td></tr>
</table>    

<a name="test-option-v1-test-updates"></a>
### Updates

---
<a name="test-option-v1-test-updatewithinput-update"></a>
### test.option.v1.Test.UpdateWithInput



**Input:** [test.option.v1.UpdateWithInputRequest](#test-option-v1-updatewithinputrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>wait_policy</td><td><pre>WAIT_POLICY_ADMITTED</pre></td></tr>
</table>

<a name="test-option-v1-test-activities"></a>
### Activities

---
<a name="test-option-v1-test-activitywithinput-activity"></a>
### test.option.v1.Test.ActivityWithInput



**Input:** [test.option.v1.ActivityWithInputRequest](#test-option-v1-activitywithinputrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>

**Output:** [test.option.v1.ActivityWithInputResponse](#test-option-v1-activitywithinputresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>result</td>
<td>string</td>
<td><pre>
json_name: result
go_name: Result</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>heartbeat_timeout</td><td>30 seconds</td></tr>
<tr><td>retry_policy.max_interval</td><td>5 seconds</td></tr>
<tr><td>schedule_to_close_timeout</td><td>30 seconds</td></tr>
<tr><td>schedule_to_start_timeout</td><td>10 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>1 minute</td></tr>
<tr><td>wait_for_cancellation</td><td>true</td></tr>
<tr><td>task_queue</td><td>option-v2</td></tr>
</table>   

<a name="test-option-v1-messages"></a>
## Messages

<a name="test-option-v1-activitywithinputrequest"></a>
### test.option.v1.ActivityWithInputRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="test-option-v1-activitywithinputresponse"></a>
### test.option.v1.ActivityWithInputResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>result</td>
<td>string</td>
<td><pre>
json_name: result
go_name: Result</pre></td>
</tr>
</table>



<a name="test-option-v1-updatewithinputrequest"></a>
### test.option.v1.UpdateWithInputRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="test-option-v1-workflowwithinputrequest"></a>
### test.option.v1.WorkflowWithInputRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="test-proto3optional"></a>
# test.proto3optional

<a name="test-proto3optional-services"></a>
## Services

<a name="test-proto3optional-fooservice"></a>
## test.proto3optional.FooService

<a name="test-proto3optional-fooservice-workflows"></a>
### Workflows

---
<a name="test-proto3optional-fooservice-foo-workflow"></a>
### test.proto3optional.FooService.Foo

**Input:** [test.proto3optional.FooInput](#test-proto3optional-fooinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>optional_bool</td>
<td>bool</td>
<td><pre>
json_name: optionalBool
go_name: OptionalBool</pre></td>
</tr><tr>
<td>optional_bytes</td>
<td>bytes</td>
<td><pre>
json_name: optionalBytes
go_name: OptionalBytes</pre></td>
</tr><tr>
<td>optional_double</td>
<td>double</td>
<td><pre>
json_name: optionalDouble
go_name: OptionalDouble</pre></td>
</tr><tr>
<td>optional_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: optionalFixed32
go_name: OptionalFixed32</pre></td>
</tr><tr>
<td>optional_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: optionalFixed64
go_name: OptionalFixed64</pre></td>
</tr><tr>
<td>optional_float</td>
<td>float</td>
<td><pre>
json_name: optionalFloat
go_name: OptionalFloat</pre></td>
</tr><tr>
<td>optional_int32</td>
<td>int32</td>
<td><pre>
json_name: optionalInt32
go_name: OptionalInt32</pre></td>
</tr><tr>
<td>optional_int64</td>
<td>int64</td>
<td><pre>
json_name: optionalInt64
go_name: OptionalInt64</pre></td>
</tr><tr>
<td>optional_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: optionalSfixed32
go_name: OptionalSfixed32</pre></td>
</tr><tr>
<td>optional_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: optionalSfixed64
go_name: OptionalSfixed64</pre></td>
</tr><tr>
<td>optional_sint32</td>
<td>sint32</td>
<td><pre>
json_name: optionalSint32
go_name: OptionalSint32</pre></td>
</tr><tr>
<td>optional_sint64</td>
<td>sint64</td>
<td><pre>
json_name: optionalSint64
go_name: OptionalSint64</pre></td>
</tr><tr>
<td>optional_string</td>
<td>string</td>
<td><pre>
json_name: optionalString
go_name: OptionalString</pre></td>
</tr><tr>
<td>optional_uint32</td>
<td>uint32</td>
<td><pre>
json_name: optionalUint32
go_name: OptionalUint32</pre></td>
</tr><tr>
<td>optional_uint64</td>
<td>uint64</td>
<td><pre>
json_name: optionalUint64
go_name: OptionalUint64</pre></td>
</tr>
</table>

**Output:** [test.proto3optional.FooOutput](#test-proto3optional-foooutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td>string</td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>    

<a name="test-proto3optional-fooservice-activities"></a>
### Activities

---
<a name="test-proto3optional-fooservice-foo-activity"></a>
### test.proto3optional.FooService.Foo



**Input:** [test.proto3optional.FooInput](#test-proto3optional-fooinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>optional_bool</td>
<td>bool</td>
<td><pre>
json_name: optionalBool
go_name: OptionalBool</pre></td>
</tr><tr>
<td>optional_bytes</td>
<td>bytes</td>
<td><pre>
json_name: optionalBytes
go_name: OptionalBytes</pre></td>
</tr><tr>
<td>optional_double</td>
<td>double</td>
<td><pre>
json_name: optionalDouble
go_name: OptionalDouble</pre></td>
</tr><tr>
<td>optional_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: optionalFixed32
go_name: OptionalFixed32</pre></td>
</tr><tr>
<td>optional_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: optionalFixed64
go_name: OptionalFixed64</pre></td>
</tr><tr>
<td>optional_float</td>
<td>float</td>
<td><pre>
json_name: optionalFloat
go_name: OptionalFloat</pre></td>
</tr><tr>
<td>optional_int32</td>
<td>int32</td>
<td><pre>
json_name: optionalInt32
go_name: OptionalInt32</pre></td>
</tr><tr>
<td>optional_int64</td>
<td>int64</td>
<td><pre>
json_name: optionalInt64
go_name: OptionalInt64</pre></td>
</tr><tr>
<td>optional_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: optionalSfixed32
go_name: OptionalSfixed32</pre></td>
</tr><tr>
<td>optional_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: optionalSfixed64
go_name: OptionalSfixed64</pre></td>
</tr><tr>
<td>optional_sint32</td>
<td>sint32</td>
<td><pre>
json_name: optionalSint32
go_name: OptionalSint32</pre></td>
</tr><tr>
<td>optional_sint64</td>
<td>sint64</td>
<td><pre>
json_name: optionalSint64
go_name: OptionalSint64</pre></td>
</tr><tr>
<td>optional_string</td>
<td>string</td>
<td><pre>
json_name: optionalString
go_name: OptionalString</pre></td>
</tr><tr>
<td>optional_uint32</td>
<td>uint32</td>
<td><pre>
json_name: optionalUint32
go_name: OptionalUint32</pre></td>
</tr><tr>
<td>optional_uint64</td>
<td>uint64</td>
<td><pre>
json_name: optionalUint64
go_name: OptionalUint64</pre></td>
</tr>
</table>

**Output:** [test.proto3optional.FooOutput](#test-proto3optional-foooutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td>string</td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>2 seconds</td></tr>
</table>   

<a name="test-proto3optional-messages"></a>
## Messages

<a name="test-proto3optional-fooinput"></a>
### test.proto3optional.FooInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>optional_bool</td>
<td>bool</td>
<td><pre>
json_name: optionalBool
go_name: OptionalBool</pre></td>
</tr><tr>
<td>optional_bytes</td>
<td>bytes</td>
<td><pre>
json_name: optionalBytes
go_name: OptionalBytes</pre></td>
</tr><tr>
<td>optional_double</td>
<td>double</td>
<td><pre>
json_name: optionalDouble
go_name: OptionalDouble</pre></td>
</tr><tr>
<td>optional_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: optionalFixed32
go_name: OptionalFixed32</pre></td>
</tr><tr>
<td>optional_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: optionalFixed64
go_name: OptionalFixed64</pre></td>
</tr><tr>
<td>optional_float</td>
<td>float</td>
<td><pre>
json_name: optionalFloat
go_name: OptionalFloat</pre></td>
</tr><tr>
<td>optional_int32</td>
<td>int32</td>
<td><pre>
json_name: optionalInt32
go_name: OptionalInt32</pre></td>
</tr><tr>
<td>optional_int64</td>
<td>int64</td>
<td><pre>
json_name: optionalInt64
go_name: OptionalInt64</pre></td>
</tr><tr>
<td>optional_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: optionalSfixed32
go_name: OptionalSfixed32</pre></td>
</tr><tr>
<td>optional_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: optionalSfixed64
go_name: OptionalSfixed64</pre></td>
</tr><tr>
<td>optional_sint32</td>
<td>sint32</td>
<td><pre>
json_name: optionalSint32
go_name: OptionalSint32</pre></td>
</tr><tr>
<td>optional_sint64</td>
<td>sint64</td>
<td><pre>
json_name: optionalSint64
go_name: OptionalSint64</pre></td>
</tr><tr>
<td>optional_string</td>
<td>string</td>
<td><pre>
json_name: optionalString
go_name: OptionalString</pre></td>
</tr><tr>
<td>optional_uint32</td>
<td>uint32</td>
<td><pre>
json_name: optionalUint32
go_name: OptionalUint32</pre></td>
</tr><tr>
<td>optional_uint64</td>
<td>uint64</td>
<td><pre>
json_name: optionalUint64
go_name: OptionalUint64</pre></td>
</tr>
</table>



<a name="test-proto3optional-foooutput"></a>
### test.proto3optional.FooOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td>string</td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr>
</table>



<a name="test-xnserr-v1"></a>
# test.xnserr.v1

<a name="test-xnserr-v1-services"></a>
## Services

<a name="test-xnserr-v1-server"></a>
## test.xnserr.v1.Server

<a name="test-xnserr-v1-server-workflows"></a>
### Workflows

---
<a name="test-xnserr-v1-server-sleep-workflow"></a>
### test.xnserr.v1.Server.Sleep

**Input:** [test.xnserr.v1.SleepRequest](#test-xnserr-v1-sleeprequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE</code></pre></td></tr>
<tr><td>parent_close_policy</td><td><pre><code>PARENT_CLOSE_POLICY_REQUEST_CANCEL</code></pre></td></tr>
<tr><td>xns.heartbeat_interval</td><td>10 seconds</td></tr>
<tr><td>xns.heartbeat_timeout</td><td>30 seconds</td></tr>
</table>     

<a name="test-xnserr-v1-client"></a>
## test.xnserr.v1.Client

<a name="test-xnserr-v1-client-workflows"></a>
### Workflows

---
<a name="test-xnserr-v1-client-callsleep-workflow"></a>
### test.xnserr.v1.Client.CallSleep

**Input:** [test.xnserr.v1.CallSleepRequest](#test-xnserr-v1-callsleeprequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr><tr>
<td>start_workflow_options</td>
<td><a href="#temporal-xns-v1-startworkflowoptions">temporal.xns.v1.StartWorkflowOptions</a></td>
<td><pre>
json_name: startWorkflowOptions
go_name: StartWorkflowOptions</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="test-xnserr-v1-messages"></a>
## Messages

<a name="test-xnserr-v1-callsleeprequest"></a>
### test.xnserr.v1.CallSleepRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr><tr>
<td>start_workflow_options</td>
<td><a href="#temporal-xns-v1-startworkflowoptions">temporal.xns.v1.StartWorkflowOptions</a></td>
<td><pre>
json_name: startWorkflowOptions
go_name: StartWorkflowOptions</pre></td>
</tr>
</table>



<a name="test-xnserr-v1-failure"></a>
### test.xnserr.v1.Failure

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>application_error_type</td>
<td>string</td>
<td><pre>
json_name: applicationErrorType
go_name: ApplicationErrorType</pre></td>
</tr><tr>
<td>info</td>
<td><a href="#test-xnserr-v1-failureinfo">test.xnserr.v1.FailureInfo</a></td>
<td><pre>
json_name: info
go_name: Info</pre></td>
</tr><tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr><tr>
<td>non_retryable</td>
<td>bool</td>
<td><pre>
json_name: nonRetryable
go_name: NonRetryable</pre></td>
</tr>
</table>



<a name="test-xnserr-v1-failureinfo"></a>
### test.xnserr.v1.FailureInfo

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>FAILURE_INFO_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_APPLICATION_ERROR</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_TIMEOUT</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_CANCELED</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_TERMINATED</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_ACTIVITY</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_WORKFLOW_EXECUTION</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_CHILD_WORKFLOW_EXECUTION</td>
<td></td>
</tr>
</table>

<a name="test-xnserr-v1-sleeprequest"></a>
### test.xnserr.v1.SleepRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>




<a name="google-protobuf"></a>
# google.protobuf

<a name="google-protobuf-messages"></a>
## Messages

<a name="google-protobuf-any"></a>
### google.protobuf.Any

<pre>
`Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }
    // or ...
    if (any.isSameTypeAs(Foo.getDefaultInstance())) {
      foo = any.unpack(Foo.getDefaultInstance());
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".

JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>type_url</td>
<td>string</td>
<td><pre>
A URL/resource name that uniquely identifies the type of the serialized
protocol buffer message. This string must contain at least
one "/" character. The last segment of the URL's path must represent
the fully qualified name of the type (as in
`path/google.protobuf.Duration`). The name should be in a canonical form
(e.g., leading "." is not accepted).

In practice, teams usually precompile into the binary all types that they
expect it to use in the context of Any. However, for URLs which use the
scheme `http`, `https`, or no scheme, one can optionally set up a type
server that maps type URLs to message definitions as follows:

* If no scheme is provided, `https` is assumed.
* An HTTP GET on the URL must yield a [google.protobuf.Type][]
  value in binary format, or produce an error.
* Applications are allowed to cache lookup results based on the
  URL, or have them precompiled into a binary to avoid any
  lookup. Therefore, binary compatibility needs to be preserved
  on changes to types. (Use versioned type names to manage
  breaking changes.)

Note: this functionality is not currently available in the official
protobuf release, and it is not used for type URLs beginning with
type.googleapis.com. As of May 2023, there are no widely used type server
implementations and no plans to implement one.

Schemes other than `http`, `https` (or the empty scheme) might be
used with implementation specific semantics.<br>

json_name: typeUrl
go_name: TypeUrl</pre></td>
</tr><tr>
<td>value</td>
<td>bytes</td>
<td><pre>
Must be a valid serialized protocol buffer of the above specified type.<br>

json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="google-protobuf-duration"></a>
### google.protobuf.Duration

<pre>
A Duration represents a signed, fixed-length span of time represented
as a count of seconds and fractions of seconds at nanosecond
resolution. It is independent of any calendar and concepts like "day"
or "month". It is related to Timestamp in that the difference between
two Timestamp values is a Duration and it can be added or subtracted
from a Timestamp. Range is approximately +-10,000 years.

# Examples

Example 1: Compute Duration from two Timestamps in pseudo code.

    Timestamp start = ...;
    Timestamp end = ...;
    Duration duration = ...;

    duration.seconds = end.seconds - start.seconds;
    duration.nanos = end.nanos - start.nanos;

    if (duration.seconds < 0 && duration.nanos > 0) {
      duration.seconds += 1;
      duration.nanos -= 1000000000;
    } else if (duration.seconds > 0 && duration.nanos < 0) {
      duration.seconds -= 1;
      duration.nanos += 1000000000;
    }

Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.

    Timestamp start = ...;
    Duration duration = ...;
    Timestamp end = ...;

    end.seconds = start.seconds + duration.seconds;
    end.nanos = start.nanos + duration.nanos;

    if (end.nanos < 0) {
      end.seconds -= 1;
      end.nanos += 1000000000;
    } else if (end.nanos >= 1000000000) {
      end.seconds += 1;
      end.nanos -= 1000000000;
    }

Example 3: Compute Duration from datetime.timedelta in Python.

    td = datetime.timedelta(days=3, minutes=10)
    duration = Duration()
    duration.FromTimedelta(td)

# JSON Mapping

In JSON format, the Duration type is encoded as a string rather than an
object, where the string ends in the suffix "s" (indicating seconds) and
is preceded by the number of seconds, with nanoseconds expressed as
fractional seconds. For example, 3 seconds with 0 nanoseconds should be
encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
be expressed in JSON format as "3.000000001s", and 3 seconds and 1
microsecond should be expressed in JSON format as "3.000001s".
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>nanos</td>
<td>int32</td>
<td><pre>
Signed fractions of a second at nanosecond resolution of the span
of time. Durations less than one second are represented with a 0
`seconds` field and a positive or negative `nanos` field. For durations
of one second or more, a non-zero value for the `nanos` field must be
of the same sign as the `seconds` field. Must be from -999,999,999
to +999,999,999 inclusive.<br>

json_name: nanos
go_name: Nanos</pre></td>
</tr><tr>
<td>seconds</td>
<td>int64</td>
<td><pre>
Signed seconds of the span of time. Must be from -315,576,000,000
to +315,576,000,000 inclusive. Note: these bounds are computed from:
60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years<br>

json_name: seconds
go_name: Seconds</pre></td>
</tr>
</table>



<a name="google-protobuf-listvalue"></a>
### google.protobuf.ListValue

<pre>
`ListValue` is a wrapper around a repeated field of values.

The JSON representation for `ListValue` is JSON array.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>values</td>
<td><a href="#google-protobuf-value">google.protobuf.Value</a></td>
<td><pre>
Repeated field of dynamically typed values.<br>

json_name: values
go_name: Values</pre></td>
</tr>
</table>



<a name="google-protobuf-nullvalue"></a>
### google.protobuf.NullValue

<pre>
`NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

The JSON representation for `NullValue` is JSON `null`.
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>NULL_VALUE</td>
<td><pre>
Null value.
</pre></td>
</tr>
</table>

<a name="google-protobuf-struct"></a>
### google.protobuf.Struct

<pre>
`Struct` represents a structured data value, consisting of fields
which map to dynamically typed values. In some languages, `Struct`
might be supported by a native representation. For example, in
scripting languages like JS a struct is represented as an
object. The details of that representation are described together
with the proto support for the language.

The JSON representation for `Struct` is JSON object.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>fields</td>
<td><a href="#google-protobuf-struct-fieldsentry">google.protobuf.Struct.FieldsEntry</a></td>
<td><pre>
Unordered map of dynamically typed values.<br>

json_name: fields
go_name: Fields</pre></td>
</tr>
</table>



<a name="google-protobuf-struct-fieldsentry"></a>
### google.protobuf.Struct.FieldsEntry

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>key</td>
<td>string</td>
<td><pre>
json_name: key
go_name: Key</pre></td>
</tr><tr>
<td>value</td>
<td><a href="#google-protobuf-value">google.protobuf.Value</a></td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="google-protobuf-timestamp"></a>
### google.protobuf.Timestamp

<pre>
A Timestamp represents a point in time independent of any time zone or local
calendar, encoded as a count of seconds and fractions of seconds at
nanosecond resolution. The count is relative to an epoch at UTC midnight on
January 1, 1970, in the proleptic Gregorian calendar which extends the
Gregorian calendar backwards to year one.

All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
second table is needed for interpretation, using a [24-hour linear
smear](https://developers.google.com/time/smear).

The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
restricting to that range, we ensure that we can convert to and from [RFC
3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

# Examples

Example 1: Compute Timestamp from POSIX `time()`.

    Timestamp timestamp;
    timestamp.set_seconds(time(NULL));
    timestamp.set_nanos(0);

Example 2: Compute Timestamp from POSIX `gettimeofday()`.

    struct timeval tv;
    gettimeofday(&tv, NULL);

    Timestamp timestamp;
    timestamp.set_seconds(tv.tv_sec);
    timestamp.set_nanos(tv.tv_usec * 1000);

Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

    FILETIME ft;
    GetSystemTimeAsFileTime(&ft);
    UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

    // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
    // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
    Timestamp timestamp;
    timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
    timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

    long millis = System.currentTimeMillis();

    Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
        .setNanos((int) ((millis % 1000) * 1000000)).build();

Example 5: Compute Timestamp from Java `Instant.now()`.

    Instant now = Instant.now();

    Timestamp timestamp =
        Timestamp.newBuilder().setSeconds(now.getEpochSecond())
            .setNanos(now.getNano()).build();

Example 6: Compute Timestamp from current time in Python.

    timestamp = Timestamp()
    timestamp.GetCurrentTime()

# JSON Mapping

In JSON format, the Timestamp type is encoded as a string in the
[RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
where {year} is always expressed using four digits while {month}, {day},
{hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
is required. A proto3 JSON serializer should always use UTC (as indicated by
"Z") when printing the Timestamp type and a proto3 JSON parser should be
able to accept both UTC and other timezones (as indicated by an offset).

For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
01:30 UTC on January 15, 2017.

In JavaScript, one can convert a Date object to this format using the
standard
[toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
method. In Python, a standard `datetime.datetime` object can be converted
to this format using
[`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
the Joda Time's [`ISODateTimeFormat.dateTime()`](
http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()
) to obtain a formatter capable of generating timestamps in this format.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>nanos</td>
<td>int32</td>
<td><pre>
Non-negative fractions of a second at nanosecond resolution. Negative
second values with fractions must still have non-negative nanos values
that count forward in time. Must be from 0 to 999,999,999
inclusive.<br>

json_name: nanos
go_name: Nanos</pre></td>
</tr><tr>
<td>seconds</td>
<td>int64</td>
<td><pre>
Represents seconds of UTC time since Unix epoch
1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
9999-12-31T23:59:59Z inclusive.<br>

json_name: seconds
go_name: Seconds</pre></td>
</tr>
</table>



<a name="google-protobuf-value"></a>
### google.protobuf.Value

<pre>
`Value` represents a dynamically typed value which can be either
null, a number, a string, a boolean, a recursive struct value, or a
list of values. A producer of value is expected to set one of these
variants. Absence of any variant indicates an error.

The JSON representation for `Value` is JSON value.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bool_value</td>
<td>bool</td>
<td><pre>
Represents a boolean value.<br>

json_name: boolValue
go_name: BoolValue</pre></td>
</tr><tr>
<td>list_value</td>
<td><a href="#google-protobuf-listvalue">google.protobuf.ListValue</a></td>
<td><pre>
Represents a repeated `Value`.<br>

json_name: listValue
go_name: ListValue</pre></td>
</tr><tr>
<td>null_value</td>
<td><a href="#google-protobuf-nullvalue">google.protobuf.NullValue</a></td>
<td><pre>
Represents a null value.<br>

json_name: nullValue
go_name: NullValue</pre></td>
</tr><tr>
<td>number_value</td>
<td>double</td>
<td><pre>
Represents a double value.<br>

json_name: numberValue
go_name: NumberValue</pre></td>
</tr><tr>
<td>string_value</td>
<td>string</td>
<td><pre>
Represents a string value.<br>

json_name: stringValue
go_name: StringValue</pre></td>
</tr><tr>
<td>struct_value</td>
<td><a href="#google-protobuf-struct">google.protobuf.Struct</a></td>
<td><pre>
Represents a structured value.<br>

json_name: structValue
go_name: StructValue</pre></td>
</tr>
</table>




<a name="mycompany-simple-common-v1"></a>
# mycompany.simple.common.v1

<a name="mycompany-simple-common-v1-messages"></a>
## Messages

<a name="mycompany-simple-common-v1-example"></a>
### mycompany.simple.common.v1.Example

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>EXAMPLE_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>EXAMPLE_FOO</td>
<td></td>
</tr>
</table>

<a name="mycompany-simple-common-v1-paginatedrequest"></a>
### mycompany.simple.common.v1.PaginatedRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cursor</td>
<td>bytes</td>
<td><pre>
json_name: cursor
go_name: Cursor</pre></td>
</tr><tr>
<td>limit</td>
<td>uint32</td>
<td><pre>
json_name: limit
go_name: Limit</pre></td>
</tr>
</table>



<a name="mycompany-simple-common-v1-paginatedresponse"></a>
### mycompany.simple.common.v1.PaginatedResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="#google-protobuf-any">google.protobuf.Any</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>next_cursor</td>
<td>bytes</td>
<td><pre>
json_name: nextCursor
go_name: NextCursor</pre></td>
</tr>
</table>




<a name="temporal-api-enums-v1"></a>
# temporal.api.enums.v1

<a name="temporal-api-enums-v1-messages"></a>
## Messages

<a name="temporal-api-enums-v1-workflowidconflictpolicy"></a>
### temporal.api.enums.v1.WorkflowIdConflictPolicy

<pre>
Defines what to do when trying to start a workflow with the same workflow id as a *running* workflow.
Note that it is *never* valid to have two actively running instances of the same workflow id.

See `WorkflowIdReusePolicy` for handling workflow id duplication with a *closed* workflow.
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_FAIL</td>
<td><pre>
Don't start a new workflow; instead return `WorkflowExecutionAlreadyStartedFailure`.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING</td>
<td><pre>
Don't start a new workflow; instead return a workflow handle for the running workflow.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING</td>
<td><pre>
Terminate the running workflow before starting a new one.
</pre></td>
</tr>
</table>


<a name="temporal-xns-v1"></a>
# temporal.xns.v1

<a name="temporal-xns-v1-messages"></a>
## Messages

<a name="temporal-xns-v1-idreusepolicy"></a>
### temporal.xns.v1.IDReusePolicy

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE</td>
<td><pre>
Allow starting a workflow execution using the same workflow id.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY</td>
<td><pre>
Allow starting a workflow execution using the same workflow id, only when the last
execution's final state is one of [terminated, cancelled, timed out, failed].
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE</td>
<td><pre>
Do not permit re-use of the workflow id for this workflow. Future start workflow requests
could potentially change the policy, allowing re-use of the workflow id.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING</td>
<td><pre>
If a workflow is running using the same workflow ID, terminate it and start a new one.
If no running workflow, then the behavior is the same as ALLOW_DUPLICATE
</pre></td>
</tr>
</table>

<a name="temporal-xns-v1-retrypolicy"></a>
### temporal.xns.v1.RetryPolicy

<pre>
RetryPolicy describes configuration for activity or child workflow retries
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>backoff_coefficient</td>
<td>double</td>
<td><pre>
json_name: backoffCoefficient
go_name: BackoffCoefficient</pre></td>
</tr><tr>
<td>initial_interval</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: initialInterval
go_name: InitialInterval</pre></td>
</tr><tr>
<td>max_attempts</td>
<td>int32</td>
<td><pre>
json_name: maxAttempts
go_name: MaxAttempts</pre></td>
</tr><tr>
<td>max_interval</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: maxInterval
go_name: MaxInterval</pre></td>
</tr><tr>
<td>non_retryable_error_types</td>
<td>string</td>
<td><pre>
json_name: nonRetryableErrorTypes
go_name: NonRetryableErrorTypes</pre></td>
</tr>
</table>



<a name="temporal-xns-v1-startworkflowoptions"></a>
### temporal.xns.v1.StartWorkflowOptions

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>enable_eager_start</td>
<td>bool</td>
<td><pre>
json_name: enableEagerStart
go_name: EnableEagerStart</pre></td>
</tr><tr>
<td>error_when_already_started</td>
<td>bool</td>
<td><pre>
json_name: errorWhenAlreadyStarted
go_name: ErrorWhenAlreadyStarted</pre></td>
</tr><tr>
<td>execution_timeout</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: executionTimeout
go_name: ExecutionTimeout</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>id_reuse_policy</td>
<td><a href="#temporal-xns-v1-idreusepolicy">temporal.xns.v1.IDReusePolicy</a></td>
<td><pre>
json_name: idReusePolicy
go_name: IdReusePolicy</pre></td>
</tr><tr>
<td>memo</td>
<td><a href="#google-protobuf-struct">google.protobuf.Struct</a></td>
<td><pre>
json_name: memo
go_name: Memo</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>run_timeout</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: runTimeout
go_name: RunTimeout</pre></td>
</tr><tr>
<td>search_attirbutes</td>
<td><a href="#google-protobuf-struct">google.protobuf.Struct</a></td>
<td><pre>
json_name: searchAttirbutes
go_name: SearchAttirbutes</pre></td>
</tr><tr>
<td>start_delay</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: startDelay
go_name: StartDelay</pre></td>
</tr><tr>
<td>task_queue</td>
<td>string</td>
<td><pre>
json_name: taskQueue
go_name: TaskQueue</pre></td>
</tr><tr>
<td>task_timeout</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: taskTimeout
go_name: TaskTimeout</pre></td>
</tr><tr>
<td>workflow_id_conflict_policy</td>
<td><a href="#temporal-api-enums-v1-workflowidconflictpolicy">temporal.api.enums.v1.WorkflowIdConflictPolicy</a></td>
<td><pre>
json_name: workflowIdConflictPolicy
go_name: WorkflowIdConflictPolicy</pre></td>
</tr>
</table>

