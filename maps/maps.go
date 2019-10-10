package main

import "fmt"
/*
	A map is declared using the following syntax -
	var m map[KeyType]ValueType

*/
type Vertex struct {
	Lat, Long float64
}

func main() {
	/*
	A map maps keys to values.
	The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	The make function returns a map of the given type, initialized and ready for use.
	*/	
	m := make(map[int]string);
	m[1] = "One";
	m[2] = "Two";
	fmt.Println(m[1]);
	fmt.Println(m[2]);
	fmt.Println(m);

	mLatLong := make(map[string]Vertex);
	mLatLong["VTS"] = Vertex{
		12.4565, -123454,
	}
	fmt.Println(mLatLong["VTS"]);
	
	for i,v := range m {
		fmt.Println("m[",i,"]=",v);
	} 	
	
	
}
