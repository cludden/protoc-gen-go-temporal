

<a name="test-xnsheartbeat-v1"></a>
# test.xnsheartbeat.v1

## Table of Contents
- [test.xnsheartbeat.v1.XnsHeartbeatCallerService](#test-xnsheartbeat-v1-xnsheartbeatcallerservice)
  - [Workflows](#test-xnsheartbeat-v1-xnsheartbeatcallerservice-workflows)
    - [test.xnsheartbeat.v1.XnsHeartbeatCallerService.CallTestWorkflow](#test-xnsheartbeat-v1-xnsheartbeatcallerservice-calltestworkflow-workflow)
- [test.xnsheartbeat.v1.XnsHeartbeatService](#test-xnsheartbeat-v1-xnsheartbeatservice)
  - [Workflows](#test-xnsheartbeat-v1-xnsheartbeatservice-workflows)
    - [test.xnsheartbeat.v1.XnsHeartbeatService.TestWorkflow](#test-xnsheartbeat-v1-xnsheartbeatservice-testworkflow-workflow)
  - [Signals](#test-xnsheartbeat-v1-xnsheartbeatservice-signals)
    - [test.xnsheartbeat.v1.XnsHeartbeatService.TestSignal](#test-xnsheartbeat-v1-xnsheartbeatservice-testsignal-signal)
  - [Updates](#test-xnsheartbeat-v1-xnsheartbeatservice-updates)
    - [test.xnsheartbeat.v1.XnsHeartbeatService.TestUpdate](#test-xnsheartbeat-v1-xnsheartbeatservice-testupdate-update)
- Messages
  - [test.xnsheartbeat.v1.TestSignalInput](#test-xnsheartbeat-v1-testsignalinput)
  - [test.xnsheartbeat.v1.TestUpdateInput](#test-xnsheartbeat-v1-testupdateinput)
  - [test.xnsheartbeat.v1.TestUpdateOutput](#test-xnsheartbeat-v1-testupdateoutput)
  - [test.xnsheartbeat.v1.TestWorkflowInput](#test-xnsheartbeat-v1-testworkflowinput)
  - [test.xnsheartbeat.v1.TestWorkflowOutput](#test-xnsheartbeat-v1-testworkflowoutput)

<a name="test-xnsheartbeat-v1-services"></a>
## Services

<a name="test-xnsheartbeat-v1-xnsheartbeatcallerservice"></a>
## test.xnsheartbeat.v1.XnsHeartbeatCallerService

<a name="test-xnsheartbeat-v1-xnsheartbeatcallerservice-workflows"></a>
### Workflows

---
<a name="test-xnsheartbeat-v1-xnsheartbeatcallerservice-calltestworkflow-workflow"></a>
### test.xnsheartbeat.v1.XnsHeartbeatCallerService.CallTestWorkflow

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="test-xnsheartbeat-v1-xnsheartbeatservice"></a>
## test.xnsheartbeat.v1.XnsHeartbeatService

<a name="test-xnsheartbeat-v1-xnsheartbeatservice-workflows"></a>
### Workflows

---
<a name="test-xnsheartbeat-v1-xnsheartbeatservice-testworkflow-workflow"></a>
### test.xnsheartbeat.v1.XnsHeartbeatService.TestWorkflow

**Input:** [test.xnsheartbeat.v1.TestWorkflowInput](#test-xnsheartbeat-v1-testworkflowinput)



**Output:** [test.xnsheartbeat.v1.TestWorkflowOutput](#test-xnsheartbeat-v1-testworkflowoutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
<tr><td>xns.heartbeat_interval</td><td>10 seconds</td></tr>
<tr><td>xns.heartbeat_timeout</td><td>30 seconds</td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-xnsheartbeat-v1-xnsheartbeatservice-testsignal-signal">test.xnsheartbeat.v1.XnsHeartbeatService.TestSignal</a></td><td>false</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#test-xnsheartbeat-v1-xnsheartbeatservice-testupdate-update">test.xnsheartbeat.v1.XnsHeartbeatService.TestUpdate</a></td></tr>
</table>   

<a name="test-xnsheartbeat-v1-xnsheartbeatservice-signals"></a>
### Signals

---
<a name="test-xnsheartbeat-v1-xnsheartbeatservice-testsignal-signal"></a>
### test.xnsheartbeat.v1.XnsHeartbeatService.TestSignal



**Input:** [test.xnsheartbeat.v1.TestSignalInput](#test-xnsheartbeat-v1-testsignalinput)

  

<a name="test-xnsheartbeat-v1-xnsheartbeatservice-updates"></a>
### Updates

---
<a name="test-xnsheartbeat-v1-xnsheartbeatservice-testupdate-update"></a>
### test.xnsheartbeat.v1.XnsHeartbeatService.TestUpdate



**Input:** [test.xnsheartbeat.v1.TestUpdateInput](#test-xnsheartbeat-v1-testupdateinput)



**Output:** [test.xnsheartbeat.v1.TestUpdateOutput](#test-xnsheartbeat-v1-testupdateoutput)

 

<a name="test-xnsheartbeat-v1-messages"></a>
## Messages

<a name="test-xnsheartbeat-v1-testsignalinput"></a>
### test.xnsheartbeat.v1.TestSignalInput



<a name="test-xnsheartbeat-v1-testupdateinput"></a>
### test.xnsheartbeat.v1.TestUpdateInput



<a name="test-xnsheartbeat-v1-testupdateoutput"></a>
### test.xnsheartbeat.v1.TestUpdateOutput



<a name="test-xnsheartbeat-v1-testworkflowinput"></a>
### test.xnsheartbeat.v1.TestWorkflowInput



<a name="test-xnsheartbeat-v1-testworkflowoutput"></a>
### test.xnsheartbeat.v1.TestWorkflowOutput

