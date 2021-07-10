package main

import "fmt"

func main() {
	/*
		   Defer statement is used to execute a function call just before the surrounding function where the defer statement is present returns. The definition might seem complex but it's pretty simple to understand
	     1. Defer fallow STACK --> LIFO(last in first out) order
	*/

	a()
	cal()
}
func a() {
	fmt.Println("in a starts")
	defer b()
	defer c()
	fmt.Println("A ends")
}
func b() {
	fmt.Println("In b")
}

func c() {
	fmt.Println("In c")
}

func cal() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // it will print in reverse order because of we use multiple defers and defer also fallow LIFO
	}
	fmt.Println("Bye") // executes 1st because defer function
}
