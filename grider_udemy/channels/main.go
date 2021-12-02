package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://microsoft.com",
		"http://apple.com",
		"http://golang.org",
	}

	channel1 := make(chan string)

	for _, link := range links {
		go httpRequest(link, channel1)
	}

	for channelContent := range channel1 {
		//function literal(like an anonymous function in javascript)
		go func(channelContent string) {
			time.Sleep(3 * time.Second)
			httpRequest(channelContent, channel1)
		}(channelContent)
		//the extra () are part of the function literal syntax, you use them to invoke the function
		//before we had go function(args), now we have go, define a function and then invoke it with ()
	}

}

func httpRequest(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Response: ", link)
		channel <- link
		return
	}
	fmt.Println(link, " is online")

	channel <- link
}
