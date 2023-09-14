package main

import "fmt"

func main() {
	RangeMap()
}
func RangeMap() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		delete(m, "two")
		m["four"] = 4
		m["four1"] = 4
		m["four2"] = 4
		m["four3"] = 4
		m["four4"] = 4
		m["four5"] = 4
		fmt.Printf("%v: %v\n", k, v)
	}
}
