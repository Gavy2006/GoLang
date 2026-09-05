package main

import "fmt"

var Public = "it is public" // can be access from outside

// agr kisi variable ko bahr se access krna hai to us var ke name ko capital mei rkhan pdega not in small public -> Public

// same applicable for func names also
//PublicFunction() -> can access from diff packages
//publicFunction() -> can not access from diff packages

func main() {

	fmt.Println("learn go language")

	printMsg()

	var name string = "Gavy"
	fmt.Println(name)

	var age int = 20
	fmt.Println(age)

	var check bool = true
	fmt.Println(check)

	var money float32 = 902002.32
	fmt.Println(money)

	fmt.Println("money =>>", money)

	const pii = 67.12
	fmt.Println(pii)

	person := 123
	fmt.Println(person)

	user := "gavy"
	fmt.Println(user)

	// user = 90 error *** samae data type
	user = "545"
	fmt.Println(user)

	person = 90
	var public = "it is public"
	var private = "it is private"

	fmt.Println(public)
	fmt.Println(private)

	fmt.Println(Public)

}
