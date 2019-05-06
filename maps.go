package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// // Original Version
// func WordCount(s string) map[string]int {
// 	f := strings.Fields(s)
// 	m := make(map[string]int, len(f))
// 	for _, word := range f {
// 		if _, ok := m[word]; ok == false {
// 			m[word] = 1
// 		} else {
// 			m[word]++
// 		}
// 	}
// 	return m
// }

// Refactor Version
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
