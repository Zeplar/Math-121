package main

import "fmt"

func main() {

	fmt.Printf("This is printf. It takes a string with any number of verbs starting with %%, and arguments matching those verbs")
	fmt.Println()
	fmt.Println()

	fmt.Printf("%s and %q\n", "This is a string", "This is a string with the quotes")

	fmt.Printf("This is a number: %d\n", 100)

	fmt.Printf("This is a float: %f\n", 3.14)
	fmt.Printf("This is a float with 2 decimal places: %.2f\n", 3.14)

	fmt.Printf("This is a boolean: %t\n", true)

	fmt.Printf("This can be anything: %v, %v, %v\n", "Hello World", 3, true)

	fmt.Printf("\n\n\n")

	fmt.Printf("Scanf is the opposite of printf. It has the same \"%%\" verbs, but it scans the string into the arguments instead of printing the arguments into the string.\n")
	fmt.Println()
	var s string
	fmt.Print("Give me a string: ")
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nYou gave me %q\n", s)
	fmt.Println()
}
