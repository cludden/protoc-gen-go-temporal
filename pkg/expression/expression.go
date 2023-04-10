package expression

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/benthosdev/benthos/v4/public/bloblang"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Declare simple grammer for bloblang expressions
type (
	Expression struct {
		Fragments []*Fragment `parser:"@@*"`
	}

	Fragment struct {
		Ident string `parser:"( @String"`
		Expr  *Query `parser:"  | '${!' @@ '}' )"`
	}

	Query struct {
		Mapping string `parser:"@Mapping"`
		m       *bloblang.Executor
	}
)

// Initialize lexer & parser for bloblang expressions
var (
	lexr = lexer.MustStateful(lexer.Rules{
		"Root": []lexer.Rule{
			{Name: "String", Pattern: `([^\$]|$[^{]|${[^!])+`, Action: nil},
			{Name: "Expr", Pattern: `\${!`, Action: lexer.Push("Expr")},
		},
		"Expr": []lexer.Rule{
			{Name: "Mapping", Pattern: `[^}]+`, Action: nil},
			{Name: "ExprEnd", Pattern: `}`, Action: lexer.Pop()},
		},
	})

	parser = participle.MustBuild[Expression](participle.Lexer(lexr))
)

// EvalExpression evaluates an expression against a proto message
func EvalExpression(expr *Expression, msg protoreflect.Message) (string, error) {
	var structured any
	var err error
	if msg != nil {
		structured, err = marshalMessage(msg)
		if err != nil {
			return "", fmt.Errorf("error serializing message for expression evaluation: %w", err)
		}
	}

	id := strings.Builder{}
	for _, fragment := range expr.Fragments {
		if fragment.Expr != nil {
			r, err := fragment.Expr.m.Query(structured)
			if err != nil {
				return "", fmt.Errorf("error querying expression: %w", err)
			}
			switch v := r.(type) {
			case string:
				id.WriteString(v)
			case []byte:
				id.Write(v)
			default:
				return "", fmt.Errorf("expected string result from expression, got: %T", r)
			}
			continue
		}
		id.WriteString(fragment.Ident)
	}
	return id.String(), nil
}

// MustParseExpression attempts to parse an Expression and panics on error
func MustParseExpression(input string) *Expression {
	expr, err := ParseExpression(input)
	if err != nil {
		panic(err)
	}
	return expr
}

// ParseExpression parses an Expression value from the provided string
func ParseExpression(input string) (*Expression, error) {
	expr, err := parser.ParseString("", input)
	if err != nil {
		return nil, fmt.Errorf("error parsing expression: %w", err)
	}
	for _, fragment := range expr.Fragments {
		if e := fragment.Expr; e != nil {
			m, err := bloblang.Parse(fmt.Sprintf(`root = %s`, e.Mapping))
			if err != nil {
				return nil, fmt.Errorf("Error parsing bloblang mapping: %q, %w", e.Mapping, err)
			}
			e.m = m
		}
	}
	return expr, nil
}

// marshalMessage marshals a proto message into a map[string]any value
func marshalMessage(msg protoreflect.Message) (any, error) {
	structured := make(map[string]any)
	var err error
	msg.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		val, merr := marshalValue(v, fd)
		if merr != nil {
			err = merr
			return false
		}
		structured[fd.JSONName()] = val
		return true
	})
	return structured, err
}

// marshalValue marshals a proto value into any
func marshalValue(val protoreflect.Value, fd protoreflect.FieldDescriptor) (any, error) {
	switch {
	case fd.IsList():
		return marshalList(val.List(), fd)
	case fd.IsMap():
		return marshalMap(val.Map(), fd)
	default:
		return marshalSingular(val, fd)
	}
}

// marshalSingular marshals a non-list, non-map proto value into its json-compatible scalar value
func marshalSingular(val protoreflect.Value, fd protoreflect.FieldDescriptor) (any, error) {
	switch kind := fd.Kind(); kind {
	case protoreflect.BoolKind:
		return val.Bool(), nil
	case protoreflect.StringKind:
		return val.String(), nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return val.Int(), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return val.Uint(), nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return val.Float(), nil
	case protoreflect.BytesKind:
		return base64.StdEncoding.EncodeToString(val.Bytes()), nil
	case protoreflect.EnumKind:
		if val.Enum() == protoreflect.EnumNumber(0) {
			return nil, nil
		} else {
			return string(fd.Enum().Values().ByNumber(val.Enum()).Name()), nil
		}
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return marshalMessage(val.Message())
	default:
		return nil, fmt.Errorf("unsupported proto kind: %v", kind)
	}
}

// marshalList marshals a proto list value into []any
func marshalList(list protoreflect.List, fd protoreflect.FieldDescriptor) ([]any, error) {
	structured := make([]any, list.Len())
	for i := 0; i < list.Len(); i++ {
		item := list.Get(i)
		val, err := marshalSingular(item, fd)
		if err != nil {
			return nil, err
		}
		structured[i] = val
	}
	return structured, nil
}

// marshalMap marshals a proto mmap value in map[string]any
func marshalMap(mmap protoreflect.Map, fd protoreflect.FieldDescriptor) (map[string]any, error) {
	structured := make(map[string]any)
	var err error
	mmap.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
		val, merr := marshalSingular(v, fd.MapValue())
		if merr != nil {
			err = merr
			return false
		}
		structured[mk.String()] = val
		return true
	})
	return structured, err
}
