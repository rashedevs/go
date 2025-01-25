package main

import (
	"fmt"

	"example.com/mathlib"
)

var c = 40
var d = 20
func main() {

fmt.Println("Starting....")

res := mathlib.Summation(c , d)
jogfol := jog(c, d)

fmt.Println("Result is from another package and parallel packages are", res, "and", jogfol)

}