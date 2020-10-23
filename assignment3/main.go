package main

import (
	"bytes"
	"fmt"
	"os"
)

type reader struct{}

func main() {

	argsInProg := os.Args
	for i := range argsInProg {
		if i != 0 {
			myFile, err := os.Open(argsInProg[i])
			if err != nil {
				fmt.Printf("Error opening file %s", argsInProg[i])
			} else {

				defer myFile.Close()

				buf := new(bytes.Buffer)
				buf.ReadFrom(myFile)
				contents := buf.String()

				fmt.Printf("File contents: %s", contents)
			}
		}
	}
}
