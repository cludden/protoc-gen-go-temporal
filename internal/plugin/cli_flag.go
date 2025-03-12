package plugin

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/gofeaturespb"
)

const (
	CliPkg       = "github.com/urfave/cli/v2"
	ConvertPkg   = "github.com/cludden/protoc-gen-go-temporal/pkg/convert"
	ProtojsonPkg = "google.golang.org/protobuf/encoding/protojson"
)

var (
	castFuncs = map[protoreflect.Kind]struct {
		from *jen.Statement
		to   *jen.Statement
	}{
		protoreflect.Fixed32Kind:  {jen.Uint64(), jen.Uint32()},
		protoreflect.FloatKind:    {jen.Float64(), jen.Float32()},
		protoreflect.Int32Kind:    {jen.Int64(), jen.Int32()},
		protoreflect.Sfixed32Kind: {jen.Int64(), jen.Int32()},
		protoreflect.Sint32Kind:   {jen.Int64(), jen.Int32()},
		protoreflect.Uint32Kind:   {jen.Uint64(), jen.Uint32()},
	}
)

func (svc *Manifest) flagName(field *protogen.Field, prefix string) string {
	opts := proto.GetExtension(field.Desc.Options(), temporalv1.E_Field).(*temporalv1.FieldOptions)

	name := opts.GetCli().GetName()
	if name == "" {
		name = svc.caser.ToKebab(field.GoName)
	}
	if prefix != "" {
		name = fmt.Sprintf("%s-%s", svc.caser.ToKebab(prefix), name)
	}
	return name
}

func flagType(field *protogen.Field) (t string, additionalUsage string, err error) {
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		if field.Desc.IsList() {
			return "StringSlice", "", nil
		}
		return "Bool", "", nil

	case protoreflect.DoubleKind, protoreflect.FloatKind:
		if field.Desc.IsList() {
			return "Float64Slice", "", nil
		}
		return "Float64", "", nil

	case protoreflect.Fixed32Kind, protoreflect.Fixed64Kind, protoreflect.Uint32Kind, protoreflect.Uint64Kind:
		if field.Desc.IsList() {
			return "Uint64Slice", "", nil
		}
		return "Uint64", "", nil

	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind:
		if field.Desc.IsList() {
			return "Int64Slice", "", nil
		}
		return "Int64", "", nil

	case protoreflect.MessageKind:
		switch field.Desc.Message().FullName() {
		case "google.protobuf.Duration":
			additionalUsage = " (e.g. \"3.000000001s\")"
			if !field.Desc.IsList() {
				return "Duration", additionalUsage, nil
			}
		case "google.protobuf.Timestamp":
			additionalUsage = " (e.g. \"2017-01-15T01:30:15.01Z\")"
			if !field.Desc.IsList() {
				return "Timestamp", additionalUsage, nil
			}
		default:
			var b bytes.Buffer
			fmt.Fprint(&b, " (json-encoded: {")
			var fieldDocs []string
			for _, f := range field.Message.Fields {
				kind := f.Desc.Kind().String()
				if f.Message != nil {
					kind = string(f.Message.Desc.FullName())
				} else if f.Enum != nil {
					kind = string(f.Enum.Desc.FullName())
				}
				fieldDocs = append(fieldDocs, fmt.Sprintf("%s: <%s>", f.Desc.JSONName(), kind))
			}
			fmt.Fprint(&b, strings.Join(fieldDocs, ", "))
			fmt.Fprint(&b, "})")
			additionalUsage = b.String()
		}

		fallthrough

	case protoreflect.BytesKind, protoreflect.EnumKind, protoreflect.StringKind:
		if k := field.Desc.Kind(); k == protoreflect.BytesKind {
			additionalUsage = " (base64-encoded)"
		} else if k == protoreflect.EnumKind {
			var values []string
			for _, v := range field.Enum.Values {
				values = append(values, string(v.Desc.Name()))
			}
			additionalUsage = fmt.Sprintf(" (%s)", strings.Join(values, ", "))
		}
		if field.Desc.IsList() {
			return "StringSlice", additionalUsage, nil
		}
		return "String", additionalUsage, nil
	}
	return "", "", fmt.Errorf("unsupported field type: %s", field.Desc.Kind())
}

