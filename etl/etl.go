package etl

import "strings"

// Function Transform transforms the given map from key, [values] to value, key
func Transform(old_map map[int][]string) map[string]int {
	result := make(map[string]int)
	for k, v := range old_map {
		for _, c := range v {
			c = strings.ToLower(c)
			result[c] = k
		}
	}
	return result
}
