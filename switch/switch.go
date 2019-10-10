package main

import "fmt"
import "runtime"
import "time"

func main() {
	os := runtime.GOOS
	switch os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s\n",os)
	}
	
	//Switch without a condition is the same as switch true.
	t := time.Now()
	switch {
		case t.Hour() < 12 :
			fmt.Println("Good Morning")
		case t.Hour() < 17 :
			fmt.Println("Good afternoon")
		default :
			fmt.Println("Good evening.")
	}

	start := time.Now()
	fmt.Println(start);
	t1 := time.Now()
	fmt.Println(t1)
	elapsed := t1.Sub(start)
	fmt.Println(elapsed)

	//Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0 :
			fmt.Println("Today!")
		case today + 1 :
			fmt.Println("Tomorrow!")
		case today + 2:
			fmt.Println("In two days")
		default :
			fmt.Println("Too far away...")
	}

}