// genCliFlagForField generates a cli flag for a message field
func (svc *Manifest) genCliFlagForField(flags *jen.Group, field *protogen.Field, category string, prefix string) {
	name := svc.getFieldName(field)
	opts := proto.GetExtension(field.Desc.Options(), temporalv1.E_Field).(*temporalv1.FieldOptions)

	flagName := svc.flagName(field, prefix)

	usage := opts.GetCli().GetUsage()
	if usage == "" {
		usage = strings.TrimSpace(strings.ReplaceAll(strings.TrimPrefix(field.Comments.Leading.String(), "//"), "\n//", ""))
	}
	if usage == "" {
		usage = fmt.Sprintf("set the value of the operation's %q parameter", name)
	}

	// determine cli flag type
	flagType, additionalusage, err := flagType(field)
	if err != nil {
		svc.Plugin.Error(fmt.Errorf("unable to generate cli flag for field %q: %w", field.Desc.Name(), err))
		return
	}
	usage += additionalusage
	flagType += "Flag"

	// generate flag
	flags.Op("&").Qual(cliPkg, flagType).CustomFunc(multiLineValues, func(flag *jen.Group) {
		flag.Id("Name").Op(":").Lit(flagName)
		flag.Id("Usage").Op(":").Lit(strings.TrimSpace(usage))
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			flag.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *jen.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if svc.cfg.CliCategories && category != "" {
			flag.Id("Category").Op(":").Lit(category)
		}
		if flagType == "TimestampFlag" {
			flag.Id("Layout").Op(":").Qual("time", "RFC3339Nano")
		}
	})
}

