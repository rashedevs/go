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

// var boyos = 30

//standard / named function
// func mul(x int, y int) int {
//     return x * y
// }

// func main() {
//     fmt.Println("Result is:", mul(10, 20))
//     fmt.Println("Now boyos is:", boyos)
// }

//--------------------------------------for class 20
// init function - (we cant call this, computer calls it automatically)
// no arguments - no return
// It runs automatically before the main function (main) or any other functions in the package.
// init functions in imported packages are executed before the init function in the main package.
// A single package can have multiple init functions, even in different files within the same package.

// func init(){
// fmt.Println("I am the init function and will be executed first-- boyos silo", boyos);
// boyos = 31
// }

// ----------------------------------------class 21--------------------------------
//-------------anonymous function & IIFE - imediately invoked function expression---------------

// An anonymous function in Go is a function without a name. It can be assigned to a variable,
//used as an inline function, or passed as an argument to another function.

// viva topic (anonymous function / expression / iiFE / invoke)

var m = 33 // an expression
var n = 34 // another expression


func addition(x , y int){   //an expression can be a var, a function and if/else statement with block
    fmt.Println("Result is:", x + y)
}
func main() {

    
    fmt.Println("I am the main function.")
    addition(m, n)
    //anonymous function with iife
    func (c int, d int){ // anonymous function expression
        res := c * d
        fmt.Println("Multiplication result is:", res)
    }(20 , 40) // immediately invoked (called)



// ------------------------------------------------------clss 22------------------------------------------
// --------------------------------------function expression or assign function in variable--------------------------
       
//................................................................................................................................
    // Gives error bcoz
    // You can define a function inside main, but it must be an anonymous function 
    // assigned to a variable or invoked immediately.
    // However, you cannot define a named function (func sum(...) {}) inside main
    // because Go does not allow function declarations inside other functions.

    // func sum (x, y int){
    //     fmt.Println("Sum is:", x + y)
    // }
    // sum(30,40)

// Points to be Noted: 

// ✅ Named functions must be declared at the package level (outside main).
// ✅ Anonymous functions can be used inside main, either by assigning them to a variable or using IIFE.
// ❌ You cannot declare a named function inside main or any other function.
//------------------------------------------------------------------------------------------------------------------------------------


    // div(m, n)   // gives undefined
    div := func(f, g int){
        fmt.Println(f/g)
    }

    div(40, 20)
}

func init(){
fmt.Println("I am the init function and will be executed first--")
}

