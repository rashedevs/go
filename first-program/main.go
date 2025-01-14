package main

import (
	"fmt"
)

// "first-program/practice"
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


func add(num1 int, num2 int) int {
	sum := num1 + num2
	// fmt.Println("Sum of", num1, "and", num2, "is", sum)
	return sum
}

// func allOps(num1 int, num2 int) (int, int) {
// 	sum := num1 + num2
// 	mul := num1 * num2
// 	// fmt.Println("Sum of", num1, "and", num2, "is", sum)
// 	return sum, mul
// }

// func myFunction(x int, y string) (result int, txt1 string) {
// 	result = x + x
// 	txt1 = y + " World!"
// 	return
//   }


//if-else or switch................................................................
// func main(){

// age :=20
// a :=9

//if-else
// if age >18{
// 	fmt.Println("You are an adult")
// } else if age < 18 {
// 	fmt.Println("You are a minor")
// } else {
// 	fmt.Println("You are 18 so wait...")
// }

//switch
// switch a {
// case 1, 3, 5:
//     fmt.Println("a is 1 or 3 or 5")
// case 2, 4:
//     fmt.Println("a is either 2 or 4")
// case 6, 7:
//     fmt.Println("a is either 6 or 7")
// default:
//     fmt.Println("a is for something else")
// }
// var gunfol  = practice.Gun(30, 30)
// var biyogfol = practice.Biyog(300, 120)



// class-10 (functions):

// x :=11
// y :=33

// result := add(x,y)

// re1, res2 := allOps(x, y)

// d, e := myFunction(10, "Hello")
// _, z := myFunction(5, "Hi") //ommit y returned value


// fmt.Println("result: ", result, re1, res2)
// fmt.Println(d,e)
// fmt.Println(z)

//applying S of SOLID................................

func printWelcomeMessage(){
	fmt.Println("Welcome to my application")
}

func getUserName() string { 
	var userName string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&userName)
	return userName
}

//take input from user
func main(){
// fmt.Println("Welcome to my application")
printWelcomeMessage()

// var userName string
// fmt.Println("Enter your name: ")
// fmt.Scanln(&userName)
// fmt.Println("your name is -----", userName)
name := getUserName()

var numOne int
var numTwo int
fmt.Println("Enter number one: ")
fmt.Scanln(&numOne)

fmt.Println("Enter number two: ")
fmt.Scanln(&numTwo)  //by &numTwo we are assigning numTwo to the variable

sum := add(numOne, numTwo)
fmt.Println("Hi",name,"Your Sum of two input numbers is: ", sum)

fmt.Println("Thank you for using my application")
fmt.Println("Good bye")
}

// SOLID >>> S = Single Responsibility Principle >>> will apply to upper main function

