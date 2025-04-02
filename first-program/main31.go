package main

import "fmt"

func main(){
	arr := [6]string{"This","is","a","go","interview","question"}
    fmt.Println(arr)

	// 1. slice from an array
	s := arr[1:4]              // ["is","a","go"]
	fmt.Println(s)

	// 2. slice from a slice
	s1 := s[1:2]              // ["a"]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	// 3. Defining a new slice with slice literal
	slice := []int{1,2,5}     // slice literal - (an array without size)
	fmt.Println("slice",slice,"len",len(slice),"cap",cap(slice))

	// 4. using make function
	sliceMake := make([]int, 5)
	fmt.Println("slice with make function", sliceMake)
}