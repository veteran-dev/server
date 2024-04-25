package user

import (
	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/user"
	userReq "github.com/veteran-dev/server/model/user/request"
)

type UserService struct {
}

// CreateUser 创建用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UserService) CreateUser(u *user.User) (err error) {
	err = global.GVA_DB.Create(u).Error
	return err
}

// DeleteUser 删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UserService) DeleteUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&user.User{}, "id = ?", ID).Error
	return err
}

// DeleteUserByIds 批量删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UserService) DeleteUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]user.User{}, "id in ?", IDs).Error
	return err
}

// UpdateUser 更新用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UserService) UpdateUser(u user.User) (err error) {
	err = global.GVA_DB.Save(&u).Error
	return err
}

// GetUser 根据ID获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UserService) GetUser(ID string) (u user.User, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&u).Error
	return
}

// FindOrCreate 创建或者更新Token
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UserService) FindOrCreateUser(u *user.User) (err error) {
	find, err := uService.GetUser(u.UserId)
	if find.UserId == "" {
		err = global.GVA_DB.Where("user_id = ?", u.UserId).Update("token", u.Token).Error
	} else {
		err = global.GVA_DB.Create(u).Error

	}
	return
}

// GetUserInfoList 分页获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UserService) GetUserInfoList(info userReq.UserSearch) (list []user.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.User{})
	var us []user.User
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&us).Error
	return us, total, err
}
