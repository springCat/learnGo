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

type CanFly interface {
		Fly() string
}

type FlyBird struct {
		Bird
}

func (self FlyBird ) Fly() string {
		return fmt.Sprintf("%s,is flying ", self.name)
}

type GuaGuaBird struct {
		Duck
}

func (self GuaGuaBird ) GuaGua() string {
		return fmt.Sprintf("gua gua")
}