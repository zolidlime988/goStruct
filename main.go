package main

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Smith",
		ContactInfo: ContactInfo{
			email:   "jim@example.com",
			zipCode: 123456,
		},
	}
	pJim := &jim
	pJim.updateFirstName("newJim")
	jim.print()
}
