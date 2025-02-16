// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello world", name)
// }

// to access and test the package level scope, wee need to run all the related files together.
// like this: go run main.go add.go

// practice test
// summation of nth number
package main

import (
	"fmt"
	"time"
)

func Sum(x int) int {
    sum := x*(x+1)/2
    return sum
}

func main() {
    x := 100
    start := time.Now()
    result := Sum(x)
    elapsed := time.Since(start)
    fmt.Println("Sum of first", x, "numbers:", result)
    fmt.Println("Execution Time:", elapsed)
}