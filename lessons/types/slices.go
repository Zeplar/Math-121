package main

import (
	"fmt"
)

func main() {
	// arrayExample()
	// sliceExample()
	// mapSlice()
	// reduceSlice()
}

// Array

func arrayExample() {
	// An array is a fixed-size collection of elements of the same type.
	var _ [5]int = [5]int{1, 2, 3, 4, 5}

	// An array's size is part of its type, so [3]int and [4]int are different types.
	// var _ = [3]int{1, 2, 3, 4} // illegal

	var implicit0 = [3]int{1, 2} // the third element is implicitly 0
	fmt.Println(implicit0[2])

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	a[0] = 100
	fmt.Println(a)

	b := a // b is a copy of a; changes to b do not affect a
	b[1] = 200
	fmt.Println(a, b)
}

func sliceExample() {
	// A slice is a dynamically-sized array. It is actually a view into an underlying fixed-size array.

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)

	s[0] = 100
	fmt.Println(s)

	var t = s // t and s share the same underlying array; changes to one affect the other
	t[1] = 200
	fmt.Println(s, t)

	// Slicing. The underlying array is the same, we're just seeing less of it.
	s = s[2:]
	fmt.Println(s)
	s[0] = 300
	fmt.Println(s, t) // t is also affected, but not in the same place

	// We can slice from either end
	s = s[:len(s)-1]
	fmt.Println(s)

	// Slices are less safe than arrays-- you can go out of bounds
	// _ = s[15]

	// Slices can be grown
	s = append(s, 5, 6, 7, 8)
	fmt.Println(s)
	// However, if the window grows larger than the underlying array, it will be moved to a new array.
	s[1] = 400
	fmt.Println(s, t) // t is not affected

	// A string is a slice of bytes. However, it is immutable.
	str := "Hello, World!"
	fmt.Println(str[0:5])

}

func mapSlice() {
	s := []string{"this", "is", "a", "sentence"}

	// map in place
	for i, word := range s {
		s[i] = word + "!"
	}
	fmt.Println(s)

	// map to a new slice
	t := []int{}
	for _, word := range s {
		t = append(t, len(word))
	}
	fmt.Println(t)
}

func reduceSlice() {
	s := []string{"This", "is", "a", "sentence."}

	var reduced string
	for _, word := range s {
		reduced += word + " "
	}
	fmt.Println(reduced)

	var reduce func([]string, string) string
	reduce = func(slice []string, out string) string {
		if len(slice) == 0 {
			return out
		}
		return reduce(slice[1:], out+slice[0]+" ")
	}

	fmt.Println(reduce(s, ""))
}
