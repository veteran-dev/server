package request

import (
	"time"

	"github.com/veteran-dev/server/model/common/request"
)

type CityCarCombinationSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}

type GetCarListReq struct {
	To        int    `json:"to"`
	From      int    `json:"from"`
	StartTime string `json:"startTime"` //出发时间
	Luggage   int64  `json:"luggage"`   //行李
	Child     int64  `json:"child"`     //小孩
	Aldult    int64  `json:"aldult"`    //成人
	request.PageInfo
}

type GetCarReq struct {
	ID    int   `json:"id"`
	ToID  int64 `json:"toId"`
	ToPID int64 `json:"toPid"`

	FromID    int64  `json:"fromId"`    //一级父id
	FromPID   int64  `json:"fromPid"`   //一级父id
	StartTime string `json:"startTime"` //出发时间
	Luggage   int64  `json:"luggage"`   //行李
	Child     int64  `json:"child"`     //小孩
	Aldult    int64  `json:"aldult"`    //成人
}

type PricePreviewReq struct {
	To        int    `json:"to"`
	From      int    `json:"from"`
	StartTime string `json:"startTime"` //出发时间
	Luggage   int64  `json:"luggage"`   //行李
	Child     int64  `json:"child"`     //小孩
	Aldult    int64  `json:"aldult"`    //成人
}

type PriceRulesReq struct {
}
