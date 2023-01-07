package main

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}
type ContactInfo struct {
	email   string
	zipCode int
}
