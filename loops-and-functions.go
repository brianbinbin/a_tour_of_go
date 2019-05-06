package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(z*z-x) > delta {
		z = (z + x/z) / 2
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
