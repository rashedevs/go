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


var (
	a1=20
	b1=30
)

func printNum(n int) {
	fmt.Println("Number is:", n)
}

func add1(x int, y int) {
    res1 := x + y
    printNum(res1)
}
func main(){
	add(a1, b1)
}