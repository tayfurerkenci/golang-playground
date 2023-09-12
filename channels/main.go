package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// make a channel
	c := make(chan string)

	for _, link := range links {
		// go keyword creates a new go routine
		go checkLink(link, c)
	}

	// Never try to access the same variable from different go routines
	for l := range c {
		// function literal (anonymous function from javascript or lambda from c#)
		go func(link string) {
			time.Sleep(5 * time.Second)
			// receive a message from the channel
			checkLink(link, c)

			// func literal send l to the func
			// l is copied in memory and then the go routine
			// has access to the copy as opposed to the original value of L
			// now, child go routine does not reference the same memory address as the parent go routine does !!!!! IMPORTANT !!!!!
		}(l)
	}

	// Alternative way to receive a message from the channel for the eternity
	// for {
	// 	go checkLink(<-c, c)
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
