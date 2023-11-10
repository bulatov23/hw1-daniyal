package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	m := make(map[string]int)
	x := strings.Split(str, " ")
	var sort_str []string
	for _, v := range x {
		m[v] += 1
	}
	type key_value struct {
		Key   string
		Value int
	}
	var sorted_struct []key_value
	for key, value := range m {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}
	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value > sorted_struct[j].Value
	})
	for _, key_value := range sorted_struct {
		sort_str = append(sort_str, key_value.Key)
	}
	return sort_str
}
