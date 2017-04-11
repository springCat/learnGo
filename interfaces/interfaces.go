package interfaces

import "fmt"

type Duck interface {
	Swim() string
}

type Bird struct {
	id   int
	name string
}

func (self Bird ) Swim() string {
	return fmt.Sprintf("%s,is swimming ", self.name)
}
