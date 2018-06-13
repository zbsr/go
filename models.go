package main

type UserModel struct {
	Id            int64  `xorm:"pk autoincr"`
	Name          string
}

func (u *UserModel) TableName() string {
	return "user"
}
