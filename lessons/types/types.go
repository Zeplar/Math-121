package main

import "fmt"

// Primitive types
var _ int = 5
var _ int32 = 0
var _ float64 = 0.25343
var _ bool = false
var _ string = "Hello, World!" // Not primitive, but treated as such

// Function types
var _ func() = func() {
	fmt.Println("Hello, World!")
}
var _ func(int) int = func(i int) int {
	return i + 1
}

// Structs

// Arrays

// Slices
