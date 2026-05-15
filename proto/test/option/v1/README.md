

<a name="test-option-v1"></a>
# test.option.v1

## Table of Contents
- [test.option.v1.Test](#test-option-v1-test)
  - [Workflows](#test-option-v1-test-workflows)
    - [test.option.v1.Test.WorkflowWithInput](#test-option-v1-test-workflowwithinput-workflow)
  - [Updates](#test-option-v1-test-updates)
    - [test.option.v1.Test.UpdateWithInput](#test-option-v1-test-updatewithinput-update)
  - [Activities](#test-option-v1-test-activities)
    - [test.option.v1.Test.ActivityWithInput](#test-option-v1-test-activitywithinput-activity)
- Messages
  - [test.option.v1.ActivityWithInputRequest](#test-option-v1-activitywithinputrequest)
  - [test.option.v1.ActivityWithInputResponse](#test-option-v1-activitywithinputresponse)
  - [test.option.v1.UpdateWithInputRequest](#test-option-v1-updatewithinputrequest)
  - [test.option.v1.WorkflowWithInputRequest](#test-option-v1-workflowwithinputrequest)

<a name="test-option-v1-services"></a>
## Services

<a name="test-option-v1-test"></a>
## test.option.v1.Test

<a name="test-option-v1-test-workflows"></a>
### Workflows

---
<a name="test-option-v1-test-workflowwithinput-workflow"></a>
### test.option.v1.Test.WorkflowWithInput

**Input:** [test.option.v1.WorkflowWithInputRequest](#test-option-v1-workflowwithinputrequest)

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

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>execution_timeout</td><td>10 minutes</td></tr>
<tr><td>id</td><td><pre><code>workflow-with-input:${! name.or(throw("name is required")) }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY</code></pre></td></tr>
<tr><td>parent_close_policy</td><td><pre><code>PARENT_CLOSE_POLICY_REQUEST_CANCEL</code></pre></td></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>run_timeout</td><td>5 minutes</td></tr>
<tr><td>search_attributes</td><td><pre><code>name = name</code></pre></td></tr>
<tr><td>task_queue</td><td><pre><code>option-v2</code></pre></td></tr>
<tr><td>task_timeout</td><td>10 seconds</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#test-option-v1-test-updatewithinput-update">test.option.v1.Test.UpdateWithInput</a></td></tr>
</table>    

<a name="test-option-v1-test-updates"></a>
### Updates

---
<a name="test-option-v1-test-updatewithinput-update"></a>
### test.option.v1.Test.UpdateWithInput



**Input:** [test.option.v1.UpdateWithInputRequest](#test-option-v1-updatewithinputrequest)

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

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>wait_policy</td><td><pre>WAIT_POLICY_ADMITTED</pre></td></tr>
</table>

<a name="test-option-v1-test-activities"></a>
### Activities

---
<a name="test-option-v1-test-activitywithinput-activity"></a>
### test.option.v1.Test.ActivityWithInput



**Input:** [test.option.v1.ActivityWithInputRequest](#test-option-v1-activitywithinputrequest)

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

**Output:** [test.option.v1.ActivityWithInputResponse](#test-option-v1-activitywithinputresponse)

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
<tr><td>heartbeat_timeout</td><td>30 seconds</td></tr>
<tr><td>retry_policy.max_interval</td><td>5 seconds</td></tr>
<tr><td>schedule_to_close_timeout</td><td>30 seconds</td></tr>
<tr><td>schedule_to_start_timeout</td><td>10 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>1 minute</td></tr>
<tr><td>wait_for_cancellation</td><td>true</td></tr>
<tr><td>task_queue</td><td>option-v2</td></tr>
</table>   

<a name="test-option-v1-messages"></a>
## Messages

<a name="test-option-v1-activitywithinputrequest"></a>
### test.option.v1.ActivityWithInputRequest

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



<a name="test-option-v1-activitywithinputresponse"></a>
### test.option.v1.ActivityWithInputResponse

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



<a name="test-option-v1-updatewithinputrequest"></a>
### test.option.v1.UpdateWithInputRequest

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



<a name="test-option-v1-workflowwithinputrequest"></a>
### test.option.v1.WorkflowWithInputRequest

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

