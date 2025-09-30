package plugin

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/gofeaturespb"
)

var (
	castFuncs = map[protoreflect.Kind]struct {
		from *j.Statement
		to   *j.Statement
	}{
		protoreflect.Fixed32Kind:  {j.Uint64(), j.Uint32()},
		protoreflect.FloatKind:    {j.Float64(), j.Float32()},
		protoreflect.Int32Kind:    {j.Int64(), j.Int32()},
		protoreflect.Sfixed32Kind: {j.Int64(), j.Int32()},
		protoreflect.Sint32Kind:   {j.Int64(), j.Int32()},
		protoreflect.Uint32Kind:   {j.Uint64(), j.Uint32()},
	}
)

func (m *Manifest) flagName(field *protogen.Field, prefix string) string {
	opts := proto.GetExtension(field.Desc.Options(), temporalv1.E_Field).(*temporalv1.FieldOptions)

	name := opts.GetCli().GetName()
	if name == "" {
		name = m.caser.ToKebab(field.GoName)
	}
	if prefix != "" {
		name = fmt.Sprintf("%s-%s", m.caser.ToKebab(prefix), name)
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
func (m *Manifest) genCliFlagForField(g *j.Group, field *protogen.Field, category string, prefix string) {
	name := m.getFieldName(field)
	opts := proto.GetExtension(field.Desc.Options(), temporalv1.E_Field).(*temporalv1.FieldOptions)

	flagName := m.flagName(field, prefix)

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
		m.Plugin.Error(fmt.Errorf("unable to generate cli flag for field %q: %w", field.Desc.Name(), err))
		return
	}
	usage += additionalusage
	flagType += "Flag"

	pkg := cliPkg
	if m.cfg.CliV3Enabled {
		pkg = cliV3Pkg
	}

	// generate flag
	g.Op("&").Qual(pkg, flagType).CustomFunc(multiLineValues, func(flag *j.Group) {
		flag.Id("Name").Op(":").Lit(flagName)
		flag.Id("Usage").Op(":").Lit(strings.TrimSpace(usage))
		if aliases := opts.GetCli().GetAliases(); len(aliases) > 0 {
			flag.Id("Aliases").Op(":").Index().String().ValuesFunc(func(g *j.Group) {
				for _, alias := range aliases {
					g.Lit(alias)
				}
			})
		}
		if m.cfg.CliCategories && category != "" {
			flag.Id("Category").Op(":").Lit(category)
		}
		if flagType == "TimestampFlag" {
			if m.cfg.CliV3Enabled {
				flag.Id("Config").Op(":").Qual(cliV3Pkg, "TimestampConfig").Values(j.DictFunc(func(g j.Dict) {
					g[j.Id("Layouts")] = j.Index().String().Values(j.Qual("time", "RFC3339Nano"))
				}))
			} else {
				flag.Id("Layout").Op(":").Qual("time", "RFC3339Nano")
			}
		}
	})
}

func (m *Manifest) genCliUnmarshalMessage(f *j.File, msg *protogen.Message) {
	goName := m.getMessageName(msg)
	fnName := fmt.Sprintf("UnmarshalCliFlagsTo%s", goName)

	f.Commentf("%s unmarshals a %s from command line flags", fnName, goName)
	f.Func().
		Id(fnName).
		ParamsFunc(func(g *j.Group) {
			if m.cfg.CliV3Enabled {
				g.Id("cmd").Op("*").Qual(cliV3Pkg, "Command")
			} else {
				g.Id("cmd").Op("*").Qual(cliPkg, "Context")
			}
			g.Id("options").Op("...").Qual(helpersPkg, "UnmarshalCliFlagsOptions")
		}).
		ParamsFunc(func(g *j.Group) {
			g.Id("*").Qual(string(msg.GoIdent.GoImportPath), goName)
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// initialize unmarshal options
			g.Id("opts").Op(":=").Qual(helpersPkg, "FlattenUnmarshalCliFlagsOptions").CallFunc(func(g *j.Group) {
				g.Id("options").Op("...")
			})

			// Initialize a new instance of the message
			g.Var().Id("result").Qual(string(msg.GoIdent.GoImportPath), goName)

			// if opts.FromFile is set, read the file and unmarshal the json
			g.If(j.Id("opts").Dot("FromFile").Op("!=").Lit("").Op("&&").Id("cmd").Dot("IsSet").Call(j.Id("opts").Dot("FromFile"))).Block(
				j.List(j.Id("f"), j.Err()).Op(":=").Qual(homedirPkg, "Expand").Call(j.Id("cmd").Dot("String").Call(j.Id("opts").Dot("FromFile"))),
				j.If(j.Err().Op("!=").Nil()).Block(
					j.Id("f").Op("=").Id("cmd").Dot("String").Call(j.Id("opts").Dot("FromFile")),
				),
				j.List(j.Id("b"), j.Err()).Op(":=").Qual("os", "ReadFile").Call(j.Id("f")),
				j.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error reading %s: %w"), j.Id("opts").Dot("FromFile"), j.Err())),
				),
				j.If(
					j.Err().Op(":=").Qual(protojsonPkg, "Unmarshal").Call(j.Id("b"), j.Op("&").Id("result")),
					j.Err().Op("!=").Nil(),
				).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error parsing %s json: %w"), j.Id("opts").Dot("FromFile"), j.Err())),
				),
			)

			// Iterate over each field of the message
			var fields []*j.Statement
		FIELDS:
			for _, field := range msg.Fields {
				fieldName := m.getFieldName(field)
				flag := m.flagName(field, "")
				useSetter := slices.Contains([]gofeaturespb.GoFeatures_APILevel{gofeaturespb.GoFeatures_API_OPAQUE, gofeaturespb.GoFeatures_API_HYBRID}, m.File.APILevel)

				// Determine the type of the flag
				fieldType, _, err := flagType(field)
				if err != nil {
					continue // Skip unsupported field types
				}

				// Generate the expression to get the value from cmd based on the field type
				var getValueExpr *j.Statement
				var isPtr bool
				switch fieldType {
				case "Bool":
					getValueExpr = j.Id("cmd").Dot("Bool")
				case "Duration":
					getValueExpr = j.Id("cmd").Dot("Duration")
				case "Timestamp":
					getValueExpr = j.Id("cmd").Dot("Timestamp")
				case "String":
					getValueExpr = j.Id("cmd").Dot("String")
				case "Int64":
					getValueExpr = j.Id("cmd").Dot("Int64")
				case "Uint64":
					getValueExpr = j.Id("cmd").Dot("Uint64")
				case "Float64":
					getValueExpr = j.Id("cmd").Dot("Float64")
				case "StringSlice":
					getValueExpr, isPtr = j.Id("cmd").Dot("StringSlice"), true
				case "Int64Slice":
					getValueExpr, isPtr = j.Id("cmd").Dot("Int64Slice"), true
				case "Uint64Slice":
					getValueExpr, isPtr = j.Id("cmd").Dot("Uint64Slice"), true
				case "Float64Slice":
					getValueExpr, isPtr = j.Id("cmd").Dot("Float64Slice"), true
				default:
					continue // Skip unsupported types
				}
				getValueExpr = getValueExpr.Call(j.Id("flag"))

				// Transform the value based on the field type if necessary
				var transforms []func(*j.Group)
				valueName := "value"
				assignmentName := valueName

				// Wrap the value expression in a cast expression if necessary
				if cast, ok := castFuncs[field.Desc.Kind()]; ok {
					unmarshal := func(g *j.Group, valueExpr, errReturn *j.Statement) {
						g.List(j.Id(valueName), j.Err()).Op(":=").Qual(convertPkg, "SafeCast").Types(cast.from, cast.to).Call(valueExpr)
						g.If(j.Err().Op("!=").Nil()).Block(
							j.Return(errReturn, j.Err()),
						)
					}
					transforms = append(transforms, func(g *j.Group) {
						if field.Desc.IsList() {
							g.List(j.Id(valueName), j.Err()).Op(":=").Qual(convertPkg, "MapSliceFunc").CallFunc(func(g *j.Group) {
								g.Add(getValueExpr)
								g.Func().
									Params(j.Id("v").Add(cast.from)).
									Params(
										j.Id("result").Add(cast.to),
										j.Err().Error(),
									).
									BlockFunc(func(g *j.Group) {
										unmarshal(g, j.Id("v"), j.Id("result"))
										g.Return(j.Id(valueName), j.Nil())
									})
							})
							g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
								g.Return(j.Nil(), j.Err())
							})
						} else {
							unmarshal(g, getValueExpr, j.Nil())
						}
					})
				}

				switch {
				case field.Desc.Kind() == protoreflect.BoolKind && field.Desc.IsList():
					isPtr = true
					transforms = append(transforms, func(g *j.Group) {
						g.List(j.Id(valueName), j.Err()).Op(":=").Qual(convertPkg, "MapSliceFunc").CallFunc(func(g *j.Group) {
							g.Add(getValueExpr)
							g.Func().
								Params(j.Id("v").String()).
								Params(
									j.Id("result").Bool(),
									j.Err().Error(),
								).
								BlockFunc(func(g *j.Group) {
									g.Return(j.Qual("strconv", "ParseBool").Call(j.Id("v")))
								})
						})
						g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
							g.Return(j.Nil(), j.Err())
						})
					})

				// Base64 decode bytes fields from string flags
				case field.Desc.Kind() == protoreflect.BytesKind:
					isPtr = true
					if field.Desc.IsList() {
						transforms = append(transforms, func(g *j.Group) {
							g.List(j.Id(valueName), j.Err()).Op(":=").Qual(convertPkg, "MapSliceFunc").CallFunc(func(g *j.Group) {
								g.Add(getValueExpr)
								g.Func().
									Params(j.Id("v").String()).
									Params(
										j.Id("result").Index().Byte(),
										j.Err().Error(),
									).
									BlockFunc(func(g *j.Group) {
										g.Return(j.Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(j.Id("v")))
									})
							})
							g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
								g.Return(j.Nil(), j.Err())
							})
						})
					} else {
						transforms = append(transforms, func(g *j.Group) {
							g.List(j.Id(valueName), j.Err()).Op(":=").Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(getValueExpr)
							g.If(j.Err().Op("!=").Nil()).Block(
								j.Return(j.Nil(), j.Err()),
							)
						})
					}

				// Convert enum strings to their corresponding enum type
				case field.Desc.Kind() == protoreflect.EnumKind:
					if field.Desc.IsList() {
						isPtr = true
						transforms = append(transforms, func(g *j.Group) {
							g.List(j.Id(valueName), j.Err()).Op(":=").Qual(convertPkg, "MapSliceFunc").CallFunc(func(g *j.Group) {
								g.Add(getValueExpr)
								g.Func().
									Params(j.Id("v").String()).
									Params(
										j.Id("result").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName),
										j.Err().Error(),
									).
									BlockFunc(func(g *j.Group) {
										g.List(j.Id("enumID"), j.Id("ok")).Op(":=").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName+"_value").Index(j.Id("v"))
										g.If(j.Op("!").Id("ok")).Block(
											j.Return(
												j.Id("result"), j.Qual("fmt", "Errorf").Call(j.Lit("invalid value for enum field %s"), j.Lit(fieldName)),
											),
										)
										g.Return(j.Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName).Call(j.Id("enumID")), j.Nil())
									})
							})
							g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
								g.Return(j.Nil(), j.Err())
							})
						})
					} else {
						transforms = append(transforms, func(g *j.Group) {
							g.List(j.Id("enumID"), j.Id("ok")).Op(":=").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName+"_value").Index(getValueExpr)
							g.If(j.Op("!").Id("ok")).Block(
								j.Return(
									j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("invalid value for enum field %s"), j.Lit(fieldName)),
								),
							)
							g.Id(valueName).Op(":=").Qual(string(field.Enum.GoIdent.GoImportPath), field.Enum.GoIdent.GoName).Call(j.Id("enumID"))
						})
					}

				// Unmarshal protojson serialized maps to map fields
				case field.Desc.IsMap():
					isPtr = true
					transforms = append(transforms, func(g *j.Group) {
						g.Var().Id("tmp").Qual(string(msg.GoIdent.GoImportPath), goName)
						g.If(j.Err().Op(":=").Qual(protojsonPkg, "Unmarshal").Call(
							j.Index().Byte().Call(j.Qual("fmt", "Sprintf").Call(j.Lit(fmt.Sprintf(`{"%s":%%s}`, field.Desc.JSONName())), getValueExpr)),
							j.Op("&").Id("tmp"),
						), j.Err().Op("!=").Nil()).Block(
							j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("error unmarshalling %q map flag: %%w", flag)), j.Err())),
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
						transforms = append(transforms, func(g *j.Group) {
							g.Id(valueName).Op(":=").Qual(durationpbPkg, "New").Call(getValueExpr)
						})

					// Convert time.Time to *timestamppb.Timestamp
					case "google.protobuf.Timestamp":
						transforms = append(transforms, func(g *j.Group) {
							g.Id("v").Op(":=").Add(getValueExpr)
							g.Id(valueName).Op(":=").Qual(timestamppbPkg, "New").CallFunc(func(g *j.Group) {
								if m.cfg.CliV3Enabled {
									g.Id("v")
								} else {
									g.Op("*").Id("v")
								}
							})
						})

					// Unmarshal protojson serialized messages
					default:
						// Unmarshal a protojson serialized message returned by valueExpr into a new instance of the message
						// assigned to a tmp variable
						unmarshal := func(g *j.Group, valueExpr *j.Statement) {
							g.Var().Id("tmp").Qual(string(field.Message.GoIdent.GoImportPath), field.Message.GoIdent.GoName)
							g.If(
								j.Err().Op(":=").Qual(protojsonPkg, "Unmarshal").Call(j.Index().Byte().Call(valueExpr), j.Op("&").Id("tmp")),
								j.Err().Op("!=").Nil(),
							).Block(
								j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("error unmarshalling %q flag: %%w", flag)), j.Err())),
							)
						}

						transforms = append(transforms, func(g *j.Group) {
							if field.Desc.IsList() {
								g.List(j.Id(valueName), j.Err()).Op(":=").Qual(convertPkg, "MapSliceFunc").CallFunc(func(g *j.Group) {
									g.Add(getValueExpr)
									g.Func().
										Params(j.Id("v").String()).
										Params(
											j.Op("*").Qual(string(field.Message.GoIdent.GoImportPath), field.Message.GoIdent.GoName),
											j.Error(),
										).
										BlockFunc(func(g *j.Group) {
											unmarshal(g, j.Id("v"))
											g.Return(j.Op("&").Id("tmp"), j.Nil())
										})
								})
								g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
									g.Return(j.Nil(), j.Err())
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
						transforms = append(transforms, func(g *j.Group) {
							g.Id(valueName).Op(":=").Add(getValueExpr)
						})
					}
					transforms = append(transforms, func(g *j.Group) {
						if len(transforms) == 0 {
							transforms = append(transforms, func(g *j.Group) {
								g.Id(valueName).Op(":=").Add(getValueExpr)
							})
						}
						g.Id(oneOfValueName).Op(":=").Op("&").Qual(string(field.GoIdent.GoImportPath), field.GoIdent.GoName).Values(j.Dict{
							j.Id(field.GoName): j.Id(valueName),
						})
					})
				}

				if len(transforms) == 0 {
					transforms = append(transforms, func(g *j.Group) {
						g.Id(valueName).Op(":=").Add(getValueExpr)
					})
				}

				// Check if the flag is set
				fields = append(fields,
					j.If(
						j.Id("flag").Op(":=").Id("opts").Dot("FlagName").Call(j.Lit(flag)),
						j.Id("cmd").Dot("IsSet").Call(j.Id("flag")),
					).BlockFunc(func(g *j.Group) {
						// Extract and transform the value from the flag as `valueName` variable
						for _, transform := range transforms {
							transform(g)
						}

						// Determine if we should use a setter method or set the field directly
						if useSetter {
							// Use the appropriate setter method
							setterName := "Set" + fieldName
							g.Id("result").Dot(setterName).Call(j.Id(assignmentName))
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
				for _, f := range fields {
					g.Add(f)
				}
			}

			// Return the populated message and nil error
			g.Return(j.Op("&").Id("result"), j.Nil())
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
