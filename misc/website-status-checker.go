package main

import (
	"fmt"
	"net/http"
)

func main(){
	
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for range links {
		fmt.Println(<-c)
	}
}

func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down"
		return
	}

	c <- link + " is up"
}
