package expression

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	cases := []struct {
		input    string
		expected []*Fragment
		err      []string
	}{
		{
			input:    "foo",
			expected: []*Fragment{{Ident: "foo"}},
		},
		{
			input:    `foo\${!}`,
			expected: []*Fragment{{Ident: `foo${!}`}},
		},
		{
			input: `foo${!bar`,
			err:   []string{"detected partial expression"},
		},
		{
			input: `foo${!bar}`,
			expected: []*Fragment{
				{Ident: `foo`},
				{Expr: &Query{Mapping: `bar`}},
			},
		},
		{
			input: `foo/${! foo.re_find_all("[abc]+").join("/") }/bar`,
			expected: []*Fragment{
				{Ident: `foo/`},
				{Expr: &Query{Mapping: ` foo.re_find_all("[abc]+").join("/") `}},
				{Ident: `/bar`},
			},
		},
		{
			input: `foo/${! foo.re_find_all("[abc]{1,5}").join("/") }/bar`,
			expected: []*Fragment{
				{Ident: `foo/`},
				{Expr: &Query{Mapping: ` foo.re_find_all("[abc]{1,5}").join("/") `}},
				{Ident: `/bar`},
			},
		},
		{
			input: `zzzzzz${! a\{{b{c{d{e\}}}}} }zzzzzz`,
			expected: []*Fragment{
				{Ident: "zzzzzz"},
				{Expr: &Query{Mapping: ` a{{b{c{d{e}}}}} `}},
				{Ident: "zzzzzz"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			expr, err := Lex(c.input)
			if len(c.err) > 0 {
				for _, msg := range c.err {
					require.ErrorContains(t, err, msg)
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, expr)
				require.Len(t, expr.Fragments, len(c.expected))
				for i, f := range c.expected {
					if f.Ident != "" {
						require.Equal(t, f.Ident, expr.Fragments[i].Ident)
					} else {
						require.NotNil(t, expr.Fragments[i].Expr)
						require.Equal(t, f.Expr.Mapping, expr.Fragments[i].Expr.Mapping)
					}
				}
				require.Equal(t, c.expected, expr.Fragments)
			}
		})
	}
}
