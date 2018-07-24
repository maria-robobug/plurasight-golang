package main

import (
	"fmt"
)

type gopher struct {
	name string
	age  int
}

func main() {
	gopher1 := gopher{name: "Pepp", age: 25}
	gopher2 := gopher{name: "Noodles", age: 90}

	fmt.Println(gopher1.jump())
	fmt.Println(gopher2.jump())

}

func (g gopher) jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	} else {
		return g.name + " can still jump"
	}
}
