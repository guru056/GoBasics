package main

import "fmt"

func main () {
	sum := 0
	//for loop syntax : separated by ; and no small brackets
	for i :=0 ; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for  ; sum < 1000 ; {
		sum += sum
	}
	
	fmt.Println(sum)

	//For is Go's "while"
	sum = 0
	i := 1
	//If you omit the loop condition , it loops forever
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
