package main

import "fmt"

func main() {
	number1 := 10
	number2 := 10
	add , sub  := addSUb( number1, number2)
	fmt.Println("Addition --->" , add , "Substartion--->" , sub)
	mul , div   := mulDiv( number1, number2)
	fmt.Println("Multiplication --->" , mul  , "Division--->" , div)
}
/* A function with multiple return value */
func addSUb(x int, y int) (int , int) {

	addition := x + y
	sub := x - y
	return addition,sub
}

/* A function with multiple return values */
// ------> mulDiv(x int, y int) --> parameter list
// ------->  (multiplication ,Division int) ----> var declaration and return type what we are expecting
func mulDiv(x int, y int) (multiplication ,Division int) {

	multiplication  = x * y
	Division = x / y
	return
}
