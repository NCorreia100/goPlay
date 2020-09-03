package main

import "fmt"

func printColors(c map[string]string) {
	for key := range c {
		fmt.Println("color", key, "is", c[key])
	}
}

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "fggddf",
	}

	colors["white"] = "#ffffff"

	printColors(colors)

}
