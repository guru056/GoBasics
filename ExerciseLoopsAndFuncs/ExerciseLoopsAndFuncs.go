package main

import (
	"fmt"
	"math"
)


func Sqrt(x float64) float64 {
	z := 1.0
	temp := z
	for i:=1; i<10;i++ {
		temp = z
		z -= (z*z - x) / (2*z) 
		fmt.Println("z for  ",i," : ",z)
		if (math.Abs(temp-z) <= 0.0001){
			return z
		}
	}
	return z
}

func main(){
	x := 1.0
    fmt.Println("--- x = ",x," : ----")
    fmt.Println(Sqrt(x))
	x = 2.0
	fmt.Println("--- x = ",x," : ----")
	fmt.Println(Sqrt(x))
	x = 3.0
    fmt.Println("--- x = ",x," : ----")
    fmt.Println(Sqrt(x))
	x = 4.0
    fmt.Println("--- x = ",x," : ----")
    fmt.Println(Sqrt(x))
}
