package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Creating custom type 'person' for use only in out program.
// type person struct {
// 	firstName string
// 	lastName  string
// 	contact contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//3 ways of declaring a struct.
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Println(alex)
	//fmt.Printf("%+v\n", alex)
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName: "Jam",
	// 	contact: contactInfo{
	// 		email: "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }
	jim := person{
		firstName: "Jim",
		lastName:  "Jam",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// JimPointer is a reference to the memory address that the struct exists at. 
	//jimPointer := &jim

	jim.updateName("jimmy")

	jim.print()

}

//*person = a pointer that points at a person, this is a type description different to *pointerToPerson which would give us the values.
func (pointerToPerson *person) updateName(newFirstName string) {
	//*pointerToPerson = This is an operator - it means we want to manipulate the value the pointer is referencing. 
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}


