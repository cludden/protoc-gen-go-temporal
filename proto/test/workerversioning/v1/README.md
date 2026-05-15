

<a name="test-workerversioning-v1"></a>
# test.workerversioning.v1

## Table of Contents
- [test.workerversioning.v1.Example](#test-workerversioning-v1-example)
  - [Workflows](#test-workerversioning-v1-example-workflows)
    - [test.workerversioning.v1.Example.Bar](#test-workerversioning-v1-example-bar-workflow)
    - [test.workerversioning.v1.Example.Baz](#test-workerversioning-v1-example-baz-workflow)
    - [test.workerversioning.v1.Example.Foo](#test-workerversioning-v1-example-foo-workflow)
    - [test.workerversioning.v1.Example.Qux](#test-workerversioning-v1-example-qux-workflow)
- Messages
  - [test.workerversioning.v1.BarInput](#test-workerversioning-v1-barinput)
  - [test.workerversioning.v1.BarOutput](#test-workerversioning-v1-baroutput)
  - [test.workerversioning.v1.BazInput](#test-workerversioning-v1-bazinput)
  - [test.workerversioning.v1.BazOutput](#test-workerversioning-v1-bazoutput)
  - [test.workerversioning.v1.FooInput](#test-workerversioning-v1-fooinput)
  - [test.workerversioning.v1.FooOutput](#test-workerversioning-v1-foooutput)
  - [test.workerversioning.v1.QuxInput](#test-workerversioning-v1-quxinput)
  - [test.workerversioning.v1.QuxOutput](#test-workerversioning-v1-quxoutput)

<a name="test-workerversioning-v1-services"></a>
## Services

<a name="test-workerversioning-v1-example"></a>
## test.workerversioning.v1.Example

<a name="test-workerversioning-v1-example-workflows"></a>
### Workflows

---
<a name="test-workerversioning-v1-example-bar-workflow"></a>
### test.workerversioning.v1.Example.Bar

**Input:** [test.workerversioning.v1.BarInput](#test-workerversioning-v1-barinput)



**Output:** [test.workerversioning.v1.BarOutput](#test-workerversioning-v1-baroutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>workflowversioning.v1.Bar/${! uuid_v4()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="test-workerversioning-v1-example-baz-workflow"></a>
### test.workerversioning.v1.Example.Baz

**Input:** [test.workerversioning.v1.BazInput](#test-workerversioning-v1-bazinput)



**Output:** [test.workerversioning.v1.BazOutput](#test-workerversioning-v1-bazoutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>workflowversioning.v1.Baz/${! uuid_v4()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="test-workerversioning-v1-example-foo-workflow"></a>
### test.workerversioning.v1.Example.Foo

**Input:** [test.workerversioning.v1.FooInput](#test-workerversioning-v1-fooinput)



**Output:** [test.workerversioning.v1.FooOutput](#test-workerversioning-v1-foooutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>workflowversioning.v1.Foo/${! uuid_v4()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

---
<a name="test-workerversioning-v1-example-qux-workflow"></a>
### test.workerversioning.v1.Example.Qux

**Input:** [test.workerversioning.v1.QuxInput](#test-workerversioning-v1-quxinput)



**Output:** [test.workerversioning.v1.QuxOutput](#test-workerversioning-v1-quxoutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>workflowversioning.v1.Qux/${! uuid_v4()}</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>     

<a name="test-workerversioning-v1-messages"></a>
## Messages

<a name="test-workerversioning-v1-barinput"></a>
### test.workerversioning.v1.BarInput



<a name="test-workerversioning-v1-baroutput"></a>
### test.workerversioning.v1.BarOutput



<a name="test-workerversioning-v1-bazinput"></a>
### test.workerversioning.v1.BazInput



<a name="test-workerversioning-v1-bazoutput"></a>
### test.workerversioning.v1.BazOutput



<a name="test-workerversioning-v1-fooinput"></a>
### test.workerversioning.v1.FooInput



<a name="test-workerversioning-v1-foooutput"></a>
### test.workerversioning.v1.FooOutput



<a name="test-workerversioning-v1-quxinput"></a>
### test.workerversioning.v1.QuxInput



<a name="test-workerversioning-v1-quxoutput"></a>
### test.workerversioning.v1.QuxOutput

