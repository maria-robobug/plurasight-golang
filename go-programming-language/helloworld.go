package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota
	B
	C
)

func main() {

	var s = Salutation{}
	s.name = "Bob"
	s.greeting = "Hello World!"

	fmt.Println(A, B, C)
	fmt.Println(PI)
	fmt.Println(Language)
}
