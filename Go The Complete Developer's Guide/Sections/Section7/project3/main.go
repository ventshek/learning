package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// A slice of strings.
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// Allocates and initiates an object of type channel, called 'c'.
	c := make(chan string)
	// Loop for calling the checkLink function for each of the strings in slice 'links'. The channel c is passed to the checkLink function also. The 'go' keyword calls the function under a new routine each time.
	for _, link := range links {
		go checkLink(link, c)
	}
	// Infinite loop that passes a link to checkLink that was passed to the chanel from checkLink that originated from the slice of strings "links" via the above loop :). This will loop each time a link is recieved via the chanel, which happens after each http request has been completed in the checkLink function.
	//	for {
	//		go checkLink(<-c, c)
	//	}
	// Alternative syntax for the above function. Wait for the channel to return some value, when it does assign it to the variable "l". It uses a lambda function to first sleep and then pass the link to the checkLink function. 
	for l := range c {
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// The function "checkLink" takes in the variable "link" of type string and the variable "c" that is both a channel and a string.
func checkLink(link string, c chan string) {
	// http.Get is called on the link that has been passed from the main function. This would return resp and err however we are only interested in err so "_" is used as a placeholder.
	_, err := http.Get(link)
	// If the err output is not equal to nil then the link and a message are printed, also the link is passed to the channel.
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	// The link and a string are printed if err is equal to nil and the link is passed to the channel.
	fmt.Println(link, "is up!")
	c <- link
}
