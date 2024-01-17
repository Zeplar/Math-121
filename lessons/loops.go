package main

import "fmt"

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

func main() {
	fmt.Println("Type something: ")
	n := inputN()
	fmt.Println("You typed", n)
}
