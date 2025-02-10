package main

import "fmt"
  
func main() {
	fmt.Println("hello world", name)
}

// to access and test the package level scope, wee need to run all the related files together.
// like this: go run main.go add.go