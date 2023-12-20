# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [patch/go.proto](#patch_go-proto)
    - [LintOptions](#go-LintOptions)
    - [Options](#go-Options)
  
    - [File-level Extensions](#patch_go-proto-extensions)
    - [File-level Extensions](#patch_go-proto-extensions)
    - [File-level Extensions](#patch_go-proto-extensions)
    - [File-level Extensions](#patch_go-proto-extensions)
    - [File-level Extensions](#patch_go-proto-extensions)
    - [File-level Extensions](#patch_go-proto-extensions)
  
- [Scalar Value Types](#scalar-value-types)



<a name="patch_go-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## patch/go.proto



<a name="go-LintOptions"></a>

### LintOptions
LintOptions represent options for linting a generated Go file.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| all | [bool](#bool) | optional | Set all to true if all generated Go symbols should be linted. This option affects generated structs, struct fields, enum types, and value constants. |
| messages | [bool](#bool) | optional | Set messages to true if message names should be linted. This does not affect message fields. |
| fields | [bool](#bool) | optional | Set messages to true if message field names should be linted. This does not affect message fields. |
| enums | [bool](#bool) | optional | Set enums to true if generated enum names should be linted. This does not affect enum values. |
| values | [bool](#bool) | optional | Set values to true if generated enum value constants should be linted. |
| extensions | [bool](#bool) | optional | Set extensions to true if generated extension names should be linted. |
| initialisms | [string](#string) | repeated | The initialisms option lets you specify strings that should not be generated as mixed-case, Examples: ID, URL, HTTP, etc. |






<a name="go-Options"></a>

### Options
Options represent Go-specific options for Protobuf messages, fields, oneofs, enums, or enum values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) | optional | The name option renames the generated Go identifier and related identifiers. For a message, this renames the generated Go struct and nested messages or enums, if any. For a message field, this renames the generated Go struct field and getter method. For a oneof field, this renames the generated Go struct field, getter method, interface type, and wrapper types. For an enum, this renames the generated Go type. For an enum value, this renames the generated Go const. |
| embed | [bool](#bool) | optional | The embed option indicates the field should be embedded in the generated Go struct. Only message types can be embedded. Oneof fields cannot be embedded. See https://golang.org/ref/spec#Struct_types. |
| type | [string](#string) | optional | The type option changes the generated field type. All generated code assumes that this type is castable to the protocol buffer field type. |
| getter | [string](#string) | optional | The getter option renames the generated getter method (default: Get&lt;Field&gt;) so a custom getter can be implemented in its place.

TODO: implement this |
| tags | [string](#string) | optional | The tags option specifies additional struct tags which are appended a generated Go struct field. This option may be specified on a message field or a oneof field. The value should omit the enclosing backticks. |
| stringer | [string](#string) | optional | The stringer option renames a generated String() method (if any) so a custom String() method can be implemented in its place.

TODO: implement for messages |
| stringer_name | [string](#string) | optional | The stringer_name option is a deprecated alias for stringer. It will be removed in a future version of this package. |





 

 


<a name="patch_go-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| enum | Options | .google.protobuf.EnumOptions | 7001 |  |
| value | Options | .google.protobuf.EnumValueOptions | 7001 |  |
| field | Options | .google.protobuf.FieldOptions | 7001 |  |
| lint | LintOptions | .google.protobuf.FileOptions | 7001 |  |
| message | Options | .google.protobuf.MessageOptions | 7001 |  |
| oneof | Options | .google.protobuf.OneofOptions | 7001 |  |

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

