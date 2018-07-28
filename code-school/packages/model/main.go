package model

type Jumper interface {
	Jump() string
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

func (g gopher) Jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	}

	return g.name + " can still jump"
}

func (h horse) Jump() string {
	if h.weight > 2500 {
		return h.name + " is too heavy to jump!"
	}

	return h.name + " can jump, neigh!!"
}

// GetList - Returns a list of jumpers
func GetList() []Jumper {
	pepp := &gopher{name: "Pepp", age: 25}
	noodles := &gopher{name: "Noodles", age: 90}
	storm := &horse{name: "Storm", weight: 2000}

	return []Jumper{pepp, noodles, storm}
}
