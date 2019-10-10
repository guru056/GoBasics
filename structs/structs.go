package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main(){
	//A struct is a collection of fields.
	v := Vertex{1,2}
	fmt.Println(v)
	fmt.Printf("%+v\n",v)

	//Struct fields are accessed using a dot.
	fmt.Println(v.X)
	
}
