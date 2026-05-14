

<a name="temporal-v1"></a>
# temporal.v1

<a name="temporal-v1-messages"></a>
## Messages

<a name="temporal-v1-retrypolicy"></a>
### temporal.v1.RetryPolicy

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
<td><a href="../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
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
<td><a href="../../google/protobuf/README.md#google-protobuf-duration">google.protobuf.Duration</a></td>
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

