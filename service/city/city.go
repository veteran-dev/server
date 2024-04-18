package city

import (
	"sort"

	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/city"
	cityReq "github.com/veteran-dev/server/model/city/request"
	"gorm.io/gorm"
)

type CityDataService struct {
}

// CreateCityData 创建城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) CreateCityData(cd *city.City) (err error) {
	err = global.GVA_DB.Create(cd).Error
	return err
}

// DeleteCityData 删除城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) DeleteCityData(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&city.City{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&city.City{}, "id = ?", ID).Error; err != nil {
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
		if err := tx.Model(&city.City{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&city.City{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCityData 更新城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) UpdateCityData(cd city.City) (err error) {
	err = global.GVA_DB.Save(&cd).Error
	return err
}

// GetCityData 根据ID获取城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) GetCityData(ID string) (cd city.City, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cd).Error
	return
}

// GetCityData 根据Name获取城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) GetCityDataByNameAndPid(name []string, pid string) (list []city.City, err error) {
	err = global.GVA_DB.Debug().Where("name IN ? and pid = ?", name, pid).Find(&list).Error
	return
}

// GetCityDataInfoList 分页获取城市数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (cdService *CityDataService) GetCityDataInfoList(info cityReq.CityDataSearch) (list []city.City, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&city.City{})
	var cds []city.City
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
	db := global.GVA_DB.Model(&city.City{})
	var cds []city.City

	err = db.Find(&cds).Error
	if len(cds) > 0 {
		result = make(map[uint]string)
		for _, v := range cds {
			result[v.ID] = v.Name
		}
	}
	return result, err
}
func (cdService *CityDataService) City() (result city.CityDataList, err error) {
	// 创建db
	db := global.GVA_DB.Model(&city.City{})
	var cds []city.City
	err = db.Debug().Find(&cds).Error
	if len(cds) > 0 {
		cityDataListMap := make(map[string][]city.Cities)
		alphabetList := make([]string, 0)
		recommends := make([]city.Cities, 0)
		for _, v := range cds {
			if v.Letter != "" {
				if _, ok := cityDataListMap[v.Letter]; !ok {
					alphabetList = append(alphabetList, v.Letter)
				}
				if v.Hot == 1 {
					recommends = append(recommends, city.Cities{
						ID:        int(v.ID),
						Name:      v.Name,
						Letter:    v.Letter,
						Latitude:  v.Lat,
						Longitude: v.Lng,
					})
				}
				cityDataListMap[v.Letter] = append(cityDataListMap[v.Letter], city.Cities{
					ID:        int(v.ID),
					Name:      v.Name,
					Letter:    v.Letter,
					Latitude:  v.Lat,
					Longitude: v.Lng,
				})
			}
		}
		sort.Strings(alphabetList)

		cityListResult := make([]city.CityList, 0)

		for _, alphabet := range alphabetList {
			cityListResult = append(cityListResult, city.CityList{
				Idx:    alphabet,
				Cities: cityDataListMap[alphabet],
			})
		}

		cityDataListResult := city.CityDataList{
			Alphabet:  alphabetList,
			Recommend: recommends,
			CityList:  cityListResult,
		}

		return cityDataListResult, nil
	}
	return city.CityDataList{}, err
}

// func (cdService *CityDataService) GetParentCity(Pid int64) (cityName string) {

// 	db := global.GVA_DB.Model(&city.CityData{})
// 	db = db.Where("parent_id = ?", 0)
// 	var cds []city.CityData
// 	db.Debug().Find(&cds)
// 	cityMap := make(map[int64]string)
// 	if len(cds) != 0 {
// 		for _, v := range cds {
// 			cityMap[int64(v.ID)] = v.Name
// 		}

// 		if value, ok := cityMap[Pid]; ok {
// 			return value
// 		} else {
// 			top := &city.CityData{}
// 			db.Where("id = ?", Pid).First(top)
// 			return
// 		}
// 	}

// 	return
// }
