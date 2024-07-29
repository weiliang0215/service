package model

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	Age       int64     `json:"age"`
	Sex       int32     `json:"sex"`
}
