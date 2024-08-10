# Table of Contents

- [example.helloworld.v1](#example-helloworld-v1)
  - Services
    - [example.helloworld.v1.Example](#example-helloworld-v1-example)
      - [Workflows](#example-helloworld-v1-example-workflows)
        - [example.v1.Hello](#example-v1-hello-workflow)
      - [Signals](#example-helloworld-v1-example-signals)
        - [example.helloworld.v1.Example.Goodbye](#example-helloworld-v1-example-goodbye-signal)
  - Messages
    - [example.helloworld.v1.GoodbyeRequest](#example-helloworld-v1-goodbyerequest)
    - [example.helloworld.v1.HelloRequest](#example-helloworld-v1-hellorequest)
    - [example.helloworld.v1.HelloResponse](#example-helloworld-v1-helloresponse)
- [example.mutex.v1](#example-mutex-v1)
  - Services
    - [example.mutex.v1.Example](#example-mutex-v1-example)
      - [Workflows](#example-mutex-v1-example-workflows)
        - [example.mutex.v1.Example.Mutex](#example-mutex-v1-example-mutex-workflow)
        - [example.mutex.v1.Example.SampleWorkflowWithMutex](#example-mutex-v1-example-sampleworkflowwithmutex-workflow)
      - [Signals](#example-mutex-v1-example-signals)
        - [example.mutex.v1.Example.AcquireLock](#example-mutex-v1-example-acquirelock-signal)
        - [example.mutex.v1.Example.LockAcquired](#example-mutex-v1-example-lockacquired-signal)
        - [example.mutex.v1.Example.ReleaseLock](#example-mutex-v1-example-releaselock-signal)
      - [Activities](#example-mutex-v1-example-activities)
        - [example.mutex.v1.Example.Mutex](#example-mutex-v1-example-mutex-activity)
  - Messages
    - [example.mutex.v1.AcquireLockInput](#example-mutex-v1-acquirelockinput)
    - [example.mutex.v1.LockAcquiredInput](#example-mutex-v1-lockacquiredinput)
    - [example.mutex.v1.MutexInput](#example-mutex-v1-mutexinput)
    - [example.mutex.v1.ReleaseLockInput](#example-mutex-v1-releaselockinput)
    - [example.mutex.v1.SampleWorkflowWithMutexInput](#example-mutex-v1-sampleworkflowwithmutexinput)
- [example.schedule.v1](#example-schedule-v1)
  - Services
    - [example.schedule.v1.Example](#example-schedule-v1-example)
      - [Workflows](#example-schedule-v1-example-workflows)
        - [example.schedule.v1.Schedule](#example-schedule-v1-schedule-workflow)
  - Messages
    - [example.schedule.v1.ScheduleInput](#example-schedule-v1-scheduleinput)
    - [example.schedule.v1.ScheduleOutput](#example-schedule-v1-scheduleoutput)
- [example.searchattributes.v1](#example-searchattributes-v1)
  - Services
    - [example.searchattributes.v1.Example](#example-searchattributes-v1-example)
      - [Workflows](#example-searchattributes-v1-example-workflows)
        - [example.searchattributes.v1.Example.SearchAttributes](#example-searchattributes-v1-example-searchattributes-workflow)
  - Messages
    - [example.searchattributes.v1.SearchAttributesInput](#example-searchattributes-v1-searchattributesinput)
- [example.updatabletimer.v1](#example-updatabletimer-v1)
  - Services
    - [example.updatabletimer.v1.Example](#example-updatabletimer-v1-example)
      - [Workflows](#example-updatabletimer-v1-example-workflows)
        - [UpdatableTimer](#updatabletimer-workflow)
      - [Queries](#example-updatabletimer-v1-example-queries)
        - [example.updatabletimer.v1.Example.GetWakeUpTime](#example-updatabletimer-v1-example-getwakeuptime-query)
      - [Signals](#example-updatabletimer-v1-example-signals)
        - [example.updatabletimer.v1.Example.UpdateWakeUpTime](#example-updatabletimer-v1-example-updatewakeuptime-signal)
  - Messages
    - [example.updatabletimer.v1.GetWakeUpTimeOutput](#example-updatabletimer-v1-getwakeuptimeoutput)
    - [example.updatabletimer.v1.UpdatableTimerInput](#example-updatabletimer-v1-updatabletimerinput)
    - [example.updatabletimer.v1.UpdateWakeUpTimeInput](#example-updatabletimer-v1-updatewakeuptimeinput)
- [example.v1](#example-v1)
  - Services
    - [example.v1.Example](#example-v1-example)
      - [Workflows](#example-v1-example-workflows)
        - [example.v1.Example.CreateFoo](#example-v1-example-createfoo-workflow)
      - [Queries](#example-v1-example-queries)
        - [example.v1.Example.GetFooProgress](#example-v1-example-getfooprogress-query)
      - [Signals](#example-v1-example-signals)
        - [example.v1.Example.SetFooProgress](#example-v1-example-setfooprogress-signal)
      - [Updates](#example-v1-example-updates)
        - [example.v1.Example.UpdateFooProgress](#example-v1-example-updatefooprogress-update)
      - [Activities](#example-v1-example-activities)
        - [example.v1.Example.Notify](#example-v1-example-notify-activity)
  - Messages
    - [example.v1.CreateFooRequest](#example-v1-createfoorequest)
    - [example.v1.CreateFooResponse](#example-v1-createfooresponse)
    - [example.v1.Foo](#example-v1-foo)
    - [example.v1.Foo.Status](#example-v1-foo-status)
    - [example.v1.GetFooProgressResponse](#example-v1-getfooprogressresponse)
    - [example.v1.NotifyRequest](#example-v1-notifyrequest)
    - [example.v1.SetFooProgressRequest](#example-v1-setfooprogressrequest)
- [example.xns.v1](#example-xns-v1)
  - Services
    - [example.xns.v1.Xns](#example-xns-v1-xns)
      - [Workflows](#example-xns-v1-xns-workflows)
        - [example.xns.v1.Xns.ProvisionFoo](#example-xns-v1-xns-provisionfoo-workflow)
    - [example.xns.v1.Example](#example-xns-v1-example)
      - [Workflows](#example-xns-v1-example-workflows)
        - [example.xns.v1.Example.CreateFoo](#example-xns-v1-example-createfoo-workflow)
      - [Queries](#example-xns-v1-example-queries)
        - [example.xns.v1.Example.GetFooProgress](#example-xns-v1-example-getfooprogress-query)
      - [Signals](#example-xns-v1-example-signals)
        - [example.xns.v1.Example.SetFooProgress](#example-xns-v1-example-setfooprogress-signal)
      - [Updates](#example-xns-v1-example-updates)
        - [example.xns.v1.Example.UpdateFooProgress](#example-xns-v1-example-updatefooprogress-update)
      - [Activities](#example-xns-v1-example-activities)
        - [example.xns.v1.Example.Notify](#example-xns-v1-example-notify-activity)
  - Messages
    - [example.xns.v1.CreateFooRequest](#example-xns-v1-createfoorequest)
    - [example.xns.v1.CreateFooResponse](#example-xns-v1-createfooresponse)
    - [example.xns.v1.Foo](#example-xns-v1-foo)
    - [example.xns.v1.Foo.Status](#example-xns-v1-foo-status)
    - [example.xns.v1.GetFooProgressResponse](#example-xns-v1-getfooprogressresponse)
    - [example.xns.v1.NotifyRequest](#example-xns-v1-notifyrequest)
    - [example.xns.v1.ProvisionFooRequest](#example-xns-v1-provisionfoorequest)
    - [example.xns.v1.ProvisionFooResponse](#example-xns-v1-provisionfooresponse)
    - [example.xns.v1.SetFooProgressRequest](#example-xns-v1-setfooprogressrequest)
- [mycompany.simple](#mycompany-simple)
  - Services
    - [mycompany.simple.Simple](#mycompany-simple-simple)
      - [Workflows](#mycompany-simple-simple-workflows)
        - [mycompany.simple.SomeWorkflow1](#mycompany-simple-someworkflow1-workflow)
        - [mycompany.simple.SomeWorkflow2](#mycompany-simple-someworkflow2-workflow)
        - [mycompany.simple.Simple.SomeWorkflow3](#mycompany-simple-simple-someworkflow3-workflow)
        - [mycompany.simple.Simple.SomeWorkflow4](#mycompany-simple-simple-someworkflow4-workflow)
      - [Queries](#mycompany-simple-simple-queries)
        - [mycompany.simple.Simple.SomeQuery1](#mycompany-simple-simple-somequery1-query)
        - [mycompany.simple.Simple.SomeQuery2](#mycompany-simple-simple-somequery2-query)
      - [Signals](#mycompany-simple-simple-signals)
        - [mycompany.simple.Simple.SomeSignal1](#mycompany-simple-simple-somesignal1-signal)
        - [mycompany.simple.Simple.SomeSignal2](#mycompany-simple-simple-somesignal2-signal)
        - [mycompany.simple.Simple.SomeSignal3](#mycompany-simple-simple-somesignal3-signal)
      - [Updates](#mycompany-simple-simple-updates)
        - [mycompany.simple.Simple.SomeUpdate1](#mycompany-simple-simple-someupdate1-update)
      - [Activities](#mycompany-simple-simple-activities)
        - [mycompany.simple.SomeActivity1](#mycompany-simple-someactivity1-activity)
        - [mycompany.simple.Simple.SomeActivity2](#mycompany-simple-simple-someactivity2-activity)
        - [mycompany.simple.Simple.SomeActivity3](#mycompany-simple-simple-someactivity3-activity)
        - [mycompany.simple.Simple.SomeActivity4](#mycompany-simple-simple-someactivity4-activity)
        - [mycompany.simple.Simple.SomeSignal1](#mycompany-simple-simple-somesignal1-activity)
        - [mycompany.simple.Simple.SomeSignal2](#mycompany-simple-simple-somesignal2-activity)
        - [mycompany.simple.Simple.SomeSignal3](#mycompany-simple-simple-somesignal3-activity)
        - [mycompany.simple.Simple.SomeUpdate1](#mycompany-simple-simple-someupdate1-activity)
    - [mycompany.simple.Other](#mycompany-simple-other)
      - [Workflows](#mycompany-simple-other-workflows)
        - [mycompany.simple.Other.OtherWorkflow](#mycompany-simple-other-otherworkflow-workflow)
      - [Queries](#mycompany-simple-other-queries)
        - [mycompany.simple.Other.OtherQuery](#mycompany-simple-other-otherquery-query)
      - [Signals](#mycompany-simple-other-signals)
        - [mycompany.simple.Other.OtherSignal](#mycompany-simple-other-othersignal-signal)
      - [Updates](#mycompany-simple-other-updates)
        - [mycompany.simple.Other.OtherUpdate](#mycompany-simple-other-otherupdate-update)
      - [Activities](#mycompany-simple-other-activities)
        - [mycompany.simple.Other.OtherWorkflow](#mycompany-simple-other-otherworkflow-activity)
    - [mycompany.simple.Ignored](#mycompany-simple-ignored)
      - [Workflows](#mycompany-simple-ignored-workflows)
        - [mycompany.simple.Ignored.What](#mycompany-simple-ignored-what-workflow)
    - [mycompany.simple.OnlyActivities](#mycompany-simple-onlyactivities)
      - [Activities](#mycompany-simple-onlyactivities-activities)
        - [mycompany.simple.OnlyActivities.LonelyActivity1](#mycompany-simple-onlyactivities-lonelyactivity1-activity)
    - [mycompany.simple.Deprecated](#mycompany-simple-deprecated)
      - [Workflows](#mycompany-simple-deprecated-workflows)
        - [mycompany.simple.Deprecated.SomeDeprecatedWorkflow1](#mycompany-simple-deprecated-somedeprecatedworkflow1-workflow)
        - [mycompany.simple.Deprecated.SomeDeprecatedWorkflow2](#mycompany-simple-deprecated-somedeprecatedworkflow2-workflow)
      - [Queries](#mycompany-simple-deprecated-queries)
        - [mycompany.simple.Deprecated.SomeDeprecatedQuery1](#mycompany-simple-deprecated-somedeprecatedquery1-query)
        - [mycompany.simple.Deprecated.SomeDeprecatedQuery2](#mycompany-simple-deprecated-somedeprecatedquery2-query)
      - [Signals](#mycompany-simple-deprecated-signals)
        - [mycompany.simple.Deprecated.SomeDeprecatedSignal1](#mycompany-simple-deprecated-somedeprecatedsignal1-signal)
        - [mycompany.simple.Deprecated.SomeDeprecatedSignal2](#mycompany-simple-deprecated-somedeprecatedsignal2-signal)
      - [Updates](#mycompany-simple-deprecated-updates)
        - [mycompany.simple.Deprecated.SomeDeprecatedUpdate1](#mycompany-simple-deprecated-somedeprecatedupdate1-update)
        - [mycompany.simple.Deprecated.SomeDeprecatedUpdate2](#mycompany-simple-deprecated-somedeprecatedupdate2-update)
      - [Activities](#mycompany-simple-deprecated-activities)
        - [mycompany.simple.Deprecated.SomeDeprecatedActivity1](#mycompany-simple-deprecated-somedeprecatedactivity1-activity)
        - [mycompany.simple.Deprecated.SomeDeprecatedActivity2](#mycompany-simple-deprecated-somedeprecatedactivity2-activity)
  - Messages
    - [mycompany.simple.Foo](#mycompany-simple-foo)
    - [mycompany.simple.LonelyActivity1Request](#mycompany-simple-lonelyactivity1request)
    - [mycompany.simple.LonelyActivity1Response](#mycompany-simple-lonelyactivity1response)
    - [mycompany.simple.OtherEnum](#mycompany-simple-otherenum)
    - [mycompany.simple.OtherQueryResponse](#mycompany-simple-otherqueryresponse)
    - [mycompany.simple.OtherSignalRequest](#mycompany-simple-othersignalrequest)
    - [mycompany.simple.OtherUpdateRequest](#mycompany-simple-otherupdaterequest)
    - [mycompany.simple.OtherUpdateResponse](#mycompany-simple-otherupdateresponse)
    - [mycompany.simple.OtherWorkflowRequest](#mycompany-simple-otherworkflowrequest)
    - [mycompany.simple.OtherWorkflowRequest.Bar](#mycompany-simple-otherworkflowrequest-bar)
    - [mycompany.simple.OtherWorkflowRequest.Baz](#mycompany-simple-otherworkflowrequest-baz)
    - [mycompany.simple.OtherWorkflowResponse](#mycompany-simple-otherworkflowresponse)
    - [mycompany.simple.Qux](#mycompany-simple-qux)
    - [mycompany.simple.SomeActivity2Request](#mycompany-simple-someactivity2request)
    - [mycompany.simple.SomeActivity3Request](#mycompany-simple-someactivity3request)
    - [mycompany.simple.SomeActivity3Response](#mycompany-simple-someactivity3response)
    - [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)
    - [mycompany.simple.SomeQuery1Response](#mycompany-simple-somequery1response)
    - [mycompany.simple.SomeQuery2Request](#mycompany-simple-somequery2request)
    - [mycompany.simple.SomeQuery2Response](#mycompany-simple-somequery2response)
    - [mycompany.simple.SomeSignal2Request](#mycompany-simple-somesignal2request)
    - [mycompany.simple.SomeSignal3Request](#mycompany-simple-somesignal3request)
    - [mycompany.simple.SomeSignal3Response](#mycompany-simple-somesignal3response)
    - [mycompany.simple.SomeUpdate1Request](#mycompany-simple-someupdate1request)
    - [mycompany.simple.SomeUpdate1Response](#mycompany-simple-someupdate1response)
    - [mycompany.simple.SomeWorkflow1Request](#mycompany-simple-someworkflow1request)
    - [mycompany.simple.SomeWorkflow1Response](#mycompany-simple-someworkflow1response)
    - [mycompany.simple.SomeWorkflow3Request](#mycompany-simple-someworkflow3request)
    - [mycompany.simple.WhatRequest](#mycompany-simple-whatrequest)
- [test.option.v1](#test-option-v1)
  - Services
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
- [test.xnserr.v1](#test-xnserr-v1)
  - Services
    - [test.xnserr.v1.Server](#test-xnserr-v1-server)
      - [Workflows](#test-xnserr-v1-server-workflows)
        - [test.xnserr.v1.Server.Sleep](#test-xnserr-v1-server-sleep-workflow)
    - [test.xnserr.v1.Client](#test-xnserr-v1-client)
      - [Workflows](#test-xnserr-v1-client-workflows)
        - [test.xnserr.v1.Client.CallSleep](#test-xnserr-v1-client-callsleep-workflow)
  - Messages
    - [test.xnserr.v1.CallSleepRequest](#test-xnserr-v1-callsleeprequest)
    - [test.xnserr.v1.Failure](#test-xnserr-v1-failure)
    - [test.xnserr.v1.FailureInfo](#test-xnserr-v1-failureinfo)
    - [test.xnserr.v1.SleepRequest](#test-xnserr-v1-sleeprequest)
- [google.protobuf](#google-protobuf)
  - Messages
    - [google.protobuf.Any](#google-protobuf-any)
    - [google.protobuf.Duration](#google-protobuf-duration)
    - [google.protobuf.ListValue](#google-protobuf-listvalue)
    - [google.protobuf.NullValue](#google-protobuf-nullvalue)
    - [google.protobuf.Struct](#google-protobuf-struct)
    - [google.protobuf.Struct.FieldsEntry](#google-protobuf-struct-fieldsentry)
    - [google.protobuf.Timestamp](#google-protobuf-timestamp)
    - [google.protobuf.Value](#google-protobuf-value)
- [mycompany.simple.common.v1](#mycompany-simple-common-v1)
  - Messages
    - [mycompany.simple.common.v1.Example](#mycompany-simple-common-v1-example)
    - [mycompany.simple.common.v1.PaginatedRequest](#mycompany-simple-common-v1-paginatedrequest)
    - [mycompany.simple.common.v1.PaginatedResponse](#mycompany-simple-common-v1-paginatedresponse)
- [temporal.xns.v1](#temporal-xns-v1)
  - Messages
    - [temporal.xns.v1.IDReusePolicy](#temporal-xns-v1-idreusepolicy)
    - [temporal.xns.v1.RetryPolicy](#temporal-xns-v1-retrypolicy)
    - [temporal.xns.v1.StartWorkflowOptions](#temporal-xns-v1-startworkflowoptions)

<a name="example-helloworld-v1"></a>
# example.helloworld.v1

<a name="example-helloworld-v1-services"></a>
## Services

<a name="example-helloworld-v1-example"></a>
## example.helloworld.v1.Example

<a name="example-helloworld-v1-example-workflows"></a>
### Workflows

---
<a name="example-v1-hello-workflow"></a>
### example.v1.Hello

<pre>
Hello prints a friendly greeting and waits for goodbye
</pre>

**Input:** [example.helloworld.v1.HelloRequest](#example-helloworld-v1-hellorequest)

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

**Output:** [example.helloworld.v1.HelloResponse](#example-helloworld-v1-helloresponse)

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
<tr><td>id</td><td><pre><code>hello/${! name.or("World") }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-helloworld-v1-example-goodbye-signal">example.helloworld.v1.Example.Goodbye</a></td><td>false</td></tr>
</table>   

<a name="example-helloworld-v1-example-signals"></a>
### Signals

---
<a name="example-helloworld-v1-example-goodbye-signal"></a>
### example.helloworld.v1.Example.Goodbye

<pre>
Goodbye signals a running workflow to exit
</pre>

**Input:** [example.helloworld.v1.GoodbyeRequest](#example-helloworld-v1-goodbyerequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>   

<a name="example-helloworld-v1-messages"></a>
## Messages

<a name="example-helloworld-v1-goodbyerequest"></a>
### example.helloworld.v1.GoodbyeRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>



<a name="example-helloworld-v1-hellorequest"></a>
### example.helloworld.v1.HelloRequest

<pre>
HelloRequest describes the input to a Hello workflow
</pre>

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



<a name="example-helloworld-v1-helloresponse"></a>
### example.helloworld.v1.HelloResponse

<pre>
HelloResponse describes the output from a Hello workflow
</pre>

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



<a name="example-mutex-v1"></a>
# example.mutex.v1

<a name="example-mutex-v1-services"></a>
## Services

<a name="example-mutex-v1-example"></a>
## example.mutex.v1.Example

<a name="example-mutex-v1-example-workflows"></a>
### Workflows

---
<a name="example-mutex-v1-example-mutex-workflow"></a>
### example.mutex.v1.Example.Mutex

**Input:** [example.mutex.v1.MutexInput](#example-mutex-v1-mutexinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>mutex:${! resourceId }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
<tr><td>retry_policy.backoff_coefficient</td><td>2</td></tr>
<tr><td>retry_policy.initial_interval</td><td>1 second</td></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>retry_policy.max_interval</td><td>1 minute</td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-mutex-v1-example-acquirelock-signal">example.mutex.v1.Example.AcquireLock</a></td><td>true</td></tr>
<tr><td><a href="#example-mutex-v1-example-releaselock-signal">example.mutex.v1.Example.ReleaseLock</a></td><td>false</td></tr>
</table>

---
<a name="example-mutex-v1-example-sampleworkflowwithmutex-workflow"></a>
### example.mutex.v1.Example.SampleWorkflowWithMutex

**Input:** [example.mutex.v1.SampleWorkflowWithMutexInput](#example-mutex-v1-sampleworkflowwithmutexinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>SampleWorkflow1WithMutex_${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-mutex-v1-example-lockacquired-signal">example.mutex.v1.Example.LockAcquired</a></td><td>false</td></tr>
</table>   

<a name="example-mutex-v1-example-signals"></a>
### Signals

---
<a name="example-mutex-v1-example-acquirelock-signal"></a>
### example.mutex.v1.Example.AcquireLock



**Input:** [example.mutex.v1.AcquireLockInput](#example-mutex-v1-acquirelockinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>timeout</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: timeout
go_name: Timeout</pre></td>
</tr><tr>
<td>workflow_id</td>
<td>string</td>
<td><pre>
json_name: workflowId
go_name: WorkflowId</pre></td>
</tr>
</table>

---
<a name="example-mutex-v1-example-lockacquired-signal"></a>
### example.mutex.v1.Example.LockAcquired



**Input:** [example.mutex.v1.LockAcquiredInput](#example-mutex-v1-lockacquiredinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>

---
<a name="example-mutex-v1-example-releaselock-signal"></a>
### example.mutex.v1.Example.ReleaseLock



**Input:** [example.mutex.v1.ReleaseLockInput](#example-mutex-v1-releaselockinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>  

<a name="example-mutex-v1-example-activities"></a>
### Activities

---
<a name="example-mutex-v1-example-mutex-activity"></a>
### example.mutex.v1.Example.Mutex



**Input:** [example.mutex.v1.MutexInput](#example-mutex-v1-mutexinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table>   

<a name="example-mutex-v1-messages"></a>
## Messages

<a name="example-mutex-v1-acquirelockinput"></a>
### example.mutex.v1.AcquireLockInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>timeout</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: timeout
go_name: Timeout</pre></td>
</tr><tr>
<td>workflow_id</td>
<td>string</td>
<td><pre>
json_name: workflowId
go_name: WorkflowId</pre></td>
</tr>
</table>



<a name="example-mutex-v1-lockacquiredinput"></a>
### example.mutex.v1.LockAcquiredInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>



<a name="example-mutex-v1-mutexinput"></a>
### example.mutex.v1.MutexInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr>
</table>



<a name="example-mutex-v1-releaselockinput"></a>
### example.mutex.v1.ReleaseLockInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>lease_id</td>
<td>string</td>
<td><pre>
json_name: leaseId
go_name: LeaseId</pre></td>
</tr>
</table>



<a name="example-mutex-v1-sampleworkflowwithmutexinput"></a>
### example.mutex.v1.SampleWorkflowWithMutexInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>resource_id</td>
<td>string</td>
<td><pre>
json_name: resourceId
go_name: ResourceId</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>



<a name="example-schedule-v1"></a>
# example.schedule.v1

<a name="example-schedule-v1-services"></a>
## Services

<a name="example-schedule-v1-example"></a>
## example.schedule.v1.Example

<a name="example-schedule-v1-example-workflows"></a>
### Workflows

---
<a name="example-schedule-v1-schedule-workflow"></a>
### example.schedule.v1.Schedule

**Input:** [example.schedule.v1.ScheduleInput](#example-schedule-v1-scheduleinput)



**Output:** [example.schedule.v1.ScheduleOutput](#example-schedule-v1-scheduleoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>started_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: startedAt
go_name: StartedAt</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="example-schedule-v1-messages"></a>
## Messages

<a name="example-schedule-v1-scheduleinput"></a>
### example.schedule.v1.ScheduleInput



<a name="example-schedule-v1-scheduleoutput"></a>
### example.schedule.v1.ScheduleOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>started_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: startedAt
go_name: StartedAt</pre></td>
</tr>
</table>



<a name="example-searchattributes-v1"></a>
# example.searchattributes.v1

<a name="example-searchattributes-v1-services"></a>
## Services

<a name="example-searchattributes-v1-example"></a>
## example.searchattributes.v1.Example

<a name="example-searchattributes-v1-example-workflows"></a>
### Workflows

---
<a name="example-searchattributes-v1-example-searchattributes-workflow"></a>
### example.searchattributes.v1.Example.SearchAttributes

**Input:** [example.searchattributes.v1.SearchAttributesInput](#example-searchattributes-v1-searchattributesinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>search_attributes_${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
<tr><td>search_attributes</td><td><pre><code>CustomKeywordField = customKeywordField 
CustomTextField = customTextField 
CustomIntField = customIntField.int64() 
CustomDoubleField = customDoubleField 
CustomBoolField = customBoolField 
CustomDatetimeField = customDatetimeField.ts_parse("2006-01-02T15:04:05Z")</code></pre></td></tr>
</table>     

<a name="example-searchattributes-v1-messages"></a>
## Messages

<a name="example-searchattributes-v1-searchattributesinput"></a>
### example.searchattributes.v1.SearchAttributesInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>custom_bool_field</td>
<td>bool</td>
<td><pre>
json_name: customBoolField
go_name: CustomBoolField</pre></td>
</tr><tr>
<td>custom_datetime_field</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: customDatetimeField
go_name: CustomDatetimeField</pre></td>
</tr><tr>
<td>custom_double_field</td>
<td>double</td>
<td><pre>
json_name: customDoubleField
go_name: CustomDoubleField</pre></td>
</tr><tr>
<td>custom_int_field</td>
<td>int64</td>
<td><pre>
json_name: customIntField
go_name: CustomIntField</pre></td>
</tr><tr>
<td>custom_keyword_field</td>
<td>string</td>
<td><pre>
json_name: customKeywordField
go_name: CustomKeywordField</pre></td>
</tr><tr>
<td>custom_text_field</td>
<td>string</td>
<td><pre>
json_name: customTextField
go_name: CustomTextField</pre></td>
</tr>
</table>



<a name="example-updatabletimer-v1"></a>
# example.updatabletimer.v1

<a name="example-updatabletimer-v1-services"></a>
## Services

<a name="example-updatabletimer-v1-example"></a>
## example.updatabletimer.v1.Example

<a name="example-updatabletimer-v1-example-workflows"></a>
### Workflows

---
<a name="updatabletimer-workflow"></a>
### UpdatableTimer

<pre>
UpdatableTimer describes an updatable timer workflow
</pre>

**Input:** [example.updatabletimer.v1.UpdatableTimerInput](#example-updatabletimer-v1-updatabletimerinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>initial_wake_up_time</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: initialWakeUpTime
go_name: InitialWakeUpTime</pre></td>
</tr><tr>
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
<tr><td>id</td><td><pre><code>updatable-timer/${! name.or(uuid_v4()) }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#example-updatabletimer-v1-example-getwakeuptime-query">example.updatabletimer.v1.Example.GetWakeUpTime</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-updatabletimer-v1-example-updatewakeuptime-signal">example.updatabletimer.v1.Example.UpdateWakeUpTime</a></td><td>false</td></tr>
</table>  

<a name="example-updatabletimer-v1-example-queries"></a>
### Queries

---
<a name="example-updatabletimer-v1-example-getwakeuptime-query"></a>
### example.updatabletimer.v1.Example.GetWakeUpTime

<pre>
GetWakeUpTime retrieves the current timer expiration timestamp
</pre>

**Output:** [example.updatabletimer.v1.GetWakeUpTimeOutput](#example-updatabletimer-v1-getwakeuptimeoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>  

<a name="example-updatabletimer-v1-example-signals"></a>
### Signals

---
<a name="example-updatabletimer-v1-example-updatewakeuptime-signal"></a>
### example.updatabletimer.v1.Example.UpdateWakeUpTime

<pre>
UpdateWakeUpTime updates the timer expiration timestamp
</pre>

**Input:** [example.updatabletimer.v1.UpdateWakeUpTimeInput](#example-updatabletimer-v1-updatewakeuptimeinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>   

<a name="example-updatabletimer-v1-messages"></a>
## Messages

<a name="example-updatabletimer-v1-getwakeuptimeoutput"></a>
### example.updatabletimer.v1.GetWakeUpTimeOutput

<pre>
GetWakeUpTimeOutput describes the input to a GetWakeUpTime query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>



<a name="example-updatabletimer-v1-updatabletimerinput"></a>
### example.updatabletimer.v1.UpdatableTimerInput

<pre>
UpdatableTimerInput describes the input to a UpdatableTimer workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>initial_wake_up_time</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: initialWakeUpTime
go_name: InitialWakeUpTime</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="example-updatabletimer-v1-updatewakeuptimeinput"></a>
### example.updatabletimer.v1.UpdateWakeUpTimeInput

<pre>
UpdateWakeUpTimeInput describes the input to a UpdateWakeUpTime signal
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>wake_up_time</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: wakeUpTime
go_name: WakeUpTime</pre></td>
</tr>
</table>



<a name="example-v1"></a>
# example.v1

<a name="example-v1-services"></a>
## Services

<a name="example-v1-example"></a>
## example.v1.Example

<a name="example-v1-example-workflows"></a>
### Workflows

---
<a name="example-v1-example-createfoo-workflow"></a>
### example.v1.Example.CreateFoo

<pre>
CreateFoo creates a new foo operation
</pre>

**Input:** [example.v1.CreateFooRequest](#example-v1-createfoorequest)

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
unique foo name<br>

json_name: name
go_name: Name</pre></td>
</tr>
</table>

**Output:** [example.v1.CreateFooResponse](#example-v1-createfooresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-v1-foo">example.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>execution_timeout</td><td>1 hour</td></tr>
<tr><td>id</td><td><pre><code>create-foo/${! name.slug() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#example-v1-example-getfooprogress-query">example.v1.Example.GetFooProgress</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-v1-example-setfooprogress-signal">example.v1.Example.SetFooProgress</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#example-v1-example-updatefooprogress-update">example.v1.Example.UpdateFooProgress</a></td></tr>
</table>  

<a name="example-v1-example-queries"></a>
### Queries

---
<a name="example-v1-example-getfooprogress-query"></a>
### example.v1.Example.GetFooProgress

<pre>
GetFooProgress returns the status of a CreateFoo operation
</pre>

**Output:** [example.v1.GetFooProgressResponse](#example-v1-getfooprogressresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>  

<a name="example-v1-example-signals"></a>
### Signals

---
<a name="example-v1-example-setfooprogress-signal"></a>
### example.v1.Example.SetFooProgress

<pre>
SetFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.v1.SetFooProgressRequest](#example-v1-setfooprogressrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>  

<a name="example-v1-example-updates"></a>
### Updates

---
<a name="example-v1-example-updatefooprogress-update"></a>
### example.v1.Example.UpdateFooProgress

<pre>
UpdateFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.v1.SetFooProgressRequest](#example-v1-setfooprogressrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>

**Output:** [example.v1.GetFooProgressResponse](#example-v1-getfooprogressresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

<a name="example-v1-example-activities"></a>
### Activities

---
<a name="example-v1-example-notify-activity"></a>
### example.v1.Example.Notify

<pre>
Notify sends a notification
</pre>

**Input:** [example.v1.NotifyRequest](#example-v1-notifyrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="example-v1-messages"></a>
## Messages

<a name="example-v1-createfoorequest"></a>
### example.v1.CreateFooRequest

<pre>
CreateFooRequest describes the input to a CreateFoo workflow
</pre>

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
unique foo name<br>

json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="example-v1-createfooresponse"></a>
### example.v1.CreateFooResponse

<pre>
SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-v1-foo">example.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="example-v1-foo"></a>
### example.v1.Foo

<pre>
Foo describes an illustrative foo resource
</pre>

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
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-v1-foo-status"></a>
### example.v1.Foo.Status

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>FOO_STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>FOO_STATUS_READY</td>
<td></td>
</tr><tr>
<td>FOO_STATUS_CREATING</td>
<td></td>
</tr>
</table>

<a name="example-v1-getfooprogressresponse"></a>
### example.v1.GetFooProgressResponse

<pre>
GetFooProgressResponse describes the output from a GetFooProgress query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-v1-foo-status">example.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-v1-notifyrequest"></a>
### example.v1.NotifyRequest

<pre>
NotifyRequest describes the input to a Notify activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>



<a name="example-v1-setfooprogressrequest"></a>
### example.v1.SetFooProgressRequest

<pre>
SetFooProgressRequest describes the input to a SetFooProgress signal
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>



<a name="example-xns-v1"></a>
# example.xns.v1

<a name="example-xns-v1-services"></a>
## Services

<a name="example-xns-v1-xns"></a>
## example.xns.v1.Xns

<a name="example-xns-v1-xns-workflows"></a>
### Workflows

---
<a name="example-xns-v1-xns-provisionfoo-workflow"></a>
### example.xns.v1.Xns.ProvisionFoo

**Input:** [example.xns.v1.ProvisionFooRequest](#example-xns-v1-provisionfoorequest)

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
unique foo name<br>

json_name: name
go_name: Name</pre></td>
</tr>
</table>

**Output:** [example.xns.v1.ProvisionFooResponse](#example-xns-v1-provisionfooresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>provision-foo/${! name.slug() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="example-xns-v1-example"></a>
## example.xns.v1.Example

<a name="example-xns-v1-example-workflows"></a>
### Workflows

---
<a name="example-xns-v1-example-createfoo-workflow"></a>
### example.xns.v1.Example.CreateFoo

<pre>
CreateFoo creates a new foo operation
</pre>

**Input:** [example.xns.v1.CreateFooRequest](#example-xns-v1-createfoorequest)

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
unique foo name<br>

json_name: name
go_name: Name</pre></td>
</tr>
</table>

**Output:** [example.xns.v1.CreateFooResponse](#example-xns-v1-createfooresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>execution_timeout</td><td>1 hour</td></tr>
<tr><td>id</td><td><pre><code>create-foo/${! name.slug() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE</code></pre></td></tr>
<tr><td>xns.heartbeat_interval</td><td>10 seconds</td></tr>
<tr><td>xns.heartbeat_timeout</td><td>20 seconds</td></tr>
<tr><td>xns.start_to_close_timeout</td><td>1 hour 30 seconds</td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#example-xns-v1-example-getfooprogress-query">example.xns.v1.Example.GetFooProgress</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-xns-v1-example-setfooprogress-signal">example.xns.v1.Example.SetFooProgress</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#example-xns-v1-example-updatefooprogress-update">example.xns.v1.Example.UpdateFooProgress</a></td></tr>
</table>  

<a name="example-xns-v1-example-queries"></a>
### Queries

---
<a name="example-xns-v1-example-getfooprogress-query"></a>
### example.xns.v1.Example.GetFooProgress

<pre>
GetFooProgress returns the status of a CreateFoo operation
</pre>

**Output:** [example.xns.v1.GetFooProgressResponse](#example-xns-v1-getfooprogressresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>  

<a name="example-xns-v1-example-signals"></a>
### Signals

---
<a name="example-xns-v1-example-setfooprogress-signal"></a>
### example.xns.v1.Example.SetFooProgress

<pre>
SetFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.xns.v1.SetFooProgressRequest](#example-xns-v1-setfooprogressrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>  

<a name="example-xns-v1-example-updates"></a>
### Updates

---
<a name="example-xns-v1-example-updatefooprogress-update"></a>
### example.xns.v1.Example.UpdateFooProgress

<pre>
UpdateFooProgress sets the current status of a CreateFoo operation
</pre>

**Input:** [example.xns.v1.SetFooProgressRequest](#example-xns-v1-setfooprogressrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>

**Output:** [example.xns.v1.GetFooProgressResponse](#example-xns-v1-getfooprogressresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
</table>

<a name="example-xns-v1-example-activities"></a>
### Activities

---
<a name="example-xns-v1-example-notify-activity"></a>
### example.xns.v1.Example.Notify

<pre>
Notify sends a notification
</pre>

**Input:** [example.xns.v1.NotifyRequest](#example-xns-v1-notifyrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="example-xns-v1-messages"></a>
## Messages

<a name="example-xns-v1-createfoorequest"></a>
### example.xns.v1.CreateFooRequest

<pre>
CreateFooRequest describes the input to a CreateFoo workflow
</pre>

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
unique foo name<br>

json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="example-xns-v1-createfooresponse"></a>
### example.xns.v1.CreateFooResponse

<pre>
SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="example-xns-v1-foo"></a>
### example.xns.v1.Foo

<pre>
Foo describes an illustrative foo resource
</pre>

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
</tr><tr>
<td>status</td>
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-xns-v1-foo-status"></a>
### example.xns.v1.Foo.Status

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>FOO_STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>FOO_STATUS_READY</td>
<td></td>
</tr><tr>
<td>FOO_STATUS_CREATING</td>
<td></td>
</tr>
</table>

<a name="example-xns-v1-getfooprogressresponse"></a>
### example.xns.v1.GetFooProgressResponse

<pre>
GetFooProgressResponse describes the output from a GetFooProgress query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
json_name: progress
go_name: Progress</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#example-xns-v1-foo-status">example.xns.v1.Foo.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="example-xns-v1-notifyrequest"></a>
### example.xns.v1.NotifyRequest

<pre>
NotifyRequest describes the input to a Notify activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>message</td>
<td>string</td>
<td><pre>
json_name: message
go_name: Message</pre></td>
</tr>
</table>



<a name="example-xns-v1-provisionfoorequest"></a>
### example.xns.v1.ProvisionFooRequest

<pre>
ProvisionFooRequest describes the input to a ProvisionFoo workflow
</pre>

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
unique foo name<br>

json_name: name
go_name: Name</pre></td>
</tr>
</table>



<a name="example-xns-v1-provisionfooresponse"></a>
### example.xns.v1.ProvisionFooResponse

<pre>
SampleWorkflowWithMutexResponse describes the output from a ProvisionFoo workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td><a href="#example-xns-v1-foo">example.xns.v1.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="example-xns-v1-setfooprogressrequest"></a>
### example.xns.v1.SetFooProgressRequest

<pre>
SetFooProgressRequest describes the input to a SetFooProgress signal
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>progress</td>
<td>float</td>
<td><pre>
value of current workflow progress<br>

json_name: progress
go_name: Progress</pre></td>
</tr>
</table>



<a name="mycompany-simple"></a>
# mycompany.simple

<a name="mycompany-simple-services"></a>
## Services

<a name="mycompany-simple-simple"></a>
## mycompany.simple.Simple

<a name="mycompany-simple-simple-workflows"></a>
### Workflows

---
<a name="mycompany-simple-someworkflow1-workflow"></a>
### mycompany.simple.SomeWorkflow1

<pre>
SomeWorkflow1 does some workflow thing.
</pre>

**Input:** [mycompany.simple.SomeWorkflow1Request](#mycompany-simple-someworkflow1request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeWorkflow1Response](#mycompany-simple-someworkflow1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>some-workflow-1/${! id }/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#mycompany-simple-simple-somequery1-query">mycompany.simple.Simple.SomeQuery1</a></td></tr>
<tr><td><a href="#mycompany-simple-simple-somequery2-query">mycompany.simple.Simple.SomeQuery2</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal1-signal">mycompany.simple.Simple.SomeSignal1</a></td><td>false</td></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal2-signal">mycompany.simple.Simple.SomeSignal2</a></td><td>false</td></tr>
</table>

---
<a name="mycompany-simple-someworkflow2-workflow"></a>
### mycompany.simple.SomeWorkflow2

<pre>
SomeWorkflow2 does some workflow thing.
</pre>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>some-workflow-2/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal1-signal">mycompany.simple.Simple.SomeSignal1</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#mycompany-simple-simple-someupdate1-update">mycompany.simple.Simple.SomeUpdate1</a></td></tr>
</table>

---
<a name="mycompany-simple-simple-someworkflow3-workflow"></a>
### mycompany.simple.Simple.SomeWorkflow3

<pre>
SomeWorkflow3 does some workflow thing.
Deprecated: Use SomeWorkflow2 instead.
</pre>

**Input:** [mycompany.simple.SomeWorkflow3Request](#mycompany-simple-someworkflow3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>execution_timeout</td><td>1 hour</td></tr>
<tr><td>id</td><td><pre><code>some-workflow-3/${! id }/${! requestVal }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE</code></pre></td></tr>
<tr><td>retry_policy.max_attempts</td><td>2</td></tr>
<tr><td>task_queue</td><td><pre><code>my-task-queue-2</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-simple-somesignal2-signal">mycompany.simple.Simple.SomeSignal2</a></td><td>true</td></tr>
</table>

---
<a name="mycompany-simple-simple-someworkflow4-workflow"></a>
### mycompany.simple.Simple.SomeWorkflow4

<pre>
SomeWorkflow4 retrieves a paginated list of items
</pre>

**Input:** [mycompany.simple.common.v1.PaginatedRequest](#mycompany-simple-common-v1-paginatedrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cursor</td>
<td>bytes</td>
<td><pre>
json_name: cursor
go_name: Cursor</pre></td>
</tr><tr>
<td>limit</td>
<td>uint32</td>
<td><pre>
json_name: limit
go_name: Limit</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.common.v1.PaginatedResponse](#mycompany-simple-common-v1-paginatedresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="#google-protobuf-any">google.protobuf.Any</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>next_cursor</td>
<td>bytes</td>
<td><pre>
json_name: nextCursor
go_name: NextCursor</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>some-workflow-4/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>  

<a name="mycompany-simple-simple-queries"></a>
### Queries

---
<a name="mycompany-simple-simple-somequery1-query"></a>
### mycompany.simple.Simple.SomeQuery1

<pre>
SomeQuery1 queries some thing.
</pre>

**Output:** [mycompany.simple.SomeQuery1Response](#mycompany-simple-somequery1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

---
<a name="mycompany-simple-simple-somequery2-query"></a>
### mycompany.simple.Simple.SomeQuery2

<pre>
SomeQuery2 queries some thing.
</pre>

**Input:** [mycompany.simple.SomeQuery2Request](#mycompany-simple-somequery2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeQuery2Response](#mycompany-simple-somequery2response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>  

<a name="mycompany-simple-simple-signals"></a>
### Signals

---
<a name="mycompany-simple-simple-somesignal1-signal"></a>
### mycompany.simple.Simple.SomeSignal1

<pre>
SomeSignal1 is a signal.
</pre>

---
<a name="mycompany-simple-simple-somesignal2-signal"></a>
### mycompany.simple.Simple.SomeSignal2

<pre>
SomeSignal2 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal2Request](#mycompany-simple-somesignal2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

---
<a name="mycompany-simple-simple-somesignal3-signal"></a>
### mycompany.simple.Simple.SomeSignal3

<pre>
SomeSignal3 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal3Request](#mycompany-simple-somesignal3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>  

<a name="mycompany-simple-simple-updates"></a>
### Updates

---
<a name="mycompany-simple-simple-someupdate1-update"></a>
### mycompany.simple.Simple.SomeUpdate1

<pre>
SomeUpdate1 updates a SomeWorkflow2
</pre>

**Input:** [mycompany.simple.SomeUpdate1Request](#mycompany-simple-someupdate1request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeUpdate1Response](#mycompany-simple-someupdate1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>validate</td><td>true</td></tr>
<tr><td>wait_policy</td><td><pre>WAIT_POLICY_COMPLETED</pre></td></tr>
</table>

<a name="mycompany-simple-simple-activities"></a>
### Activities

---
<a name="mycompany-simple-someactivity1-activity"></a>
### mycompany.simple.SomeActivity1

<pre>
SomeActivity1 does some activity thing.
</pre> 

---
<a name="mycompany-simple-simple-someactivity2-activity"></a>
### mycompany.simple.Simple.SomeActivity2

<pre>
SomeActivity2 does some activity thing.
</pre>

**Input:** [mycompany.simple.SomeActivity2Request](#mycompany-simple-someactivity2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_interval</td><td>30 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
<tr><td>wait_for_cancellation</td><td>true</td></tr>
</table> 

---
<a name="mycompany-simple-simple-someactivity3-activity"></a>
### mycompany.simple.Simple.SomeActivity3

<pre>
SomeActivity3 does some activity thing.
</pre>

**Input:** [mycompany.simple.SomeActivity3Request](#mycompany-simple-someactivity3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeActivity3Response](#mycompany-simple-someactivity3response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
<tr><td>task_queue</td><td>some-other-task-queue</td></tr>
</table> 

---
<a name="mycompany-simple-simple-someactivity4-activity"></a>
### mycompany.simple.Simple.SomeActivity4

<pre>
SomeActivity4 does some activity thing.
</pre>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>heartbeat_timeout</td><td>30 seconds</td></tr>
<tr><td>retry_policy.max_attempts</td><td>5</td></tr>
<tr><td>schedule_to_close_timeout</td><td>30 seconds</td></tr>
<tr><td>schedule_to_start_timeout</td><td>5 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>1 minute</td></tr>
<tr><td>wait_for_cancellation</td><td>true</td></tr>
<tr><td>task_queue</td><td>some-other-task-queue</td></tr>
</table> 

---
<a name="mycompany-simple-simple-somesignal1-activity"></a>
### mycompany.simple.Simple.SomeSignal1

<pre>
SomeSignal1 is a signal.
</pre>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-simple-somesignal2-activity"></a>
### mycompany.simple.Simple.SomeSignal2

<pre>
SomeSignal2 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal2Request](#mycompany-simple-somesignal2request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-simple-somesignal3-activity"></a>
### mycompany.simple.Simple.SomeSignal3

<pre>
SomeSignal3 is a signal.
</pre>

**Input:** [mycompany.simple.SomeSignal3Request](#mycompany-simple-somesignal3request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeSignal3Response](#mycompany-simple-somesignal3response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-simple-someupdate1-activity"></a>
### mycompany.simple.Simple.SomeUpdate1

<pre>
SomeUpdate1 updates a SomeWorkflow2
</pre>

**Input:** [mycompany.simple.SomeUpdate1Request](#mycompany-simple-someupdate1request)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.SomeUpdate1Response](#mycompany-simple-someupdate1response)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>retry_policy.non_retryable_error_types</td><td>something</td></tr>
<tr><td>start_to_close_timeout</td><td>10 seconds</td></tr>
</table>   

<a name="mycompany-simple-other"></a>
## mycompany.simple.Other

<a name="mycompany-simple-other-workflows"></a>
### Workflows

---
<a name="mycompany-simple-other-otherworkflow-workflow"></a>
### mycompany.simple.Other.OtherWorkflow

**Input:** [mycompany.simple.OtherWorkflowRequest](#mycompany-simple-otherworkflowrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td><a href="#mycompany-simple-otherworkflowrequest-bar">mycompany.simple.OtherWorkflowRequest.Bar</a></td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr><tr>
<td>baz</td>
<td><a href="#mycompany-simple-otherworkflowrequest-baz">mycompany.simple.OtherWorkflowRequest.Baz</a></td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr><tr>
<td>common_enum</td>
<td><a href="#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
<td><pre>
json_name: commonEnum
go_name: CommonEnum</pre></td>
</tr><tr>
<td>example_bool</td>
<td>bool</td>
<td><pre>
json_name: exampleBool
go_name: ExampleBool</pre></td>
</tr><tr>
<td>example_bytes</td>
<td>bytes</td>
<td><pre>
json_name: exampleBytes
go_name: ExampleBytes</pre></td>
</tr><tr>
<td>example_double</td>
<td>double</td>
<td><pre>
json_name: exampleDouble
go_name: ExampleDouble</pre></td>
</tr><tr>
<td>example_duration</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="#google-protobuf-empty">google.protobuf.Empty</a></td>
<td><pre>
json_name: exampleEmpty
go_name: ExampleEmpty</pre></td>
</tr><tr>
<td>example_enum</td>
<td><a href="#mycompany-simple-otherenum">mycompany.simple.OtherEnum</a></td>
<td><pre>
json_name: exampleEnum
go_name: ExampleEnum</pre></td>
</tr><tr>
<td>example_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: exampleFixed32
go_name: ExampleFixed32</pre></td>
</tr><tr>
<td>example_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: exampleFixed64
go_name: ExampleFixed64</pre></td>
</tr><tr>
<td>example_float</td>
<td>float</td>
<td><pre>
json_name: exampleFloat
go_name: ExampleFloat</pre></td>
</tr><tr>
<td>example_int32</td>
<td>int32</td>
<td><pre>
json_name: exampleInt32
go_name: ExampleInt32</pre></td>
</tr><tr>
<td>example_int64</td>
<td>int64</td>
<td><pre>
json_name: exampleInt64
go_name: ExampleInt64</pre></td>
</tr><tr>
<td>example_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: exampleSfixed32
go_name: ExampleSfixed32</pre></td>
</tr><tr>
<td>example_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: exampleSfixed64
go_name: ExampleSfixed64</pre></td>
</tr><tr>
<td>example_sint32</td>
<td>sint32</td>
<td><pre>
json_name: exampleSint32
go_name: ExampleSint32</pre></td>
</tr><tr>
<td>example_sint64</td>
<td>sint64</td>
<td><pre>
json_name: exampleSint64
go_name: ExampleSint64</pre></td>
</tr><tr>
<td>example_timestamp</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: exampleTimestamp
go_name: ExampleTimestamp</pre></td>
</tr><tr>
<td>example_uint32</td>
<td>uint32</td>
<td><pre>
json_name: exampleUint32
go_name: ExampleUint32</pre></td>
</tr><tr>
<td>example_uint64</td>
<td>uint64</td>
<td><pre>
json_name: exampleUint64
go_name: ExampleUint64</pre></td>
</tr><tr>
<td>foo</td>
<td><a href="#mycompany-simple-foo">mycompany.simple.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr><tr>
<td>quux</td>
<td>string</td>
<td><pre>
json_name: quux
go_name: Quux</pre></td>
</tr><tr>
<td>qux</td>
<td><a href="#mycompany-simple-qux">mycompany.simple.Qux</a></td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr><tr>
<td>some_val</td>
<td>string</td>
<td><pre>
json_name: someVal
go_name: SomeVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.OtherWorkflowResponse](#mycompany-simple-otherworkflowresponse)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>other-workflow/${!uuid_v4()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>  

<a name="mycompany-simple-other-queries"></a>
### Queries

---
<a name="mycompany-simple-other-otherquery-query"></a>
### mycompany.simple.Other.OtherQuery



**Output:** [mycompany.simple.OtherQueryResponse](#mycompany-simple-otherqueryresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>filter</td>
<td>string</td>
<td><pre>
json_name: filter
go_name: Filter</pre></td>
</tr>
</table>  

<a name="mycompany-simple-other-signals"></a>
### Signals

---
<a name="mycompany-simple-other-othersignal-signal"></a>
### mycompany.simple.Other.OtherSignal



**Input:** [mycompany.simple.OtherSignalRequest](#mycompany-simple-othersignalrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>type</td>
<td>string</td>
<td><pre>
json_name: type
go_name: Type</pre></td>
</tr>
</table>  

<a name="mycompany-simple-other-updates"></a>
### Updates

---
<a name="mycompany-simple-other-otherupdate-update"></a>
### mycompany.simple.Other.OtherUpdate



**Input:** [mycompany.simple.OtherUpdateRequest](#mycompany-simple-otherupdaterequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>mode</td>
<td>string</td>
<td><pre>
json_name: mode
go_name: Mode</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.OtherUpdateResponse](#mycompany-simple-otherupdateresponse)



<a name="mycompany-simple-other-activities"></a>
### Activities

---
<a name="mycompany-simple-other-otherworkflow-activity"></a>
### mycompany.simple.Other.OtherWorkflow



**Input:** [mycompany.simple.OtherWorkflowRequest](#mycompany-simple-otherworkflowrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td><a href="#mycompany-simple-otherworkflowrequest-bar">mycompany.simple.OtherWorkflowRequest.Bar</a></td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr><tr>
<td>baz</td>
<td><a href="#mycompany-simple-otherworkflowrequest-baz">mycompany.simple.OtherWorkflowRequest.Baz</a></td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr><tr>
<td>common_enum</td>
<td><a href="#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
<td><pre>
json_name: commonEnum
go_name: CommonEnum</pre></td>
</tr><tr>
<td>example_bool</td>
<td>bool</td>
<td><pre>
json_name: exampleBool
go_name: ExampleBool</pre></td>
</tr><tr>
<td>example_bytes</td>
<td>bytes</td>
<td><pre>
json_name: exampleBytes
go_name: ExampleBytes</pre></td>
</tr><tr>
<td>example_double</td>
<td>double</td>
<td><pre>
json_name: exampleDouble
go_name: ExampleDouble</pre></td>
</tr><tr>
<td>example_duration</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="#google-protobuf-empty">google.protobuf.Empty</a></td>
<td><pre>
json_name: exampleEmpty
go_name: ExampleEmpty</pre></td>
</tr><tr>
<td>example_enum</td>
<td><a href="#mycompany-simple-otherenum">mycompany.simple.OtherEnum</a></td>
<td><pre>
json_name: exampleEnum
go_name: ExampleEnum</pre></td>
</tr><tr>
<td>example_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: exampleFixed32
go_name: ExampleFixed32</pre></td>
</tr><tr>
<td>example_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: exampleFixed64
go_name: ExampleFixed64</pre></td>
</tr><tr>
<td>example_float</td>
<td>float</td>
<td><pre>
json_name: exampleFloat
go_name: ExampleFloat</pre></td>
</tr><tr>
<td>example_int32</td>
<td>int32</td>
<td><pre>
json_name: exampleInt32
go_name: ExampleInt32</pre></td>
</tr><tr>
<td>example_int64</td>
<td>int64</td>
<td><pre>
json_name: exampleInt64
go_name: ExampleInt64</pre></td>
</tr><tr>
<td>example_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: exampleSfixed32
go_name: ExampleSfixed32</pre></td>
</tr><tr>
<td>example_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: exampleSfixed64
go_name: ExampleSfixed64</pre></td>
</tr><tr>
<td>example_sint32</td>
<td>sint32</td>
<td><pre>
json_name: exampleSint32
go_name: ExampleSint32</pre></td>
</tr><tr>
<td>example_sint64</td>
<td>sint64</td>
<td><pre>
json_name: exampleSint64
go_name: ExampleSint64</pre></td>
</tr><tr>
<td>example_timestamp</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: exampleTimestamp
go_name: ExampleTimestamp</pre></td>
</tr><tr>
<td>example_uint32</td>
<td>uint32</td>
<td><pre>
json_name: exampleUint32
go_name: ExampleUint32</pre></td>
</tr><tr>
<td>example_uint64</td>
<td>uint64</td>
<td><pre>
json_name: exampleUint64
go_name: ExampleUint64</pre></td>
</tr><tr>
<td>foo</td>
<td><a href="#mycompany-simple-foo">mycompany.simple.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr><tr>
<td>quux</td>
<td>string</td>
<td><pre>
json_name: quux
go_name: Quux</pre></td>
</tr><tr>
<td>qux</td>
<td><a href="#mycompany-simple-qux">mycompany.simple.Qux</a></td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr><tr>
<td>some_val</td>
<td>string</td>
<td><pre>
json_name: someVal
go_name: SomeVal</pre></td>
</tr>
</table>

**Output:** [mycompany.simple.OtherWorkflowResponse](#mycompany-simple-otherworkflowresponse)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="mycompany-simple-ignored"></a>
## mycompany.simple.Ignored

<a name="mycompany-simple-ignored-workflows"></a>
### Workflows

---
<a name="mycompany-simple-ignored-what-workflow"></a>
### mycompany.simple.Ignored.What

**Input:** [mycompany.simple.WhatRequest](#mycompany-simple-whatrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>blah</td>
<td>string</td>
<td><pre>
json_name: blah
go_name: Blah</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>what/${!ksuid()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="mycompany-simple-onlyactivities"></a>
## mycompany.simple.OnlyActivities   

<a name="mycompany-simple-onlyactivities-activities"></a>
### Activities

---
<a name="mycompany-simple-onlyactivities-lonelyactivity1-activity"></a>
### mycompany.simple.OnlyActivities.LonelyActivity1



**Input:** [mycompany.simple.LonelyActivity1Request](#mycompany-simple-lonelyactivity1request)



**Output:** [mycompany.simple.LonelyActivity1Response](#mycompany-simple-lonelyactivity1response)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table>   

<a name="mycompany-simple-deprecated"></a>
## mycompany.simple.Deprecated

<a name="mycompany-simple-deprecated-workflows"></a>
### Workflows

---
<a name="mycompany-simple-deprecated-somedeprecatedworkflow1-workflow"></a>
### mycompany.simple.Deprecated.SomeDeprecatedWorkflow1

<pre>
SomeDeprecatedWorkflow1 does something
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedquery1-query">mycompany.simple.Deprecated.SomeDeprecatedQuery1</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedsignal1-signal">mycompany.simple.Deprecated.SomeDeprecatedSignal1</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedupdate1-update">mycompany.simple.Deprecated.SomeDeprecatedUpdate1</a></td></tr>
</table>

---
<a name="mycompany-simple-deprecated-somedeprecatedworkflow2-workflow"></a>
### mycompany.simple.Deprecated.SomeDeprecatedWorkflow2

<pre>
SomeDeprecatedWorkflow2 does something else

Deprecated: a custom workflow deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedquery2-query">mycompany.simple.Deprecated.SomeDeprecatedQuery2</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedsignal2-signal">mycompany.simple.Deprecated.SomeDeprecatedSignal2</a></td><td>true</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#mycompany-simple-deprecated-somedeprecatedupdate2-update">mycompany.simple.Deprecated.SomeDeprecatedUpdate2</a></td></tr>
</table>  

<a name="mycompany-simple-deprecated-queries"></a>
### Queries

---
<a name="mycompany-simple-deprecated-somedeprecatedquery1-query"></a>
### mycompany.simple.Deprecated.SomeDeprecatedQuery1

<pre>
SomeDeprecatedQuery1 does something else
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



---
<a name="mycompany-simple-deprecated-somedeprecatedquery2-query"></a>
### mycompany.simple.Deprecated.SomeDeprecatedQuery2

<pre>
SomeDeprecatedQuery2 does something else

Deprecated: a custom query deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)

  

<a name="mycompany-simple-deprecated-signals"></a>
### Signals

---
<a name="mycompany-simple-deprecated-somedeprecatedsignal1-signal"></a>
### mycompany.simple.Deprecated.SomeDeprecatedSignal1

<pre>
SomeDeprecatedSignal1 does something else
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



---
<a name="mycompany-simple-deprecated-somedeprecatedsignal2-signal"></a>
### mycompany.simple.Deprecated.SomeDeprecatedSignal2

<pre>
SomeDeprecatedSignal2 does something else

Deprecated: a custom signal deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)

  

<a name="mycompany-simple-deprecated-updates"></a>
### Updates

---
<a name="mycompany-simple-deprecated-somedeprecatedupdate1-update"></a>
### mycompany.simple.Deprecated.SomeDeprecatedUpdate1

<pre>
SomeDeprecatedUpdate1 does something else
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



---
<a name="mycompany-simple-deprecated-somedeprecatedupdate2-update"></a>
### mycompany.simple.Deprecated.SomeDeprecatedUpdate2

<pre>
SomeDeprecatedUpdate2 does something else

Deprecated: a custom signal deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



<a name="mycompany-simple-deprecated-activities"></a>
### Activities

---
<a name="mycompany-simple-deprecated-somedeprecatedactivity1-activity"></a>
### mycompany.simple.Deprecated.SomeDeprecatedActivity1

<pre>
SomeDeprecatedActivity1 does something
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table> 

---
<a name="mycompany-simple-deprecated-somedeprecatedactivity2-activity"></a>
### mycompany.simple.Deprecated.SomeDeprecatedActivity2

<pre>
SomeDeprecatedActivity2 does something else

Deprecated: a custom activity deprecation message.
</pre>

**Input:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Output:** [mycompany.simple.SomeDeprecatedMessage](#mycompany-simple-somedeprecatedmessage)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table>   

<a name="mycompany-simple-messages"></a>
## Messages

<a name="mycompany-simple-foo"></a>
### mycompany.simple.Foo

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>foo</td>
<td>string</td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr>
</table>



<a name="mycompany-simple-lonelyactivity1request"></a>
### mycompany.simple.LonelyActivity1Request



<a name="mycompany-simple-lonelyactivity1response"></a>
### mycompany.simple.LonelyActivity1Response



<a name="mycompany-simple-otherenum"></a>
### mycompany.simple.OtherEnum

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>OTHER_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>OTHER_FOO</td>
<td></td>
</tr><tr>
<td>OTHER_BAR</td>
<td></td>
</tr>
</table>

<a name="mycompany-simple-otherqueryresponse"></a>
### mycompany.simple.OtherQueryResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>filter</td>
<td>string</td>
<td><pre>
json_name: filter
go_name: Filter</pre></td>
</tr>
</table>



<a name="mycompany-simple-othersignalrequest"></a>
### mycompany.simple.OtherSignalRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>type</td>
<td>string</td>
<td><pre>
json_name: type
go_name: Type</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherupdaterequest"></a>
### mycompany.simple.OtherUpdateRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>mode</td>
<td>string</td>
<td><pre>
json_name: mode
go_name: Mode</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherupdateresponse"></a>
### mycompany.simple.OtherUpdateResponse



<a name="mycompany-simple-otherworkflowrequest"></a>
### mycompany.simple.OtherWorkflowRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td><a href="#mycompany-simple-otherworkflowrequest-bar">mycompany.simple.OtherWorkflowRequest.Bar</a></td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr><tr>
<td>baz</td>
<td><a href="#mycompany-simple-otherworkflowrequest-baz">mycompany.simple.OtherWorkflowRequest.Baz</a></td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr><tr>
<td>common_enum</td>
<td><a href="#mycompany-simple-common-v1-example">mycompany.simple.common.v1.Example</a></td>
<td><pre>
json_name: commonEnum
go_name: CommonEnum</pre></td>
</tr><tr>
<td>example_bool</td>
<td>bool</td>
<td><pre>
json_name: exampleBool
go_name: ExampleBool</pre></td>
</tr><tr>
<td>example_bytes</td>
<td>bytes</td>
<td><pre>
json_name: exampleBytes
go_name: ExampleBytes</pre></td>
</tr><tr>
<td>example_double</td>
<td>double</td>
<td><pre>
json_name: exampleDouble
go_name: ExampleDouble</pre></td>
</tr><tr>
<td>example_duration</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: exampleDuration
go_name: ExampleDuration</pre></td>
</tr><tr>
<td>example_empty</td>
<td><a href="#google-protobuf-empty">google.protobuf.Empty</a></td>
<td><pre>
json_name: exampleEmpty
go_name: ExampleEmpty</pre></td>
</tr><tr>
<td>example_enum</td>
<td><a href="#mycompany-simple-otherenum">mycompany.simple.OtherEnum</a></td>
<td><pre>
json_name: exampleEnum
go_name: ExampleEnum</pre></td>
</tr><tr>
<td>example_fixed32</td>
<td>fixed32</td>
<td><pre>
json_name: exampleFixed32
go_name: ExampleFixed32</pre></td>
</tr><tr>
<td>example_fixed64</td>
<td>fixed64</td>
<td><pre>
json_name: exampleFixed64
go_name: ExampleFixed64</pre></td>
</tr><tr>
<td>example_float</td>
<td>float</td>
<td><pre>
json_name: exampleFloat
go_name: ExampleFloat</pre></td>
</tr><tr>
<td>example_int32</td>
<td>int32</td>
<td><pre>
json_name: exampleInt32
go_name: ExampleInt32</pre></td>
</tr><tr>
<td>example_int64</td>
<td>int64</td>
<td><pre>
json_name: exampleInt64
go_name: ExampleInt64</pre></td>
</tr><tr>
<td>example_sfixed32</td>
<td>sfixed32</td>
<td><pre>
json_name: exampleSfixed32
go_name: ExampleSfixed32</pre></td>
</tr><tr>
<td>example_sfixed64</td>
<td>sfixed64</td>
<td><pre>
json_name: exampleSfixed64
go_name: ExampleSfixed64</pre></td>
</tr><tr>
<td>example_sint32</td>
<td>sint32</td>
<td><pre>
json_name: exampleSint32
go_name: ExampleSint32</pre></td>
</tr><tr>
<td>example_sint64</td>
<td>sint64</td>
<td><pre>
json_name: exampleSint64
go_name: ExampleSint64</pre></td>
</tr><tr>
<td>example_timestamp</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: exampleTimestamp
go_name: ExampleTimestamp</pre></td>
</tr><tr>
<td>example_uint32</td>
<td>uint32</td>
<td><pre>
json_name: exampleUint32
go_name: ExampleUint32</pre></td>
</tr><tr>
<td>example_uint64</td>
<td>uint64</td>
<td><pre>
json_name: exampleUint64
go_name: ExampleUint64</pre></td>
</tr><tr>
<td>foo</td>
<td><a href="#mycompany-simple-foo">mycompany.simple.Foo</a></td>
<td><pre>
json_name: foo
go_name: Foo</pre></td>
</tr><tr>
<td>quux</td>
<td>string</td>
<td><pre>
json_name: quux
go_name: Quux</pre></td>
</tr><tr>
<td>qux</td>
<td><a href="#mycompany-simple-qux">mycompany.simple.Qux</a></td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr><tr>
<td>some_val</td>
<td>string</td>
<td><pre>
json_name: someVal
go_name: SomeVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherworkflowrequest-bar"></a>
### mycompany.simple.OtherWorkflowRequest.Bar

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bar</td>
<td>string</td>
<td><pre>
json_name: bar
go_name: Bar</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherworkflowrequest-baz"></a>
### mycompany.simple.OtherWorkflowRequest.Baz

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>baz</td>
<td>string</td>
<td><pre>
json_name: baz
go_name: Baz</pre></td>
</tr>
</table>



<a name="mycompany-simple-otherworkflowresponse"></a>
### mycompany.simple.OtherWorkflowResponse



<a name="mycompany-simple-qux"></a>
### mycompany.simple.Qux

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>qux</td>
<td>string</td>
<td><pre>
json_name: qux
go_name: Qux</pre></td>
</tr>
</table>



<a name="mycompany-simple-someactivity2request"></a>
### mycompany.simple.SomeActivity2Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someactivity3request"></a>
### mycompany.simple.SomeActivity3Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someactivity3response"></a>
### mycompany.simple.SomeActivity3Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somedeprecatedmessage"></a>
### mycompany.simple.SomeDeprecatedMessage



<a name="mycompany-simple-somequery1response"></a>
### mycompany.simple.SomeQuery1Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somequery2request"></a>
### mycompany.simple.SomeQuery2Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somequery2response"></a>
### mycompany.simple.SomeQuery2Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somesignal2request"></a>
### mycompany.simple.SomeSignal2Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somesignal3request"></a>
### mycompany.simple.SomeSignal3Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-somesignal3response"></a>
### mycompany.simple.SomeSignal3Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someupdate1request"></a>
### mycompany.simple.SomeUpdate1Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someupdate1response"></a>
### mycompany.simple.SomeUpdate1Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someworkflow1request"></a>
### mycompany.simple.SomeWorkflow1Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someworkflow1response"></a>
### mycompany.simple.SomeWorkflow1Response

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>response_val</td>
<td>string</td>
<td><pre>
json_name: responseVal
go_name: ResponseVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-someworkflow3request"></a>
### mycompany.simple.SomeWorkflow3Request

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>request_val</td>
<td>string</td>
<td><pre>
json_name: requestVal
go_name: RequestVal</pre></td>
</tr>
</table>



<a name="mycompany-simple-whatrequest"></a>
### mycompany.simple.WhatRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>blah</td>
<td>string</td>
<td><pre>
json_name: blah
go_name: Blah</pre></td>
</tr>
</table>



<a name="test-option-v1"></a>
# test.option.v1

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



<a name="test-xnserr-v1"></a>
# test.xnserr.v1

<a name="test-xnserr-v1-services"></a>
## Services

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
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
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
<td><a href="#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr><tr>
<td>start_workflow_options</td>
<td><a href="#temporal-xns-v1-startworkflowoptions">temporal.xns.v1.StartWorkflowOptions</a></td>
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
<td><a href="#temporal-xns-v1-retrypolicy">temporal.xns.v1.RetryPolicy</a></td>
<td><pre>
json_name: retryPolicy
go_name: RetryPolicy</pre></td>
</tr><tr>
<td>sleep</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr><tr>
<td>start_workflow_options</td>
<td><a href="#temporal-xns-v1-startworkflowoptions">temporal.xns.v1.StartWorkflowOptions</a></td>
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
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: sleep
go_name: Sleep</pre></td>
</tr>
</table>




<a name="google-protobuf"></a>
# google.protobuf

<a name="google-protobuf-messages"></a>
## Messages

<a name="google-protobuf-any"></a>
### google.protobuf.Any

<pre>
`Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }
    // or ...
    if (any.isSameTypeAs(Foo.getDefaultInstance())) {
      foo = any.unpack(Foo.getDefaultInstance());
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".

JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>type_url</td>
<td>string</td>
<td><pre>
A URL/resource name that uniquely identifies the type of the serialized
protocol buffer message. This string must contain at least
one "/" character. The last segment of the URL's path must represent
the fully qualified name of the type (as in
`path/google.protobuf.Duration`). The name should be in a canonical form
(e.g., leading "." is not accepted).

In practice, teams usually precompile into the binary all types that they
expect it to use in the context of Any. However, for URLs which use the
scheme `http`, `https`, or no scheme, one can optionally set up a type
server that maps type URLs to message definitions as follows:

* If no scheme is provided, `https` is assumed.
* An HTTP GET on the URL must yield a [google.protobuf.Type][]
  value in binary format, or produce an error.
* Applications are allowed to cache lookup results based on the
  URL, or have them precompiled into a binary to avoid any
  lookup. Therefore, binary compatibility needs to be preserved
  on changes to types. (Use versioned type names to manage
  breaking changes.)

Note: this functionality is not currently available in the official
protobuf release, and it is not used for type URLs beginning with
type.googleapis.com. As of May 2023, there are no widely used type server
implementations and no plans to implement one.

Schemes other than `http`, `https` (or the empty scheme) might be
used with implementation specific semantics.<br>

json_name: typeUrl
go_name: TypeUrl</pre></td>
</tr><tr>
<td>value</td>
<td>bytes</td>
<td><pre>
Must be a valid serialized protocol buffer of the above specified type.<br>

json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="google-protobuf-duration"></a>
### google.protobuf.Duration

<pre>
A Duration represents a signed, fixed-length span of time represented
as a count of seconds and fractions of seconds at nanosecond
resolution. It is independent of any calendar and concepts like "day"
or "month". It is related to Timestamp in that the difference between
two Timestamp values is a Duration and it can be added or subtracted
from a Timestamp. Range is approximately +-10,000 years.

# Examples

Example 1: Compute Duration from two Timestamps in pseudo code.

    Timestamp start = ...;
    Timestamp end = ...;
    Duration duration = ...;

    duration.seconds = end.seconds - start.seconds;
    duration.nanos = end.nanos - start.nanos;

    if (duration.seconds < 0 && duration.nanos > 0) {
      duration.seconds += 1;
      duration.nanos -= 1000000000;
    } else if (duration.seconds > 0 && duration.nanos < 0) {
      duration.seconds -= 1;
      duration.nanos += 1000000000;
    }

Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.

    Timestamp start = ...;
    Duration duration = ...;
    Timestamp end = ...;

    end.seconds = start.seconds + duration.seconds;
    end.nanos = start.nanos + duration.nanos;

    if (end.nanos < 0) {
      end.seconds -= 1;
      end.nanos += 1000000000;
    } else if (end.nanos >= 1000000000) {
      end.seconds += 1;
      end.nanos -= 1000000000;
    }

Example 3: Compute Duration from datetime.timedelta in Python.

    td = datetime.timedelta(days=3, minutes=10)
    duration = Duration()
    duration.FromTimedelta(td)

# JSON Mapping

In JSON format, the Duration type is encoded as a string rather than an
object, where the string ends in the suffix "s" (indicating seconds) and
is preceded by the number of seconds, with nanoseconds expressed as
fractional seconds. For example, 3 seconds with 0 nanoseconds should be
encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
be expressed in JSON format as "3.000000001s", and 3 seconds and 1
microsecond should be expressed in JSON format as "3.000001s".
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>nanos</td>
<td>int32</td>
<td><pre>
Signed fractions of a second at nanosecond resolution of the span
of time. Durations less than one second are represented with a 0
`seconds` field and a positive or negative `nanos` field. For durations
of one second or more, a non-zero value for the `nanos` field must be
of the same sign as the `seconds` field. Must be from -999,999,999
to +999,999,999 inclusive.<br>

json_name: nanos
go_name: Nanos</pre></td>
</tr><tr>
<td>seconds</td>
<td>int64</td>
<td><pre>
Signed seconds of the span of time. Must be from -315,576,000,000
to +315,576,000,000 inclusive. Note: these bounds are computed from:
60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years<br>

json_name: seconds
go_name: Seconds</pre></td>
</tr>
</table>



<a name="google-protobuf-listvalue"></a>
### google.protobuf.ListValue

<pre>
`ListValue` is a wrapper around a repeated field of values.

The JSON representation for `ListValue` is JSON array.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>values</td>
<td><a href="#google-protobuf-value">google.protobuf.Value</a></td>
<td><pre>
Repeated field of dynamically typed values.<br>

json_name: values
go_name: Values</pre></td>
</tr>
</table>



<a name="google-protobuf-nullvalue"></a>
### google.protobuf.NullValue

<pre>
`NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

The JSON representation for `NullValue` is JSON `null`.
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>NULL_VALUE</td>
<td><pre>
Null value.
</pre></td>
</tr>
</table>

<a name="google-protobuf-struct"></a>
### google.protobuf.Struct

<pre>
`Struct` represents a structured data value, consisting of fields
which map to dynamically typed values. In some languages, `Struct`
might be supported by a native representation. For example, in
scripting languages like JS a struct is represented as an
object. The details of that representation are described together
with the proto support for the language.

The JSON representation for `Struct` is JSON object.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>fields</td>
<td><a href="#google-protobuf-struct-fieldsentry">google.protobuf.Struct.FieldsEntry</a></td>
<td><pre>
Unordered map of dynamically typed values.<br>

json_name: fields
go_name: Fields</pre></td>
</tr>
</table>



<a name="google-protobuf-struct-fieldsentry"></a>
### google.protobuf.Struct.FieldsEntry

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>key</td>
<td>string</td>
<td><pre>
json_name: key
go_name: Key</pre></td>
</tr><tr>
<td>value</td>
<td><a href="#google-protobuf-value">google.protobuf.Value</a></td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="google-protobuf-timestamp"></a>
### google.protobuf.Timestamp

<pre>
A Timestamp represents a point in time independent of any time zone or local
calendar, encoded as a count of seconds and fractions of seconds at
nanosecond resolution. The count is relative to an epoch at UTC midnight on
January 1, 1970, in the proleptic Gregorian calendar which extends the
Gregorian calendar backwards to year one.

All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
second table is needed for interpretation, using a [24-hour linear
smear](https://developers.google.com/time/smear).

The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
restricting to that range, we ensure that we can convert to and from [RFC
3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

# Examples

Example 1: Compute Timestamp from POSIX `time()`.

    Timestamp timestamp;
    timestamp.set_seconds(time(NULL));
    timestamp.set_nanos(0);

Example 2: Compute Timestamp from POSIX `gettimeofday()`.

    struct timeval tv;
    gettimeofday(&tv, NULL);

    Timestamp timestamp;
    timestamp.set_seconds(tv.tv_sec);
    timestamp.set_nanos(tv.tv_usec * 1000);

Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

    FILETIME ft;
    GetSystemTimeAsFileTime(&ft);
    UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

    // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
    // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
    Timestamp timestamp;
    timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
    timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

    long millis = System.currentTimeMillis();

    Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
        .setNanos((int) ((millis % 1000) * 1000000)).build();

Example 5: Compute Timestamp from Java `Instant.now()`.

    Instant now = Instant.now();

    Timestamp timestamp =
        Timestamp.newBuilder().setSeconds(now.getEpochSecond())
            .setNanos(now.getNano()).build();

Example 6: Compute Timestamp from current time in Python.

    timestamp = Timestamp()
    timestamp.GetCurrentTime()

# JSON Mapping

In JSON format, the Timestamp type is encoded as a string in the
[RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
where {year} is always expressed using four digits while {month}, {day},
{hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
is required. A proto3 JSON serializer should always use UTC (as indicated by
"Z") when printing the Timestamp type and a proto3 JSON parser should be
able to accept both UTC and other timezones (as indicated by an offset).

For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
01:30 UTC on January 15, 2017.

In JavaScript, one can convert a Date object to this format using the
standard
[toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
method. In Python, a standard `datetime.datetime` object can be converted
to this format using
[`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
the Joda Time's [`ISODateTimeFormat.dateTime()`](
http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()
) to obtain a formatter capable of generating timestamps in this format.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>nanos</td>
<td>int32</td>
<td><pre>
Non-negative fractions of a second at nanosecond resolution. Negative
second values with fractions must still have non-negative nanos values
that count forward in time. Must be from 0 to 999,999,999
inclusive.<br>

json_name: nanos
go_name: Nanos</pre></td>
</tr><tr>
<td>seconds</td>
<td>int64</td>
<td><pre>
Represents seconds of UTC time since Unix epoch
1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
9999-12-31T23:59:59Z inclusive.<br>

json_name: seconds
go_name: Seconds</pre></td>
</tr>
</table>



<a name="google-protobuf-value"></a>
### google.protobuf.Value

<pre>
`Value` represents a dynamically typed value which can be either
null, a number, a string, a boolean, a recursive struct value, or a
list of values. A producer of value is expected to set one of these
variants. Absence of any variant indicates an error.

The JSON representation for `Value` is JSON value.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>bool_value</td>
<td>bool</td>
<td><pre>
Represents a boolean value.<br>

json_name: boolValue
go_name: BoolValue</pre></td>
</tr><tr>
<td>list_value</td>
<td><a href="#google-protobuf-listvalue">google.protobuf.ListValue</a></td>
<td><pre>
Represents a repeated `Value`.<br>

json_name: listValue
go_name: ListValue</pre></td>
</tr><tr>
<td>null_value</td>
<td><a href="#google-protobuf-nullvalue">google.protobuf.NullValue</a></td>
<td><pre>
Represents a null value.<br>

json_name: nullValue
go_name: NullValue</pre></td>
</tr><tr>
<td>number_value</td>
<td>double</td>
<td><pre>
Represents a double value.<br>

json_name: numberValue
go_name: NumberValue</pre></td>
</tr><tr>
<td>string_value</td>
<td>string</td>
<td><pre>
Represents a string value.<br>

json_name: stringValue
go_name: StringValue</pre></td>
</tr><tr>
<td>struct_value</td>
<td><a href="#google-protobuf-struct">google.protobuf.Struct</a></td>
<td><pre>
Represents a structured value.<br>

json_name: structValue
go_name: StructValue</pre></td>
</tr>
</table>




<a name="mycompany-simple-common-v1"></a>
# mycompany.simple.common.v1

<a name="mycompany-simple-common-v1-messages"></a>
## Messages

<a name="mycompany-simple-common-v1-example"></a>
### mycompany.simple.common.v1.Example

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>EXAMPLE_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>EXAMPLE_FOO</td>
<td></td>
</tr>
</table>

<a name="mycompany-simple-common-v1-paginatedrequest"></a>
### mycompany.simple.common.v1.PaginatedRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cursor</td>
<td>bytes</td>
<td><pre>
json_name: cursor
go_name: Cursor</pre></td>
</tr><tr>
<td>limit</td>
<td>uint32</td>
<td><pre>
json_name: limit
go_name: Limit</pre></td>
</tr>
</table>



<a name="mycompany-simple-common-v1-paginatedresponse"></a>
### mycompany.simple.common.v1.PaginatedResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="#google-protobuf-any">google.protobuf.Any</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>next_cursor</td>
<td>bytes</td>
<td><pre>
json_name: nextCursor
go_name: NextCursor</pre></td>
</tr>
</table>




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
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
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
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
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
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
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
<td><a href="#google-protobuf-struct">google.protobuf.Struct</a></td>
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
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: runTimeout
go_name: RunTimeout</pre></td>
</tr><tr>
<td>search_attirbutes</td>
<td><a href="#google-protobuf-struct">google.protobuf.Struct</a></td>
<td><pre>
json_name: searchAttirbutes
go_name: SearchAttirbutes</pre></td>
</tr><tr>
<td>start_delay</td>
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
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
<td><a href="#google-protobuf-duration">google.protobuf.Duration</a></td>
<td><pre>
json_name: taskTimeout
go_name: TaskTimeout</pre></td>
</tr>
</table>

