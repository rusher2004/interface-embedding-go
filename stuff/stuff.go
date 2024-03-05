package stuff

import "fmt"

type Thinger interface {
	DoThisThing() (Thing, error)
	DoThatThing() error
}

type Thing struct {
	ID string
}

func DoSomeStuff(t Thinger) (Thing, error) {
	out, err := t.DoThisThing()
	if err != nil {
		return Thing{}, fmt.Errorf("error doing some stuff: %w", err)
	}

	return out, nil
}

func DoOtherStuff(t Thinger) error {
	err := t.DoThatThing()
	if err != nil {
		return fmt.Errorf("error doing that thing: %w", err)
	}

	return nil
}
