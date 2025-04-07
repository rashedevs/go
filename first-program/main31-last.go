package main

import "fmt"

// first example - (see before : 1.55 min)...........
/*

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

*/

// second example - (see video : 1.55 min)..........
/*
func changeSlice(p []int) []int{
    p[0]=10
	p = append(p, 11)
	return p
}

func main(){

    x := []int{1,2,3,4,5}
    x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)      // [1,2,3,4,10,6,7] len=7, cap=10
    fmt.Println(x[0:8]) // [1,2,3,4,10,6,7,11] - (by force printing the next element of x)
	fmt.Println(y)      // [10,6,7,11]

}
*/

// variadic function..........(video : 02.13 min)
// a function that can receives unknown amount of same type params
// we can't use a normal function bcoz there will be fixed number of predefined params

func printLotsNumbers(numbers ...int){  // numbers will be recieved as a slice
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main(){
	printLotsNumbers(5,6,7,8,9)
}