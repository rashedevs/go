package main

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

import "fmt"

// Global variable (Data Segment)
var globalVar int = 10

// Global constant (Data Segment or Code Segment, depending on usage)
const globalConst = 100

func main() {
    // Local variable (Stack)
    localVar := 5

    // Local constant (Stack or Code Segment, optimized by compiler)
    const localConst = 50

    fmt.Println("Global Variable:", globalVar)
    fmt.Println("Global Constant:", globalConst)
    fmt.Println("Local Variable:", localVar)
    fmt.Println("Local Constant:", localConst)

    // Inner function (Code Segment, stack frame access)
    innerFunction := func() {
        innerVar := localVar + localConst // Stored in the stack
        fmt.Println("Inner Variable:", innerVar)
    }

    innerFunction() // Creates a stack frame for this function call
}
