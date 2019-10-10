package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if(f < 0){
		return float64(-f)
	}
	return float64(f)

	//No ternary operators in go
	//return (f<0)?float64(-f):float64(f) --> this will throw error
}

func main(){
	f := MyFloat(-math.Sqrt(4));
	fmt.Println(f.Abs());
	f = MyFloat(2.0);
	fmt.Println(f.Abs());

}
