package main

import (
	"fmt"
	"time"
)

//goroutine in complex way

var a = 10
const pi = 3.1416

// func add(a,b int){
// 	fmt.Println(a+b)
// }

func printHello(num int){
	time.Sleep(5*time.Second)
	fmt.Println("Hello rashed",num)
	// add(4,5)
}

func main() {

slc := []int{}

slc = append(slc, 10)

fmt.Print(cap(slc))

slc = append(slc, 13, 14)

fmt.Print(cap(slc))

slc = append(slc, 16)

fmt.Print(cap(slc))
}