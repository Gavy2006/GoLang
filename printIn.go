package main

import "fmt"

func printIn() {

	name := "Gavy"
	age := 20
	class := "Eco"
	Salary := 2000000.1212

	fmt.Println("name", name, "age", age, "class", class)

	fmt.Printf("Name is %s", name)       //  s  for  string
	fmt.Printf("Age is %d", age)         //  d  for int
	fmt.Printf("Age is %T", age)         //  Type of var
	fmt.Printf("Salary is %f", Salary)   // font type
	fmt.Printf("Salary is %.3f", Salary) // font type

}
