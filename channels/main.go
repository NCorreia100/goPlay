package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " is down!")
	} else {
		fmt.Println(link + " OK!")
	}
	time.Sleep(time.Second * 2) //wait 5 min in child go routine (not main routine)
	ch <- link
}
func main() {
	// msec := time.Now().UnixNano() / 1000000
	var ch = make(chan string)
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, l := range links {
		go checkLink(l, ch)
		// fmt.Println(<-ch) ---> if I print the channel here isntead, it will block the execution of the main routine
	}

	// qty := 0
	for c := range ch {
		// if qty%5 == 0 {
		// time.Sleep(time.Second * 2)
		// }
		// qty++

		//restart go routines
		go checkLink(c, ch)
	}
	// fmt.Println("time passed", (time.Now().UnixNano()/1000000)-msec)
}
