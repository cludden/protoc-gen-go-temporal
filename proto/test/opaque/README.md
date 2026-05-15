

<a name="test-opaque"></a>
# test.opaque

## Table of Contents
- [test.opaque.Hybrid](#test-opaque-hybrid)
  - [Workflows](#test-opaque-hybrid-workflows)
    - [test.opaque.Hybrid.PutHybridExample](#test-opaque-hybrid-puthybridexample-workflow)
  - [Signals](#test-opaque-hybrid-signals)
    - [test.opaque.Hybrid.SignalHybrid](#test-opaque-hybrid-signalhybrid-signal)
- [test.opaque.Opaque](#test-opaque-opaque)
  - [Workflows](#test-opaque-opaque-workflows)
    - [test.opaque.Opaque.PutOpaqueExample](#test-opaque-opaque-putopaqueexample-workflow)
  - [Signals](#test-opaque-opaque-signals)
    - [test.opaque.Opaque.SignalOpaque](#test-opaque-opaque-signalopaque-signal)
- [test.opaque.Open](#test-opaque-open)
  - [Workflows](#test-opaque-open-workflows)
    - [test.opaque.Open.PutOpenExample](#test-opaque-open-putopenexample-workflow)
  - [Signals](#test-opaque-open-signals)
    - [test.opaque.Open.SignalOpen](#test-opaque-open-signalopen-signal)
- [test.opaque.Optional](#test-opaque-optional)
  - [Workflows](#test-opaque-optional-workflows)
    - [test.opaque.Optional.PutOptionalExample](#test-opaque-optional-putoptionalexample-workflow)
  - [Signals](#test-opaque-optional-signals)
    - [test.opaque.Optional.SignalOptional](#test-opaque-optional-signaloptional-signal)
- Messages
  - [test.opaque.Address](#test-opaque-address)
  - [test.opaque.HybridExample](#test-opaque-hybridexample)
  - [test.opaque.HybridExample.ExtraEntry](#test-opaque-hybridexample-extraentry)
  - [test.opaque.OpaqueExample](#test-opaque-opaqueexample)
  - [test.opaque.OpaqueExample.ExtraEntry](#test-opaque-opaqueexample-extraentry)
  - [test.opaque.OpenExample](#test-opaque-openexample)
  - [test.opaque.OpenExample.ExtraEntry](#test-opaque-openexample-extraentry)
  - [test.opaque.OptionalExample](#test-opaque-optionalexample)
  - [test.opaque.OptionalExample.ExtraEntry](#test-opaque-optionalexample-extraentry)
  - [test.opaque.Status](#test-opaque-status)

<a name="test-opaque-services"></a>
## Services

<a name="test-opaque-hybrid"></a>
## test.opaque.Hybrid

<a name="test-opaque-hybrid-workflows"></a>
### Workflows

---
<a name="test-opaque-hybrid-puthybridexample-workflow"></a>
### test.opaque.Hybrid.PutHybridExample

**Input:** [test.opaque.HybridExample](#test-opaque-hybridexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.HybridExample](#test-opaque-hybridexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-hybrid-signalhybrid-signal">test.opaque.Hybrid.SignalHybrid</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-hybrid-signals"></a>
### Signals

---
<a name="test-opaque-hybrid-signalhybrid-signal"></a>
### test.opaque.Hybrid.SignalHybrid



**Input:** [test.opaque.HybridExample](#test-opaque-hybridexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-opaque"></a>
## test.opaque.Opaque

<a name="test-opaque-opaque-workflows"></a>
### Workflows

---
<a name="test-opaque-opaque-putopaqueexample-workflow"></a>
### test.opaque.Opaque.PutOpaqueExample

**Input:** [test.opaque.OpaqueExample](#test-opaque-opaqueexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.OpaqueExample](#test-opaque-opaqueexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-opaque-signalopaque-signal">test.opaque.Opaque.SignalOpaque</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-opaque-signals"></a>
### Signals

---
<a name="test-opaque-opaque-signalopaque-signal"></a>
### test.opaque.Opaque.SignalOpaque



**Input:** [test.opaque.OpaqueExample](#test-opaque-opaqueexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-open"></a>
## test.opaque.Open

<a name="test-opaque-open-workflows"></a>
### Workflows

---
<a name="test-opaque-open-putopenexample-workflow"></a>
### test.opaque.Open.PutOpenExample

**Input:** [test.opaque.OpenExample](#test-opaque-openexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.OpenExample](#test-opaque-openexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-open-signalopen-signal">test.opaque.Open.SignalOpen</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-open-signals"></a>
### Signals

---
<a name="test-opaque-open-signalopen-signal"></a>
### test.opaque.Open.SignalOpen



**Input:** [test.opaque.OpenExample](#test-opaque-openexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-optional"></a>
## test.opaque.Optional

<a name="test-opaque-optional-workflows"></a>
### Workflows

---
<a name="test-opaque-optional-putoptionalexample-workflow"></a>
### test.opaque.Optional.PutOptionalExample

**Input:** [test.opaque.OptionalExample](#test-opaque-optionalexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Output:** [test.opaque.OptionalExample](#test-opaque-optionalexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#test-opaque-optional-signaloptional-signal">test.opaque.Optional.SignalOptional</a></td><td>true</td></tr>
</table>   

<a name="test-opaque-optional-signals"></a>
### Signals

---
<a name="test-opaque-optional-signaloptional-signal"></a>
### test.opaque.Optional.SignalOptional



**Input:** [test.opaque.OptionalExample](#test-opaque-optionalexample)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>   

<a name="test-opaque-messages"></a>
## Messages

<a name="test-opaque-address"></a>
### test.opaque.Address

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>city</td>
<td>string</td>
<td><pre>
json_name: city
go_name: City</pre></td>
</tr><tr>
<td>state</td>
<td>string</td>
<td><pre>
json_name: state
go_name: State</pre></td>
</tr><tr>
<td>street</td>
<td>string</td>
<td><pre>
json_name: street
go_name: Street</pre></td>
</tr><tr>
<td>zip</td>
<td>string</td>
<td><pre>
json_name: zip
go_name: Zip</pre></td>
</tr>
</table>



<a name="test-opaque-hybridexample"></a>
### test.opaque.HybridExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-hybridexample-extraentry">test.opaque.HybridExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-hybridexample-extraentry"></a>
### test.opaque.HybridExample.ExtraEntry

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
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-opaqueexample"></a>
### test.opaque.OpaqueExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-opaqueexample-extraentry">test.opaque.OpaqueExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-opaqueexample-extraentry"></a>
### test.opaque.OpaqueExample.ExtraEntry

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
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-openexample"></a>
### test.opaque.OpenExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-openexample-extraentry">test.opaque.OpenExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-openexample-extraentry"></a>
### test.opaque.OpenExample.ExtraEntry

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
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-optionalexample"></a>
### test.opaque.OptionalExample

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: address
go_name: Address</pre></td>
</tr><tr>
<td>age</td>
<td>int32</td>
<td><pre>
json_name: age
go_name: Age</pre></td>
</tr><tr>
<td>ages</td>
<td>int32</td>
<td><pre>
json_name: ages
go_name: Ages</pre></td>
</tr><tr>
<td>connection_id</td>
<td>sint32</td>
<td><pre>
json_name: connectionId
go_name: ConnectionId</pre></td>
</tr><tr>
<td>connection_ids</td>
<td>sint32</td>
<td><pre>
json_name: connectionIds
go_name: ConnectionIds</pre></td>
</tr><tr>
<td>data</td>
<td>bytes</td>
<td><pre>
json_name: data
go_name: Data</pre></td>
</tr><tr>
<td>datas</td>
<td>bytes</td>
<td><pre>
json_name: datas
go_name: Datas</pre></td>
</tr><tr>
<td>emails</td>
<td>string</td>
<td><pre>
json_name: emails
go_name: Emails</pre></td>
</tr><tr>
<td>extra</td>
<td><a href="#test-opaque-optionalexample-extraentry">test.opaque.OptionalExample.ExtraEntry</a></td>
<td><pre>
json_name: extra
go_name: Extra</pre></td>
</tr><tr>
<td>fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLength
go_name: FixedLength</pre></td>
</tr><tr>
<td>fixed_lengths</td>
<td>fixed64</td>
<td><pre>
json_name: fixedLengths
go_name: FixedLengths</pre></td>
</tr><tr>
<td>fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSize
go_name: FixedSize</pre></td>
</tr><tr>
<td>fixed_sizes</td>
<td>fixed32</td>
<td><pre>
json_name: fixedSizes
go_name: FixedSizes</pre></td>
</tr><tr>
<td>id</td>
<td>int64</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>ids</td>
<td>int64</td>
<td><pre>
json_name: ids
go_name: Ids</pre></td>
</tr><tr>
<td>is_active</td>
<td>bool</td>
<td><pre>
json_name: isActive
go_name: IsActive</pre></td>
</tr><tr>
<td>is_actives</td>
<td>bool</td>
<td><pre>
json_name: isActives
go_name: IsActives</pre></td>
</tr><tr>
<td>length</td>
<td>uint64</td>
<td><pre>
json_name: length
go_name: Length</pre></td>
</tr><tr>
<td>lengths</td>
<td>uint64</td>
<td><pre>
json_name: lengths
go_name: Lengths</pre></td>
</tr><tr>
<td>name</td>
<td>string</td>
<td><pre>
json_name: name
go_name: Name</pre></td>
</tr><tr>
<td>oneof_address</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: oneofAddress
go_name: OneofAddress</pre></td>
</tr><tr>
<td>oneof_age</td>
<td>int32</td>
<td><pre>
json_name: oneofAge
go_name: OneofAge</pre></td>
</tr><tr>
<td>oneof_connection_id</td>
<td>sint32</td>
<td><pre>
json_name: oneofConnectionId
go_name: OneofConnectionId</pre></td>
</tr><tr>
<td>oneof_data</td>
<td>bytes</td>
<td><pre>
json_name: oneofData
go_name: OneofData</pre></td>
</tr><tr>
<td>oneof_fixed_length</td>
<td>fixed64</td>
<td><pre>
json_name: oneofFixedLength
go_name: OneofFixedLength</pre></td>
</tr><tr>
<td>oneof_fixed_size</td>
<td>fixed32</td>
<td><pre>
json_name: oneofFixedSize
go_name: OneofFixedSize</pre></td>
</tr><tr>
<td>oneof_id</td>
<td>int64</td>
<td><pre>
json_name: oneofId
go_name: OneofId</pre></td>
</tr><tr>
<td>oneof_is_active</td>
<td>bool</td>
<td><pre>
json_name: oneofIsActive
go_name: OneofIsActive</pre></td>
</tr><tr>
<td>oneof_length</td>
<td>uint64</td>
<td><pre>
json_name: oneofLength
go_name: OneofLength</pre></td>
</tr><tr>
<td>oneof_name</td>
<td>string</td>
<td><pre>
json_name: oneofName
go_name: OneofName</pre></td>
</tr><tr>
<td>oneof_ratio</td>
<td>float</td>
<td><pre>
json_name: oneofRatio
go_name: OneofRatio</pre></td>
</tr><tr>
<td>oneof_score</td>
<td>double</td>
<td><pre>
json_name: oneofScore
go_name: OneofScore</pre></td>
</tr><tr>
<td>oneof_session_id</td>
<td>sint64</td>
<td><pre>
json_name: oneofSessionId
go_name: OneofSessionId</pre></td>
</tr><tr>
<td>oneof_sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: oneofSfixedLength
go_name: OneofSfixedLength</pre></td>
</tr><tr>
<td>oneof_sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: oneofSfixedSize
go_name: OneofSfixedSize</pre></td>
</tr><tr>
<td>oneof_size</td>
<td>uint32</td>
<td><pre>
json_name: oneofSize
go_name: OneofSize</pre></td>
</tr><tr>
<td>oneof_status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: oneofStatus
go_name: OneofStatus</pre></td>
</tr><tr>
<td>previous_addresses</td>
<td><a href="#test-opaque-address">test.opaque.Address</a></td>
<td><pre>
json_name: previousAddresses
go_name: PreviousAddresses</pre></td>
</tr><tr>
<td>ratio</td>
<td>float</td>
<td><pre>
json_name: ratio
go_name: Ratio</pre></td>
</tr><tr>
<td>ratios</td>
<td>float</td>
<td><pre>
json_name: ratios
go_name: Ratios</pre></td>
</tr><tr>
<td>score</td>
<td>double</td>
<td><pre>
json_name: score
go_name: Score</pre></td>
</tr><tr>
<td>scores</td>
<td>double</td>
<td><pre>
json_name: scores
go_name: Scores</pre></td>
</tr><tr>
<td>session_id</td>
<td>sint64</td>
<td><pre>
json_name: sessionId
go_name: SessionId</pre></td>
</tr><tr>
<td>session_ids</td>
<td>sint64</td>
<td><pre>
json_name: sessionIds
go_name: SessionIds</pre></td>
</tr><tr>
<td>sfixed_length</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLength
go_name: SfixedLength</pre></td>
</tr><tr>
<td>sfixed_lengths</td>
<td>sfixed64</td>
<td><pre>
json_name: sfixedLengths
go_name: SfixedLengths</pre></td>
</tr><tr>
<td>sfixed_size</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSize
go_name: SfixedSize</pre></td>
</tr><tr>
<td>sfixed_sizes</td>
<td>sfixed32</td>
<td><pre>
json_name: sfixedSizes
go_name: SfixedSizes</pre></td>
</tr><tr>
<td>size</td>
<td>uint32</td>
<td><pre>
json_name: size
go_name: Size</pre></td>
</tr><tr>
<td>sizes</td>
<td>uint32</td>
<td><pre>
json_name: sizes
go_name: Sizes</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>statuses</td>
<td><a href="#test-opaque-status">test.opaque.Status</a></td>
<td><pre>
json_name: statuses
go_name: Statuses</pre></td>
</tr>
</table>



<a name="test-opaque-optionalexample-extraentry"></a>
### test.opaque.OptionalExample.ExtraEntry

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
<td>string</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="test-opaque-status"></a>
### test.opaque.Status

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>STATUS_OK</td>
<td></td>
</tr><tr>
<td>STATUS_ERROR</td>
<td></td>
</tr>
</table>