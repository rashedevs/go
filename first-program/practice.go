package main

import "fmt"

// func modify(arr []int){

// arr = append(arr, 100)
// fmt.Println("inside modify",arr)
// // fmt.Printf("%p\n",arr[0])
// // fmt.Printf("%p\n",&arr)
// fmt.Printf("Address of slice header(ptr, len, cap): %p\n", &arr)
// fmt.Printf("Address of underlying array (data[0]): %p\n", &arr[0])
// fmt.Println(len(arr))
// fmt.Println(cap(arr))

// }

// func main(){

// // data := []int{1,2,3}
// data := make([]int, 3, 5)
// data[2] = 99
// fmt.Println("inside main",data)
// // fmt.Println(&data)
// fmt.Printf("Address of main slice header(ptr, len, cap): %p\n", &data)
// fmt.Printf("Address of main underlying array (data[0]): %p\n", &data[0])
// fmt.Println(len(data))
// fmt.Println(cap(data))

// modify(data)
// }



func main() {
	a := 6*2
	a++
	var b bool = len("Hello, গো") == a
	fmt.Println(b)
}



// sl := make([]int, 3, 5)
// sl[0]=5
// sl[1]=15
// sl[2]=25
// sl[3]=35   // >>> runtime err but append works why?

// sl = append(sl, 35)


func init(){
	fmt.Println("testing...")
}