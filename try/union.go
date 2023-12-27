package main

import "fmt"

func main() {
	a := []interface{}{}
	b := []interface{}{}
	fmt.Println(Union(a, b))
}

// Union string for merge code_mapping @puyu @pulin
// jsonSchema 确保了子项为string
// 求并集
func Union(a, b []any) []interface{} {
	totalLen := len(a) + len(b)
	set := make(map[string]struct{}, totalLen)
	union := make([]interface{}, 0, totalLen)

	// Add all elements from 'a' to the set
	for _, item := range a {
		vs, ok := item.(string)
		if !ok {
			continue
		}
		if _, found := set[vs]; !found {
			set[vs] = struct{}{}
			union = append(union, item)
		}
	}

	// Add all elements from 'b' to the set if they're not already in it
	for _, item := range b {
		vs, ok := item.(string)
		if !ok {
			continue
		}
		if _, found := set[vs]; !found {
			set[vs] = struct{}{}
			union = append(union, item)
		}
	}

	return union
}
