package main

import "fmt"

// A struct is a collection of fields that can be of any type.
type Person struct {
	Name string
	Age  int
}

var Mitchell = Person{
	Name: "Mitchell",
	Age:  29,
}

// A struct can be used as a function parameter like any other type.

func sayHello(p Person) {
	fmt.Println("Hello, " + p.Name)
}

func main() {
	var Mitchell = Person{
		Name: "Mitchell",
		Age:  29,
	}

	sayHello(Mitchell)

	// Copying a struct copies all of the fields.
	var Josh = Mitchell
	Josh.Name = "Josh"
	Josh.Age = 30
	fmt.Printf("Mitchell: %v\n", Mitchell)
	fmt.Printf("Josh: %v\n", Josh)

	// Like other types, a struct can be created and used anonymously (without setting to a variable)
	sayHello(Person{
		Name: "Josh",
		Age:  30,
	})

}
