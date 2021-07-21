package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom Writer
type logWriter struct {
}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : --->", err)
		os.Exit(1)
	}
	//fmt.Println("Response --------------->" , resp)
	// bs:= []byte // Previous declaration
	//bs := make([]byte , 99999) // make is a default func/method that is created with an byte slice with 99999 spaces
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body) // above 3 line and this same

	// custom writer implementation
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("-------------------")
	return len(bs), nil
}
