package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	m := make(map[string]int)
	wordsArr := strings.Fields(str)
	sortStr := make([]string, 0, 20)
	for _, v := range wordsArr {
		m[v]++
	}
	type keyValue struct {
		Key   string
		Value int
	}
	sortedStruct := make([]keyValue, 0, 20)
	for key, value := range m {
		sortedStruct = append(sortedStruct, keyValue{key, value})
	}
	sort.Slice(sortedStruct, func(i, j int) bool {
		if sortedStruct[i].Value == sortedStruct[j].Value {
			return sortedStruct[i].Key < sortedStruct[j].Key
		}
		return sortedStruct[i].Value > sortedStruct[j].Value
	})
	for i, keyValue := range sortedStruct {
		if i > 9 {
			break
		}
		sortStr = append(sortStr, keyValue.Key)
	}
	return sortStr
}
