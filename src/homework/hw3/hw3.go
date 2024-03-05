package main

import (
	"fmt"
)

type poly1 func(int) int
type poly2 func(int) bool

func max_of_two(fn1 poly1, fn2 poly1, x int) int {
	if fn1(x) > fn2(x) {
		return fn1(x)
	} else {
		return fn2(x)
	}
}

func max_of_funcs(fn1 poly1, fn2 poly1) poly1 {
	return func(x int) int {
		if fn1(x) > fn2(x) {
			return fn1(x)
		} else {
			return fn2(x)
		}
	}
}

func evenX(x int) bool {
	return x%2 == 0
}

var is_even = func(x int) bool {
	return x%2 == 0
}

func conditional_printchan(test func(int) bool) func(int) {
	return func(x int) {
		if test(x) {
			fmt.Println(x)
		}
	}
}

var print_evenchan = conditional_printchan(evenX)

func prints_evenchan() {
	conditional_printchan(evenX)
	f := conditional_printchan(evenX)
	fmt.Println(f)
}

func main() {

}

//Excercise 6:Im not sure why it would take one function "g" over the other,
//				just know that when I run it, it takes the first one, "x+3" and spits out a = 5 ; b = 6
