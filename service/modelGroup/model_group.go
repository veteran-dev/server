package modelGroup

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/modelGroup"
	modelGroupReq "github.com/5asp/gin-vue-admin/server/model/modelGroup/request"
	"gorm.io/gorm"
)

type ModelGroupService struct {
}

// CreateModelGroup 创建车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (mgService *ModelGroupService) CreateModelGroup(mg *modelGroup.ModelGroup) (err error) {
	err = global.GVA_DB.Create(mg).Error
	return err
}

// DeleteModelGroup 删除车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (mgService *ModelGroupService) DeleteModelGroup(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&modelGroup.ModelGroup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&modelGroup.ModelGroup{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteModelGroupByIds 批量删除车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (mgService *ModelGroupService) DeleteModelGroupByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&modelGroup.ModelGroup{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&modelGroup.ModelGroup{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateModelGroup 更新车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (mgService *ModelGroupService) UpdateModelGroup(mg modelGroup.ModelGroup) (err error) {
	err = global.GVA_DB.Save(&mg).Error
	return err
}

// GetModelGroup 根据ID获取车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (mgService *ModelGroupService) GetModelGroup(ID string) (mg modelGroup.ModelGroup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mg).Error
	return
}

// GetModelGroupInfoList 分页获取车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (mgService *ModelGroupService) GetModelGroupInfoList(info modelGroupReq.ModelGroupSearch) (list []modelGroup.ModelGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&modelGroup.ModelGroup{})
	var mgs []modelGroup.ModelGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["seat"] = true
	orderMap["baggage"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mgs).Error
	return mgs, total, err
}
