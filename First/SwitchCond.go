package main

import "fmt"

func switchcond(day int) {

	switch day { // ya to yaha day nhi to simple switch or cases mein day = 1 , day >= 0 ye vo yaha switch{} krke he hoga got it

	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")

	case 9, 8, 10:
		fmt.Println("holidayyyyy")

		// cant use day here simple switch switch{} mein isem aese ase define kr skte hain

	// case (day > 10 && day <= 15):
	// 	fmt.Println("holidayyyyy 2.0")

	default:
		fmt.Println("No day exist , go die")
	}

}
