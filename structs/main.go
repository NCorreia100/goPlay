package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	guy := person{"Jimmy", 16}
	girl := person{name: "Monica", age: 30}

	fmt.Println(guy)
	fmt.Printf("%+v", girl)

}
