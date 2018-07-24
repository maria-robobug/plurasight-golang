package main

import (
	"fmt"
)

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func main() {
	gopher1 := &gopher{name: "Pepp", age: 18}
	gopher2 := &gopher{name: "Noodles", age: 90}

	validateAge(gopher1)
	validateAge(gopher2)

	fmt.Println(gopher1)
	fmt.Println(gopher2)
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}
