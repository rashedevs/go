package main

import "fmt"

func main(){
	// pointer or address of memory (ram)
	x := 28
    
	// p holds the address of x
	p := &x //ampersand x = address of x || 0xc000010170 (hex) ==> 211106249404416
	fmt.Println("p holds the address of x variable =", p)
	fmt.Println("value (of x) at the p address =", *p)   // * => value at address
	fmt.Println("in short",*&x)
}