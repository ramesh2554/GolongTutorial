package main

import (
	"fmt"
	"net/http"
)

// website status checker
/*
	step 1 : Take a first url from string of slice
	step 2 : Make a request ---> Get ("http://google.com")
	step 3 : Wait for a response , log it --->
		Take next Link repeat above steps for 2nd link also and this process is continuous until all links in a slice completed
		---> for achieve this drawback  we use Go Routines
		Go Routines Syntax : - go func_name() --> Example :- go checkLine(link)
		---> go --> creates a new thread go routine
		---> func_name() --> run this function with it
		when executes a go func_name() main Routine is exited automatically so to achieve this draw back we use channels


Channels implementation :
	creation : c := make(chan string)

*/

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}
	//  create a channel
	c := make(chan string)
	for _, link := range links {
		// 	checkLink(link)
		go checkLink(link, c) // when we go routines use syntax  ---> go checkLink(link)

	}
	fmt.Println(<-c)
}
func checkLink(link string, c chan string) {
	_, err := http.Get(link) //  resp  ,  err := http.Get(link) --> Here we dont need to print resp so we use underscore(_)

	if err != nil {
		fmt.Println(link, "Might down ")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "Is up ")
	c <- "Yep its up"
}
