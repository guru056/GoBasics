package main

import "fmt"

func main() {
	//	A slice is a dynamically-sized, flexible view into the elements of an array
	numbers := [6]int {1,2,3,4,5,6}
	fmt.Println(numbers)
		
	var s []int = numbers[1:4] // [1,4)
	fmt.Println(s)

/**	Slices are like references to arrays
 **	A slice does not store any data, it just describes a section of an underlying array.
 **
 **	Changing the elements of a slice modifies the corresponding elements of its underlying array.
 **
 **	Other slices that share the same underlying array will see those changes.
**/	
	s[0] = 7
	fmt.Println(s)
	fmt.Println(numbers)
	
}
