package main

import "fmt"

var s string = "package scope"

func main() {
	var s string = "function scope"
	fmt.Println(s)

	shadowFunction()

	angryFunction("I'm happy")

	fmt.Println(s)
}

func shadowFunction() {
	var s string = "helper function scope"
	fmt.Println(s)
	anotherHelperFunction()

}

func anotherHelperFunction() {
	fmt.Println(s)
}

func angryFunction(s string) {
	fmt.Println(s)

	s = "I'm angry"
}
