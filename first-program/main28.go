package main

import "fmt"


type User struct{
	Name string
	Age int
}

// normal function to print user details

// func printUserDetails(usr User){
// 	fmt.Println("Name:", usr.Name)
// 	fmt.Println("Age:", usr.Age)
// }


func (usr User) printUserDetails(otherParam string){
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
user1.printUserDetails("Hello, i am a parameter for receiver function")

user2 := User{                                     
	Name: "uzzaman",    
	Age: 32,
}

// printUserDetails(user2)
user2.printUserDetails("Hello, i am a parameter for receiver function")

}