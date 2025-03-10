package main

import "fmt"


type User struct{
	Name string
	Age int
}

func printUserDetails(usr User){
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}


func main() {

var user1 User
user1 = User{
	Name: "rashed",      
	Age: 31,
}

printUserDetails(user1)


user2 := User{                                     
	Name: "uzzaman",    
	Age: 32,
}

printUserDetails(user2)

}