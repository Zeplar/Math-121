package main

import (
	"fmt"
)

func input() string {
	var i string
	fmt.Scanln(&i)
	return i
}

func inputN() int {
	var n int
	fmt.Scanln(&n)
	return n
}

// IsPrime returns true if the argument is prime.
func IsPrime(x int) bool {
	c := 2
	for {
		if c >= x {
			break
		}
		if (x % c) == 0 {
			return false
		} else {
			c++
		}
	}
	return true
}

// SumCubes returns the sum of the cubes of the first n numbers.
func SumCubes(x int) int {
	sum := 0
	for i := 1; i <= x; i++ {
		sum += (i * i * i)
	}
	return sum
}

// Fibonacci returns the nth Fibonacci3 number.
// 1) Clear answer cache
// 2) Relocate first and second most recent numbers
// 3) Add them together
// 4) Relabel first and second most numbers
// 5) Repeat # of times
func Fibonacci3(x int) int {
	v1 := 1
	v2 := 1
	v3 := 1
	fib := 1
	for i := 3; i < x; i++ {
		fib = v1 + v2 + v3
		v1 = v2
		v2 = v3
		v3 = fib
	}
	return fib
}

func poly1(x int) int {
	return x*x*x*x - 8*x*x*x + 6*x - 4
}

func poly2(x int, y int) int {
	return -(x * x * x * x) + 3*(x*x) - (y * y * y * y) + 5*(y*y)
}

func Polymin(x int) int {
	i := 1
	lowY := poly1(0)
	lowX := 0
	for i <= 10 {
		if poly1(i) < lowY {
			lowX = i
			lowY = poly1(i)
		}
		i++
	}
	fmt.Println(lowX)
	return lowX
}

// Polymax returns the maximum of a polynomial function -x^4 + 3x^2 -y^4 + 5y^2.
// 1) Establish counter at the start of a range, taking in 4 integers representing the minX, maxX, minY, maxY
// 2) Set maxZ to be equal to a poly(x & y within the range)
// 3) Everytime there is a higher poly(i), replace the current maxZ
// 4) Increase i by 1, repeat
func Polymax(minX int, maxX int, minY int, maxY int) (int, int, int) {
	i := minX
	c := minY
	var x int
	var y int
	maxZ := poly2(minX, minY)
	for c <= maxY {
		for i <= maxX {
			if maxZ < poly2(i, c) {
				maxZ = poly2(i, c)
				x = i
				y = c
			}
			i++
		}
		i = minX
		c++
	}
	return x, y, maxZ
}

// Collatz returns the number of steps it takes to reach 1 under Collatz's conjecture.

func Collatz(n int) int {
	i := 1
	var x int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		i++
		x = i
	}
	return x
}

func main() {
	fmt.Println(Polymax(0, 5, 0, 10))
	fmt.Println(Polymax(0, 0, 0, 0))

}
