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

	jim.print()
	jim.updateFirstName("newJim")
	jim.print()
}
