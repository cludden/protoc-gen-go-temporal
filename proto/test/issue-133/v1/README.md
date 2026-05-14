

<a name="issue_133-v1"></a>
# issue_133.v1

<a name="issue_133-v1-services"></a>
## Services

<a name="issue_133-v1-createuserworkflowservice"></a>
## issue_133.v1.CreateUserWorkflowService   

<a name="issue_133-v1-createuserworkflowservice-activities"></a>
### Activities

---
<a name="createuserinapiactivity-activity"></a>
### CreateUserInAPIActivity



**Input:** [issue_133.v1.CreateUserInAPIRequest](#issue_133-v1-createuserinapirequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>email</td>
<td>string</td>
<td><pre>
json_name: email
go_name: Email</pre></td>
</tr>
</table>

**Output:** [issue_133.v1.CreateUserInAPIResponse](#issue_133-v1-createuserinapiresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>success</td>
<td>bool</td>
<td><pre>
json_name: success
go_name: Success</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="issue_133-v1-messages"></a>
## Messages

<a name="issue_133-v1-createuserinapirequest"></a>
### issue_133.v1.CreateUserInAPIRequest

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>email</td>
<td>string</td>
<td><pre>
json_name: email
go_name: Email</pre></td>
</tr>
</table>



<a name="issue_133-v1-createuserinapiresponse"></a>
### issue_133.v1.CreateUserInAPIResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>success</td>
<td>bool</td>
<td><pre>
json_name: success
go_name: Success</pre></td>
</tr>
</table>

