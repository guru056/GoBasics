package main

import "fmt"

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 +1 }
func needFloat(x float64) float64 { return x*0.1}

func main() {
	const Truth = true
	fmt.Println(Pi,Truth)
	fmt.Println(Small)
	//fmt.Println(Big) // Integer overflow error

	//An untyped constant takes the type needed by its context.
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