func (svc *Manifest) genCliUnmarshalMessage(f *jen.File, msg *protogen.Message) {
	goName := svc.getMessageName(msg)
	fnName := fmt.Sprintf("UnmarshalCliFlagsTo%s", goName)

	f.Commentf("%s unmarshals a %s from command line flags", fnName, goName)
	f.Func().
		Id(fnName).
		ParamsFunc(func(g *jen.Group) {
			g.Id("cmd").Op("*").Qual(CliPkg, "Context")
			g.Id("options").Op("...").Qual(helpersPkg, "UnmarshalCliFlagsOptions")
		}).
		ParamsFunc(func(g *jen.Group) {
			g.Id("*").Qual(string(msg.GoIdent.GoImportPath), goName)
			g.Error()
		}).
		BlockFunc(func(g *jen.Group) {
			// Initialize a new instance of the message
			g.Var().Id("result").Qual(string(msg.GoIdent.GoImportPath), goName)
			g.If(jen.Id("cmd").Dot("IsSet").Call(jen.Lit("input-file"))).Block(
				jen.List(jen.Id("inputFile"), jen.Err()).Op(":=").Qual(homedirPkg, "Expand").Call(jen.Id("cmd").Dot("String").Call(jen.Lit("input-file"))),
				jen.If(jen.Err().Op("!=").Nil()).Block(
					jen.Id("inputFile").Op("=").Id("cmd").Dot("String").Call(jen.Lit("input-file")),
				),
				jen.List(jen.Id("b"), jen.Err()).Op(":=").Qual("os", "ReadFile").Call(jen.Id("inputFile")),
				jen.If(jen.Err().Op("!=").Nil()).Block(
					jen.Return(jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit("error reading input-file: %w"), jen.Err())),
				),
				jen.If(
					jen.Err().Op(":=").Qual(protojsonPkg, "Unmarshal").Call(jen.Id("b"), jen.Op("&").Id("result")),
					jen.Err().Op("!=").Nil(),
				).Block(
					jen.Return(jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit("error parsing input-file json: %w"), jen.Err())),
				),
			)

			// Iterate over each field of the message
			var fields []*jen.Statement
		FIELDS:
			for _, field := range msg.Fields {
				fieldName := svc.getFieldName(field)
				flag := svc.flagName(field, "")
				useSetter := slices.Contains([]gofeaturespb.GoFeatures_APILevel{gofeaturespb.GoFeatures_API_OPAQUE, gofeaturespb.GoFeatures_API_HYBRID}, svc.File.APILevel)

				// Determine the type of the flag
				fieldType, _, err := flagType(field)
				if err != nil {
					continue // Skip unsupported field types
				}

				// Generate the expression to get the value from cmd based on the field type
				var getValueExpr *jen.Statement
				var isPtr bool
				switch fieldType {
				case "Bool":
					getValueExpr = jen.Id("cmd").Dot("Bool")
				case "Duration":
					getValueExpr = jen.Id("cmd").Dot("Duration")
				case "Timestamp":
					getValueExpr = jen.Id("cmd").Dot("Timestamp")
				case "String":
					getValueExpr = jen.Id("cmd").Dot("String")
				case "Int64":
					getValueExpr = jen.Id("cmd").Dot("Int64")
				case "Uint64":
					getValueExpr = jen.Id("cmd").Dot("Uint64")
				case "Float64":
					getValueExpr = jen.Id("cmd").Dot("Float64")
				case "StringSlice":
					getValueExpr, isPtr = jen.Id("cmd").Dot("StringSlice"), true
				case "Int64Slice":
					getValueExpr, isPtr = jen.Id("cmd").Dot("Int64Slice"), true
				case "Uint64Slice":
					getValueExpr, isPtr = jen.Id("cmd").Dot("Uint64Slice"), true
				case "Float64Slice":
					getValueExpr, isPtr = jen.Id("cmd").Dot("Float64Slice"), true
				default:
					continue // Skip unsupported types
				}
				getValueExpr = getValueExpr.Call(jen.Id("flag"))

				// Transform the value based on the field type if necessary
				var transforms []func(*jen.Group)
				valueName := "value"
				assignmentName := valueName

				// Wrap the value expression in a cast expression if necessary
				if cast, ok := castFuncs[field.Desc.Kind()]; ok {
					unmarshal := func(g *jen.Group, valueExpr, errReturn *jen.Statement) {
						g.List(jen.Id(valueName), jen.Err()).Op(":=").Qual(ConvertPkg, "SafeCast").Types(cast.from, cast.to).Call(valueExpr)
						g.If(jen.Err().Op("!=").Nil()).Block(
							jen.Return(errReturn, jen.Err()),
						)
					}
					transforms = append(transforms, func(g *jen.Group) {
						if field.Desc.IsList() {
							g.List(jen.Id(valueName), jen.Err()).Op(":=").Qual(ConvertPkg, "MapSliceFunc").CallFunc(func(g *jen.Group) {
								g.Add(getValueExpr)
								g.Func().
									Params(jen.Id("v").Add(cast.from)).
									Params(
										jen.Id("result").Add(cast.to),
										jen.Err().Error(),
									).
									BlockFunc(func(g *jen.Group) {
										unmarshal(g, jen.Id("v"), jen.Id("result"))
										g.Return(jen.Id(valueName), jen.Nil())
									})
							})
							g.If(jen.Err().Op("!=").Nil()).BlockFunc(func(g *jen.Group) {
								g.Return(jen.Nil(), jen.Err())
							})
						} else {
							unmarshal(g, getValueExpr, jen.Nil())
						}
					})
				}

				switch {
				case field.Desc.Kind() == protoreflect.BoolKind && field.Desc.IsList():
					isPtr = true
					transforms = append(transforms, func(g *jen.Group) {
						g.List(jen.Id(valueName), jen.Err()).Op(":=").Qual(ConvertPkg, "MapSliceFunc").CallFunc(func(g *jen.Group) {
							g.Add(getValueExpr)
							g.Func().
								Params(jen.Id("v").String()).
								Params(
									jen.Id("result").Bool(),
									jen.Err().Error(),
								).
								BlockFunc(func(g *jen.Group) {
									g.Return(jen.Qual("strconv", "ParseBool").Call(jen.Id("v")))
								})
						})
						g.If(jen.Err().Op("!=").Nil()).BlockFunc(func(g *jen.Group) {
							g.Return(jen.Nil(), jen.Err())
						})
					})

				// Base64 decode bytes fields from string flags
				case field.Desc.Kind() == protoreflect.BytesKind:
					isPtr = true
					if field.Desc.IsList() {
						transforms = append(transforms, func(g *jen.Group) {
							g.List(jen.Id(valueName), jen.Err()).Op(":=").Qual(ConvertPkg, "MapSliceFunc").CallFunc(func(g *jen.Group) {
								g.Add(getValueExpr)
								g.Func().
									Params(jen.Id("v").String()).
									Params(
										jen.Id("result").Index().Byte(),
										jen.Err().Error(),
									).
									BlockFunc(func(g *jen.Group) {
										g.Return(jen.Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(jen.Id("v")))
									})
							})
							g.If(jen.Err().Op("!=").Nil()).BlockFunc(func(g *jen.Group) {
								g.Return(jen.Nil(), jen.Err())
							})
						})
					} else {
						transforms = append(transforms, func(g *jen.Group) {
							g.List(jen.Id(valueName), jen.Err()).Op(":=").Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(getValueExpr)
							g.If(jen.Err().Op("!=").Nil()).Block(
								jen.Return(jen.Nil(), jen.Err()),
							)
						})
					}

				// Convert enum strings to their corresponding enum type
				case field.Desc.Kind() == protoreflect.EnumKind:
					if field.Desc.IsList() {
						isPtr = true
						transforms = append(transforms, func(g *jen.Group) {
							g.List(jen.Id(valueName), jen.Err()).Op(":=").Qual(ConvertPkg, "MapSliceFunc").CallFunc(func(g *jen.Group) {
								g.Add(getValueExpr)
								g.Func().
									Params(jen.Id("v").String()).
									Params(
										jen.Id("result").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName),
										jen.Err().Error(),
									).
									BlockFunc(func(g *jen.Group) {
										g.List(jen.Id("enumID"), jen.Id("ok")).Op(":=").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName+"_value").Index(jen.Id("v"))
										g.If(jen.Op("!").Id("ok")).Block(
											jen.Return(
												jen.Id("result"), jen.Qual("fmt", "Errorf").Call(jen.Lit("invalid value for enum field %s"), jen.Lit(fieldName)),
											),
										)
										g.Return(jen.Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName).Call(jen.Id("enumID")), jen.Nil())
									})
							})
							g.If(jen.Err().Op("!=").Nil()).BlockFunc(func(g *jen.Group) {
								g.Return(jen.Nil(), jen.Err())
							})
						})
					} else {
						transforms = append(transforms, func(g *jen.Group) {
							g.List(jen.Id("enumID"), jen.Id("ok")).Op(":=").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName+"_value").Index(getValueExpr)
							g.If(jen.Op("!").Id("ok")).Block(
								jen.Return(
									jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit("invalid value for enum field %s"), jen.Lit(fieldName)),
								),
							)
							g.Id(valueName).Op(":=").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName).Call(jen.Id("enumID"))
						})
					}

				// Unmarshal protojson serialized maps to map fields
				case field.Desc.IsMap():
					isPtr = true
					transforms = append(transforms, func(g *jen.Group) {
						g.Var().Id("tmp").Qual(string(msg.GoIdent.GoImportPath), goName)
						g.If(jen.Err().Op(":=").Qual(ProtojsonPkg, "Unmarshal").Call(
							jen.Index().Byte().Call(jen.Qual("fmt", "Sprintf").Call(jen.Lit(fmt.Sprintf(`{"%s":%%s}`, field.Desc.JSONName())), getValueExpr)),
							jen.Op("&").Id("tmp"),
						), jen.Err().Op("!=").Nil()).Block(
							jen.Return(jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("error unmarshalling %q map flag: %%w", flag)), jen.Err())),
						)
						g.Id(valueName).Op(":=").Id("tmp").Dot(fmt.Sprintf("Get%s", fieldName)).Call()
					})

				case field.Desc.Kind() == protoreflect.MessageKind:
					isPtr = true
					switch field.Desc.Message().FullName() {
					// Skip empty fields
					case "google.protobuf.Empty":
						continue FIELDS

					// Convert time.Duration to *durationpb.Duration
					case "google.protobuf.Duration":
						transforms = append(transforms, func(g *jen.Group) {
							g.Id(valueName).Op(":=").Qual("google.golang.org/protobuf/types/known/durationpb", "New").Call(getValueExpr)
						})

					// Convert time.Time to *timestamppb.Timestamp
					case "google.protobuf.Timestamp":
						transforms = append(transforms, func(g *jen.Group) {
							g.Id("v").Op(":=").Add(getValueExpr)
							g.Id(valueName).Op(":=").Qual("google.golang.org/protobuf/types/known/timestamppb", "New").Call(jen.Op("*").Id("v"))
						})

					// Unmarshal protojson serialized messages
					default:
						// Unmarshal a protojson serialized message returned by valueExpr into a new instance of the message
						// assigned to a tmp variable
						unmarshal := func(g *jen.Group, valueExpr *jen.Statement) {
							g.Var().Id("tmp").Qual(string(field.Message.GoIdent.GoImportPath), field.Message.GoIdent.GoName)
							g.If(
								jen.Err().Op(":=").Qual(ProtojsonPkg, "Unmarshal").Call(jen.Index().Byte().Call(valueExpr), jen.Op("&").Id("tmp")),
								jen.Err().Op("!=").Nil(),
							).Block(
								jen.Return(jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("error unmarshalling %q flag: %%w", flag)), jen.Err())),
							)
						}

						transforms = append(transforms, func(g *jen.Group) {
							if field.Desc.IsList() {
								g.List(jen.Id(valueName), jen.Err()).Op(":=").Qual(ConvertPkg, "MapSliceFunc").CallFunc(func(g *jen.Group) {
									g.Add(getValueExpr)
									g.Func().
										Params(jen.Id("v").String()).
										Params(
											jen.Op("*").Qual(string(field.Message.GoIdent.GoImportPath), field.Message.GoIdent.GoName),
											jen.Error(),
										).
										BlockFunc(func(g *jen.Group) {
											unmarshal(g, jen.Id("v"))
											g.Return(jen.Op("&").Id("tmp"), jen.Nil())
										})
								})
								g.If(jen.Err().Op("!=").Nil()).BlockFunc(func(g *jen.Group) {
									g.Return(jen.Nil(), jen.Err())
								})
							} else {
								unmarshal(g, getValueExpr)
								g.Id(valueName).Op(":=").Op("&").Id("tmp")
							}
						})
					}
				}

				// If the field is part of a genuine oneof, we need to create a new instance of the oneof
				if isGenuineOneof(field) && !useSetter {
					oneOfValueName := "oneOfValue"
					assignmentName = oneOfValueName
					fieldName = field.Oneof.GoName
					isPtr = true
					if len(transforms) == 0 {
						transforms = append(transforms, func(g *jen.Group) {
							g.Id(valueName).Op(":=").Add(getValueExpr)
						})
					}
					transforms = append(transforms, func(g *jen.Group) {
						if len(transforms) == 0 {
							transforms = append(transforms, func(g *jen.Group) {
								g.Id(valueName).Op(":=").Add(getValueExpr)
							})
						}
						g.Id(oneOfValueName).Op(":=").Op("&").Qual(string(field.GoIdent.GoImportPath), field.GoIdent.GoName).Values(jen.Dict{
							jen.Id(field.GoName): jen.Id(valueName),
						})
					})
				}

				if len(transforms) == 0 {
					transforms = append(transforms, func(g *jen.Group) {
						g.Id(valueName).Op(":=").Add(getValueExpr)
					})
				}

				// Check if the flag is set
				fields = append(fields,
					jen.If(
						jen.Id("flag").Op(":=").Id("opts").Dot("FlagName").Call(jen.Lit(flag)),
						jen.Id("cmd").Dot("IsSet").Call(jen.Id("flag")),
					).BlockFunc(func(g *jen.Group) {
						// Extract and transform the value from the flag as `valueName` variable
						for _, transform := range transforms {
							transform(g)
						}

						// Determine if we should use a setter method or set the field directly
						if useSetter {
							// Use the appropriate setter method
							setterName := "Set" + fieldName
							g.Id("result").Dot(setterName).Call(jen.Id(assignmentName))
						} else if (field.Desc.HasOptionalKeyword() || field.Desc.HasPresence()) && !isPtr {
							g.Id("result").Dot(fieldName).Op("=").Op("&").Id(assignmentName)
						} else {
							// Default to setting the field directly if API level is not specified
							g.Id("result").Dot(fieldName).Op("=").Id(assignmentName)
						}
					}),
				)
			}

			if len(fields) > 0 {
				g.Id("opts").Op(":=").Qual(helpersPkg, "UnmarshalCliFlagsOptions").Values()
				g.If(jen.Len(jen.Id("options")).Op(">").Lit(0)).Block(
					jen.Id("opts").Op("=").Id("options").Index(jen.Lit(0)),
				)
				for _, f := range fields {
					g.Add(f)
				}
			}

			// Return the populated message and nil error
			g.Return(jen.Op("&").Id("result"), jen.Nil())
		})
}

// isGenuineOneof returns true if the given field is part of a user‐declared (genuine)
// oneof, and false otherwise. (Proto3 optional fields are represented as synthetic oneofs.)
func isGenuineOneof(field *protogen.Field) bool {
	if field.Oneof == nil {
		return false
	}
	// Proto3 optional fields are represented as oneof fields, but they are synthetic.
	if field.Oneof.Desc.IsSynthetic() {
		return false
	}
	return true
}
