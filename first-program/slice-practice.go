package main

// func addElement(s []int, value int) []int {
// 	fmt.Printf("[addElement] addr of s[0]: %p\n", &s[0])
// 	s = append(s, value)
// 	fmt.Printf("[addElement] after append addr of s[0]: %p\n", &s[0])
// 	return s
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

// slice class practice here below,

// func main(){
// arr := [6]string{"this","is","a","go","interview","question"}

// // 1. slice from an array
// slice1 := arr[1:4]
// fmt.Println("slice1", slice1)  //[is a go]
// fmt.Println(len(slice1))       // 3
// fmt.Println(cap(slice1))       // 5

// //2. slice from a slice
// slice2 := slice1[1:2]
// fmt.Println("slice2", slice2)  // [a]
// fmt.Println(len(slice2))       // 1
// fmt.Println(cap(slice2))       // 4

// //3. slice literals (initially len and cap same)
// slc := []int{1,2,3}
// fmt.Println("slc", slc)  // [1 2 3]
// fmt.Println(len(slc))       // 3
// fmt.Println(cap(slc))       // 3

// //4. slice using make function with length
// mslc := make([]int, 3)
// fmt.Println("mslc", mslc)  // [0 0 0]
// fmt.Println(len(mslc))       // 3
// fmt.Println(cap(mslc))       // 3

// //5. slice using make function with length & cap
// mslc2 := make([]int, 3, 5)
// fmt.Println("mslc2", mslc2)  // [0 0 0]
// fmt.Println(len(mslc2))       // 3
// fmt.Println(maincap(mslc2))       // 3

// //6. empty / nil slice (initially, ptr = nil, length = 0 & cap = 0)
// nilslc := []int{}
// fmt.Println("nilslc", nilslc)  // []
// fmt.Println(len(nilslc))       // 0
// fmt.Println(cap(nilslc))       // 0

// }

