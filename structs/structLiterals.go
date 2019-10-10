package main 

import "fmt"

type Vertex struct {
	X int
	Y int
} 

func main(){
	
	//A struct literal denotes a newly allocated struct value by listing the values of its fields.
	//You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	
	var (
		v1 = Vertex{1,2}
		v2 = Vertex{X:1}
		v3 = Vertex{}
		p = &Vertex{1,2}
	)
	
	fmt.Println(v1, v2, v3, p)
}
