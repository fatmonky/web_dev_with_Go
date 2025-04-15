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

func (p person) pSpeak() {
	fmt.Println("This is a person speaking")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("I will kill you!")
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

	fmt.Println(p1.firstname)
	p1.pSpeak()
	fmt.Println(sa1.licenseToKill)
	fmt.Println(sa1.firstname)
	sa1.saSpeak()
	sa1.person.pSpeak()
}
