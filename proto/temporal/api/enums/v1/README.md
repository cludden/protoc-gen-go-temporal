

<a name="temporal-api-enums-v1"></a>
# temporal.api.enums.v1

## Table of Contents
- Messages
  - [temporal.api.enums.v1.VersioningBehavior](#temporal-api-enums-v1-versioningbehavior)
  - [temporal.api.enums.v1.WorkflowIdConflictPolicy](#temporal-api-enums-v1-workflowidconflictpolicy)

<a name="temporal-api-enums-v1-messages"></a>
## Messages

<a name="temporal-api-enums-v1-versioningbehavior"></a>
### temporal.api.enums.v1.VersioningBehavior

<pre>
Versioning Behavior specifies if and how a workflow execution moves between Worker Deployment
Versions. The Versioning Behavior of a workflow execution is typically specified by the worker
who completes the first task of the execution, but is also overridable manually for new and
existing workflows (see VersioningOverride).
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>VERSIONING_BEHAVIOR_UNSPECIFIED</td>
<td><pre>
Workflow execution does not have a Versioning Behavior and is called Unversioned. This is the
legacy behavior. An Unversioned workflow's task can go to any Unversioned worker (see
`WorkerVersioningMode`.)
User needs to use Patching to keep the new code compatible with prior versions when dealing
with Unversioned workflows.
</pre></td>
</tr><tr>
<td>VERSIONING_BEHAVIOR_PINNED</td>
<td><pre>
Workflow will start on its Target Version and then will be pinned to that same Deployment
Version until completion (the Version that this Workflow is pinned to is specified in
`versioning_info.version` and is the Pinned Version of the Workflow).

The workflow's Target Version is the Current Version of its Task Queue, or, if the
Task Queue has a Ramping Version with non-zero Ramp Percentage `P`, the workflow's Target
Version has a P% chance of being the Ramping Version. Whether a workflow falls into the
Ramping group depends on its Workflow ID and and the Ramp Percentage.

This behavior eliminates most of compatibility concerns users face when changing their code.
Patching is not needed when pinned workflows code change.
Can be overridden explicitly via `UpdateWorkflowExecutionOptions` API to move the
execution to another Deployment Version.
Activities of `PINNED` workflows are sent to the same Deployment Version. Exception to this
would be when the activity Task Queue workers are not present in the workflow's Deployment
Version, in which case the activity will be sent to the Current Deployment Version of its own
task queue.
</pre></td>
</tr><tr>
<td>VERSIONING_BEHAVIOR_AUTO_UPGRADE</td>
<td><pre>
Workflow will automatically move to its Target Version when the next workflow task is dispatched.

The workflow's Target Version is the Current Version of its Task Queue, or, if the
Task Queue has a Ramping Version with non-zero Ramp Percentage `P`, the workflow's Target
Version has a P% chance of being the Ramping Version. Whether a workflow falls into the
Ramping group depends on its Workflow ID and and the Ramp Percentage.

AutoUpgrade behavior is suitable for long-running workflows as it allows them to move to the
latest Deployment Version, but the user still needs to use Patching to keep the new code
compatible with prior versions for changed workflow types.
Activities of `AUTO_UPGRADE` workflows are sent to the Deployment Version of the workflow
execution (as specified in versioning_info.version based on the last completed
workflow task). Exception to this would be when the activity Task Queue workers are not
present in the workflow's Deployment Version, in which case, the activity will be sent to a
different Deployment Version according to the Current or Ramping Deployment Version of its own
Task Queue.
Workflows stuck on a backlogged activity will still auto-upgrade if their Target Version
changes, without having to wait for the backlogged activity to complete on the old Version.
</pre></td>
</tr>
</table>

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