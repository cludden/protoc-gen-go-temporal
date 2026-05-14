

<a name="example-xns-v1"></a>
# example.xns.v1

<a name="example-xns-v1-services"></a>
## Services

<a name="example-xns-v1-xns"></a>
## example.xns.v1.Xns

<a name="example-xns-v1-xns-workflows"></a>
### Workflows

---
<a name="example-xns-v1-xns-provisionfoo-workflow"></a>
### example.xns.v1.Xns.ProvisionFoo

**Input:** [example.xns.v1.ProvisionFooRequest](#example-xns-v1-provisionfoorequest)

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
go_name: Name</pre></td>
</tr>
</table>

**Output:** [example.xns.v1.ProvisionFooResponse](#example-xns-v1-provisionfooresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
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

<a name="example-xns-v1-example"></a>
## example.xns.v1.Example

<a name="example-xns-v1-example-workflows"></a>
### Workflows

---
<a name="example-xns-v1-example-createfoo-workflow"></a>
### example.xns.v1.Example.CreateFoo

<pre>
CreateFoo creates a new foo operation
</pre>

**Input:** [example.xns.v1.CreateFooRequest](#example-xns-v1-createfoorequest)

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
go_name: Name</pre></td>
</tr>
</table>

**Output:** [example.xns.v1.CreateFooResponse](#example-xns-v1-createfooresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
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
<tr><td>xns.heartbeat_interval</td><td>10 seconds</td></tr>
<tr><td>xns.heartbeat_timeout</td><td>20 seconds</td></tr>
<tr><td>xns.start_to_close_timeout</td><td>1 hour 30 seconds</td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#example-xns-v1-example-getfooprogress-query">example.xns.v1.Example.GetFooProgress</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-xns-v1-example-setfooprogress-signal">example.xns.v1.Example.SetFooProgress</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#example-xns-v1-example-updatefooprogress-update">example.xns.v1.Example.UpdateFooProgress</a></td></tr>
</table>  

<a name="example-xns-v1-example-queries"></a>
### Queries

---
<a name="example-xns-v1-example-getfooprogress-query"></a>
### example.xns.v1.Example.GetFooProgress

<pre>
GetFooProgress returns the status of a CreateFoo operation
</pre>

**Output:** [example.xns.v1.GetFooProgressResponse](#example-xns-v1-getfooprogressresponse)

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
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>  

<a name="example-xns-v1-example-signals"></a>
### Signals

---
<a name="example-xns-v1-example-setfooprogress-signal"></a>
### example.xns.v1.Example.SetFooProgress

<pre>
SetFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.xns.v1.SetFooProgressRequest](#example-xns-v1-setfooprogressrequest)

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

<a name="example-xns-v1-example-updates"></a>
### Updates

---
<a name="example-xns-v1-example-updatefooprogress-update"></a>
### example.xns.v1.Example.UpdateFooProgress

<pre>
UpdateFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.xns.v1.SetFooProgressRequest](#example-xns-v1-setfooprogressrequest)

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

**Output:** [example.xns.v1.GetFooProgressResponse](#example-xns-v1-getfooprogressresponse)

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
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
</table>

<a name="example-xns-v1-example-activities"></a>
### Activities

---
<a name="example-xns-v1-example-notify-activity"></a>
### example.xns.v1.Example.Notify

<pre>
Notify sends a notification
</pre>

**Input:** [example.xns.v1.NotifyRequest](#example-xns-v1-notifyrequest)

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

<a name="example-xns-v1-messages"></a>
## Messages

<a name="example-xns-v1-createfoorequest"></a>
### example.xns.v1.CreateFooRequest

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
go_name: Name</pre></td>
</tr>
</table>



<a name="example-xns-v1-createfooresponse"></a>
### example.xns.v1.CreateFooResponse

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
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="example-xns-v1-foo"></a>
### example.xns.v1.Foo

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
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-xns-v1-foo-status"></a>
### example.xns.v1.Foo.Status

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

<a name="example-xns-v1-getfooprogressresponse"></a>
### example.xns.v1.GetFooProgressResponse

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
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-xns-v1-notifyrequest"></a>
### example.xns.v1.NotifyRequest

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



<a name="example-xns-v1-provisionfoorequest"></a>
### example.xns.v1.ProvisionFooRequest

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
go_name: Name</pre></td>
</tr>
</table>



<a name="example-xns-v1-provisionfooresponse"></a>
### example.xns.v1.ProvisionFooResponse

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
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="example-xns-v1-setfooprogressrequest"></a>
### example.xns.v1.SetFooProgressRequest

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

