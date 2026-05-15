

<a name="example-helloworld-v1"></a>
# example.helloworld.v1

## Table of Contents
- [example.helloworld.v1.Example](#example-helloworld-v1-example)
  - [Workflows](#example-helloworld-v1-example-workflows)
    - [example.v1.Hello](#example-v1-hello-workflow)
  - [Signals](#example-helloworld-v1-example-signals)
    - [example.helloworld.v1.Example.Goodbye](#example-helloworld-v1-example-goodbye-signal)
- Messages
  - [example.helloworld.v1.GoodbyeRequest](#example-helloworld-v1-goodbyerequest)
  - [example.helloworld.v1.HelloRequest](#example-helloworld-v1-hellorequest)
  - [example.helloworld.v1.HelloResponse](#example-helloworld-v1-helloresponse)

<a name="example-helloworld-v1-services"></a>
## Services

<a name="example-helloworld-v1-example"></a>
## example.helloworld.v1.Example

<a name="example-helloworld-v1-example-workflows"></a>
### Workflows

---
<a name="example-v1-hello-workflow"></a>
### example.v1.Hello

<pre>
Hello prints a friendly greeting and waits for goodbye
</pre>

**Input:** [example.helloworld.v1.HelloRequest](#example-helloworld-v1-hellorequest)

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

**Output:** [example.helloworld.v1.HelloResponse](#example-helloworld-v1-helloresponse)

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
<tr><td>id</td><td><pre><code>hello/${! name.or("World") }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-helloworld-v1-example-goodbye-signal">example.helloworld.v1.Example.Goodbye</a></td><td>false</td></tr>
</table>   

<a name="example-helloworld-v1-example-signals"></a>
### Signals

---
<a name="example-helloworld-v1-example-goodbye-signal"></a>
### example.helloworld.v1.Example.Goodbye

<pre>
Goodbye signals a running workflow to exit
</pre>

**Input:** [example.helloworld.v1.GoodbyeRequest](#example-helloworld-v1-goodbyerequest)

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

<a name="example-helloworld-v1-messages"></a>
## Messages

<a name="example-helloworld-v1-goodbyerequest"></a>
### example.helloworld.v1.GoodbyeRequest

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



<a name="example-helloworld-v1-hellorequest"></a>
### example.helloworld.v1.HelloRequest

<pre>
HelloRequest describes the input to a Hello workflow
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
</tr>
</table>



<a name="example-helloworld-v1-helloresponse"></a>
### example.helloworld.v1.HelloResponse

<pre>
HelloResponse describes the output from a Hello workflow
</pre>

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

