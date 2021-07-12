package main

import "fmt"

/*
Struct is a Data Structure ,  Collection of properties that are related to together

The basic idea behind the structures is to keep things together.
Let's say you have data about a particular Student or about a particular person.
What is that everything has particular attributes example, let's say student, A student will have own Name, Roll Number. Marks, Address same thing goes for bank customer will have Name, Account Number, Amount, Address so every entity will have some parameters and of-course to represent them we can use variables, In side your function you create your variables.

*/

type person struct { // -----> create a struct
	firstName string
	lastName  string
}

type personEmbeddedStruct struct { // -------> Embedded Struct
	roll    int
	name    string
	contact contactInfo // uses contactInfo struct
}
type contactInfo struct {
	email   string
	zipcode int
}

func main() {

	fmt.Println(person{"Ramesh", "Konakalla"}) // ---> 1st way

	ramesh := person{firstName: "ram", lastName: "k"} // ---> 2nd way
	fmt.Println(ramesh)

	var ayyapa person // -------> 3rd way but it returns empty value
	fmt.Println(ayyapa)
	fmt.Printf("%+v", ayyapa) // ----> prints person feild names

	ayyapa.firstName = "Ayyappa Naveen Kumar " //  update a value into struct
	ayyapa.lastName = "Konakalla"

	fmt.Println("After update values in Struct ", ayyapa)

	fmt.Println("-------------------------------------------------")

	fmt.Println("Embedded Struct")

	details := personEmbeddedStruct{
		roll: 101,
		name: "rameswara",
		contact: contactInfo{
			email:   "temp@gmail.com",
			zipcode: 53445,
		},
	}

	details.updateName("Konakalla Rameswara (kR)")
	details.print()
}

// receiver function used for print
func (p personEmbeddedStruct) print() {
	fmt.Printf("Embedded struct -----> %+v", p)
}

func (p *personEmbeddedStruct) updateName(newName string) {
	(p).name = newName
}
