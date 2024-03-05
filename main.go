package main

import (
	"fmt"
	"log"

	"github.com/rusher2004/interface-embedding-go/stuff"
)

type myThing struct {
	id string
}

func (m myThing) DoThisThing() (stuff.Thing, error) {
	return stuff.Thing{ID: m.id}, nil
}

func (m myThing) DoThatThing() error {
	return nil
}

func main() {
	m := myThing{id: "123"}

	t, err := stuff.DoSomeStuff(m)
	if err != nil {
		log.Fatalf("error doing some stuff: %v", err)
	}

	fmt.Println("did some stuff and got thing:", t.ID)

	if err := stuff.DoOtherStuff(m); err != nil {
		log.Fatalf("error doing other stuff: %v", err)
	}

	fmt.Println("did other stuff")
}
