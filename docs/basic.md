# Table of Contents

- [example.v1](#example-v1)
  - Services
    - [example.v1.Example](#example-v1-example)
      - [Workflows](#example-v1-example-workflows)
        - [example.v1.Example.CreateFoo](#example-v1-example-createfoo-workflow)
      - [Queries](#example-v1-example-queries)
        - [example.v1.Example.GetFooProgress](#example-v1-example-getfooprogress-query)
      - [Signals](#example-v1-example-signals)
        - [example.v1.Example.SetFooProgress](#example-v1-example-setfooprogress-signal)
      - [Updates](#example-v1-example-updates)
        - [example.v1.Example.UpdateFooProgress](#example-v1-example-updatefooprogress-update)
      - [Activities](#example-v1-example-activities)
        - [example.v1.Example.Notify](#example-v1-example-notify-activity)
    - [example.v1.External](#example-v1-external)
      - [Workflows](#example-v1-external-workflows)
        - [example.v1.External.ProvisionFoo](#example-v1-external-provisionfoo-workflow)
  - Messages
    - [example.v1.CreateFooRequest](#example-v1-createfoorequest)
    - [example.v1.CreateFooResponse](#example-v1-createfooresponse)
    - [example.v1.Foo](#example-v1-foo)
    - [example.v1.Foo.Status](#example-v1-foo-status)
    - [example.v1.GetFooProgressResponse](#example-v1-getfooprogressresponse)
    - [example.v1.NotifyRequest](#example-v1-notifyrequest)
    - [example.v1.ProvisionFooRequest](#example-v1-provisionfoorequest)
    - [example.v1.ProvisionFooResponse](#example-v1-provisionfooresponse)
    - [example.v1.SetFooProgressRequest](#example-v1-setfooprogressrequest)
- [mycompany.simple](#mycompany-simple)
  - Services
    - [mycompany.simple.Simple](#mycompany-simple-simple)
      - [Workflows](#mycompany-simple-simple-workflows)
        - [mycompany.simple.SomeWorkflow1](#mycompany-simple-someworkflow1-workflow)
        - [mycompany.simple.SomeWorkflow2](#mycompany-simple-someworkflow2-workflow)
        - [mycompany.simple.Simple.SomeWorkflow3](#mycompany-simple-simple-someworkflow3-workflow)
      - [Queries](#mycompany-simple-simple-queries)
        - [mycompany.simple.Simple.SomeQuery1](#mycompany-simple-simple-somequery1-query)
        - [mycompany.simple.Simple.SomeQuery2](#mycompany-simple-simple-somequery2-query)
      - [Signals](#mycompany-simple-simple-signals)
        - [mycompany.simple.Simple.SomeSignal1](#mycompany-simple-simple-somesignal1-signal)
        - [mycompany.simple.Simple.SomeSignal2](#mycompany-simple-simple-somesignal2-signal)
        - [mycompany.simple.Simple.SomeSignal3](#mycompany-simple-simple-somesignal3-signal)
      - [Updates](#mycompany-simple-simple-updates)
        - [mycompany.simple.Simple.SomeUpdate1](#mycompany-simple-simple-someupdate1-update)
      - [Activities](#mycompany-simple-simple-activities)
        - [mycompany.simple.SomeActivity1](#mycompany-simple-someactivity1-activity)
        - [mycompany.simple.Simple.SomeActivity2](#mycompany-simple-simple-someactivity2-activity)
        - [mycompany.simple.Simple.SomeActivity3](#mycompany-simple-simple-someactivity3-activity)
        - [mycompany.simple.Simple.SomeSignal1](#mycompany-simple-simple-somesignal1-activity)
        - [mycompany.simple.Simple.SomeSignal2](#mycompany-simple-simple-somesignal2-activity)
        - [mycompany.simple.Simple.SomeSignal3](#mycompany-simple-simple-somesignal3-activity)
        - [mycompany.simple.Simple.SomeUpdate1](#mycompany-simple-simple-someupdate1-activity)
    - [mycompany.simple.Other](#mycompany-simple-other)
      - [Workflows](#mycompany-simple-other-workflows)
        - [mycompany.simple.Other.OtherWorkflow](#mycompany-simple-other-otherworkflow-workflow)
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
    - [mycompany.simple.SomeWorkflow1Request](#mycompany-simple-someworkflow1request)
    - [mycompany.simple.SomeWorkflow1Response](#mycompany-simple-someworkflow1response)
    - [mycompany.simple.SomeWorkflow3Request](#mycompany-simple-someworkflow3request)
    - [mycompany.simple.WhatRequest](#mycompany-simple-whatrequest)
- [google.protobuf](#google-protobuf)
  - Messages
    - [google.protobuf.Duration](#google-protobuf-duration)
    - [google.protobuf.Timestamp](#google-protobuf-timestamp)

<a name="example-v1"></a>
# example.v1

<a name="example-v1-services"></a>
## Services

<a name="example-v1-example"></a>
## example.v1.Example

<a name="example-v1-example-workflows"></a>
### Workflows

---
<a name="example-v1-example-createfoo-workflow"></a>
### example.v1.Example.CreateFoo

<pre>
CreateFoo creates a new foo operation
</pre>

**Input:** [example.v1.CreateFooRequest](#example-v1-createfoorequest)

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
unique foo name<br>

json_name: name
go_name: RequestName</pre></td>
</tr>
</table>

**Output:** [example.v1.CreateFooResponse](#example-v1-createfooresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-v1-foo">example.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>execution_timeout</td><td>1 hour</td></tr>
<tr><td>id</td><td><pre><code>create-foo/${! name.slug() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE</code></pre></td></tr>
<tr><td>search_attributes</td><td><pre><code>foo = name
created_at = now().ts_tz("UTC")</code></pre></td></tr>
<tr><td>xns.heartbeat_interval</td><td>10 seconds</td></tr>
<tr><td>xns.heartbeat_timeout</td><td>20 seconds</td></tr>
<tr><td>xns.start_to_close_timeout</td><td>1 hour 30 seconds</td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#example-v1-example-getfooprogress-query">example.v1.Example.GetFooProgress</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-v1-example-setfooprogress-signal">example.v1.Example.SetFooProgress</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#example-v1-example-updatefooprogress-update">example.v1.Example.UpdateFooProgress</a></td></tr>
</table>  

<a name="example-v1-example-queries"></a>
### Queries

---
<a name="example-v1-example-getfooprogress-query"></a>
### example.v1.Example.GetFooProgress

<pre>
GetFooProgress returns the status of a CreateFoo operation
</pre>

**Output:** [example.v1.GetFooProgressResponse](#example-v1-getfooprogressresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>  

<a name="example-v1-example-signals"></a>
### Signals

---
<a name="example-v1-example-setfooprogress-signal"></a>
### example.v1.Example.SetFooProgress

<pre>
SetFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.v1.SetFooProgressRequest](#example-v1-setfooprogressrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>  

<a name="example-v1-example-updates"></a>
### Updates

---
<a name="example-v1-example-updatefooprogress-update"></a>
### example.v1.Example.UpdateFooProgress

<pre>
UpdateFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.v1.SetFooProgressRequest](#example-v1-setfooprogressrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>

**Output:** [example.v1.GetFooProgressResponse](#example-v1-getfooprogressresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
</table>

<a name="example-v1-example-activities"></a>
### Activities

---
<a name="example-v1-example-notify-activity"></a>
### example.v1.Example.Notify

<pre>
Notify sends a notification
</pre>

**Input:** [example.v1.NotifyRequest](#example-v1-notifyrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="example-v1-external"></a>
## example.v1.External

<a name="example-v1-external-workflows"></a>
### Workflows

---
<a name="example-v1-external-provisionfoo-workflow"></a>
### example.v1.External.ProvisionFoo

**Input:** [example.v1.ProvisionFooRequest](#example-v1-provisionfoorequest)

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
unique foo name<br>

json_name: name
go_name: RequestName</pre></td>
</tr>
</table>

**Output:** [example.v1.ProvisionFooResponse](#example-v1-provisionfooresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-v1-foo">example.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>provision-foo/${! name.slug() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-v1-external-example-v1-example-setfooprogress-signal">example.v1.External.example.v1.Example.SetFooProgress</a></td><td>false</td></tr>
<tr><td><a href="#example-v1-external-mycompany-simple-simple-somesignal3-signal">example.v1.External.mycompany.simple.Simple.SomeSignal3</a></td><td>false</td></tr>
</table>     

<a name="example-v1-messages"></a>
## Messages

<a name="example-v1-createfoorequest"></a>
### example.v1.CreateFooRequest

<pre>
CreateFooRequest describes the input to a CreateFoo workflow
</pre>

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
unique foo name<br>

json_name: name
go_name: RequestName</pre></td>
</tr>
</table>



<a name="example-v1-createfooresponse"></a>
### example.v1.CreateFooResponse

<pre>
SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-v1-foo">example.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="example-v1-foo"></a>
### example.v1.Foo

<pre>
Foo describes an illustrative foo resource
</pre>

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
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-v1-foo-status"></a>
### example.v1.Foo.Status

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>FOO_STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>FOO_STATUS_READY</td>
<td></td>
</tr><tr>
<td>FOO_STATUS_CREATING</td>
<td></td>
</tr>
</table>

<a name="example-v1-getfooprogressresponse"></a>
### example.v1.GetFooProgressResponse

<pre>
GetFooProgressResponse describes the output from a GetFooProgress query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-v1-notifyrequest"></a>
### example.v1.NotifyRequest

<pre>
NotifyRequest describes the input to a Notify activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>



<a name="example-v1-provisionfoorequest"></a>
### example.v1.ProvisionFooRequest

<pre>
ProvisionFooRequest describes the input to a ProvisionFoo workflow
</pre>

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
unique foo name<br>

json_name: name
go_name: RequestName</pre></td>
</tr>
</table>



<a name="example-v1-provisionfooresponse"></a>
### example.v1.ProvisionFooResponse

<pre>
SampleWorkflowWithMutexResponse describes the output from a ProvisionFoo workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-v1-foo">example.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="example-v1-setfooprogressrequest"></a>
### example.v1.SetFooProgressRequest

<pre>
SetFooProgressRequest describes the input to a SetFooProgress signal
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>



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

---
<a name="mycompany-simple-someworkflow2-workflow"></a>
### mycompany.simple.SomeWorkflow2

<pre>
SomeWorkflow2 does some workflow thing.
</pre>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
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




<a name="google-protobuf"></a>
# google.protobuf

<a name="google-protobuf-messages"></a>
## Messages

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

