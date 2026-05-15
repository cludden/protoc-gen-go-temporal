

<a name="issue_125-v1"></a>
# issue_125.v1

## Table of Contents
- [issue_125.v1.Issue125Service](#issue_125-v1-issue125service)
  - [Workflows](#issue_125-v1-issue125service-workflows)
    - [issue_125.v1.Issue125Service.Foo](#issue_125-v1-issue125service-foo-workflow)
  - [Updates](#issue_125-v1-issue125service-updates)
    - [issue_125.v1.Issue125Service.Bar](#issue_125-v1-issue125service-bar-update)
    - [issue_125.v1.Issue125Service.Baz](#issue_125-v1-issue125service-baz-update)
- Messages
  - [issue_125.v1.BarInput](#issue_125-v1-barinput)
  - [issue_125.v1.BarOutput](#issue_125-v1-baroutput)
  - [issue_125.v1.BazInput](#issue_125-v1-bazinput)
  - [issue_125.v1.BazOutput](#issue_125-v1-bazoutput)
  - [issue_125.v1.FooInput](#issue_125-v1-fooinput)
  - [issue_125.v1.FooOutput](#issue_125-v1-foooutput)

<a name="issue_125-v1-services"></a>
## Services

<a name="issue_125-v1-issue125service"></a>
## issue_125.v1.Issue125Service

<a name="issue_125-v1-issue125service-workflows"></a>
### Workflows

---
<a name="issue_125-v1-issue125service-foo-workflow"></a>
### issue_125.v1.Issue125Service.Foo

**Input:** [issue_125.v1.FooInput](#issue_125-v1-fooinput)

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
</tr>
</table>

**Output:** [issue_125.v1.FooOutput](#issue_125-v1-foooutput)

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
<tr><td>id</td><td><pre><code>foo/${! id }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#issue_125-v1-issue125service-bar-update">issue_125.v1.Issue125Service.Bar</a></td></tr>
<tr><td><a href="#issue_125-v1-issue125service-baz-update">issue_125.v1.Issue125Service.Baz</a></td></tr>
</table>    

<a name="issue_125-v1-issue125service-updates"></a>
### Updates

---
<a name="issue_125-v1-issue125service-bar-update"></a>
### issue_125.v1.Issue125Service.Bar



**Input:** [issue_125.v1.BarInput](#issue_125-v1-barinput)

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
</tr>
</table>

**Output:** [issue_125.v1.BarOutput](#issue_125-v1-baroutput)

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

---
<a name="issue_125-v1-issue125service-baz-update"></a>
### issue_125.v1.Issue125Service.Baz



**Input:** [issue_125.v1.BazInput](#issue_125-v1-bazinput)

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
</tr>
</table>

**Output:** [issue_125.v1.BazOutput](#issue_125-v1-bazoutput)

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

<a name="issue_125-v1-messages"></a>
## Messages

<a name="issue_125-v1-barinput"></a>
### issue_125.v1.BarInput

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
</tr>
</table>



<a name="issue_125-v1-baroutput"></a>
### issue_125.v1.BarOutput

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



<a name="issue_125-v1-bazinput"></a>
### issue_125.v1.BazInput

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
</tr>
</table>



<a name="issue_125-v1-bazoutput"></a>
### issue_125.v1.BazOutput

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



<a name="issue_125-v1-fooinput"></a>
### issue_125.v1.FooInput

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
</tr>
</table>



<a name="issue_125-v1-foooutput"></a>
### issue_125.v1.FooOutput

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

