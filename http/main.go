package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("http request failed", err)
		os.Exit(1)
	} else {
		if res.StatusCode >= 400 {
			fmt.Println("Client faillure", res.StatusCode)
		} else {

			// io.Copy(os.Stdout, res.Body)
			lw := logWritter{}
			io.Copy(lw, res.Body)
		}

	}

}
