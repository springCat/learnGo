package structs

type User struct {
	Id   int64
	Name string
	Age  int32
}

func (u User) getId() int64 {
	return u.Id
}

func (u User) getName() string {
	return u.Name
}

func (u User) getAge() int32 {
	return u.Age
}

type SuperUser struct {
		User
		Permerssion string
}
//type SuperUser struct {
//		User  User
//		Permerssion string
//}

func (u SuperUser) getPermerssion() string {
		return u.Permerssion
}


