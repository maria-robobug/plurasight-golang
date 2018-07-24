package main

import (
	"fmt"
)

type jumper interface {
	jump() string
}

type (
	gopher struct {
		name string
		age  int
	}

	horse struct {
		name   string
		weight int
	}
)

func main() {
	jumperList := getList()

	for _, jumper := range jumperList {
		fmt.Println(jumper.jump())
	}
}

func (g gopher) jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	}

	return g.name + " can still jump"
}

func (h horse) jump() string {
	if h.weight > 2500 {
		return h.name + " is too heavy to jump!"
	}

	return h.name + " can jump, neigh!!"
}

func getList() []jumper {
	pepp := &gopher{name: "Pepp", age: 25}
	noodles := &gopher{name: "Noodles", age: 90}
	storm := &horse{name: "Storm", weight: 2000}

	return []jumper{pepp, noodles, storm}
}
