package main

import (
	"first-program/practice"
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
a :=9

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
case 1, 3, 5:
    fmt.Println("a is 1 or 3 or 5")
case 2, 4:
    fmt.Println("a is either 2 or 4")
case 6, 7:
    fmt.Println("a is either 6 or 7")
default:
    fmt.Println("a is for something else")
}
var gunfol  = practice.Gun(30, 30)
var biyogfol = practice.Biyog(300, 120)

fmt.Print("gun----------", gunfol, "-------biyog--------", biyogfol)
}