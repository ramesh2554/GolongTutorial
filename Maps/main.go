package main

import (
	"fmt"
)

/*
Map is a collection of value pairs ex: key --> value
syntax : variable_name := map[typeofkey ex:-string] valuetype ex string { " " : " " }
*/
func main() {

	// 1st method
	colors := map[string]string{
		"red":    "#FF0000",
		"green":  "#00FF00",
		"Yellow": "#ABCD0",
	}
	// 2nd method
	// var colors map[string]string     // -------> prints and empty map like , map[]

	// 3rd Method

	//  colors := make(map[string]string) // --------> prints and empty map like , map[]
	fmt.Println(colors)
	colors["white"] = "#FFFFF" // add elements to existing map
	fmt.Println("After update --------> ", colors)
	fmt.Println("Access Individual Elements in maps ----->", colors["red"])

	delete(colors, "white")
	fmt.Println("After deleting into map elements -------> ", colors)
	printMap(colors)
	// Iterating Over a maps
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
