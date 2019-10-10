package main

import (
	 "fmt"
	 "math"
)

type Vertex struct {
	X,Y float64
}

/*Normal function*/
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


/*
*
*A method is a function with a special receiver argument.
*The receiver appears in its own argument list between the func keyword and the method name.
*In this example, the Abs method has a receiver of type Vertex named v.
*
*/
func (v Vertex) Hypot() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	
	v := Vertex{3,4}
	fmt.Println(Abs(Vertex{3,4}));
	fmt.Println(v.Hypot());

}
