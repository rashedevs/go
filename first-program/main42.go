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
    // var x int = 10
	// fmt.Println("hello", x)

	// go printHello(1)

	// go printHello(2)p[10]

	// go printHello(3)

	// go printHello(4)

	// go printHello(5)

	// fmt.Println(a," ",pi)

	// time.Sleep(5*time.Second)

    // fmt.Println("haha")

	p := []int{1,2,3,4,5}
	p = append(p, 6)
	p = append(p, 7)
	q := p[4:]
	q = append(q, 11)
	q = append(q, 12)
	q = append(q, 13)
	q = append(q, 14)

fmt.Printf("ok %p",&p)
fmt.Println(len(p))
fmt.Println(q)
fmt.Println(cap(q))
fmt.Println(p)

fmt.Printf("okp %p\n",&p)
fmt.Printf("okq %p",&q)
}