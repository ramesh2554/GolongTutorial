package main

import "fmt"

/*
In the main function we call calEvenOrOdd()
 in calEvenOrOdd we create values of integer Slice with range 0 to 10 and check each number is even or odd
Example formatted output :
1 odd
2 even
3 odd
4 even

*/
func calEvenOrOdd() {

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	/*number := make([]int , 11)
	number[1]=1 // insted of assign values to  number number[1] =1  use for loop to assign values
	for i=0; i<11 ; i++ {
	     number[i] = i
	}
	*/
	for _, value := range values {

		if value%2 == 0 {
			fmt.Println(value, "even")
		} else {
			fmt.Println(value, "odd")
		}

	}
}
func main() {
	calEvenOrOdd()
}
