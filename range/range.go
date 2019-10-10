package main

import "fmt"

func main() {
	/**The range form of the for loop iterates over a slice or map.

	**When ranging over a slice, two values are returned for each iteration.	** The first is the index, and the second is a copy of the element at that index.
	**/	
	var pow = []int{1,2,4,8,16,32,64,128,256,512}
	for i,v := range pow {
		fmt.Printf("2^%d = %d\n",i,v)
	}

	//You can skip the index or value by assigning to _.
		
	var exp = make([]int,10)
	for i,_ := range exp {
		exp[i] = 1 << uint(i)
	}
	
	//If you only want the index, you can omit the second variable.
	for i:= range exp{
		exp[i] = 1 << uint(i)
	}

	for _,v := range exp {
		fmt.Println(v)
	}
}
