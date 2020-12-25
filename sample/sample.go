package main

import "fmt"

type VampireProps interface {
	getAge() int
	calcAge() int
}

type Vampire struct {
	location  string
	birthDate int
	deathDate int
	weakness  []string
}

func (v Vampire) getAge() int {
	return v.calcAge()
}

func (v Vampire) calcAge() int {
	return v.deathDate - v.birthDate
}

func main() {
	dracula := Vampire{
		location:  "Transylvania",
		birthDate: 1428,
		deathDate: 1476,
		weakness:  []string{"Sunlight", "Garlic"},
	}
	fmt.Println(dracula)
	fmt.Println(dracula.getAge())
}
