package main

import "fmt"

func PrintSlice(s []int){
	fmt.Printf("len=%d, cap=%d, slice=%v \n", len(s), cap(s), s)
}

func main() {

	s := []int{1,2,3,4,5,6,7}
	PrintSlice(s)

	//set length to 4
	s = s[:4]   //len = 4 , cap = 7
	PrintSlice(s)
	
	//Drop first two values
	s = s[2:]   // len = (4-2=2) and cap = (7-2)=5
	PrintSlice(s)

	//Extend length to 4 
	s = s[:4]
	PrintSlice(s)
	
	//Can't extend length more than capacity
}
