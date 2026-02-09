package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	fmt.Println("Learning Go Routine")
	// create a slice of string
	links := []string{
		"https://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a brand new channel
	// use make inbuilt function, the first parameter is channel(short chan) of type string
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	// a single line receive a single message, written to the channeel
	for l := range c {
		go func(lnk string) {
			time.Sleep(5 * time.Second)
			checkLink(lnk, c)
		}(l)
	}

	// fmt.Println("main done")
}

/**
* c is a channel of string type
 */
func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		c <- link
		os.Exit(1)
	}
	fmt.Println(link + " is up")
	c <- link
	// io.Copy(os.Stdout, resp.Body)

}
