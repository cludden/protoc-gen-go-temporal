# Table of Contents

- [example.protocgengonexus](#example-protocgengonexus)
  - Services
    - [example.protocgengonexus.Greeting](#example-protocgengonexus-greeting)
      - [Workflows](#example-protocgengonexus-greeting-workflows)
        - [example.protocgengonexus.Greeting.Greet](#example-protocgengonexus-greeting-greet-workflow)
      - [Activities](#example-protocgengonexus-greeting-activities)
        - [example.protocgengonexus.Greeting.GenerateGreeting](#example-protocgengonexus-greeting-generategreeting-activity)
    - [example.protocgengonexus.Caller](#example-protocgengonexus-caller)
      - [Workflows](#example-protocgengonexus-caller-workflows)
        - [example.protocgengonexus.Caller.CallGreet](#example-protocgengonexus-caller-callgreet-workflow)
  - Messages
    - [example.protocgengonexus.CallGreetInput](#example-protocgengonexus-callgreetinput)
    - [example.protocgengonexus.CallGreetOutput](#example-protocgengonexus-callgreetoutput)
    - [example.protocgengonexus.GenerateGreetingInput](#example-protocgengonexus-generategreetinginput)
    - [example.protocgengonexus.GenerateGreetingOutput](#example-protocgengonexus-generategreetingoutput)
    - [example.protocgengonexus.GreetInput](#example-protocgengonexus-greetinput)
    - [example.protocgengonexus.GreetOutput](#example-protocgengonexus-greetoutput)

<a name="example-protocgengonexus"></a>
# example.protocgengonexus

<a name="example-protocgengonexus-services"></a>
## Services

<a name="example-protocgengonexus-greeting"></a>
## example.protocgengonexus.Greeting

<a name="example-protocgengonexus-greeting-workflows"></a>
### Workflows

---
<a name="example-protocgengonexus-greeting-greet-workflow"></a>
### example.protocgengonexus.Greeting.Greet

**Input:** [example.protocgengonexus.GreetInput](#example-protocgengonexus-greetinput)

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

**Output:** [example.protocgengonexus.GreetOutput](#example-protocgengonexus-greetoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>greeting</td>
<td>string</td>
<td><pre>
json_name: greeting
go_name: Greeting</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>Greet/${! name.or("World") }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>    

<a name="example-protocgengonexus-greeting-activities"></a>
### Activities

---
<a name="example-protocgengonexus-greeting-generategreeting-activity"></a>
### example.protocgengonexus.Greeting.GenerateGreeting



**Input:** [example.protocgengonexus.GenerateGreetingInput](#example-protocgengonexus-generategreetinginput)

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

**Output:** [example.protocgengonexus.GenerateGreetingOutput](#example-protocgengonexus-generategreetingoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>greeting</td>
<td>string</td>
<td><pre>
json_name: greeting
go_name: Greeting</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table>   

<a name="example-protocgengonexus-caller"></a>
## example.protocgengonexus.Caller

<a name="example-protocgengonexus-caller-workflows"></a>
### Workflows

---
<a name="example-protocgengonexus-caller-callgreet-workflow"></a>
### example.protocgengonexus.Caller.CallGreet

**Input:** [example.protocgengonexus.CallGreetInput](#example-protocgengonexus-callgreetinput)

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

**Output:** [example.protocgengonexus.CallGreetOutput](#example-protocgengonexus-callgreetoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>greeting</td>
<td>string</td>
<td><pre>
json_name: greeting
go_name: Greeting</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>CallGreet/${! name.or("World") }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="example-protocgengonexus-messages"></a>
## Messages

<a name="example-protocgengonexus-callgreetinput"></a>
### example.protocgengonexus.CallGreetInput

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



<a name="example-protocgengonexus-callgreetoutput"></a>
### example.protocgengonexus.CallGreetOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>greeting</td>
<td>string</td>
<td><pre>
json_name: greeting
go_name: Greeting</pre></td>
</tr>
</table>



<a name="example-protocgengonexus-generategreetinginput"></a>
### example.protocgengonexus.GenerateGreetingInput

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



<a name="example-protocgengonexus-generategreetingoutput"></a>
### example.protocgengonexus.GenerateGreetingOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>greeting</td>
<td>string</td>
<td><pre>
json_name: greeting
go_name: Greeting</pre></td>
</tr>
</table>



<a name="example-protocgengonexus-greetinput"></a>
### example.protocgengonexus.GreetInput

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



<a name="example-protocgengonexus-greetoutput"></a>
### example.protocgengonexus.GreetOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>greeting</td>
<td>string</td>
<td><pre>
json_name: greeting
go_name: Greeting</pre></td>
</tr>
</table>

