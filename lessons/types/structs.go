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
// The fields of the struct can be accessed using the dot operator.
func sayHello(p Person) {
	fmt.Println("Hello, " + p.Name)
}

// Like with other types, a struct is copied when passed to a function.
func birthday(p Person) {
	p.Age++
	fmt.Println(p.Name + " is now " + fmt.Sprint(p.Age))
}

func main() {
	var Mitchell = Person{
		Name: "Mitchell",
		Age:  29,
	}

	sayHello(Mitchell)
	birthday(Mitchell)
	fmt.Println(Mitchell.Name + " is still " + fmt.Sprint(Mitchell.Age))

	// Copying a struct copies all of the fields.
	var Josh = Mitchell
	Josh.Name = "Josh"
	Josh.Age = 30
	fmt.Printf("Mitchell: %v\n", Mitchell)
	fmt.Printf("Josh: %v\n", Josh)

	// Like other types, a struct can be created and used on the spot, without setting to a variable.
	sayHello(Person{
		Name: "Josh",
		Age:  30,
	})

}
