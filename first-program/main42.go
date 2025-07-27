package main

import "fmt"

// "time"

//goroutine in complex way

// var a = 10
// const pi = 3.1416

// func add(a,b int){
// 	fmt.Println(a+b)
// }

// func printHello(num int){
// 	time.Sleep(5*time.Second)
// 	fmt.Println("Hello rashed",num)
// 	// add(4,5)
// }


func main() {
	// var x []int          // [], ptr = nil, len=0, cap=0

	x := make([]int, 512)

	// x = append(x, 1)   
	// x = append(x, 2)
	// x = append(x, 3)
	// x = append(x, 4)
	// x = append(x, 5)
	// x = append(x, 6)
	// x = append(x, 7)
	// x = append(x, 8)
	// x = append(x, 9)

	

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))





	// for i := 1; i <= 1025; i++{
	// 	x = append(x, i)
	// 	fmt.Printf("After %d: len=%d cap=%d\n", i, len(x), cap(x))
	// }


	// x = append(x, 3)	
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// y := x
	// x = append(x, 4)	
	// y = append(x, 5)	

	// x[0] = 10	

	// fmt.Println(x)  
	// fmt.Println(y)   
	
	
	// var slice []int
    // prevCap := cap(slice)
    // fmt.Printf("%-10s | %-10s | %-6s | %-8s | %-10s\n", "Length", "Capacity", "Δ", "Growth", "× Increase")
    // for i := 0; i < 25000; i++ {
    //     slice = append(slice, i)
    //     newCap := cap(slice)
    //     if newCap != prevCap {
    //         delta := newCap - prevCap
    //         growthPercent := 0.0
    //         multiplier := "-"
    //         if prevCap != 0 {
    //             growthPercent = float64(delta) / float64(prevCap) * 100
    //             multiplier = fmt.Sprintf("%.2fx", float64(newCap)/float64(prevCap))
    //         }
    //         fmt.Printf("Length: %-4d  Capacity: %-4d     |   Δ: %-4d   |   Growth: %5.1f%%    |  %-10s\n",
    //             len(slice), newCap, delta, growthPercent, multiplier)
    //         prevCap = newCap
    //     }
    // }
}