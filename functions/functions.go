package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add1(x , y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y,x
}

func add2(x, y int) (sum int){
	sum = x + y
	return
}

func split(sum int) (x, y int){
	x = ( sum * 4 ) / 9
	y = sum - x
	return 
}

func main(){
	fmt.Println(add(3,4));
	fmt.Println(add1(3,4));
	fmt.Println(add2(3,4));
	a,b := swap("Hello", "World")
	fmt.Println(a,b)
	//fmt.Println(b)
	c,d := split(9)
	fmt.Println(c,d)
//	fmt.Println(d)
}
