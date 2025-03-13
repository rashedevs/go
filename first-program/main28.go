package main

import "fmt"


type User struct{
	Name string                 //member variable or property
	Age int
}

// normal function to print user details

// func printUserDetails(usr User){
// 	fmt.Println("Name:", usr.Name)
// 	fmt.Println("Age:", usr.Age)
// }

// receiver function is a function that is associated with a type
func (usr User) printDetails(otherParam string){
	fmt.Println(otherParam)
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func (usr User) call(a int){
	fmt.Println("Amar jonmo", a)
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

func main() {

var user1 User
user1 = User{
	Name: "rashed",      
	Age: 31,
}

// printUserDetails(user1)
user1.printDetails("Hello, I am a parameter for receiver function for user1 instance")

user2 := User{                                     
	Name: "uzzaman",    
	Age: 32,
}

// printUserDetails(user2)
user2.printDetails("Hello, I am a parameter for receiver function for user2 instance")

user2.call(1994)

}