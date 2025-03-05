//class 19 [function types] : (viii) closure - close over

package main

import "fmt"

const a = 10  // constant

var p = 100

func outer() func(){
	money := 1000

	age := 30
	fmt.Println("age :", age)

	show := func(){
		money = money + a + p
		fmt.Println("money :", money)
	}
	return show
}

func call (){
	increment1 := outer()
	increment1()
	increment1()
	increment1()

    increment2 := outer()
	increment2()
	increment2()
}

func main(){
	fmt.Println("printing....", p)
	call()	
}


func init () {
	fmt.Println("Hello, init")
}