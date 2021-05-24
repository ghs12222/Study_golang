package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	school string
	grade  int
}

func (p *Student) greeting() {
	fmt.Println("hello Student")
}

func main() {
	var s Student

	s.Person.greeting()
	s.greeting()
}
