package main

import (
	"fmt"
)

func main() {

	type Person struct {
		name string
		age int
		address string
	}

	p := Person{name: "John Hopkins", age: 39, address: "Malboro street"}

	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.address)
}
