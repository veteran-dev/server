package city

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/city"
	cityReq "github.com/5asp/gin-vue-admin/server/model/city/request"
	"gorm.io/gorm"
)

type CityDataService struct {
}

// CreateCityData 创建城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) CreateCityData(cd *city.CityData) (err error) {
	err = global.GVA_DB.Create(cd).Error
	return err
}

// DeleteCityData 删除城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) DeleteCityData(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&city.CityData{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&city.CityData{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCityDataByIds 批量删除城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) DeleteCityDataByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&city.CityData{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&city.CityData{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCityData 更新城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) UpdateCityData(cd city.CityData) (err error) {
	err = global.GVA_DB.Save(&cd).Error
	return err
}

// GetCityData 根据ID获取城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) GetCityData(ID string) (cd city.CityData, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cd).Error
	return
}

// GetCityDataInfoList 分页获取城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) GetCityDataInfoList(info cityReq.CityDataSearch) (list []city.CityData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&city.CityData{})
	var cds []city.CityData
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
	orderMap["alphabet"] = true
	orderMap["hot"] = true
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

	err = db.Find(&cds).Error
	return cds, total, err
}

func (cdService *CityDataService) GetCityList() (result map[uint]string, err error) {
	// 创建db
	db := global.GVA_DB.Model(&city.CityData{})
	var cds []city.CityData

	err = db.Find(&cds).Error
	if len(cds) > 0 {
		result = make(map[uint]string)
		for _, v := range cds {
			result[v.ID] = v.Name
		}
	}
	return result, err
}

func (cdService *CityDataService) City() (result map[uint]string, err error) {
	// 创建db
	db := global.GVA_DB.Model(&city.CityData{})
	var cds []city.CityData

	err = db.Find(&cds).Error
	if len(cds) > 0 {
		result = make(map[uint]string)
		for _, v := range cds {
			result[v.ID] = v.Name
		}
	}
	return result, err
}
