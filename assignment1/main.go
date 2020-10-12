package main

import "fmt"

func main() {

	intArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range intArray {
		if intArray[i]%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("odd")
		}
	}
}
