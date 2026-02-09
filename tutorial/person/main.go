package main

import "fmt"

type ContactInfo struct {
	email string
	zip   int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	var naresh Person
	naresh.firstName = "Naresh"
	naresh.lastName = "Saw"
	naresh.contact.email = "nareshsaw5@gmail.com"
	naresh.contact.zip = 560061
	fmt.Printf("%+v", naresh)
	fmt.Println()

	var jim = Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: ContactInfo{
			email: "jim@yahoo.com",
			zip:   825301,
		},
	}

	jim.updateName("jimmy")
	jim.print()

}

/*
*
* please notice * in front of the person
 */
func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
