package main

import "fmt"

type Human struct { // this is a struct which has two fields name and age
	name string
	age int
}

func (h Human)describe() { // here we take two arguments h & human. h is the instance for human which initializes the field values
	fmt.Println(h.name, "is", h.age, "years old") // describe method uses human to take objects and print the message
}

func main() {
	h:= Human{name: "Jackson", age: 22} // initializes the values
	h.describe() //here we call describe that prints the message with h's initialize values
}
