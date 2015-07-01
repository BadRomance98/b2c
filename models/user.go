package models

import (
	"time"
)

type User struct {
	Id       int64
	Username string
	Passwd   string
	Role     int
	SubDate  time.Time
	Status   int
}

func GetUserInfo(uname string) (*User, bool, error) {
	user := new(User)
	has, err := Xorm.Where("username=?", uname).Get(user)
	return user, has, err
}
