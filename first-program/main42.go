package main

import (
	"fmt"
	"time"
)

//goroutine in complex way

var a = 10
const pi = 3.1416

func printHello(num int){
	fmt.Println("Hello rashed",num)
}

func main(){
	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a," ",pi)


	time.Sleep(5*time.Second)
}