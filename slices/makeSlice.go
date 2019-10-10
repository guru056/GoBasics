package main

import "fmt"

func PrintSlice(s []int) {
	fmt.Printf("len=%d,cap=%d,s=%v\n",len(s),cap(s),s)
}

func main() {
	a := make([]int,5)
	PrintSlice(a)
	
	s := make([]int,0,5)
	PrintSlice(s)

	s = s[:cap(s)-1]
	PrintSlice(s)

	s = s[1:]
	PrintSlice(s)
}
