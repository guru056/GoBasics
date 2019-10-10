package main

import "fmt"

//When a function is receiving an array as a parameter , the size must be specified
func PrintArray(arr [2]int){

	for i:=0;i<len(arr);i++ {
		fmt.Println(arr[i])
	}

}

func printArray(arr [6]int){
	for i:=0;i<len(arr);i++ {
		fmt.Println(arr[i])
	}
}


func main() {

	var a [2]int
	//fmt.Println(a[0])
	a[0] =1
	a[1] =2

	PrintArray(a)

	var primes = [6] int {2,3,5,7,11,13}
	printArray(primes)

}
