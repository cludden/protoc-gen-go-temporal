

<a name="test-editions"></a>
# test.editions

## Table of Contents
- [test.editions.FooService](#test-editions-fooservice)
  - [Workflows](#test-editions-fooservice-workflows)
    - [test.editions.FooService.Foo](#test-editions-fooservice-foo-workflow)
  - [Activities](#test-editions-fooservice-activities)
    - [test.editions.FooService.Foo](#test-editions-fooservice-foo-activity)
- Messages
  - [test.editions.FooInput](#test-editions-fooinput)
  - [test.editions.FooOutput](#test-editions-foooutput)

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

