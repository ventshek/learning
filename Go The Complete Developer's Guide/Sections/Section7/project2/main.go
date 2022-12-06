package main

import (
	"fmt"
	"net/http"
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
	// Loop that waits till the number of channel messages recieved are equal to the ammount of strings in slice 'links'. Once this is completed the main routine is no longer blocked.
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

// The function "checkLink" takes in the variable "link" of type string and the variable "c" that is both a channel and a string.
func checkLink(link string, c chan string) {
	// http.Get is called on the link that has been passed from the main function. This would return resp and err however we are only interested in err so "_" is used as a placeholder.
	_, err := http.Get(link)
	// If the err output is not equal to nil then the link and a message are printed, also a message is passed to the channel c.
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	// The link and a string are printed if err is equal to nil.
	fmt.Println(link, "is up!")
	c <- "Yep it's up"

}
