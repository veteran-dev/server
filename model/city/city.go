// 自动生成模板CityData
package city

import (
	"github.com/veteran-dev/server/global"
)

// 城市数据 结构体  City
type City struct {
	global.GVA_MODEL
	Name   string  `json:"name" form:"name" gorm:"column:name;comment:城市名称;"`      //城市名称
	Letter string  `json:"letter" form:"letter" gorm:"column:letter;comment:首字母;"` //首字母
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Hot    int     `json:"hot"`
	Pid    int     `json:"pid"`
}

func (City) TableName() string {
	return "citys"
}

type CityDataList struct {
	Alphabet  []string   `json:"alphabet"`
	Recommend []Cities   `json:"recommend"`
	CityList  []CityList `json:"cityList"`
}

type CityList struct {
	Idx    string   `json:"idx"`
	Cities []Cities `json:"cities"`
}
type Cities struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Letter    string  `json:"letter"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GaodeData struct {
	Suggestion struct {
		Keywords []interface{} `json:"keywords"`
		Cities   []interface{} `json:"cities"`
	} `json:"suggestion"`
	Count    string `json:"count"`
	Infocode string `json:"infocode"`
	Pois     []Pois `json:"pois"`
	Status   string `json:"status"`
	Info     string `json:"info"`
}

type Pois struct {
	Parent     []interface{} `json:"parent"`
	Address    string        `json:"address"`
	Distance   []interface{} `json:"distance"`
	Pname      string        `json:"pname"`
	Importance []interface{} `json:"importance"`
	BizExt     struct {
		Cost   []interface{} `json:"cost"`
		Rating []interface{} `json:"rating"`
	} `json:"biz_ext"`
	BizType  []interface{} `json:"biz_type"`
	Cityname string        `json:"cityname"`
	Type     string        `json:"type"`
	Photos   []struct {
		Title []interface{} `json:"title"`
		URL   string        `json:"url"`
	} `json:"photos"`
	Typecode  string        `json:"typecode"`
	Shopinfo  string        `json:"shopinfo"`
	Poiweight []interface{} `json:"poiweight"`
	Childtype []interface{} `json:"childtype"`
	Adname    string        `json:"adname"`
	Children  []struct {
		Typecode string `json:"typecode"`
		Address  string `json:"address"`
		Distance string `json:"distance"`
		Subtype  string `json:"subtype"`
		Sname    string `json:"sname"`
		Name     string `json:"name"`
		Location string `json:"location"`
		ID       string `json:"id"`
	} `json:"children,omitempty"`
	Name     string        `json:"name"`
	Location string        `json:"location"`
	Tel      string        `json:"tel"`
	Shopid   []interface{} `json:"shopid"`
	ID       string        `json:"id"`
}

type LocalArea struct {
	Type     string `json:"type"`     //poi类型
	AreaName string `json:"areaName"` //区域名称
	Name     string `json:"name"`     //poi Name
	Address  string `json:"address"`
	Count    int64  `json:"count"`
}
