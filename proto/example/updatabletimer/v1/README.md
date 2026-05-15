

<a name="example-updatabletimer-v1"></a>
# example.updatabletimer.v1

## Table of Contents
- [example.updatabletimer.v1.Example](#example-updatabletimer-v1-example)
  - [Workflows](#example-updatabletimer-v1-example-workflows)
    - [UpdatableTimer](#updatabletimer-workflow)
  - [Queries](#example-updatabletimer-v1-example-queries)
    - [example.updatabletimer.v1.Example.GetWakeUpTime](#example-updatabletimer-v1-example-getwakeuptime-query)
  - [Signals](#example-updatabletimer-v1-example-signals)
    - [example.updatabletimer.v1.Example.UpdateWakeUpTime](#example-updatabletimer-v1-example-updatewakeuptime-signal)
- Messages
  - [example.updatabletimer.v1.GetWakeUpTimeOutput](#example-updatabletimer-v1-getwakeuptimeoutput)
  - [example.updatabletimer.v1.UpdatableTimerInput](#example-updatabletimer-v1-updatabletimerinput)
  - [example.updatabletimer.v1.UpdateWakeUpTimeInput](#example-updatabletimer-v1-updatewakeuptimeinput)

<a name="example-updatabletimer-v1-services"></a>
## Services

<a name="example-updatabletimer-v1-example"></a>
## example.updatabletimer.v1.Example

<a name="example-updatabletimer-v1-example-workflows"></a>
### Workflows

---
<a name="updatabletimer-workflow"></a>
### UpdatableTimer

<pre>
UpdatableTimer describes an updatable timer workflow
</pre>

**Input:** [example.updatabletimer.v1.UpdatableTimerInput](#example-updatabletimer-v1-updatabletimerinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>initial_wake_up_time</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: initialWakeUpTime
go_name: InitialWakeUpTime</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>updatable-timer/${! name.or(uuid_v4()) }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#example-updatabletimer-v1-example-getwakeuptime-query">example.updatabletimer.v1.Example.GetWakeUpTime</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-updatabletimer-v1-example-updatewakeuptime-signal">example.updatabletimer.v1.Example.UpdateWakeUpTime</a></td><td>false</td></tr>
</table>  

<a name="example-updatabletimer-v1-example-queries"></a>
### Queries

---
<a name="example-updatabletimer-v1-example-getwakeuptime-query"></a>
### example.updatabletimer.v1.Example.GetWakeUpTime

<pre>
GetWakeUpTime retrieves the current timer expiration timestamp
</pre>

**Output:** [example.updatabletimer.v1.GetWakeUpTimeOutput](#example-updatabletimer-v1-getwakeuptimeoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>  

<a name="example-updatabletimer-v1-example-signals"></a>
### Signals

---
<a name="example-updatabletimer-v1-example-updatewakeuptime-signal"></a>
### example.updatabletimer.v1.Example.UpdateWakeUpTime

<pre>
UpdateWakeUpTime updates the timer expiration timestamp
</pre>

**Input:** [example.updatabletimer.v1.UpdateWakeUpTimeInput](#example-updatabletimer-v1-updatewakeuptimeinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>   

<a name="example-updatabletimer-v1-messages"></a>
## Messages

<a name="example-updatabletimer-v1-getwakeuptimeoutput"></a>
### example.updatabletimer.v1.GetWakeUpTimeOutput

<pre>
GetWakeUpTimeOutput describes the input to a GetWakeUpTime query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>



<a name="example-updatabletimer-v1-updatabletimerinput"></a>
### example.updatabletimer.v1.UpdatableTimerInput

<pre>
UpdatableTimerInput describes the input to a UpdatableTimer workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>initial_wake_up_time</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: initialWakeUpTime
go_name: InitialWakeUpTime</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="example-updatabletimer-v1-updatewakeuptimeinput"></a>
### example.updatabletimer.v1.UpdateWakeUpTimeInput

<pre>
UpdateWakeUpTimeInput describes the input to a UpdateWakeUpTime signal
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>

