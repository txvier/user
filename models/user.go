package models

import (
	"github.com/txvier/base/txtime"
)

// User
//
// User Entity
//
// swagger:model User
type User struct {
	UId        string          `json:"uid"`
	Name       string          `json:"name" binding:"required"`
	Pwd        string          `json:"pwd" binding:"required"`
	CreateTime txtime.JSONTime `json:"create_time"`
}
