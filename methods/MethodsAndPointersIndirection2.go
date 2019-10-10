package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}

func AbsFunc(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}

func main(){

	/*The equivalent thing happens in the reverse direction.

	Functions that take a value argument must take a value of that specific type:*/
	/*while methods with value receivers take either a value or a pointer as the receiver when they are called:*/
	v := Vertex{3,4}
	p := &v

	fmt.Println(v.Abs());
	fmt.Println(p.Abs());

	fmt.Println(AbsFunc(v));
	//fmt.Println(AbsFunc(p)); //compile error


}


