package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int32 `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Username string `gorm:"column:username;type:varchar(50);comment:用户名" json:"username"`
	Password string `gorm:"column:password;type:varchar(200);comment:密码" json:"password"`
	Mobile   string `gorm:"column:mobile;type:char(11);comment:手机号" json:"mobile"`
	Email    string `gorm:"column:email;type:varchar(100);comment:邮箱" json:"email"`
	Age      int64  `gorm:"column:age;type:int(11);comment:年龄" json:"age"`
	Sex      int32  `gorm:"column:sex;type:tinyint(2);comment:性别 1=女 2=男" json:"sex"`
}
