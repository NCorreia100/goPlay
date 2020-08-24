package main

import "fmt"

type contactInfo struct {
	phone int
	email string
}
type person struct {
	name string
	age  int
	contactInfo
}

func (p person) print() {
	fmt.Printf("\n%+v", p)
}

func (p *person) setContactInfo(contactInfo contactInfo) {
	(*p).contactInfo = contactInfo
}

func updateSlice(s *string) {
	*s = "thomas"
}

func main() {
	guy := person{"Jimmy", 16, contactInfo{9876543210, "test@email.com"}}
	// girlContact := contactInfo{12345678910, "email@email.com"}

	girl := person{
		name: "Monica",
		age:  30,
		contactInfo: contactInfo{
			phone: 1234567890,
			email: "email@email.com",
		},
	}

	guy.print()
	girl.print()

	// girlPointer := &girl
	girl.setContactInfo(guy.contactInfo)
	girl.print()

	personName := "jake"

	updateSlice(&personName)

	fmt.Println("\n", personName)
}
