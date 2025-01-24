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

fmt.Println("Result is from another package", res)

}