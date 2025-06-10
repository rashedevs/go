// Defer i dont understand

package main

import "fmt"

// func a(){
// 	i := 0
// 	fmt.Println("first",i)
// 	defer fmt.Println("second",i)
// 	i++
// 	fmt.Println("third",i)
// 	defer fmt.Println("fourth", i)
// 	return
// }


func calculate() (result int){  // named return value
    fmt.Println("first", result)

	show := func(){
		result = result+10
		fmt.Println("defer", result)
	}
   
	defer show()
	// defer fmt.Println("defer", result)

	result = 5
	// fmt.Println("second", result)

	// return

	p := func(r int){
		fmt.Println("this is r", r)
	}

	defer p(result)
	defer fmt.Println(result)

	fmt.Println("second", result)

	defer fmt.Println(5)
	return
}

// first 0, second 5, 5, 5, this is r 5, defer 15
// there is a cell inside stack frame called defer list pointer

/*
--- rules for - named return value ---

1. all codes execude
2. defer functions stored at magiic box in LIFO order
3. before return -> all defer function will execude
4. will return all named return variables values

*/


func calc() int{  

	result := 0
    fmt.Println("calc first", result)

	show := func(){
		result = result+10
		fmt.Println("calc defer", result)
	}
   
	defer show()
	// defer fmt.Println("defer", result)

	result = 5
	fmt.Println("calc second", result)

	return result
}

/*
--- rules for - just return types ---

1. all codes execude
2. defer functions stored at magiic box in LIFO order
3. before return - before execute defers -> all return values at this current stage will be evaluated
   (did store the current return values)....
4. all defer function will be execuded
4. will return all named return variables values

*/

func main(){
	a := calculate()
	fmt.Println(a)	// named return and output is 15

	b := calc()
	fmt.Println(b)	 // just return types and output is 5
}


/*
-- Defer uses Linked list but executes such like stack.
-- Every defer nodes consist of a value and next portion.
-- Defer info can be stored at both stack and heap. Most of the time at heap.
-- Defer value info got binds with variable and function pointer if closure forms. 
   Otherwise value and next. 
*/