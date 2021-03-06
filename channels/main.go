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

	c := make(chan string)

	for _, link := range links {
		go getChannelStatus(link, c)
	}
	/*for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}*/

	/*for {
		go getChannelStatus(<-c, c)
	}*/

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go getChannelStatus(link, c)
		}(l)
	}

}

func getChannelStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down")
		c <- link
		return
	}

	c <- link
	fmt.Println(link, " is up")

}
