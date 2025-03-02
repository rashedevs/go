// package main

// Go internal memory usage.

// Memory (RAM) Partioning:

// 1. Code segment: The code segment is used to store the program's instructions and const variables.
//    The code segment is read-only. Local function expressions are stored in the code segment.

// 2. Data segment: The data segment is used to store the values of re-assignable variables.
//    These are allocated to ram's data segment at run/execution time.

// 3. Stack: Place where the function call is stored. The stack is used to store
//    the values of variables that are created at compile time.

//---- Stack Frame: A stack frame is a block of memory that is used to store
//---- the values of variables that are created at runtime. Stack frames are dynamic in size.

// 4. Heap: The heap is used to store the values of variables that are created at runtime.

// -----Go uses a garbage collector to manage memory.
// -----The garbage collector is responsible for freeing memory that is no longer in use.
// -----The garbage collector runs in the background.

//example.....................

// const a = 10
// var p = 100

// func call(){
//     add := func(x int, y int){
//         z := x + y
//         fmt.Println(z)
//     }

//     add(5, 6)
//     add(p, a)
// }

// func main(){
//     call()
//     fmt.Println("global", a)  // global const 10
//     a := 20
//     fmt.Println("redeclare with shadow", a) //local redeclare with shadow 20
//     a, b := 30, 40
//     fmt.Println("redeclare with a shadow and b new declaration", a, b) //local shadow with :=, with new var b declaration
//     a = 30
//     fmt.Println("reassign", a) // reassign to local shadowed 30
// }

// func init(){
//     fmt.Println("Hello, init")
// }




