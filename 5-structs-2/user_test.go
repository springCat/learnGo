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
		//u.Permerssion = "Permerssion" error
		t.Logf("Test_user name1:%s", u.Name)
		t.Logf("Test_user name1:%s", u.getName())
		t.Log("---------------")

		u1 := new(SuperUser)
		u1.Name = "小明1"
		u1.Permerssion = "Permerssion"
		t.Logf("Test_user name2:%s", u1.Name)
		t.Logf("Test_user name2:%s", u1.getName())
		t.Logf("Test_user Permerssion:%s", u1.Permerssion)
		t.Logf("Test_user Permerssion:%s", u1.getPermerssion())
		t.Log("---------------")

		u2 := &SuperUser{
				User:User{
						Id:   1,
						Name: "小明2",
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
		t.Log("---------------")

		u3 := new(SuperUser)
		u3.Name = "小明3"
		u3.Permerssion = "Permerssion3"

		t.Logf("Test_user name3:%s", u3.Name)
		t.Logf("Test_user name3:%s", u3.User.Name)
		t.Logf("Test_user name3:%s", u3.getName())
		t.Logf("Test_user name3:%s", u3.User.getName())
		t.Logf("Test_user Permerssion:%s", u3.Permerssion)
		t.Logf("Test_user Permerssion:%s", u3.getPermerssion())
		t.Log("---------------")
}
