package interfaces

import (
		"testing"
)

func Test_MapDemo(t *testing.T) {

		bird := Bird{1, "blackBird"}
		t.Log(bird.Swim())
		t.Log("-------bird---------")

		d := Duck(bird)
		t.Log(d.Swim())
		t.Log("--------Duck--------")

		flyBird := FlyBird{
				Bird: bird,
		}
		t.Log(flyBird.Swim())
		t.Log(flyBird.Fly())
		t.Log("--------flyBird--------")

		guaGuaBird := GuaGuaBird{
				Duck:bird,
		}
		t.Log(guaGuaBird.Swim())
		//t.Log(guaGuaBird.Fly())
		t.Log(guaGuaBird.GuaGua())
		t.Log("------guaGuaBird----------")


		guaGuaBird2 := GuaGuaBird{
				Duck:flyBird,
		}
		t.Log(guaGuaBird2.Swim())
		//t.Log(guaGuaBird.Fly())
		t.Log(guaGuaBird2.GuaGua())


		if dd,ok := (guaGuaBird2.Duck).(FlyBird);ok {
				t.Log(dd.Fly())
		}
		//t.Log(guaGuaBird2.(Duck).Swim()) //error
		t.Log("-------guaGuaBird2---------")

}
