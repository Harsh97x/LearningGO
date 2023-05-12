package main

import "fmt"

type Human struct {
	name string
	age int
	gender string
}

func (h Human)describe() {
	fmt.Println(h.name, "is", h.age, "years old", "and is a", h.gender)
}

func main() {
	h:= Human{"Jason", 29, "Male"}
	h.describe()
}
