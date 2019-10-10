package main

import "fmt"

func main() {

	//The type *T is a pointer to a T value. Its zero value is nil.
	var pointer *int
	fmt.Println(pointer) // <nil>

	/*
	**The & operator generates a pointer to its operand.
	**The * operator denotes the pointer's underlying value.
	*/

	i,j := 12,15
	fmt.Println("The value of i before is : ",i)
	p := &i
	*p = 15
	fmt.Println("The value of i now is : ",i)

	fmt.Println("The value of j before is : ",j)
	p = &j
	*p /= 5
	fmt.Println("The value of j now is : ",j)

}
