package main

import (
	"fmt"
	"math"
)

// Given the polynomial x^4-8x^3+6x-4, what integer value of x between 0 and 10 results in the lowest value?
var poly func(x int) int

func polyMin(xMin int, xMax int, poly func(x int) int) int {
	var min int = math.MaxInt
	var indexOfMin = xMin
	for i := xMin; i < xMax; i++ {
		if poly(i) < min {
			indexOfMin, min = i, poly(i)
		}
	}
	return indexOfMin
}

func main() {

	poly = func(x int) int {
		return x*x*x*x - 8*x*x*x + 6*x - 4
	}
	min := polyMin(0, 10, poly)
	fmt.Println("Min:", min)
}
