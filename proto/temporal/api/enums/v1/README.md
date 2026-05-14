

<a name="temporal-api-enums-v1"></a>
# temporal.api.enums.v1

<a name="temporal-api-enums-v1-messages"></a>
## Messages

<a name="temporal-api-enums-v1-workflowidconflictpolicy"></a>
### temporal.api.enums.v1.WorkflowIdConflictPolicy

<pre>
Defines what to do when trying to start a workflow with the same workflow id as a *running* workflow.
Note that it is *never* valid to have two actively running instances of the same workflow id.

See `WorkflowIdReusePolicy` for handling workflow id duplication with a *closed* workflow.
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_FAIL</td>
<td><pre>
Don't start a new workflow; instead return `WorkflowExecutionAlreadyStartedFailure`.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING</td>
<td><pre>
Don't start a new workflow; instead return a workflow handle for the running workflow.
</pre></td>
</tr><tr>
<td>WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING</td>
<td><pre>
Terminate the running workflow before starting a new one.
</pre></td>
</tr>
</table>