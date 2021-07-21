package main

import (
	"fmt"
	"io"
	"os"
)

/*
create a program Read a contents from text file and prints on terminal
Note : File  should be read on Command Line Args

*/
func main() {

	f, err := os.Open(os.Args[1]) //--------->>>   open a file from terminal EX: go run InterfaceAssignment2.go filename
	// os.Args[0] creates  a temp file created by go  ---- os.Args[1] takes a filename as we pass and commandLIne
	if err != nil {
		fmt.Println("Error---> : ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f) // This Lines uses for Print the data in console/terminal

}
