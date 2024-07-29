package forms

type UserRegisterForm struct {
	Mobile   string `json:"mobile" xml:"mobile" form:"mobile" binding:"required,mobile"`
	Username string `json:"username" xml:"username" form:"username" binding:"required"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}

type UserInfoForm struct {
	Mobile string `json:"mobile" xml:"mobile" form:"mobile" binding:"required,mobile"`
}

type UserLoginForm struct {
	Mobile   string `json:"mobile" xml:"mobile" form:"mobile" binding:"required,mobile"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}

type UserUpdateForm struct {
	Id    int64  `json:"id" xml:"id" form:"id" binding:"required"`
	Email string `json:"email" xml:"email" form:"email" binding:"required"`
	Age   int64  `json:"age" xml:"age" form:"age" binding:"required"`
	Sex   int32  `json:"sex" xml:"sex" form:"sex" binding:"required"`
}

type UserDeleteForm struct {
	Mobile string `json:"mobile" xml:"mobile" form:"mobile" binding:"required,mobile"`
}
