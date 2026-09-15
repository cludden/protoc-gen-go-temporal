

<a name="temporal-api-common-v1"></a>
# temporal.api.common.v1

## Table of Contents
- Messages
  - [temporal.api.common.v1.Payload](#temporal-api-common-v1-payload)
  - [temporal.api.common.v1.Payload.ExternalPayloadDetails](#temporal-api-common-v1-payload-externalpayloaddetails)
  - [temporal.api.common.v1.Payload.MetadataEntry](#temporal-api-common-v1-payload-metadataentry)
  - [temporal.api.common.v1.Priority](#temporal-api-common-v1-priority)
  - [temporal.api.common.v1.SearchAttributes](#temporal-api-common-v1-searchattributes)
  - [temporal.api.common.v1.SearchAttributes.IndexedFieldsEntry](#temporal-api-common-v1-searchattributes-indexedfieldsentry)

<a name="temporal-api-common-v1-messages"></a>
## Messages

<a name="temporal-api-common-v1-payload"></a>
### temporal.api.common.v1.Payload

<pre>
Represents some binary (byte array) data (ex: activity input parameters or workflow result) with
metadata which describes this binary data (format, encoding, encryption, etc). Serialization
of the data may be user-defined.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>external_payloads</td>
<td><a href="#temporal-api-common-v1-payload-externalpayloaddetails">temporal.api.common.v1.Payload.ExternalPayloadDetails</a></td>
<td><pre>
Details about externally stored payloads associated with this payload.<br>

json_name: externalPayloads
go_name: ExternalPayloads</pre></td>
</tr><tr>
<td>metadata</td>
<td><a href="#temporal-api-common-v1-payload-metadataentry">temporal.api.common.v1.Payload.MetadataEntry</a></td>
<td><pre>
json_name: metadata
go_name: Metadata</pre></td>
</tr>
</table>



<a name="temporal-api-common-v1-payload-externalpayloaddetails"></a>
### temporal.api.common.v1.Payload.ExternalPayloadDetails

<pre>
Describes an externally stored object referenced by this payload.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>size_bytes</td>
<td>int64</td>
<td><pre>
Size in bytes of the externally stored payload<br>

json_name: sizeBytes
go_name: SizeBytes</pre></td>
</tr>
</table>



<a name="temporal-api-common-v1-payload-metadataentry"></a>
### temporal.api.common.v1.Payload.MetadataEntry

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
<td>bytes</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="temporal-api-common-v1-priority"></a>
### temporal.api.common.v1.Priority

<pre>
Priority contains metadata that controls relative ordering of task processing
when tasks are backed up in a queue. Initially, Priority will be used in
matching (workflow and activity) task queues. Later it may be used in history
task queues and in rate limiting decisions.

Priority is attached to workflows and activities. By default, activities
inherit Priority from the workflow that created them, but may override fields
when an activity is started or modified.

Despite being named "Priority", this message also contains fields that
control "fairness" mechanisms.

For all fields, the field not present or equal to zero/empty string means to
inherit the value from the calling workflow, or if there is no calling
workflow, then use the default value.

For all fields other than fairness_key, the zero value isn't meaningful so
there's no confusion between inherit/default and a meaningful value. For
fairness_key, the empty string will be interpreted as "inherit". This means
that if a workflow has a non-empty fairness key, you can't override the
fairness key of its activity to the empty string.

The overall semantics of Priority are:
1. First, consider "priority": higher priority (lower number) goes first.
2. Then, consider fairness: try to dispatch tasks for different fairness keys
   in proportion to their weight.

Applications may use any subset of mechanisms that are useful to them and
leave the other fields to use default values.

Not all queues in the system may support the "full" semantics of all priority
fields. (Currently only support in matching task queues is planned.)
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>fairness_key</td>
<td>string</td>
<td><pre>
Fairness key is a short string that's used as a key for a fairness
balancing mechanism. It may correspond to a tenant id, or to a fixed
string like "high" or "low". The default is the empty string.

The fairness mechanism attempts to dispatch tasks for a given key in
proportion to its weight. For example, using a thousand distinct tenant
ids, each with a weight of 1.0 (the default) will result in each tenant
getting a roughly equal share of task dispatch throughput.

(Note: this does not imply equal share of worker capacity! Fairness
decisions are made based on queue statistics, not
current worker load.)

As another example, using keys "high" and "low" with weight 9.0 and 1.0
respectively will prefer dispatching "high" tasks over "low" tasks at a
9:1 ratio, while allowing either key to use all worker capacity if the
other is not present.

All fairness mechanisms, including rate limits, are best-effort and
probabilistic. The results may not match what a "perfect" algorithm with
infinite resources would produce. The more unique keys are used, the less
accurate the results will be.

Fairness keys are limited to 64 bytes.<br>

json_name: fairnessKey
go_name: FairnessKey</pre></td>
</tr><tr>
<td>fairness_weight</td>
<td>float</td>
<td><pre>
Fairness weight for a task can come from multiple sources for
flexibility. From highest to lowest precedence:
1. Weights for a small set of keys can be overridden in task queue
   configuration with an API.
2. It can be attached to the workflow/activity in this field.
3. The default weight of 1.0 will be used.

Weight values are clamped to the range [0.001, 1000].<br>

json_name: fairnessWeight
go_name: FairnessWeight</pre></td>
</tr><tr>
<td>priority_key</td>
<td>int32</td>
<td><pre>
Priority key is a positive integer from 1 to n, where smaller integers
correspond to higher priorities (tasks run sooner). In general, tasks in
a queue should be processed in close to priority order, although small
deviations are possible.

The maximum priority value (minimum priority) is determined by server
configuration, and defaults to 5.

If priority is not present (or zero), then the effective priority will be
the default priority, which is calculated by (min+max)/2. With the
default max of 5, and min of 1, that comes out to 3.<br>

json_name: priorityKey
go_name: PriorityKey</pre></td>
</tr>
</table>



<a name="temporal-api-common-v1-searchattributes"></a>
### temporal.api.common.v1.SearchAttributes

<pre>
A user-defined set of *indexed* fields that are used/exposed when listing/searching workflows.
The payload is not serialized in a user-defined way.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>indexed_fields</td>
<td><a href="#temporal-api-common-v1-searchattributes-indexedfieldsentry">temporal.api.common.v1.SearchAttributes.IndexedFieldsEntry</a></td>
<td><pre>
json_name: indexedFields
go_name: IndexedFields</pre></td>
</tr>
</table>



<a name="temporal-api-common-v1-searchattributes-indexedfieldsentry"></a>
### temporal.api.common.v1.SearchAttributes.IndexedFieldsEntry

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
<td><a href="#temporal-api-common-v1-payload">temporal.api.common.v1.Payload</a></td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>

