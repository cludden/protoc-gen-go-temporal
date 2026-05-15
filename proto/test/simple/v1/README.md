

<a name="mycompany-simple"></a>
# mycompany.simple

## Table of Contents
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
- [mycompany.simple.Ignored](#mycompany-simple-ignored)
  - [Workflows](#mycompany-simple-ignored-workflows)
    - [mycompany.simple.Ignored.What](#mycompany-simple-ignored-what-workflow)
- [mycompany.simple.OnlyActivities](#mycompany-simple-onlyactivities)
  - [Activities](#mycompany-simple-onlyactivities-activities)
    - [mycompany.simple.OnlyActivities.LonelyActivity1](#mycompany-simple-onlyactivities-lonelyactivity1-activity)
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
- [mycompany.simple.Simple](#mycompany-simple-simple)
  - [Workflows](#mycompany-simple-simple-workflows)
    - [mycompany.simple.Simple.ExampleContinueAsNew](#mycompany-simple-simple-examplecontinueasnew-workflow)
    - [mycompany.simple.Simple.SomeWorkflow3](#mycompany-simple-simple-someworkflow3-workflow)
    - [mycompany.simple.Simple.SomeWorkflow4](#mycompany-simple-simple-someworkflow4-workflow)
    - [mycompany.simple.SomeWorkflow1](#mycompany-simple-someworkflow1-workflow)
    - [mycompany.simple.SomeWorkflow2](#mycompany-simple-someworkflow2-workflow)
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
    - [mycompany.simple.Simple.SomeActivity2](#mycompany-simple-simple-someactivity2-activity)
    - [mycompany.simple.Simple.SomeActivity3](#mycompany-simple-simple-someactivity3-activity)
    - [mycompany.simple.Simple.SomeActivity4](#mycompany-simple-simple-someactivity4-activity)
    - [mycompany.simple.Simple.SomeSignal1](#mycompany-simple-simple-somesignal1-activity)
    - [mycompany.simple.Simple.SomeSignal2](#mycompany-simple-simple-somesignal2-activity)
    - [mycompany.simple.Simple.SomeSignal3](#mycompany-simple-simple-somesignal3-activity)
    - [mycompany.simple.Simple.SomeUpdate1](#mycompany-simple-simple-someupdate1-activity)
    - [mycompany.simple.SomeActivity1](#mycompany-simple-someactivity1-activity)
- Messages
  - [mycompany.simple.ExampleContinueAsNewRequest](#mycompany-simple-examplecontinueasnewrequest)
  - [mycompany.simple.ExampleContinueAsNewResponse](#mycompany-simple-examplecontinueasnewresponse)
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

<a name="mycompany-simple-services"></a>
## Services

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
<td><a href="../common/v1/README.md#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
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
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-empty">google.protobuf.Empty</a></td>
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
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
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

**Input:** [mycompany.simple.common.v1.PaginatedRequest](../common/v1/README.md#mycompany-simple-common-v1-paginatedrequest)

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

**Output:** [mycompany.simple.common.v1.PaginatedResponse](../common/v1/README.md#mycompany-simple-common-v1-paginatedresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-any">google.protobuf.Any</a></td>
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
<td><a href="../common/v1/README.md#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
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
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-empty">google.protobuf.Empty</a></td>
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
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
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

<a name="mycompany-simple-simple"></a>
## mycompany.simple.Simple

<a name="mycompany-simple-simple-workflows"></a>
### Workflows

---
<a name="mycompany-simple-simple-examplecontinueasnew-workflow"></a>
### mycompany.simple.Simple.ExampleContinueAsNew

**Input:** [mycompany.simple.ExampleContinueAsNewRequest](#mycompany-simple-examplecontinueasnewrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>remaining</td>
<td>int32</td>
<td><pre>
json_name: remaining
go_name: Remaining</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="../../../temporal/v1/README.md#temporal-v1-retrypolicy">temporal.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.ExampleContinueAsNewResponse](#mycompany-simple-examplecontinueasnewresponse)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>example-continue-as-new/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
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

**Input:** [mycompany.simple.common.v1.PaginatedRequest](../common/v1/README.md#mycompany-simple-common-v1-paginatedrequest)

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

**Output:** [mycompany.simple.common.v1.PaginatedResponse](../common/v1/README.md#mycompany-simple-common-v1-paginatedresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-any">google.protobuf.Any</a></td>
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

---
<a name="mycompany-simple-someactivity1-activity"></a>
### mycompany.simple.SomeActivity1

<pre>
SomeActivity1 does some activity thing.
</pre>   

<a name="mycompany-simple-messages"></a>
## Messages

<a name="mycompany-simple-examplecontinueasnewrequest"></a>
### mycompany.simple.ExampleContinueAsNewRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>remaining</td>
<td>int32</td>
<td><pre>
json_name: remaining
go_name: Remaining</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="../../../temporal/v1/README.md#temporal-v1-retrypolicy">temporal.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr>
</table>



<a name="mycompany-simple-examplecontinueasnewresponse"></a>
### mycompany.simple.ExampleContinueAsNewResponse



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
<td><a href="../common/v1/README.md#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
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
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-empty">google.protobuf.Empty</a></td>
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
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
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

