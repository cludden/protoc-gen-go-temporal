

<a name="temporal-api-workflow-v1"></a>
# temporal.api.workflow.v1

## Table of Contents
- Messages
  - [temporal.api.workflow.v1.VersioningOverride](#temporal-api-workflow-v1-versioningoverride)
  - [temporal.api.workflow.v1.VersioningOverride.PinnedOverride](#temporal-api-workflow-v1-versioningoverride-pinnedoverride)
  - [temporal.api.workflow.v1.VersioningOverride.PinnedOverrideBehavior](#temporal-api-workflow-v1-versioningoverride-pinnedoverridebehavior)

<a name="temporal-api-workflow-v1-messages"></a>
## Messages

<a name="temporal-api-workflow-v1-versioningoverride"></a>
### temporal.api.workflow.v1.VersioningOverride

<pre>
Used to override the versioning behavior (and pinned deployment version, if applicable) of a
specific workflow execution. If set, this override takes precedence over worker-sent values.
See `WorkflowExecutionInfo.VersioningInfo` for more information.

To remove the override, call `UpdateWorkflowExecutionOptions` with a null
`VersioningOverride`, and use the `update_mask` to indicate that it should be mutated.

Pinned behavior overrides are automatically inherited by child workflows, workflow retries, continue-as-new
workflows, and cron workflows.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>auto_upgrade</td>
<td>bool</td>
<td><pre>
Override the workflow to have AutoUpgrade behavior.<br>

json_name: autoUpgrade
go_name: AutoUpgrade</pre></td>
</tr><tr>
<td>behavior</td>
<td><a href="../../enums/v1/README.md#temporal-api-enums-v1-versioningbehavior">temporal.api.enums.v1.VersioningBehavior</a></td>
<td><pre>
Required.
Deprecated. Use `override`.<br>

json_name: behavior
go_name: Behavior</pre></td>
</tr><tr>
<td>deployment</td>
<td><a href="../../deployment/v1/README.md#temporal-api-deployment-v1-deployment">temporal.api.deployment.v1.Deployment</a></td>
<td><pre>
Required if behavior is `PINNED`. Must be null if behavior is `AUTO_UPGRADE`.
Identifies the worker deployment to pin the workflow to.
Deprecated. Use `override.pinned.version`.<br>

json_name: deployment
go_name: Deployment</pre></td>
</tr><tr>
<td>pinned</td>
<td><a href="#temporal-api-workflow-v1-versioningoverride-pinnedoverride">temporal.api.workflow.v1.VersioningOverride.PinnedOverride</a></td>
<td><pre>
Override the workflow to have Pinned behavior.<br>

json_name: pinned
go_name: Pinned</pre></td>
</tr><tr>
<td>pinned_version</td>
<td>string</td>
<td><pre>
Required if behavior is `PINNED`. Must be absent if behavior is not `PINNED`.
Identifies the worker deployment version to pin the workflow to, in the format
"<deployment_name>.<build_id>".
Deprecated. Use `override.pinned.version`.<br>

json_name: pinnedVersion
go_name: PinnedVersion</pre></td>
</tr>
</table>



<a name="temporal-api-workflow-v1-versioningoverride-pinnedoverride"></a>
### temporal.api.workflow.v1.VersioningOverride.PinnedOverride

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>behavior</td>
<td><a href="#temporal-api-workflow-v1-versioningoverride-pinnedoverridebehavior">temporal.api.workflow.v1.VersioningOverride.PinnedOverrideBehavior</a></td>
<td><pre>
Defaults to PINNED_OVERRIDE_BEHAVIOR_UNSPECIFIED.
See `PinnedOverrideBehavior` for details.<br>

json_name: behavior
go_name: Behavior</pre></td>
</tr><tr>
<td>version</td>
<td><a href="../../deployment/v1/README.md#temporal-api-deployment-v1-workerdeploymentversion">temporal.api.deployment.v1.WorkerDeploymentVersion</a></td>
<td><pre>
Specifies the Worker Deployment Version to pin this workflow to.
Required if the target workflow is not already pinned to a version.

If omitted and the target workflow is already pinned, the effective
pinned version will be the existing pinned version.

If omitted and the target workflow is not pinned, the override request
will be rejected with a PreconditionFailed error.<br>

json_name: version
go_name: Version</pre></td>
</tr>
</table>



<a name="temporal-api-workflow-v1-versioningoverride-pinnedoverridebehavior"></a>
### temporal.api.workflow.v1.VersioningOverride.PinnedOverrideBehavior

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>PINNED_OVERRIDE_BEHAVIOR_UNSPECIFIED</td>
<td><pre>
Unspecified.
</pre></td>
</tr><tr>
<td>PINNED_OVERRIDE_BEHAVIOR_PINNED</td>
<td><pre>
Override workflow behavior to be Pinned.
</pre></td>
</tr>
</table>