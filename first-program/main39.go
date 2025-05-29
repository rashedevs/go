package main

import "fmt"

// vogus data types........

func main(){

	var a int = 58       // can be 32/64 bit based on pc config
	var b int8 = 127     // range -128 to 127 
	var c uint8 = 255    // can store 0 to 255

	var red float32 = 12.1510
	var ash float64 = 15.1018

	var flag bool = true   // bool got stored as 8 bit / 1 byte
    var name string = "redash"

	flower := '💐'

	fmt.Printf("%c\n", flower)  // c for character / rune type
	fmt.Printf("%f\n", red)     // f for floating number
	fmt.Printf("%.3f\n", ash)   // .3 to choose 3 numbers after point
	fmt.Printf("%v\n", flag)    // v for flag/bool type
	fmt.Printf("%s\n", name)    // s for string type

	fmt.Printf("type of rune/flower = %T\n", flower) // check type of a variable using %T
	fmt.Printf("type of name = %T\n", name)

	fmt.Println(a,b,c, red, ash, flag)
}
