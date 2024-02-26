package cityCarCombination

import (
	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/cityCarCombination"
	cityCarCombinationReq "github.com/veteran-dev/server/model/cityCarCombination/request"
	"gorm.io/gorm"
)

type CityCarCombinationService struct {
}

// CreateCityCarCombination 创建车型城市组合记录
// Author [piexlmax](https://github.com/piexlmax)
func (cccService *CityCarCombinationService) CreateCityCarCombination(ccc *cityCarCombination.CityCarCombination) (err error) {
	err = global.GVA_DB.Create(ccc).Error
	return err
}

// DeleteCityCarCombination 删除车型城市组合记录
// Author [piexlmax](https://github.com/piexlmax)
func (cccService *CityCarCombinationService) DeleteCityCarCombination(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cityCarCombination.CityCarCombination{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&cityCarCombination.CityCarCombination{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCityCarCombinationByIds 批量删除车型城市组合记录
// Author [piexlmax](https://github.com/piexlmax)
func (cccService *CityCarCombinationService) DeleteCityCarCombinationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cityCarCombination.CityCarCombination{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&cityCarCombination.CityCarCombination{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCityCarCombination 更新车型城市组合记录
// Author [piexlmax](https://github.com/piexlmax)
func (cccService *CityCarCombinationService) UpdateCityCarCombination(ccc cityCarCombination.CityCarCombination) (err error) {
	err = global.GVA_DB.Save(&ccc).Error
	return err
}

// GetCityCarCombination 根据ID获取车型城市组合记录
// Author [piexlmax](https://github.com/piexlmax)
func (cccService *CityCarCombinationService) GetCityCarCombination(ID string) (ccc cityCarCombination.CityCarCombination, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ccc).Error
	return
}

// GetCityCarCombinationInfoList 分页获取车型城市组合记录
// Author [piexlmax](https://github.com/piexlmax)
func (cccService *CityCarCombinationService) GetCityCarCombinationInfoList(info cityCarCombinationReq.CityCarCombinationSearch) (list []cityCarCombination.CityCarCombination, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cityCarCombination.CityCarCombination{})
	var cccs []cityCarCombination.CityCarCombination
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

	err = db.Find(&cccs).Error
	return cccs, total, err
}
