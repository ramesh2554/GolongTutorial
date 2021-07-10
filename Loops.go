package main

import "fmt"

/*In Golang only loop is there i.e For
 we can use 3 different ways
1. infinite for loop --> for { body of loop }
2. inside for loop index value will be increase to break a loop --> var decl + init for condition { body of loop , increment index }
3. like we did in other programming lang like java we can use same in go lang also ---> for index=value;condition;increment { body of loop }
*/
func main() {
	//1st method
	/*for{
		fmt.Println("Ramesh")
	}*/
	// 2nd method
	fmt.Println("---------------------2nd Method")
	index := 0
	for index <=5 {
		fmt.Println("ramesh" , index)
		index++
	}
	// 3nd method
	fmt.Println("---------------------3rd Method")
	for i:=0; i<=5; i++{

		fmt.Println("ramesh" , i)

	}
}
