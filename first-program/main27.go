package main

import "fmt"

/*
A struct in Go is a way to group related data together.
It acts like a custom data type that holds multiple property of different types under one name.
*/

type User struct{
	Name string
	Age int
}

func main() {

var user1 User

//Instantiate means creating an instance of something, usually an object or a struct in programming.
// user1 is an instance of the User struct

// Instantiating (creating) a struct/object

user1 = User{
	Name: "rashed",      // called property or field or member-variable.
	Age: 31,
}

fmt.Println("Name-1:", user1.Name)
fmt.Println("Age-1:", user1.Age)

// Instantiating (creating) another struct/object

user2 := User{          // called instance or object. ----(while we use a custom type and                            
	Name: "uzzaman",    // assign value to it, it is called instance or object)
	Age: 32,
}

fmt.Println("Name-2:", user2.Name)
fmt.Println("Age-2:", user2.Age)

}


/*
In Go, a struct is a composite data type that groups together variables under a single name.
These variables, called fields/properties, can have different types. Structs are commonly used to 
represent real-world entities and define the structure of objects in Go.
*/