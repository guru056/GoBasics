package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y);
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func (v Vertex) ScaleWithoutPointer(f float64) {
      v.X = v.X*f
      v.Y = v.Y*f
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v.Abs());
	v.Scale(10) // changes the values of v
	fmt.Println(v.Abs());
	v.ScaleWithoutPointer(10) // doesn't change the values of v
    fmt.Println(v.Abs());
}
