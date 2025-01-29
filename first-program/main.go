package main

import (
	"fmt"
	// "example.com/mathlib"
)

//---------------------------------------for class 16

// var c = 40
// var d = 20
// func main() {

// fmt.Println("Starting....")

// res := mathlib.Summation(c , d)
// jogfol := jog(c, d)

// fmt.Println("Result is from another package and parallel packages are", res, "and", jogfol)

// }

// --------------------------------------for class 17
// functions order ups/down isn't a problem here to access.....

// var (
// 	a1=20
// 	b1=30
// )

// func printNum(n int) {
// 	fmt.Println("Number is:", n)
// }

// func add1(x int, y int) {
//     res1 := x + y
//     printNum(res1)
// }
// func main(){
// 	add(a1, b1)
// }

// --------------------------------------for class 18
// variable shadowing.................................

// func main() {
// 	x := 20
//     if x > 18 {
//         x := 10
//         fmt.Println("Local x is after redeclared:", x)   // print 10
//     }
//     fmt.Println("Global x is:", x)                       // print 20
// }

//--------------------------------------for class 19
// standard or named functions --------------------------------

var boyos = 30

//standard function
func mul(x int, y int) int {
    return x * y
}

func main() {
    fmt.Println("Result is:", mul(10, 20))
    fmt.Println("Now boyos is:", boyos)
}

//--------------------------------------for class 20
// init function - (we cant call this, computer calls it automatically)
// no arguments - no return
// It runs automatically before the main function (main) or any other functions in the package.
// init functions in imported packages are executed before the init function in the main package.
// A single package can have multiple init functions, even in different files within the same package. 

func init(){
fmt.Println("I am the init function and will be executed first-- boyos silo", boyos);
boyos = 31
}

