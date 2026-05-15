

<a name="example-searchattributes-v1"></a>
# example.searchattributes.v1

## Table of Contents
- [example.searchattributes.v1.Example](#example-searchattributes-v1-example)
  - [Workflows](#example-searchattributes-v1-example-workflows)
    - [example.searchattributes.v1.Example.SearchAttributes](#example-searchattributes-v1-example-searchattributes-workflow)
    - [example.searchattributes.v1.Example.TypedSearchAttributes](#example-searchattributes-v1-example-typedsearchattributes-workflow)
- Messages
  - [example.searchattributes.v1.SearchAttributesInput](#example-searchattributes-v1-searchattributesinput)
  - [example.searchattributes.v1.TypedSearchAttributesInput](#example-searchattributes-v1-typedsearchattributesinput)
  - [example.searchattributes.v1.TypedSearchAttributesOutput](#example-searchattributes-v1-typedsearchattributesoutput)

<a name="example-searchattributes-v1-services"></a>
## Services

<a name="example-searchattributes-v1-example"></a>
## example.searchattributes.v1.Example

<a name="example-searchattributes-v1-example-workflows"></a>
### Workflows

---
<a name="example-searchattributes-v1-example-searchattributes-workflow"></a>
### example.searchattributes.v1.Example.SearchAttributes

**Input:** [example.searchattributes.v1.SearchAttributesInput](#example-searchattributes-v1-searchattributesinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>search_attributes_${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
<tr><td>search_attributes</td><td><pre><code>CustomKeywordField = customKeywordField 
CustomTextField = customTextField 
CustomIntField = customIntField.int64() 
CustomDoubleField = customDoubleField 
CustomBoolField = customBoolField 
CustomDatetimeField = customDatetimeField.ts_parse("2006-01-02T15:04:05Z")</code></pre></td></tr>
</table>

---
<a name="example-searchattributes-v1-example-typedsearchattributes-workflow"></a>
### example.searchattributes.v1.Example.TypedSearchAttributes

**Input:** [example.searchattributes.v1.TypedSearchAttributesInput](#example-searchattributes-v1-typedsearchattributesinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_keyword_list_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordListField
go_name: CustomKeywordListField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>

**Output:** [example.searchattributes.v1.TypedSearchAttributesOutput](#example-searchattributes-v1-typedsearchattributesoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_keyword_list_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordListField
go_name: CustomKeywordListField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>searchattributes.v1.TypedSearchAttributes/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="example-searchattributes-v1-messages"></a>
## Messages

<a name="example-searchattributes-v1-searchattributesinput"></a>
### example.searchattributes.v1.SearchAttributesInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>



<a name="example-searchattributes-v1-typedsearchattributesinput"></a>
### example.searchattributes.v1.TypedSearchAttributesInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_keyword_list_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordListField
go_name: CustomKeywordListField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>



<a name="example-searchattributes-v1-typedsearchattributesoutput"></a>
### example.searchattributes.v1.TypedSearchAttributesOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_keyword_list_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordListField
go_name: CustomKeywordListField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>

