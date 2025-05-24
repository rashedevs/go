package main

import "fmt"

// func addElement(s []int, value int) []int {
// 	fmt.Printf("[addElement] addr of s[0]: %p\n", &s[0])
// 	s = append(s, value)
// 	fmt.Printf("[addElement] after append addr of s[0]: %p\n", &s[0])
// 	return s
// }

// func main() {
// // variable shadowing
// var x = 5
// fmt.Println(x)

// {
// 	x := 10
// fmt.Println(x)

// }
// fmt.Println(x)
// }

// 1024 er pore 25% - 50% capacity increase hoite pare.
// func main() {
//     s := make([]int, 1024)
//     fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
//     s = append(s, 1)
//     fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
// }

// output:
// len: 1024, cap: 1024
// len: 1025, cap: 1536

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:3]
	c := a[1:4]

	b[2] = 99
	c[1] = 88

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}