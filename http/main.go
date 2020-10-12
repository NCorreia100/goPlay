package main

import (
	"fmt"
	"net/http"
	"os"
)

func Read(b []byte) (int, error) {
	if error != nil {
		fmt.Println("unable to read response", error)
	} else {
		fmt.Println(b)

	}
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

			resBody, err := http.ReadResponse(res)
			Read(resBody)
			// if err != nil {
			// 	fmt.Println("Error reading response", err)
			// } else {
			// 	fmt.Println(resBody)
			// }
		}

	}

}
