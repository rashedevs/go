package main

// Go internal memory usage.

// Memory (RAM) Partioning:

// 1. Code segment: The code segment is used to store the program's instructions.

// 2. Data segment: The data segment is used to store the values of variables.

// 3. Stack: Place where the function call is stored. The stack is used to store
//     the values of variables that are created at compile time.

//---- Stack Frame: A stack frame is a block of memory that is used to store
//---- the values of variables that are created at runtime.

// 4. Heap: The heap is used to store the values of variables that are created at runtime.

// -----Go uses a garbage collector to manage memory.
// -----The garbage collector is responsible for freeing memory that is no longer in use.
// -----The garbage collector runs in the background.