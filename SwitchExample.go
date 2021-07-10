package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Switch Example
		In this program we use time package for finding when was a sunday
	*/
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("None")
	}

	demoTime()

	findSaturday()
}
func demoTime() {
	fmt.Println(time.Now().Weekday())

}
func findSaturday() {
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tomorrow is Monday")
	case today + 2:
		fmt.Println("Day after tomorrow is Monday")
	default:
		fmt.Println("So far way")
	}

}
