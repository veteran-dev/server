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
	ParentID int64 `json:"parantID"`
}
