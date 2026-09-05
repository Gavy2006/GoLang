package main

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {

	fmt.Println("Enter Your name Mr..")

	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello Mr", name)

	// buffer reader ias type of reader  from any underlying source such as file or statndard input(keyboard)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // jb tk  next line nhi milta tbk tk read

	fmt.Println("Hello Mr", name)

}
