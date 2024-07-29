package handler

import (
	"0729/shop_srv/user_srv/global"
	"0729/shop_srv/user_srv/model"
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	userPb "github.com/weiliang0215/user_proto/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"
)

type UserService struct {
	userPb.UnimplementedUserServer
}

func ModelToResponse(user model.User) userPb.UserInfo {
	userInfo := userPb.UserInfo{
		Id:       int64(user.ID),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Age:      user.Age,
		Sex:      user.Sex,
	}

	return userInfo
}

func Paging(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}
		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}

}

func (u *UserService) GetUserList(ctx context.Context, in *userPb.PageInfoReq) (*userPb.GetUserListResp, error) {
	var user []model.User
	req := global.DB.Model(&model.User{}).Scopes(Paging(int(in.Page), int(in.Limit))).Find(&user)

	if req.Error != nil {
		return nil, status.Error(codes.NotFound, "数据库异常")
	}
	if req.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户列表为空")
	}
	var count int64

	res := global.DB.Table("users").Count(&count)

	if res.Error != nil {
		return nil, status.Error(codes.NotFound, "数据库异常")
	}

	userList := []*userPb.UserInfo{}
	for _, v := range user {
		response := ModelToResponse(v)

		userList = append(userList, &response)
	}
	return &userPb.GetUserListResp{
		Total: int32(count),
		Data:  userList,
	}, nil
}

func (u *UserService) GetUserInfoByMobile(ctx context.Context, in *userPb.GetUserInfoByMobileReq) (*userPb.UserInfo, error) {
	var user model.User
	req := global.DB.Model(&model.User{}).Where("mobile = ?", in.Mobile).First(&user)

	if req.Error != nil {
		return nil, status.Error(codes.NotFound, "数据库异常")
	}
	response := ModelToResponse(user)
	return &response, nil
}

func (u *UserService) CheckPassword(ctx context.Context, in *userPb.CheckPasswordReq) (*userPb.CheckPasswordResp, error) {
	options := &password.Options{10, 10000, 50, sha512.New}

	split := strings.Split(in.EncryptionPassword, "$")

	check := password.Verify(in.Password, split[2], split[3], options)

	return &userPb.CheckPasswordResp{Success: check}, nil
}
func (u *UserService) CreateUser(ctx context.Context, in *userPb.CreateUserReq) (*userPb.UserInfo, error) {
	var user model.User
	global.DB.Model(model.User{}).Where("mobile = ?", in.Mobile).First(&user)

	if user.ID > 0 {
		return nil, status.Error(codes.NotFound, "用户已存在")
	}

	options := &password.Options{10, 10000, 50, sha512.New}
	salt, encodedPwd := password.Encode(in.Password, options)
	user.Mobile = in.Mobile
	user.Username = in.Username
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	req := global.DB.Model(&model.User{}).Create(&user)
	if req.Error != nil {
		return nil, status.Error(codes.NotFound, "用户添加失败")
	}

	response := ModelToResponse(user)

	return &response, nil
}
func (u *UserService) UpdateUser(ctx context.Context, in *userPb.UpdateUserReq) (*userPb.UserInfo, error) {
	var user model.User
	req := global.DB.Model(&model.User{}).Where("id = ?", in.Id).First(&user)

	if req.Error != nil {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	user.Email = in.Email
	user.Age = in.Age
	user.Sex = in.Sex

	res := global.DB.Model(&model.User{
		Model: model.Model{ID: int32(in.Id)},
	}).Updates(&user)

	if res.Error != nil {
		return nil, status.Error(codes.NotFound, "更新用户资源失败")
	}
	response := ModelToResponse(user)

	return &response, nil
}
func (u *UserService) DeleteUser(ctx context.Context, in *userPb.DeleteUserReq) (*userPb.DeleteUserResp, error) {
	var user model.User
	req := global.DB.Model(&model.User{}).Where("mobile = ?", in.Mobile).First(&user)

	if req.Error != nil {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	res := global.DB.Model(&model.User{}).Where("mobile = ?", in.Mobile).Delete(&user)

	if res.Error != nil {
		return nil, status.Error(codes.NotFound, "用户删除失败")
	}
	return &userPb.DeleteUserResp{Success: true}, nil
}
