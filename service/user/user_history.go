package user

import (
	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/user"
    userReq "github.com/veteran-dev/server/model/user/request"
)

type UserHistoryService struct {
}

// CreateUserHistory 创建用户历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (uhService *UserHistoryService) CreateUserHistory(uh *user.UserHistory) (err error) {
	err = global.GVA_DB.Create(uh).Error
	return err
}

// DeleteUserHistory 删除用户历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (uhService *UserHistoryService)DeleteUserHistory(ID string) (err error) {
	err = global.GVA_DB.Delete(&user.UserHistory{},"id = ?",ID).Error
	return err
}

// DeleteUserHistoryByIds 批量删除用户历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (uhService *UserHistoryService)DeleteUserHistoryByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]user.UserHistory{},"id in ?",IDs).Error
	return err
}

// UpdateUserHistory 更新用户历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (uhService *UserHistoryService)UpdateUserHistory(uh user.UserHistory) (err error) {
	err = global.GVA_DB.Save(&uh).Error
	return err
}

// GetUserHistory 根据ID获取用户历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (uhService *UserHistoryService)GetUserHistory(ID string) (uh user.UserHistory, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&uh).Error
	return
}

// GetUserHistoryInfoList 分页获取用户历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (uhService *UserHistoryService)GetUserHistoryInfoList(info userReq.UserHistorySearch) (list []user.UserHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&user.UserHistory{})
    var uhs []user.UserHistory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&uhs).Error
	return  uhs, total, err
}
