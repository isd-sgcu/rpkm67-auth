package model

import constants "github.com/isd-sgcu/rpkm67-auth/internal/constant"

type User struct {
	Base
	StudentId string         `json:"student_id" gorm:"tinytext;unique"`
	Password  string         `json:"password" gorm:"tinytext"`
	Firstname string         `json:"firstname" gorm:"tinytext"`
	Lastname  string         `json:"lastname" gorm:"tinytext"`
	Tel       string         `json:"tel" gorm:"tinytext"`
	Role      constants.Role `json:"role" gorm:"tinytext"`
}
