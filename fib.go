package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// // Original Version
// func fibonacci() func() int {
// 	a, b := 0, 1
// 	n := 0
// 	return func() int {
// 		for {
// 			if n == 0 {
// 				n++
// 				return 0
// 			} else if n == 1 {
// 				n++
// 				return 1
// 			} else {
// 				a, b = b, a+b
// 				return b
// 			}
// 		}
// 	}
// }

// Refactor Version
func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b := b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
