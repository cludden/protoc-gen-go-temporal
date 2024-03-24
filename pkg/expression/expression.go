package expression

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/cludden/benthos/v4/public/bloblang"
	_ "github.com/cludden/benthos/v4/public/components/pure"
	_ "github.com/cludden/benthos/v4/public/components/pure/extended"
	"google.golang.org/protobuf/encoding/protojson"
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
		structured, err = ToStructured(msg)
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

// MustParseMapping attempts to parse a bloblang mapping and panics on error
func MustParseMapping(input string) *bloblang.Executor {
	m, err := bloblang.Parse(input)
	if err != nil {
		panic(err)
	}
	return m
}

// ParseExpression parses an Expression value from the provided string
func ParseExpression(input string) (*Expression, error) {
	expr, err := parser.ParseString("", input)
	if err != nil {
		return nil, fmt.Errorf("error parsing expression %q: %w", input, err)
	}
	for _, fragment := range expr.Fragments {
		if e := fragment.Expr; e != nil {
			m, err := bloblang.Parse(fmt.Sprintf(`root = %s`, e.Mapping))
			if err != nil {
				return nil, fmt.Errorf("error parsing bloblang mapping: %q, %w", e.Mapping, err)
			}
			e.m = m
		}
	}
	return expr, nil
}

// ToStructured marshals a proto message into a map[string]any value
func ToStructured(msg protoreflect.Message) (any, error) {
	b, err := protojson.Marshal(msg.Interface())
	if err != nil {
		return nil, err
	}
	structured := make(map[string]any)
	return structured, json.Unmarshal(b, &structured)
}
