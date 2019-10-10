package main

import "fmt"

func main() {

	s := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(s[1:4])
	fmt.Println(s[:4])
	fmt.Println(s[4:])
	fmt.Println(s[:])

}
