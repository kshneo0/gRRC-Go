package main

import (
	"fmt"
)

type speaker interface {
	speak()
}

type person struct {
	name  string
	email string
}

func (p person) speak() {
	fmt.Println(p.name)
}

type cat struct {
	name string
}

func (c cat) speak() {
	fmt.Println("Meow")
}

func saySomething(s speaker) {
	s.speak()
}

func main() {
	p := person{name: "Jack", email: "s@d.com"}
	saySomething(p)
	c := cat{}
	saySomething(c)
}
