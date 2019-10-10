package main

import (
	"fmt"
	"math"
)

//Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
//while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

type Vertex struct {
	X,Y float64
}

func Abs(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}

func (v *Vertex) ScaleMethod(f float64){
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main(){

	v := Vertex{3,4}
	v.ScaleMethod(2)
	fmt.Println(v)
	fmt.Println(Abs(v))

	(&v).ScaleMethod(0.5);
	fmt.Println(v)
	fmt.Println(Abs(v))

	p := &v
	p.ScaleMethod(2)
	fmt.Println(v)
	fmt.Println(Abs(v))

	ScaleFunc(p, 0.5)
	fmt.Println(v)
	fmt.Println(Abs(v))


}

