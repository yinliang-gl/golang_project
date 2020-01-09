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
	return fmt.Sprint("%d-%s", u.Person.Get(id), u.Person.GetName(id))
}
