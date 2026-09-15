

<a name="temporal-api-deployment-v1"></a>
# temporal.api.deployment.v1

## Table of Contents
- Messages
  - [temporal.api.deployment.v1.Deployment](#temporal-api-deployment-v1-deployment)
  - [temporal.api.deployment.v1.WorkerDeploymentVersion](#temporal-api-deployment-v1-workerdeploymentversion)

<a name="temporal-api-deployment-v1-messages"></a>
## Messages

<a name="temporal-api-deployment-v1-deployment"></a>
### temporal.api.deployment.v1.Deployment

<pre>
`Deployment` identifies a deployment of Temporal workers. The combination of deployment series
name + build ID serves as the identifier. User can use `WorkerDeploymentOptions` in their worker
programs to specify these values.
Deprecated.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>build_id</td>
<td>string</td>
<td><pre>
Build ID changes with each version of the worker when the worker program code and/or config
changes.<br>

json_name: buildId
go_name: BuildId</pre></td>
</tr><tr>
<td>series_name</td>
<td>string</td>
<td><pre>
Different versions of the same worker service/application are related together by having a
shared series name.
Out of all deployments of a series, one can be designated as the current deployment, which
receives new workflow executions and new tasks of workflows with
`VERSIONING_BEHAVIOR_AUTO_UPGRADE` versioning behavior.<br>

json_name: seriesName
go_name: SeriesName</pre></td>
</tr>
</table>



<a name="temporal-api-deployment-v1-workerdeploymentversion"></a>
### temporal.api.deployment.v1.WorkerDeploymentVersion

<pre>
A Worker Deployment Version (Version, for short) represents a
version of workers within a Worker Deployment. (see documentation of WorkerDeploymentVersionInfo)
Version records are created in Temporal server automatically when their
first poller arrives to the server.
Experimental. Worker Deployment Versions are experimental and might significantly change in the future.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>build_id</td>
<td>string</td>
<td><pre>
A unique identifier for this Version within the Deployment it is a part of.
Not necessarily unique within the namespace.
The combination of `deployment_name` and `build_id` uniquely identifies this
Version within the namespace, because Deployment names are unique within a namespace.<br>

json_name: buildId
go_name: BuildId</pre></td>
</tr><tr>
<td>deployment_name</td>
<td>string</td>
<td><pre>
Identifies the Worker Deployment this Version is part of.<br>

json_name: deploymentName
go_name: DeploymentName</pre></td>
</tr>
</table>

