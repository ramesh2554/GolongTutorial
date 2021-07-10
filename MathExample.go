package main

import (
	"fmt"
	"math"
)

func main() {

	number := 12 //  number is type of int when we use sqrt method it should be type casted

	var number1 float64 = 25
	var result = math.Sqrt(float64(number)) // -- >  math.Sqrt accepts only float number so here we type casted to float64
	fmt.Println("Variable is type casted ", result)
	fmt.Printf("Variable is %.2f type casted , and it will be Formmated using printf ", result)
	fmt.Println("Variable is type casted, and it will be rounded ", math.Round(result))
	fmt.Println("Variable is declared as float64 ", math.Sqrt(number1))

}
