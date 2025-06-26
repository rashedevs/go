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

func main(){
    var x int = 10
	fmt.Println("hello", x)

	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a," ",pi)

	time.Sleep(5*time.Second)

    fmt.Println("haha")

}