package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string{
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*Like for, the if statement can start with a short statement to execute before the condition.
**Variables declared by the statement are only in scope until the end of the if.
**Variables declared inside an if short statement are also available inside any of the else blocks
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n",v,lim)
	}
	return lim
}

func main () {
	x :=  -1
	if x >=0 {
		fmt.Println("If")
	} else {
		fmt.Println("Else If")
	}
	
	fmt.Println(sqrt(-4))
	fmt.Println(sqrt(4))
	fmt.Println(sqrt(-2))
	fmt.Println(sqrt(2))
	fmt.Println(pow(3,2,10))
	fmt.Println(pow(3,3,20))
}
