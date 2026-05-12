

<a name="temporal-xns-v1"></a>
# temporal.xns.v1

<a name="temporal-xns-v1-messages"></a>
## Messages

<a name="temporal-xns-v1-idreusepolicy"></a>
### temporal.xns.v1.IDReusePolicy

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE</td>
<td><pre>
Allow starting a workflow execution using the same workflow id.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY</td>
<td><pre>
Allow starting a workflow execution using the same workflow id, only when the last
execution's final state is one of [terminated, cancelled, timed out, failed].
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE</td>
<td><pre>
Do not permit re-use of the workflow id for this workflow. Future start workflow requests
could potentially change the policy, allowing re-use of the workflow id.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING</td>
<td><pre>
If a workflow is running using the same workflow ID, terminate it and start a new one.
If no running workflow, then the behavior is the same as ALLOW_DUPLICATE
</pre></td>
</tr>
</table>

<a name="temporal-xns-v1-retrypolicy"></a>
### temporal.xns.v1.RetryPolicy

<pre>
RetryPolicy describes configuration for activity or child workflow retries
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>backoff_coefficient</td>
<td>double</td>
<td><pre>
json_name: backoffCoefficient
go_name: BackoffCoefficient</pre></td>
</tr><tr>
<td>initial_interval</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: initialInterval
go_name: InitialInterval</pre></td>
</tr><tr>
<td>max_attempts</td>
<td>int32</td>
<td><pre>
json_name: maxAttempts
go_name: MaxAttempts</pre></td>
</tr><tr>
<td>max_interval</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: maxInterval
go_name: MaxInterval</pre></td>
</tr><tr>
<td>non_retryable_error_types</td>
<td>string</td>
<td><pre>
json_name: nonRetryableErrorTypes
go_name: NonRetryableErrorTypes</pre></td>
</tr>
</table>



<a name="temporal-xns-v1-startworkflowoptions"></a>
### temporal.xns.v1.StartWorkflowOptions

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>enable_eager_start</td>
<td>bool</td>
<td><pre>
json_name: enableEagerStart
go_name: EnableEagerStart</pre></td>
</tr><tr>
<td>error_when_already_started</td>
<td>bool</td>
<td><pre>
json_name: errorWhenAlreadyStarted
go_name: ErrorWhenAlreadyStarted</pre></td>
</tr><tr>
<td>execution_timeout</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: executionTimeout
go_name: ExecutionTimeout</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>id_reuse_policy</td>
<td><a href="#temporal-xns-v1-idreusepolicy">temporal.xns.v1.IDReusePolicy</a></td>
<td><pre>
json_name: idReusePolicy
go_name: IdReusePolicy</pre></td>
</tr><tr>
<td>memo</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-struct">google.protobuf.Struct</a></td>
<td><pre>
json_name: memo
go_name: Memo</pre></td>
</tr><tr>
<td>retry_policy</td>
<td><a href="#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>run_timeout</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: runTimeout
go_name: RunTimeout</pre></td>
</tr><tr>
<td>search_attirbutes</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-struct">google.protobuf.Struct</a></td>
<td><pre>
json_name: searchAttirbutes
go_name: SearchAttirbutes</pre></td>
</tr><tr>
<td>start_delay</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: startDelay
go_name: StartDelay</pre></td>
</tr><tr>
<td>task_queue</td>
<td>string</td>
<td><pre>
json_name: taskQueue
go_name: TaskQueue</pre></td>
</tr><tr>
<td>task_timeout</td>
<td><a href="../../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: taskTimeout
go_name: TaskTimeout</pre></td>
</tr><tr>
<td>workflow_id_conflict_policy</td>
<td><a href="../../api/enums/v1/README.md#temporal-api-enums-v1-workflowidconflictpolicy">temporal.api.enums.v1.WorkflowIdConflictPolicy</a></td>
<td><pre>
json_name: workflowIdConflictPolicy
go_name: WorkflowIdConflictPolicy</pre></td>
</tr>
</table>

