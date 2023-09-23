package tools

import "sort"

func SortKey(key string) string {
	runes := []rune(key)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
