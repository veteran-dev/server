package carCombination

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/carCombination"
	carCombinationReq "github.com/5asp/gin-vue-admin/server/model/carCombination/request"
	"gorm.io/gorm"
)

type CarCombinationService struct {
}

// CreateCarCombination 创建车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (carcService *CarCombinationService) CreateCarCombination(carc *carCombination.CarCombination) (err error) {
	err = global.GVA_DB.Create(carc).Error
	return err
}

// DeleteCarCombination 删除车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (carcService *CarCombinationService) DeleteCarCombination(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&carCombination.CarCombination{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&carCombination.CarCombination{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCarCombinationByIds 批量删除车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (carcService *CarCombinationService) DeleteCarCombinationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&carCombination.CarCombination{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&carCombination.CarCombination{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCarCombination 更新车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (carcService *CarCombinationService) UpdateCarCombination(carc carCombination.CarCombination) (err error) {
	err = global.GVA_DB.Save(&carc).Error
	return err
}

// GetCarCombination 根据ID获取车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (carcService *CarCombinationService) GetCarCombination(ID string) (carc carCombination.CarCombination, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&carc).Error
	return
}

// GetCarCombinationInfoList 分页获取车型组记录
// Author [piexlmax](https://github.com/piexlmax)
func (carcService *CarCombinationService) GetCarCombinationInfoList(info carCombinationReq.CarCombinationSearch) (list []carCombination.CarCombination, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&carCombination.CarCombination{})
	var carcs []carCombination.CarCombination
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
	orderMap["seats"] = true
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

	err = db.Find(&carcs).Error
	return carcs, total, err
}

func (carcService *CarCombinationService) GetCarList() (result map[uint]string, err error) {
	// 创建db
	db := global.GVA_DB.Model(&carCombination.CarCombination{})
	var carcs []carCombination.CarCombination

	err = db.Find(&carcs).Error
	if len(carcs) > 0 {
		result = make(map[uint]string)
		for _, v := range carcs {
			result[v.ID] = v.CombinationTitle
		}
	}
	return result, err
}
