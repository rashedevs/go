package main

import "fmt"

// func addElement(s []int, value int) []int {
// 	fmt.Printf("[addElement] addr of s[0]: %p\n", &s[0])
// 	s = append(s, value)
// 	fmt.Printf("[addElement] after append addr of s[0]: %p\n", &s[0])
// 	return s
// }

func main() {
	// s := []int{1, 2, 3}
	// fmt.Printf("[main---] before append addr of s[0]: %p\n", &s[0])
	// s = addElement(s, 4)
	// fmt.Printf("[main---] after append addr of s[0]: %p\n", &s[0])

	s := make([]int, 3, 5)

	fmt.Println("s after initialization", &s[0])

	s[0]=5
	s[1]=15
	s[2]=25
	fmt.Println("s after index based insertion", &s[0])
	// s[3]=30

	s = append(s, 30)
	fmt.Println("s after append", &s[0])
}
