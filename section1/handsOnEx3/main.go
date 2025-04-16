package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type speaker interface {
	speak()
}

func (p person) speak() {
	fmt.Println("This is a person speaking")
}

func (sa secretAgent) speak() {
	fmt.Println("I will kill you!")
}

func hooman(h speaker) {
	fmt.Printf("this %v is of type %T, saying: ", h, h)
	h.speak()
}

func main() {
	p1 := person{
		"Ah",
		"Gow",
	}

	sa1 := secretAgent{
		person{"James",
			"Bond"},
		true,
	}

	hooman(p1)
	hooman(sa1)
}
