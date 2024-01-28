package combination

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/combination"
	combinationReq "github.com/5asp/gin-vue-admin/server/model/combination/request"
	"gorm.io/gorm"
)

type CombinationService struct {
}

// CreateCombination 创建城市车型价格记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcService *CombinationService) CreateCombination(pc *combination.Combination) (err error) {
	err = global.GVA_DB.Create(pc).Error
	return err
}

// DeleteCombination 删除城市车型价格记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcService *CombinationService) DeleteCombination(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&combination.Combination{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&combination.Combination{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCombinationByIds 批量删除城市车型价格记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcService *CombinationService) DeleteCombinationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&combination.Combination{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&combination.Combination{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCombination 更新城市车型价格记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcService *CombinationService) UpdateCombination(pc combination.Combination) (err error) {
	err = global.GVA_DB.Save(&pc).Error
	return err
}

// GetCombination 根据ID获取城市车型价格记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcService *CombinationService) GetCombination(ID string) (pc combination.Combination, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pc).Error
	return
}

// GetCombinationInfoList 分页获取城市车型价格记录
// Author [piexlmax](https://github.com/piexlmax)
func (pcService *CombinationService) GetCombinationInfoList(info combinationReq.CombinationSearch) (list []combination.Combination, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&combination.Combination{})
	var pcs []combination.Combination
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

	err = db.Find(&pcs).Error
	return pcs, total, err
}
