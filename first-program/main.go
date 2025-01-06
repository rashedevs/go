package main

import (
	"fmt"
)

//variable................................................................
// func main(){
// 	// Declare and initialize an int/float32/bool/string variable
// 	// var a int = 10
// 	// a :=10

// 	a :=false
// 	a =true
//     // a="habib"
//     const p=100
// 	fmt.Println("hello, world!",a,p)
// }

//if-else or switch................................................................
func main(){

age :=20
a :=3

//if-else
if age >18{
	fmt.Println("You are an adult")
} else if age < 18 {
	fmt.Println("You are a minor")
} else {
	fmt.Println("You are 18 so wait...")
}

//switch
switch a {
case 1:
    fmt.Println("a is 1")
case 2, 3:
    fmt.Println("a is either 2 or 3")
default:
    fmt.Println("a is not 1 or 2")
}
}