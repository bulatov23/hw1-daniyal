package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	m := make(map[string]int)
	x := strings.Fields(str)
	sort_str := make([]string, 0, 20)
	for _, v := range x {
		m[v] += 1
	}
	type key_value struct {
		Key   string
		Value int
	}
	sorted_struct := make([]key_value, 0, 20)
	for key, value := range m {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}
	sort.Slice(sorted_struct, func(i, j int) bool {
		if sorted_struct[i].Value == sorted_struct[j].Value {
			return sorted_struct[i].Key < sorted_struct[j].Key
		}
		return sorted_struct[i].Value > sorted_struct[j].Value
	})
	for i, key_value := range sorted_struct {
		if i > 9 {
			break
		}
		sort_str = append(sort_str, key_value.Key)
	}
	return sort_str
}
