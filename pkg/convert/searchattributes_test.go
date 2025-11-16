package convert

import (
	"testing"
	"time"

	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/temporal"
)

func TestMarshalTypedSearchAttributes(t *testing.T) {
	tests := []struct {
		name   string
		tsa    string
		errors []string
		assert func(t *testing.T, sa temporal.SearchAttributes)
	}{
		{
			name: "success",
			tsa: `
bool.boolValue = true
bool.boolValue2 = "false"
float64.float64Value = 123.0
float64.float64Value2 = 123
float64.float64Value3 = "123.0"
int64.int64Value = 123
int64.int64Value2 = "123"
int64.int64Value3 = 123.0
keyword_list.keywordListValue = ["qux1", "qux2"]
keyword.keywordValue = "bar"
string.textValue = "test value"
time.timeValue = now().ts_tz("UTC")
time.timeValue2 = now().ts_tz("UTC").ts_format("2006-01-02T15:04:05Z")
`,
			assert: func(t *testing.T, sa temporal.SearchAttributes) {
				boolValue, ok := sa.GetBool(temporal.NewSearchAttributeKeyBool("boolValue"))
				require.True(t, ok)
				require.True(t, boolValue)
				boolValue2, ok := sa.GetBool(temporal.NewSearchAttributeKeyBool("boolValue2"))
				require.True(t, ok)
				require.False(t, boolValue2)
				float64Value, ok := sa.GetFloat64(temporal.NewSearchAttributeKeyFloat64("float64Value"))
				require.True(t, ok)
				require.Equal(t, 123.0, float64Value)
				float64Value2, ok := sa.GetFloat64(temporal.NewSearchAttributeKeyFloat64("float64Value2"))
				require.True(t, ok)
				require.Equal(t, 123.0, float64Value2)
				float64Value3, ok := sa.GetFloat64(temporal.NewSearchAttributeKeyFloat64("float64Value3"))
				require.True(t, ok)
				require.Equal(t, 123.0, float64Value3)
				int64Value, ok := sa.GetInt64(temporal.NewSearchAttributeKeyInt64("int64Value"))
				require.True(t, ok)
				require.Equal(t, int64(123), int64Value)
				int64Value2, ok := sa.GetInt64(temporal.NewSearchAttributeKeyInt64("int64Value2"))
				require.True(t, ok)
				require.Equal(t, int64(123), int64Value2)
				int64Value3, ok := sa.GetInt64(temporal.NewSearchAttributeKeyInt64("int64Value3"))
				require.True(t, ok)
				require.Equal(t, int64(123), int64Value3)
				keywordListValue, ok := sa.GetKeywordList(temporal.NewSearchAttributeKeyKeywordList("keywordListValue"))
				require.True(t, ok)
				require.Equal(t, []string{"qux1", "qux2"}, keywordListValue)
				keywordValue, ok := sa.GetKeyword(temporal.NewSearchAttributeKeyKeyword("keywordValue"))
				require.True(t, ok)
				require.Equal(t, "bar", keywordValue)
				textValue, ok := sa.GetString(temporal.NewSearchAttributeKeyString("textValue"))
				require.True(t, ok)
				require.Equal(t, "test value", textValue)
				timeValue, ok := sa.GetTime(temporal.NewSearchAttributeKeyTime("timeValue"))
				require.True(t, ok)
				require.WithinDuration(t, time.Now().UTC(), timeValue, 1*time.Second)
				timeValue2, ok := sa.GetTime(temporal.NewSearchAttributeKeyTime("timeValue2"))
				require.True(t, ok)
				require.WithinDuration(t, timeValue, timeValue2, 1*time.Second)
			},
		},
		{
			name: "invalid bool string",
			tsa:  `bool.boolValue = "invalid"`,
			errors: []string{
				"error parsing string to bool: strconv.ParseBool: parsing \"invalid\": invalid syntax",
			},
		},
		{
			name: "invalid bool float64",
			tsa:  `bool.boolValue = 123.0`,
			errors: []string{
				"expected bool or string, got float64",
			},
		},
		{
			name: "invalid float64 string",
			tsa:  `float64.float64Value = "invalid"`,
			errors: []string{
				"error parsing string to float64: strconv.ParseFloat: parsing \"invalid\": invalid syntax",
			},
		},
		{
			name: "invalid float64 bool",
			tsa:  `float64.float64Value = true`,
			errors: []string{
				"expected float64, int64, or string, got bool",
			},
		},
		{
			name: "invalid int64 string",
			tsa:  `int64.int64Value = "invalid"`,
			errors: []string{
				"error parsing string to int64: strconv.ParseInt: parsing \"invalid\": invalid syntax",
			},
		},
		{
			name: "invalid int64 bool",
			tsa:  `int64.int64Value = true`,
			errors: []string{
				"expected int64, float64, or string, got bool",
			},
		},
		{
			name: "invalid string bool",
			tsa:  `string.textValue = true`,
			errors: []string{
				"expected string, got bool",
			},
		},
		{
			name: "invalid keyword bool",
			tsa:  `keyword.keywordValue = true`,
			errors: []string{
				"expected string, got bool",
			},
		},
		{
			name: "invalid keyword list bool",
			tsa:  `keyword_list.keywordListValue = true`,
			errors: []string{
				"expected []any, got bool",
			},
		},
		{
			name: "invalid keyword list values",
			tsa:  `keyword_list.keywordListValue = [123]`,
			errors: []string{
				"expected string, got int64",
			},
		},
		{
			name: "invalid time string",
			tsa:  `time.timeValue = "invalid"`,
			errors: []string{
				"error parsing string to time.Time: parsing time",
			},
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			m := expression.MustParseMapping(c.tsa)
			o, err := m.Query(map[string]any{})
			require.NoError(t, err)
			sa, ok := o.(map[string]any)
			require.True(t, ok)
			tsa, err := MarshalTypedSearchAttributes(sa)
			if len(c.errors) > 0 {
				require.Error(t, err)
				for _, e := range c.errors {
					require.ErrorContains(t, err, e)
				}
			} else {
				require.NoError(t, err)
				if c.assert != nil {
					c.assert(t, tsa)
				}
			}
		})
	}
}
