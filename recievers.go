package main

import "fmt"

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
func (p *Person) updateFirstName(name string) {
	(*p).firstName = name
}
