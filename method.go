package main

import "fmt"

type Human struct {
	name string
	age int
}

func (h Human)describe() {
	fmt.Println(h.name, "is", h.age, "years old")
}

func main() {
	h:= Human{name: "Jackson", age: 22}
	h.describe()
}
