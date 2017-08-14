package interfaces

import (
	"testing"
)

func Test_MapDemo(t *testing.T) {
	bird := Bird{1, "blackBird"}
	d := Duck(bird)
	t.Log(d.Swim())
}
