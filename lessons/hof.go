//higher order functions

package main

import "fmt"

var polynomial func(int) int

type poly1 func(int) int
type poly2 func(poly1) int

func inputN() int {
	var n int
	fmt.Scanln(&n)
	return n
}

func squarepolynomial(x int, fn poly1) int {
	return fn(x) * fn(x)
}

func plusN(n int) poly1 {
	return func(x int) int {
		return x + n
	}
}

func mk(x int) int {
	return x*x*x*x - 8*x*x*x + 6*x - 4
}

func polymin(Xmin int, Xmax int, fn poly1) (int, int) {
	i := Xmin
	var x int
	var y int
	maxZ := poly2(minX, minY)
	for c <= maxY {
		for i <= maxX {
			if maxZ < poly2(i, c) {
				maxZ = poly2(i, c)
				x, y = i, c
			}
			i++
		}n
		i = minX
		c++
	}
	return x, y, maxZ
}

var poly3 = func(Xmin int, Xmax int) (int, int) {
	return asdf(Xmin, Xmax, mk)
}

func main() {
	fmt.Println("input for plusN")
	f := plusN(5)
	fmt.Println(f(6))
	g := plusN(12)
	fmt.Println(g(44))
}

func asdf(Xmin int, Xmax int, fn poly1) (int, int) {
	x, Ymin := Xmin, fn(Xmin)
	for i := Xmin; i <= Xmax; i++ {
		if Ymin > fn(i) {
			x, Ymin = i, fn(i)
		}
	}
	return x, Ymin
}

//generalize polymin for any fucntion
