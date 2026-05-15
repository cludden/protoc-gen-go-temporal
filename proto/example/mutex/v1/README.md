

<a name="example-mutex-v1"></a>
# example.mutex.v1

## Table of Contents
- [example.mutex.v1.Example](#example-mutex-v1-example)
  - [Workflows](#example-mutex-v1-example-workflows)
    - [example.mutex.v1.Example.SampleWorkflowWithMutex](#example-mutex-v1-example-sampleworkflowwithmutex-workflow)
    - [mutex.v1.Mutex](#mutex-v1-mutex-workflow)
  - [Signals](#example-mutex-v1-example-signals)
    - [mutex.v1.ReleaseLock](#mutex-v1-releaselock-signal)
  - [Updates](#example-mutex-v1-example-updates)
    - [mutex.v1.AcquireLock](#mutex-v1-acquirelock-update)
  - [Activities](#example-mutex-v1-example-activities)
    - [example.mutex.v1.Example.Mutex](#example-mutex-v1-example-mutex-activity)
- Messages
  - [example.mutex.v1.AcquireLockInput](#example-mutex-v1-acquirelockinput)
  - [example.mutex.v1.AcquireLockOutput](#example-mutex-v1-acquirelockoutput)
  - [example.mutex.v1.MutexInput](#example-mutex-v1-mutexinput)
  - [example.mutex.v1.ReleaseLockInput](#example-mutex-v1-releaselockinput)
  - [example.mutex.v1.SampleWorkflowWithMutexInput](#example-mutex-v1-sampleworkflowwithmutexinput)

<a name="example-mutex-v1-services"></a>
## Services

<a name="example-mutex-v1-example"></a>
## example.mutex.v1.Example

<a name="example-mutex-v1-example-workflows"></a>
### Workflows

---
<a name="example-mutex-v1-example-sampleworkflowwithmutex-workflow"></a>
### example.mutex.v1.Example.SampleWorkflowWithMutex

<pre>
SampleWorkflowWithMutex is a sample workflow that demonstrates how to
use the Mutex service.
</pre>

**Input:** [example.mutex.v1.SampleWorkflowWithMutexInput](#example-mutex-v1-sampleworkflowwithmutexinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>SampleWorkflow1WithMutex_${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="mutex-v1-mutex-workflow"></a>
### mutex.v1.Mutex

<pre>
Mutex is a workflow that manages concurrent access to a resource
identified by `resource_id`.
</pre>

**Input:** [example.mutex.v1.MutexInput](#example-mutex-v1-mutexinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>mutex:${! resourceId }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
<tr><td>retry_policy.backoff_coefficient</td><td>2</td></tr>
<tr><td>retry_policy.initial_interval</td><td>1 second</td></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>retry_policy.max_interval</td><td>1 minute</td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-mutex-v1-example-releaselock-signal">example.mutex.v1.Example.ReleaseLock</a></td><td>false</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#example-mutex-v1-example-acquirelock-update">example.mutex.v1.Example.AcquireLock</a></td></tr>
</table>   

<a name="example-mutex-v1-example-signals"></a>
### Signals

---
<a name="mutex-v1-releaselock-signal"></a>
### mutex.v1.ReleaseLock

<pre>
ReleaseLock releases a lock on a resource identified by `lease_id`.
</pre>

**Input:** [example.mutex.v1.ReleaseLockInput](#example-mutex-v1-releaselockinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>  

<a name="example-mutex-v1-example-updates"></a>
### Updates

---
<a name="mutex-v1-acquirelock-update"></a>
### mutex.v1.AcquireLock

<pre>
AcquireLock requests a lock on a resource identified by `resource_id`
and blocks until the lock is acquired, returning a `lease_id` that
can be used to release the lock.
</pre>

**Input:** [example.mutex.v1.AcquireLockInput](#example-mutex-v1-acquirelockinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>timeout</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: timeout
go_name: Timeout</pre></td>
</tr>
</table>

**Output:** [example.mutex.v1.AcquireLockOutput](#example-mutex-v1-acquirelockoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>

<a name="example-mutex-v1-example-activities"></a>
### Activities

---
<a name="example-mutex-v1-example-mutex-activity"></a>
### example.mutex.v1.Example.Mutex

<pre>
Mutex is a workflow that manages concurrent access to a resource
identified by `resource_id`.
</pre>

**Input:** [example.mutex.v1.MutexInput](#example-mutex-v1-mutexinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table>   

<a name="example-mutex-v1-messages"></a>
## Messages

<a name="example-mutex-v1-acquirelockinput"></a>
### example.mutex.v1.AcquireLockInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>timeout</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: timeout
go_name: Timeout</pre></td>
</tr>
</table>



<a name="example-mutex-v1-acquirelockoutput"></a>
### example.mutex.v1.AcquireLockOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>



<a name="example-mutex-v1-mutexinput"></a>
### example.mutex.v1.MutexInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr>
</table>



<a name="example-mutex-v1-releaselockinput"></a>
### example.mutex.v1.ReleaseLockInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>



<a name="example-mutex-v1-sampleworkflowwithmutexinput"></a>
### example.mutex.v1.SampleWorkflowWithMutexInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>

