package main

type UserModel struct {
	Id            int64  `xorm:"pk autoincr" json:"id"`
	Name          string `json:"name"`
}

func (u *UserModel) TableName() string {
	return "user"
}