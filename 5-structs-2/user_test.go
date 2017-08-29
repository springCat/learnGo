package structs

import (
	"testing"
)

func Test_user(t *testing.T) {

		u2 := &SuperUser{
				User:User{
						Id:   1,
						Name: "小明",
						Age:  16,
				},
				Permerssion:"Permerssion",
		}
		t.Logf("Test_user name3:%s", u2.Name)
		t.Logf("Test_user name3:%s", u2.User.Name)
		t.Logf("Test_user name3:%s", u2.getName())
		t.Logf("Test_user name3:%s", u2.User.getName())
		t.Logf("Test_user Permerssion:%s", u2.Permerssion)
		t.Logf("Test_user Permerssion:%s", u2.getPermerssion())
		t.Log("-------小明--------")

		u3 := new(SuperUser)
		u3.Name = "小明"
		t.Logf("Test_user name3:%s", u3.getName())
		u3.User.Name = "小明plus"
		t.Logf("Test_user name3:%s", u3.getName())

		t.Log("------小明plus---------")
}
