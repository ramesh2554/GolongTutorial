package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	message := []byte("This is an example of saveing a file to Hard drive ")
	ioutil.WriteFile("saveFileToHDD", message, 0666)
	readFile()
}

/* Retrieve data from hdd file */
func readFile() {
	content, err := ioutil.ReadFile("Loops.go")
	if err != nil {
		fmt.Println("Error : -------- > ", err)
	}
	fmt.Println(string(content))
}
