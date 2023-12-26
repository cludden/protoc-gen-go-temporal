# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [temporal/xns/v1/xns.proto](#temporal_xns_v1_xns-proto)
    - [QueryRequest](#temporal-xns-v1-QueryRequest)
    - [RetryPolicy](#temporal-xns-v1-RetryPolicy)
    - [SignalRequest](#temporal-xns-v1-SignalRequest)
    - [StartWorkflowOptions](#temporal-xns-v1-StartWorkflowOptions)
    - [UpdateRequest](#temporal-xns-v1-UpdateRequest)
    - [UpdateWorkflowWithOptionsRequest](#temporal-xns-v1-UpdateWorkflowWithOptionsRequest)
    - [WorkflowRequest](#temporal-xns-v1-WorkflowRequest)
  
    - [IDReusePolicy](#temporal-xns-v1-IDReusePolicy)
    - [ParentClosePolicy](#temporal-xns-v1-ParentClosePolicy)
    - [WaitPolicy](#temporal-xns-v1-WaitPolicy)
  
- [Scalar Value Types](#scalar-value-types)



<a name="temporal_xns_v1_xns-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## temporal/xns/v1/xns.proto



<a name="temporal-xns-v1-QueryRequest"></a>

### QueryRequest
QueryRequest can be used to configure xns query activities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| run_id | [string](#string) |  |  |
| workflow_id | [string](#string) |  |  |
| request | [google.protobuf.Any](#google-protobuf-Any) |  |  |
| heartbeat_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |






<a name="temporal-xns-v1-RetryPolicy"></a>

### RetryPolicy
RetryPolicy describes configuration for activity or child workflow retries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| initial_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| backoff_coefficient | [double](#double) |  |  |
| max_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| max_attempts | [int32](#int32) |  |  |
| non_retryable_error_types | [string](#string) | repeated |  |






<a name="temporal-xns-v1-SignalRequest"></a>

### SignalRequest
SignalRequest can be used to configure xns signal activities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| run_id | [string](#string) |  |  |
| workflow_id | [string](#string) |  |  |
| request | [google.protobuf.Any](#google-protobuf-Any) |  |  |
| heartbeat_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |






<a name="temporal-xns-v1-StartWorkflowOptions"></a>

### StartWorkflowOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| task_queue | [string](#string) |  |  |
| execution_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| run_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| task_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| id_reuse_policy | [IDReusePolicy](#temporal-xns-v1-IDReusePolicy) |  |  |
| error_when_already_started | [bool](#bool) |  |  |
| retry_policy | [RetryPolicy](#temporal-xns-v1-RetryPolicy) |  |  |
| memo | [google.protobuf.Struct](#google-protobuf-Struct) |  |  |
| search_attirbutes | [google.protobuf.Struct](#google-protobuf-Struct) |  |  |
| enable_eager_start | [bool](#bool) |  |  |
| start_delay | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |






<a name="temporal-xns-v1-UpdateRequest"></a>

### UpdateRequest
UpdateRequest can be used to configure xns update activities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| heartbeat_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| update_workflow_options | [UpdateWorkflowWithOptionsRequest](#temporal-xns-v1-UpdateWorkflowWithOptionsRequest) |  |  |
| request | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="temporal-xns-v1-UpdateWorkflowWithOptionsRequest"></a>

### UpdateWorkflowWithOptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| update_id | [string](#string) |  |  |
| workflow_id | [string](#string) |  |  |
| run_id | [string](#string) |  |  |
| first_execution_run_id | [string](#string) |  |  |
| wait_policy | [WaitPolicy](#temporal-xns-v1-WaitPolicy) |  |  |






<a name="temporal-xns-v1-WorkflowRequest"></a>

### WorkflowRequest
WorkflowRequest can be used to configure xns workflow activities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| heartbeat_interval | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| start_workflow_options | [StartWorkflowOptions](#temporal-xns-v1-StartWorkflowOptions) |  |  |
| request | [google.protobuf.Any](#google-protobuf-Any) |  |  |
| detached | [bool](#bool) |  |  |
| signal | [google.protobuf.Any](#google-protobuf-Any) |  |  |





 


<a name="temporal-xns-v1-IDReusePolicy"></a>

### IDReusePolicy


| Name | Number | Description |
| ---- | ------ | ----------- |
| WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED | 0 |  |
| WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE | 1 | Allow starting a workflow execution using the same workflow id. |
| WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY | 2 | Allow starting a workflow execution using the same workflow id, only when the last execution&#39;s final state is one of [terminated, cancelled, timed out, failed]. |
| WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE | 3 | Do not permit re-use of the workflow id for this workflow. Future start workflow requests could potentially change the policy, allowing re-use of the workflow id. |
| WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING | 4 | If a workflow is running using the same workflow ID, terminate it and start a new one. If no running workflow, then the behavior is the same as ALLOW_DUPLICATE |



<a name="temporal-xns-v1-ParentClosePolicy"></a>

### ParentClosePolicy
Defines how child workflows will react to their parent completing

| Name | Number | Description |
| ---- | ------ | ----------- |
| PARENT_CLOSE_POLICY_UNSPECIFIED | 0 |  |
| PARENT_CLOSE_POLICY_TERMINATE | 1 | The child workflow will also terminate |
| PARENT_CLOSE_POLICY_ABANDON | 2 | The child workflow will do nothing |
| PARENT_CLOSE_POLICY_REQUEST_CANCEL | 3 | Cancellation will be requested of the child workflow |



<a name="temporal-xns-v1-WaitPolicy"></a>

### WaitPolicy
WaitPolicy used to indicate to the server how long the client wishes to wait for a return 
value from an UpdateWorkflow RPC

| Name | Number | Description |
| ---- | ------ | ----------- |
| WAIT_POLICY_UNSPECIFIED | 0 |  |
| WAIT_POLICY_ADMITTED | 1 |  |
| WAIT_POLICY_ACCEPTED | 2 |  |
| WAIT_POLICY_COMPLETED | 3 |  |


 

 

 



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

