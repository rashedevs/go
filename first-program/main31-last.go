package main

import "fmt"

func main(){
	// Empty slice or nil slice with append
	
	var x []int                      // [], len=0, cap=0
	x = append(x, 1)                 // [1], len=1, cap=1
	x = append(x, 2)                 // [1,2], len=2, cap=2
	x = append(x, 3)                 // [1,2,3], len=3, cap=4

	// fmt.Println(len(x))
	// fmt.Println(cap(x))


	y := x                          // present p=r29, l=3 and c=4 of x will be copied to y

	x = append(x, 4)                // [1,2,3,4], len=4, cap=4
	// fmt.Println(x)
	// fmt.Println(y)               // [1,2,3], len=3, cap=4


	y = append(y, 5)            // Heap array's 4th element (4) will be replaced with 5 = [1,2,3,5]
	x[0]=10                     // first index value (1) will be replaced with 10 = [10,2,3,5]

	fmt.Println(x)              // [10,2,3,5]
	fmt.Println(y)              // [10,2,3,5]
}