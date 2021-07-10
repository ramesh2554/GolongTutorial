package main

import "fmt"

var num = 9 //declares global we can access in inside any method
func demo() {
	// number := 20 //Method level declaration we cannot access it
	num = 8 // num is re init to 8 so its prints 8 because it compiler checks for local var 1st if its found then its prints that
	// if its not found then check for global
	fmt.Println("in demo ", num) // its prints 8 because it compiler checks for local var 1st

}
func main() {
	a := 10 //
	demo()
	fmt.Println("in  main ", a)
}
