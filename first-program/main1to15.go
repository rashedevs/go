package main

// import (
// 	"fmt"
// )

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

// func add(num1 int, num2 int) int {
// 	sum := num1 + num2
// 	// fmt.Println("Sum of", num1, "and", num2, "is", sum)
// 	return sum
// }

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

// func printWelcomeMessage(){
// 	fmt.Println("Welcome to my application")
// }

//take input from user
// func getUserName() string {
// 	var userName string
// 	fmt.Println("Enter your name: ")
// 	fmt.Scanln(&userName)
// 	return userName
// }

// func getTwoNumbers() (int, int) {
// 	var numOne int
// 	var numTwo int
// 	fmt.Println("Enter number one: ")
// 	fmt.Scanln(&numOne)

// 	fmt.Println("Enter number two: ")
// 	fmt.Scanln(&numTwo)  //by &numTwo we are assigning numTwo to the variable
// 	return numOne, numTwo
// }

// func display(name string, sum int) {
// 	fmt.Println("Hi",name,"!","Sum of two input numbers is: ", sum)
// }

// func printGoodbyeMessage() {
// 	fmt.Println("Thank you for using my application")
//     fmt.Println("Good bye")
// }

//recursion
// func testcount(x int) int {
// 	if x == 11 {
// 	  return 0
// 	}
// 	fmt.Println(x)
// 	return testcount(x + 1)
//   }

// func main(){
// // SOLID >>> S = Single Responsibility Principle >>> will apply to main function
// printWelcomeMessage()
// name := getUserName()
// n1,n2 := getTwoNumbers()
// sum := add(n1, n2)
// display(name, sum)
// printGoodbyeMessage()

// testcount(5)

// }

//class - 14 >>> scope

// var a = 20
// var b = 30
// func add(x int , y int){
// 	z := x + y
// 	fmt.Println(z)
// }
// func main() {
// 	p := 30
// 	q := 40
// 	add(p, q)
// 	add(a, b)
// 	add(a, p)

// 	// add(b, z)
// }

//class -15 >>> local scope for if-else, function, switch

// var a = 20
// var b = 30

// func add1to15(a, b int) int {
//     result := a + b // Local to the `add` function
//     return result
// }
// func main1to15() {
// 	x := 20
// 	if x >= 18 {
// 		p := 10
//         fmt.Println("I am adult")
// 		fmt.Println("I have", p , "girlfriend")
//     }
// 	// fmt.Println("hi", p)

// // more examples for if-else, function, switch
// //>> 1. if-else................................................................
// if true {
// 	x := 10 // Local to this `if` block
// 	fmt.Println("x in if block:", x)
// }
// // fmt.Println("x outside if block:", x) // Error: x is not accessible here

// if false {
// 	y := 20 // Local to this `if` block
// 	fmt.Println("y in if block:", y)
// } else {
// 	z := 30 // Local to this `else` block
// 	fmt.Println("z in else block:", z)
// }
// // fmt.Println("z outside else block:", z) // Error: z is not accessible

// //>> 2. function................................................................

// sum := add1to15(5, 7)
//     fmt.Println("Sum:", sum)
//     // fmt.Println(result) // Error: `result` is not accessible in `main`

// //>> 3. switch................................................................

// num := 2

// switch num {
// case 1:
// 	x := "One" // Local to this `case`
// 	fmt.Println(x)
// case 2:
// 	y := "Two" // Local to this `case`
// 	fmt.Println(y)
// default:
// 	z := "Unknown" // Local to the `default` case
// 	fmt.Println(z)
// }

// // fmt.Println(x, y, z) // Error: x, y, z are not accessible outside `switch`
// }

//..................................NOTE..............................
/*

1. What is Package Scope in Go?

In Go, package scope means that any variable, constant, function, or type declared
at the top level of a file is accessible throughout the same package.

- Each folder should contain same package and parallel file should contain same package.

- These items can be shared among all files in the same package.

- To make items available outside the package,
  their names must start with an uppercase letter (exported).

Key Characteristics of Package Scope:

- Items declared at the top level of a file are available throughout the entire package.
- Only exported items (uppercase names) can be accessed outside the package.

2. Why Go Doesn't Have True Global Scope ?

  Go avoids true global variables (accessible everywhere across all packages)
  to encourage modularity and prevent name clashes. Instead, Go uses:

  A. Package Scope: For sharing data across the same package.
  B. Exported Items: For sharing specific data/functions across packages.

In other languages, a global variable might be accessible everywhere.
In Go, such variables would be scoped to a package and need to be explicitly exported.


## Key Takeaways

- Package scope is Go's way of creating "global" variables but limits them
  to the package's boundary.

- There are no "true global" variables accessible across all packages in Go.

- This design enforces encapsulation, modularity, and minimizes
  unintended interference between packages.

*/