package main

import "fmt"

func main() {
	/*
		Example of If and if else
		In this program we print given number is less than 5 print Hi else print Bye
	*/
	exampleIf()
	exampleIfElse()
	oddEven() // --> checks number is even or odd

}
func exampleIf() {

	number := 4
	if number <= 5 {
		fmt.Println("Hi")
	} else {
		fmt.Println("Bye")
	}
}
func exampleIfElse() {

	number := 1
	if number == 1 {
		fmt.Println("ONE")
	} else if number == 2 {
		fmt.Println("TWO")
	} else {
		fmt.Println("default else")
	}
}

func oddEven() {

	num := 11
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
