// 自动生成模板CityData
package city

import (
	"github.com/5asp/gin-vue-admin/server/global"
)

// 城市数据 结构体  CityData
type CityData struct {
	global.GVA_MODEL
	Name      string   `json:"name" form:"name" gorm:"column:name;comment:城市名称;"`               //城市名称
	Pinyin    string   `json:"pinyin" form:"pinyin" gorm:"column:pinyin;comment:城市名拼音;"`        //城市名拼音
	Latitude  *float64 `json:"latitude" form:"latitude" gorm:"column:latitude;comment:经纬度;"`    //经纬度
	Longitude *float64 `json:"longitude" form:"longitude" gorm:"column:longitude;comment:经纬度;"` //经纬度
	Alphabet  string   `json:"alphabet" form:"alphabet" gorm:"column:alphabet;comment:首字母;"`    //首字母
	Hot       *bool    `json:"hot" form:"hot" gorm:"column:hot;comment:推荐;"`                    //推荐
	CreatedBy uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 城市数据 CityData自定义表名 prev_city
func (CityData) TableName() string {
	return "prev_city"
}

type CityList struct {
	Alphabet  []string `json:"alphabet"`
	Recommend []Cities
	CityList  []struct {
		Idx    string `json:"idx"`
		Cities []Cities
	} `json:"cityList"`
}

type Cities struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Pinyin    string  `json:"pinyin"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hot       int     `json:"hot"`
}
