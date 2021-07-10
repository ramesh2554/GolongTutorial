package main

import "fmt"

/*
The basic idea behind the structures is to keep things together.
Let's say you have data about a particular Student or about a particular person.
What is that everything has particular attributes example, let's say student, A student will have own Name, Roll Number. Marks, Address same thing goes for bank customer will have Name, Account Number, Amount, Address so every entity will have some parameters and of-course to represent them we can use variables, In side your function you create your variables.
*/
type student struct {
	rollNumber int
	name       string
	marks      int
}

func main() {
	student1 := student{101, "Rameswara", 80} // when we use assign all values
	fmt.Println(student1)
	fmt.Println("Retrieves student1 name -----> ", student1.name)
	student2 := student{rollNumber: 102, marks: 90} // when we want initialise some values use this type
	fmt.Print(student2)
}
