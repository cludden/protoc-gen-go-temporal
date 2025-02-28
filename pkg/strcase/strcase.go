package strcase

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	lower int = iota
	upper
	number
	separator
)

// Converter holds the configuration for acronyms
type (
	Caser struct {
		acronyms map[string]struct{}
	}

	CaserOption func(*Caser)
)

func WithAcronyms(acronyms ...string) CaserOption {
	return func(c *Caser) {
		for _, a := range acronyms {
			c.acronyms[a] = struct{}{}
		}
	}
}

// NewCaser creates a new Converter with given acronyms
func NewCaser(options ...CaserOption) *Caser {
	c := &Caser{
		acronyms: make(map[string]struct{}),
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

// ToCamel converts a string to CamelCase
func (c *Caser) ToCamel(s string) string {
	return c.toCamel(s, true)
}

// ToKebab converts a string to kebab-case
func (c *Caser) ToKebab(s string) string {
	words := c.splitWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "-")
}

// ToLowerCamel converts a string to lowerCamelCase
func (c *Caser) ToLowerCamel(s string) string {
	return c.toCamel(s, false)
}

// toCamel converts a string to CamelCase or lowerCamelCase
func (c *Caser) toCamel(s string, upperFirst bool) string {
	words := c.splitWords(s)
	if len(words) == 0 {
		return ""
	}

	for i, word := range words {
		if _, ok := c.acronyms[word]; !ok {
			words[i] = cases.Title(language.English).String(strings.ToLower(word))
		}
	}

	if !upperFirst {
		words[0] = strings.ToLower(words[0])
	}

	return strings.Join(words, "")
}

// splitWords splits a string into words, discarding all separators and honoring
// case change and acronym boundaries.
func (c *Caser) splitWords(s string) []string {
	words := make([]string, 0)
	runes := []rune(s)
	var w strings.Builder

CHARS:
	for i := 0; i < len(runes); i++ {
		for a := range c.acronyms {
			if i < len(runes)-1 && strings.HasPrefix(s[i:], a) {
				if w.Len() > 0 {
					words = append(words, w.String())
					w.Reset()
				}
				words = append(words, a)
				i += len(a) - 1
				continue CHARS
			}
		}

		a := classify(runes[i])
		if a == separator {
			if w.Len() > 0 {
				words = append(words, w.String())
				w.Reset()
			}
			continue
		}
		w.WriteRune(runes[i])

		if i < len(runes)-1 {
			b := classify(runes[i+1])
			if a != b && (a == lower || a == number) {
				if w.Len() > 0 {
					words = append(words, w.String())
					w.Reset()
				}
			}
		}
	}

	if w.Len() > 0 {
		words = append(words, w.String())
		w.Reset()
	}
	return words
}

// classify classifies a rune into one of four categories: lower case, upper
// case, number, or separator.
func classify(r rune) int {
	switch {
	case isLower(r):
		return lower
	case isUpper(r):
		return upper
	case isNumber(r):
		return number
	default:
		return separator
	}
}

// isLower returns true if the rune is a lower case letter.
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// isUpper returns true if the rune is an upper case letter.
func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// isNumber returns true if the rune is a number.
func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
