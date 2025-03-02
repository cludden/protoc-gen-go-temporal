package convert

import (
	"cmp"
	"iter"
	"slices"
)

func IterMapDeterministic[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return func(yield func(K, V) bool) {
		for _, k := range keys {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
