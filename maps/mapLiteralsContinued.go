package main

import "fmt"

type Name struct {
	firstName,lastName string
} 

type Employee struct {
	EmpId int 
	EmpName Name
}

func main(){

	employee := Employee{
		42,Name{"firstName","lastName"},
	}

	fmt.Println(employee);

	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	m := map[string]Employee{
		"1" : {12,Name{"firstName","lastName"}},
		"2" : {11,Name{"firstName2","lastName2"}},
	}
	fmt.Println(m)
	//following would give error where Name is to be defined
	/*m := map[string]Employee{
          "1" : {12,{"firstName","lastName"}},
          "2" : {11,{"firstName2","lastName2"}},
      } */
}
