package expression

import "errors"

func Lex(input string) (*Expression, error) {
	runes := []rune(input)
	expr := &Expression{}
	fragment := &Fragment{}
	inside := false
	i, length := 0, 0
	brackets := 0
	for {
		switch {
		// match end of input
		case i >= len(runes):
			if inside {
				return nil, errors.New("detected partial expression")
			}
			if length > 0 {
				expr.Fragments = append(expr.Fragments, fragment)
			}
			if len(expr.Fragments) == 0 {
				return nil, errors.New("empty expression")
			}
			return expr, nil
		// match query open
		case !inside && len(runes[i:]) > 3 && string(runes[i:i+3]) == "${!" && (i == 0 || runes[i-1] != '\\'):
			if length > 0 {
				expr.Fragments = append(expr.Fragments, fragment)
			}
			fragment = &Fragment{Expr: &Query{}}
			i, length = i+3, 0
			inside = true
		// match escape
		case runes[i] == '\\' && (i == 0 || runes[i-1] != '\\'):
			i++
		// match bracket open inside query
		case inside && runes[i] == '{' && runes[i-1] != '\\':
			fragment.Expr.Mapping += string(runes[i])
			brackets++
			i++
			length++
		// match bracket close inside query
		case inside && runes[i] == '}' && runes[i-1] != '\\' && brackets > 0:
			fragment.Expr.Mapping += string(runes[i])
			brackets--
			i++
			length++
		// match query close
		case inside && runes[i] == '}' && runes[i-1] != '\\':
			if length > 0 {
				expr.Fragments = append(expr.Fragments, fragment)
			}
			fragment = &Fragment{}
			i, length = i+1, 0
			inside = false
		// match mapping
		case inside:
			fragment.Expr.Mapping += string(runes[i])
			i++
			length++
		// match identity
		default:
			fragment.Ident += string(runes[i])
			i++
			length++
		}
	}
}
