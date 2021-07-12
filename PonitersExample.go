package main

import "fmt"

func main() {
	/*var p *int
	// fmt.Println(*p) ---> error this values nothing stores in  var p
	fmt.Println(&p)*/
	// & -------> i want to hold a reference of things like address
	// * -------> when we use * it means its a pointer

	var lifebooster float64 = 99.2 // assign value
	lifeboosterRef := &lifebooster // stores address , (we can't declare var = lifeboosterRef = &lifebooster

	fmt.Println(lifebooster)     // prints value ----> 99.2
	fmt.Println(lifeboosterRef)  // prints stored address
	fmt.Println(*lifeboosterRef) // prints value because we use * for access value
}
