# Table of Contents

- [test.patch](#test-patch)
  - Services
    - [test.patch.FooService](#test-patch-fooservice)
      - [Workflows](#test-patch-fooservice-workflows)
        - [test.patch.FooService.Foo](#test-patch-fooservice-foo-workflow)
      - [Activities](#test-patch-fooservice-activities)
        - [test.patch.FooService.Foo](#test-patch-fooservice-foo-activity)
  - Messages
    - [test.patch.FooInput](#test-patch-fooinput)
    - [test.patch.FooOutput](#test-patch-foooutput)

<a name="test-patch"></a>
# test.patch

<a name="test-patch-services"></a>
## Services

<a name="test-patch-fooservice"></a>
## test.patch.FooService

<a name="test-patch-fooservice-workflows"></a>
### Workflows

---
<a name="test-patch-fooservice-foo-workflow"></a>
### test.patch.FooService.Foo

**Input:** [test.patch.FooInput](#test-patch-fooinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo_id</td>
<td>string</td>
<td><pre>
json_name: fooID
go_name: FooId
go_tags: json:"fooID"</pre></td>
</tr>
</table>

**Output:** [test.patch.FooOutput](#test-patch-foooutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>    

<a name="test-patch-fooservice-activities"></a>
### Activities

---
<a name="test-patch-fooservice-foo-activity"></a>
### test.patch.FooService.Foo



**Input:** [test.patch.FooInput](#test-patch-fooinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo_id</td>
<td>string</td>
<td><pre>
json_name: fooID
go_name: FooId
go_tags: json:"fooID"</pre></td>
</tr>
</table>

**Output:** [test.patch.FooOutput](#test-patch-foooutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>2 seconds</td></tr>
</table>   

<a name="test-patch-messages"></a>
## Messages

<a name="test-patch-fooinput"></a>
### test.patch.FooInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo_id</td>
<td>string</td>
<td><pre>
json_name: fooID
go_name: FooId
go_tags: json:"fooID"</pre></td>
</tr>
</table>



<a name="test-patch-foooutput"></a>
### test.patch.FooOutput

