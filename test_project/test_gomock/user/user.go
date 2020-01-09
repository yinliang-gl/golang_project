package user

import (
	"fmt"
	"github.com/yinliang-gl/golang_project/test_project/test_gomock/person"
)

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) string {
	res := fmt.Sprintf("%d-%s", u.Person.GetId(id), u.Person.GetName(id))
	//res := fmt.Sprintf("%d-%s", 10, "u.Person.GetName(id)")
	return res
}
