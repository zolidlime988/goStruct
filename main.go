package main

import "fmt"

func main() {
	var p1 person
	fmt.Println(p1)
	fmt.Printf("%+v\n", p1) // print value + key
	p1.firstName = "John"
	p1.lastName = "Doe"
	fmt.Printf("%+v\n", p1) // print value + key
}
