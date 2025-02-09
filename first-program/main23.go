package main

import "fmt"

// Parameter vs Arguments.
/* First order function.
   - Standard or named function
   - Anonymous function
   - Immediately invoke function expression (IIFE)
   - Function expression
*/
/* Higher order function
   - A function that accepts other functions as arguments
   - Returns a function as a result
   - Both
*/
// First class citizen
// First class function (Higher order function)

func processOperation(x int, y int, op func(k int, l int)) (func (m int, n int)){

// parameter function call which is actually the shadowed add function
op(x, y)

// anonymous function expression which will be returned
mult := func(m int, n int){
	fmt.Println(m*n)
}
return mult
}

func add(a int, b int){
	c := a+b
	fmt.Println(c)
}
func main(){
	fmt.Println("starting...")
	result := processOperation(10, 20, add) // passing add function as argument
	// result is getting the mult function as a return and we can invoke it
	result(15, 20)
}