# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [temporal/v1/temporal.proto](#temporal_v1_temporal-proto)
    - [ActivityOptions](#temporal-v1-ActivityOptions)
    - [CLIOptions](#temporal-v1-CLIOptions)
    - [CommandOptions](#temporal-v1-CommandOptions)
    - [QueryOptions](#temporal-v1-QueryOptions)
    - [RetryPolicy](#temporal-v1-RetryPolicy)
    - [ServiceOptions](#temporal-v1-ServiceOptions)
    - [SignalOptions](#temporal-v1-SignalOptions)
    - [UpdateOptions](#temporal-v1-UpdateOptions)
    - [WorkflowOptions](#temporal-v1-WorkflowOptions)
    - [WorkflowOptions.Query](#temporal-v1-WorkflowOptions-Query)
    - [WorkflowOptions.Signal](#temporal-v1-WorkflowOptions-Signal)
    - [WorkflowOptions.Update](#temporal-v1-WorkflowOptions-Update)
  
    - [CLIFeature](#temporal-v1-CLIFeature)
    - [IDReusePolicy](#temporal-v1-IDReusePolicy)
    - [ParentClosePolicy](#temporal-v1-ParentClosePolicy)
    - [WaitPolicy](#temporal-v1-WaitPolicy)
  
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
    - [File-level Extensions](#temporal_v1_temporal-proto-extensions)
  
- [Scalar Value Types](#scalar-value-types)



<a name="temporal_v1_temporal-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## temporal/v1/temporal.proto



<a name="temporal-v1-ActivityOptions"></a>

### ActivityOptions
ActivityOptions identifies an rpc method as a Temporal activity definition, and describes
available activity configuration options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Fully-qualified activity name |
| task_queue | [string](#string) |  | Override default task queue for activity |
| schedule_to_close_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | Total time that a workflow is willing to wait for Activity to complete |
| schedule_to_start_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | Time that the Activity Task can stay in the Task Queue before it is picked up by a Worker |
| start_to_close_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | Maximum time of a single Activity execution attempt |
| heartbeat_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | Heartbeat interval. Activity must call Activity.RecordHeartbeat(ctx, &#34;my-heartbeat&#34;) |
| retry_policy | [RetryPolicy](#temporal-v1-RetryPolicy) |  | Specifies how to retry an Activity if an error occurs |






<a name="temporal-v1-CLIOptions"></a>

### CLIOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ignore | [bool](#bool) |  |  |






<a name="temporal-v1-CommandOptions"></a>

### CommandOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ignore | [bool](#bool) |  |  |






<a name="temporal-v1-QueryOptions"></a>

### QueryOptions
QueryOptions identifies an rpc method as a Temporal query definition, and describes
available query configuration options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Fully-qualified query name |






<a name="temporal-v1-RetryPolicy"></a>

### RetryPolicy
RetryPolicy describes configuration for activity or child workflow retries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| initial_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| backoff_coefficient | [double](#double) |  |  |
| max_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| max_attempts | [int32](#int32) |  |  |
| non_retryable_error_types | [string](#string) | repeated |  |






<a name="temporal-v1-ServiceOptions"></a>

### ServiceOptions
ServiceOptions provides options that can be used to define common configuration
shared by all methods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | Default namespace for child workflows, activities |
| task_queue | [string](#string) |  | Default task queue for all workflows, activities |






<a name="temporal-v1-SignalOptions"></a>

### SignalOptions
SignalOptions identifies an rpc method as a Temporal singal definition, and describes
available signal configuration options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Fully-qualified signal name |






<a name="temporal-v1-UpdateOptions"></a>

### UpdateOptions
UpdateOptions identifies an rpc method as a Temporal update definition, and describes
available update configuration options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID expression |
| name | [string](#string) |  | Fully-qualified update name |
| validate | [bool](#bool) |  | Include validation hook |
| wait_policy | [WaitPolicy](#temporal-v1-WaitPolicy) |  | Default wait policy if not specified |






<a name="temporal-v1-WorkflowOptions"></a>

### WorkflowOptions
WorkflowOptions identifies an rpc method as a Temporal workflow definition, and describes
available workflow configuration options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Fully-qualified workflow name |
| query | [WorkflowOptions.Query](#temporal-v1-WorkflowOptions-Query) | repeated | Queries supported by this workflow |
| signal | [WorkflowOptions.Signal](#temporal-v1-WorkflowOptions-Signal) | repeated | Signals supported by this workflow |
| update | [WorkflowOptions.Update](#temporal-v1-WorkflowOptions-Update) | repeated | Updates supported by this workflow |
| execution_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | The timeout for duration of workflow execution. It includes retries and continue as new. Use WorkflowRunTimeout to limit execution time of a single workflow run. |
| id | [string](#string) |  | Id expression |
| id_reuse_policy | [IDReusePolicy](#temporal-v1-IDReusePolicy) |  | Whether server allow reuse of workflow ID |
| namespace | [string](#string) |  | Specifies default namespace for child workflows |
| parent_close_policy | [ParentClosePolicy](#temporal-v1-ParentClosePolicy) |  | Specifies a default parent close policy for child workflows |
| retry_policy | [RetryPolicy](#temporal-v1-RetryPolicy) |  | Specifies how to retry an Workflow if an error occurs |
| run_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | The timeout for duration of a single workflow run. |
| search_attributes | [string](#string) |  | Bloblang mapping defining default workflow search attributes |
| task_queue | [string](#string) |  | Override service task queeu |
| task_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  | The timeout for processing workflow task from the time the worker pulled this task. If a workflow task is lost, it is retried after this timeout. The resolution is seconds. |
| wait_for_cancellation | [bool](#bool) |  | WaitForCancellation specifies whether to wait for canceled child workflow to be ended (child workflow can be ended as: completed/failed/timedout/terminated/canceled) |






<a name="temporal-v1-WorkflowOptions-Query"></a>

### WorkflowOptions.Query
Query identifies a query supported by the worklow


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref | [string](#string) |  | Query name |






<a name="temporal-v1-WorkflowOptions-Signal"></a>

### WorkflowOptions.Signal
Signal identifies a signal supported by the workflow


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref | [string](#string) |  | Signal name |
| start | [bool](#bool) |  | Include convenience method for signal with start |






<a name="temporal-v1-WorkflowOptions-Update"></a>

### WorkflowOptions.Update
Update identifies an update supported by the workflow


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref | [string](#string) |  | Update name |





 


<a name="temporal-v1-CLIFeature"></a>

### CLIFeature
CLIFeature enumerates cli feature statuses

| Name | Number | Description |
| ---- | ------ | ----------- |
| CLI_FEATURE_DISALBED | 0 |  |
| CLI_FEATURE_ENABLED | 1 |  |



<a name="temporal-v1-IDReusePolicy"></a>

### IDReusePolicy
IDReusePolicy defines how new runs of a workflow with a particular ID may or 
may not be allowed. Note that it is *never* valid to have two actively 
running instances of the same workflow id.

| Name | Number | Description |
| ---- | ------ | ----------- |
| WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED | 0 |  |
| WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE | 1 | Allow starting a workflow execution using the same workflow id. |
| WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY | 2 | Allow starting a workflow execution using the same workflow id, only when the last execution&#39;s final state is one of [terminated, cancelled, timed out, failed]. |
| WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE | 3 | Do not permit re-use of the workflow id for this workflow. Future start workflow requests could potentially change the policy, allowing re-use of the workflow id. |
| WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING | 4 | If a workflow is running using the same workflow ID, terminate it and start a new one. If no running workflow, then the behavior is the same as ALLOW_DUPLICATE |



<a name="temporal-v1-ParentClosePolicy"></a>

### ParentClosePolicy
Defines how child workflows will react to their parent completing

| Name | Number | Description |
| ---- | ------ | ----------- |
| PARENT_CLOSE_POLICY_UNSPECIFIED | 0 |  |
| PARENT_CLOSE_POLICY_TERMINATE | 1 | The child workflow will also terminate |
| PARENT_CLOSE_POLICY_ABANDON | 2 | The child workflow will do nothing |
| PARENT_CLOSE_POLICY_REQUEST_CANCEL | 3 | Cancellation will be requested of the child workflow |



<a name="temporal-v1-WaitPolicy"></a>

### WaitPolicy
WaitPolicy used to indicate to the server how long the client wishes to wait for a return 
value from an UpdateWorkflow RPC

| Name | Number | Description |
| ---- | ------ | ----------- |
| WAIT_POLICY_UNSPECIFIED | 0 |  |
| WAIT_POLICY_ADMITTED | 1 |  |
| WAIT_POLICY_ACCEPTED | 2 |  |
| WAIT_POLICY_COMPLETED | 3 |  |


 


<a name="temporal_v1_temporal-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| activity | ActivityOptions | .google.protobuf.MethodOptions | 7234 |  |
| command | CommandOptions | .google.protobuf.MethodOptions | 7238 |  |
| query | QueryOptions | .google.protobuf.MethodOptions | 7235 |  |
| signal | SignalOptions | .google.protobuf.MethodOptions | 7236 |  |
| update | UpdateOptions | .google.protobuf.MethodOptions | 7237 |  |
| workflow | WorkflowOptions | .google.protobuf.MethodOptions | 7233 |  |
| cli | CLIOptions | .google.protobuf.ServiceOptions | 7234 |  |
| service | ServiceOptions | .google.protobuf.ServiceOptions | 7233 |  |

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

