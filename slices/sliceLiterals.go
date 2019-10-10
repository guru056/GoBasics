package main

import "fmt"

func main() {
	//This is an array literal
	a := [3]bool{true,true,false}
	
	//A slice literal is like an array literal without the length.
	//Examples :
	q := []int{2,3,5,7}
	fmt.Println(q)

	r := []bool{true,false,true,false}
	fmt.Println(r)

	s := []struct{
				i int
				 b bool
			}{
					{2,true},
					{3,false},
					{5,true},
					{7,false},
			 }

	fmt.Println(s)

}
