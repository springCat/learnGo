package structs

import (
	"testing"
)

func Test_user(t *testing.T) {
	u := &User{
		Id:   1,
		Name: "小明",
		Age:  16,
	}
	t.Logf("Test_user name1:", u.Name)

	u1 := new(User)
	u1.Name = "小明"
	t.Logf("Test_user name2:", u1.Name)
}
