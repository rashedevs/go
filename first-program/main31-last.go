package main

import "fmt"

func main(){
	// Empty slice or nil slice
	var x []int                      // [], len=0, cap=0
	x = append(x, 1)                 // [1], len=1, cap=1
	x = append(x, 2)                 // [1,2], len=2, cap=2
	x = append(x, 3)                 // [1,2,3], len=3, cap=4

	// fmt.Println(len(x))
	// fmt.Println(cap(x))


	y := x

	x = append(x, 4)
	fmt.Println(x)

	y = append(y, 5)

	// x[0]=10

	fmt.Println(x)
	fmt.Println(y)
}