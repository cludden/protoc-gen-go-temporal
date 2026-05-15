

<a name="test-xnserr-v1"></a>
# test.xnserr.v1

## Table of Contents
- [test.xnserr.v1.Client](#test-xnserr-v1-client)
  - [Workflows](#test-xnserr-v1-client-workflows)
    - [test.xnserr.v1.Client.CallSleep](#test-xnserr-v1-client-callsleep-workflow)
- [test.xnserr.v1.Server](#test-xnserr-v1-server)
  - [Workflows](#test-xnserr-v1-server-workflows)
    - [test.xnserr.v1.Server.Sleep](#test-xnserr-v1-server-sleep-workflow)
- Messages
  - [test.xnserr.v1.CallSleepRequest](#test-xnserr-v1-callsleeprequest)
  - [test.xnserr.v1.Failure](#test-xnserr-v1-failure)
  - [test.xnserr.v1.FailureInfo](#test-xnserr-v1-failureinfo)
  - [test.xnserr.v1.SleepRequest](#test-xnserr-v1-sleeprequest)

<a name="test-xnserr-v1-services"></a>
## Services

<a name="test-xnserr-v1-client"></a>
## test.xnserr.v1.Client

<a name="test-xnserr-v1-client-workflows"></a>
### Workflows

---
<a name="test-xnserr-v1-client-callsleep-workflow"></a>
### test.xnserr.v1.Client.CallSleep

**Input:** [test.xnserr.v1.CallSleepRequest](#test-xnserr-v1-callsleeprequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="../../../temporal/xns/v1/README.md#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr><tr>
<td>start_workflow_options</td>
<td><a href="../../../temporal/xns/v1/README.md#temporal-xns-v1-startworkflowoptions">temporal.xns.v1.StartWorkflowOptions</a></td>
<td><pre>
json_name: startWorkflowOptions
go_name: StartWorkflowOptions</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="test-xnserr-v1-server"></a>
## test.xnserr.v1.Server

<a name="test-xnserr-v1-server-workflows"></a>
### Workflows

---
<a name="test-xnserr-v1-server-sleep-workflow"></a>
### test.xnserr.v1.Server.Sleep

**Input:** [test.xnserr.v1.SleepRequest](#test-xnserr-v1-sleeprequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
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
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE</code></pre></td></tr>
<tr><td>parent_close_policy</td><td><pre><code>PARENT_CLOSE_POLICY_REQUEST_CANCEL</code></pre></td></tr>
<tr><td>xns.heartbeat_interval</td><td>10 seconds</td></tr>
<tr><td>xns.heartbeat_timeout</td><td>30 seconds</td></tr>
</table>     

<a name="test-xnserr-v1-messages"></a>
## Messages

<a name="test-xnserr-v1-callsleeprequest"></a>
### test.xnserr.v1.CallSleepRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="../../../temporal/xns/v1/README.md#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr><tr>
<td>start_workflow_options</td>
<td><a href="../../../temporal/xns/v1/README.md#temporal-xns-v1-startworkflowoptions">temporal.xns.v1.StartWorkflowOptions</a></td>
<td><pre>
json_name: startWorkflowOptions
go_name: StartWorkflowOptions</pre></td>
</tr>
</table>



<a name="test-xnserr-v1-failure"></a>
### test.xnserr.v1.Failure

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>application_error_type</td>
<td>string</td>
<td><pre>
json_name: applicationErrorType
go_name: ApplicationErrorType</pre></td>
</tr><tr>
<td>info</td>
<td><a href="#test-xnserr-v1-failureinfo">test.xnserr.v1.FailureInfo</a></td>
<td><pre>
json_name: info
go_name: Info</pre></td>
</tr><tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr><tr>
<td>non_retryable</td>
<td>bool</td>
<td><pre>
json_name: nonRetryable
go_name: NonRetryable</pre></td>
</tr>
</table>



<a name="test-xnserr-v1-failureinfo"></a>
### test.xnserr.v1.FailureInfo

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>FAILURE_INFO_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_APPLICATION_ERROR</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_TIMEOUT</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_CANCELED</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_TERMINATED</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_ACTIVITY</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_WORKFLOW_EXECUTION</td>
<td></td>
</tr><tr>
<td>FAILURE_INFO_CHILD_WORKFLOW_EXECUTION</td>
<td></td>
</tr>
</table>

<a name="test-xnserr-v1-sleeprequest"></a>
### test.xnserr.v1.SleepRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>failure</td>
<td><a href="#test-xnserr-v1-failure">test.xnserr.v1.Failure</a></td>
<td><pre>
json_name: failure
go_name: Failure</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>

