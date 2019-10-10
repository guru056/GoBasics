package main

import "fmt"

type Vertex struct {
	Lat , Long float64
}

func PrintMap(m map[int]string){
	for i,v := range m{
		fmt.Println("m[",i,"]=",v);
	}
}

func PrintMapVertex(m map[string]Vertex){
	for i,v := range m{
          fmt.Println("m[",i,"]=",v);
      }
}

func main() {
	//Map literals are like struct literals, but the keys are required.	
	m := map[int]string{
		1 : "One",
		2 : "Two",
	};
	PrintMap(m);

	mLatLong := map[string]Vertex{
		"VTS" : Vertex{12.4563,-12.4563},
		"OHC" : Vertex{13.4632,-13.2341},
	}

	PrintMapVertex(mLatLong);

	mLatLong1 := map[string]Vertex{
		"VTS" : {12.4563,-12.4563},
		"OHC" : {13.4632,-13.2341},
	}

	PrintMapVertex(mLatLong1);
}
