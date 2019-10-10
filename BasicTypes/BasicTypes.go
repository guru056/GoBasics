package main

import "fmt"
import "math"
import "math/cmplx"

var (
	ToBool bool   		= false
	MaxInt uint64 		= (1 << 64) -1
	z      complex128	= cmplx.Sqrt(-5+12i)	
)

func main() {
	fmt.Printf("Type : %T Value : %v \n", ToBool, ToBool)
	fmt.Printf("Type : %T Value : %v \n", MaxInt, MaxInt)
	fmt.Printf("Type : %T Value : %v \n", z, z)

	//Zero Values
	var i int
	var j float64
	var k bool
	var s string

	fmt.Println("%v %v %v %q \n", i, j, k, s)
	//To learn more about %q and others -- https://golang.org/pkg/fmt/


	//Type Conversions
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
	//All Type conversions are explicit
	
	//Type Interference
	a := 42
	fmt.Printf("a type of type %T\n",a)
	b := 42.0 
	fmt.Printf("b type of type %T\n",b)
	c := 3 + 4i
	fmt.Printf("c type of type %T\n",c)
}
