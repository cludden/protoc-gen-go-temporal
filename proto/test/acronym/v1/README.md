

<a name="test-acronym-v1"></a>
# test.acronym.v1

## Table of Contents
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

