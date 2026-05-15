

<a name="mycompany-simple-common-v1"></a>
# mycompany.simple.common.v1

## Table of Contents
- Messages
  - [mycompany.simple.common.v1.Example](#mycompany-simple-common-v1-example)
  - [mycompany.simple.common.v1.PaginatedRequest](#mycompany-simple-common-v1-paginatedrequest)
  - [mycompany.simple.common.v1.PaginatedResponse](#mycompany-simple-common-v1-paginatedresponse)

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
<td><a href="../../../../google/protobuf/README.md#google-protobuf-any">google.protobuf.Any</a></td>
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

