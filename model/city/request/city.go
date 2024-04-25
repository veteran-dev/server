package request

import (
	"time"

	"github.com/veteran-dev/server/model/common/request"
)

type CityDataSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type CityDataReq struct {
}

type CityLocalReq struct {
	CityID   int64  `json:"cityID" form:"cityID"`
	Address  string `json:"address" form:"address"`
	PageSize int    `json:"pageSize"`
	Page     int    `json:"page"`
}

type GaodeReq struct {
	Key      string `json:"key"`
	Keywords string `json:"keywords"`
	Types    string `json:"types"`
	City     string `json:"city"`
	Children int    `json:"children"`
	Offset   int    `json:"offset"`
	Page     int    `json:"page"`
}

type CityListReq struct {
	Keyword string `json:"keyword"` //搜索城市关键词
}
