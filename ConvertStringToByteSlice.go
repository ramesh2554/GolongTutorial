package main

import "fmt"

func main() {

	name := "Ramesh"

	convertedToSlice := []byte(name)

	convertedSliceToString := string(convertedToSlice)
	fmt.Println("Name --------> ", name) // --> string of a name
	fmt.Println(name, "Byte array(Slice) ----> ", convertedToSlice)
	fmt.Println(convertedToSlice, "convertedSliceToString--------->", convertedSliceToString)
}
