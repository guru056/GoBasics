package main

import (
	"fmt"
	"math"
)

func compute(x float64,y float64,fn func(float64,float64)float64) float64{
	return fn(x,y)
}

func main(){
	hypot := func(x,y float64) float64 {
				return math.Sqrt(x*x + y*y);
			}
	
	fmt.Println(hypot(3,4));
	fmt.Println(compute(3,4,hypot));
	fmt.Println(compute(3,4,math.Pow));

	/*
	**Functions are values too. They can be passed around just like other values.
	**Function values may be used as function arguments and return values.
	*/
}

