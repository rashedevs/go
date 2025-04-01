package main

import "fmt"

type User struct{
	Name string
	Age int
	Salary float64
}

func print(nums *[3]int){
	fmt.Println(nums)
}

func (usr *User) display(){
	fmt.Println(usr)
}

func main(){
	// pointer or address of memory (ram)
	x := 28
    
	// p holds the address of x
	p := &x           //ampersand(&) x => address of x || 0xc000010170 (hex) ==> 211106249404416
	*p = 30           // * => value at address || alternatively x = 30
	fmt.Println("which is actually x's new value =", x, *p) 

	arr := [3]int{1,2,3}
	print(&arr)

	user1 := User{
		Name: "Rashed",
		Age: 30,
		Salary: 50000,
	}

	// q := &user1
	// fmt.Println(q.Name)

	user1.display()
	// why we pass the whole struct, not with ampersand & for struct instance.?
}