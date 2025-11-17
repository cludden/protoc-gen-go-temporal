package convert

import (
	"fmt"
	"strconv"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func MarshalTypedSearchAttributes(sa map[string]any) (tsa temporal.SearchAttributes, err error) {
	var attrs []temporal.SearchAttributeUpdate
	for _, t := range workflow.DeterministicKeys(sa) {
		v, _ := sa[t].(map[string]any)
		switch t {
		case "bool":
			var x bool
			for k, vv := range v {
				switch vvv := vv.(type) {
				case bool:
					x = vvv
				case string:
					x, err = strconv.ParseBool(vvv)
					if err != nil {
						return temporal.SearchAttributes{}, fmt.Errorf("error parsing string to bool: %w", err)
					}
				default:
					return temporal.SearchAttributes{}, fmt.Errorf("expected bool or string, got %T", vvv)
				}
				attrs = append(attrs, temporal.NewSearchAttributeKeyBool(k).ValueSet(x))
			}
		case "float64":
			for k, vv := range v {
				var x float64
				switch vvv := vv.(type) {
				case float64:
					x = vvv
				case int64:
					x, err = SafeCast[int64, float64](vvv)
					if err != nil {
						return temporal.SearchAttributes{}, fmt.Errorf("error casting int64 to float64: %w", err)
					}
				case string:
					x, err = strconv.ParseFloat(vvv, 64)
					if err != nil {
						return temporal.SearchAttributes{}, fmt.Errorf("error parsing string to float64: %w", err)
					}
				default:
					return temporal.SearchAttributes{}, fmt.Errorf("expected float64, int64, or string, got %T", vvv)
				}
				attrs = append(attrs, temporal.NewSearchAttributeKeyFloat64(k).ValueSet(x))
			}
		case "int64":
			for k, vv := range v {
				var x int64
				switch vvv := vv.(type) {
				case int64:
					x = vvv
				case float64:
					x, err = SafeCast[float64, int64](vvv)
					if err != nil {
						return temporal.SearchAttributes{}, fmt.Errorf("error casting float64 to int64: %w", err)
					}
				case string:
					x, err = strconv.ParseInt(vvv, 10, 64)
					if err != nil {
						return temporal.SearchAttributes{}, fmt.Errorf("error parsing string to int64: %w", err)
					}
				default:
					return temporal.SearchAttributes{}, fmt.Errorf("expected int64, float64, or string, got %T", vvv)
				}
				attrs = append(attrs, temporal.NewSearchAttributeKeyInt64(k).ValueSet(x))
			}
		case "keyword", "string":
			for k, vv := range v {
				vvv, ok := vv.(string)
				if !ok {
					return temporal.SearchAttributes{}, fmt.Errorf("expected string, got %T", vv)
				}
				if t == "keyword" {
					attrs = append(attrs, temporal.NewSearchAttributeKeyKeyword(k).ValueSet(vvv))
				} else {
					attrs = append(attrs, temporal.NewSearchAttributeKeyString(k).ValueSet(vvv))
				}
			}
		case "keyword_list", "list":
			for k, vv := range v {
				vvv, ok := vv.([]any)
				if !ok {
					return temporal.SearchAttributes{}, fmt.Errorf("expected []any, got %T", vv)
				}
				var x []string
				for _, vvvv := range vvv {
					xx, ok := vvvv.(string)
					if !ok {
						return temporal.SearchAttributes{}, fmt.Errorf("expected string, got %T", vvvv)
					}
					x = append(x, xx)
				}
				attrs = append(attrs, temporal.NewSearchAttributeKeyKeywordList(k).ValueSet(x))
			}
		case "time":
			for k, vv := range v {
				var x time.Time
				switch vvv := vv.(type) {
				case time.Time:
					x = vvv
				case string:
					x, err = time.Parse(time.RFC3339, vvv)
					if err != nil {
						return temporal.SearchAttributes{}, fmt.Errorf("error parsing string to time.Time: %w", err)
					}
				default:
					return temporal.SearchAttributes{}, fmt.Errorf("expected time.Time or string, got %T", vvv)
				}
				attrs = append(attrs, temporal.NewSearchAttributeKeyTime(k).ValueSet(x))
			}
		default:
			return temporal.SearchAttributes{}, fmt.Errorf("unknown search attribute type: %s", t)
		}
	}
	return temporal.NewSearchAttributes(attrs...), nil
}
