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
	FromCityName string `json:"fromCityName"` //起点城市ID
	FromAreaID   int64  `json:"fromAreaID"`   //起点区域ID
	FromLocation string `json:"fromLocation"` //起点定位
	ToCityName   string `json:"toCityName"`   //到达城市ID
	ToAreaID     int64  `json:"toAreaID"`     //到达区域ID
	ToLocation   string `json:"toLocation"`   //到达定位
	StartTime    string `json:"startTime"`    //出发时间
	Luggage      int64  `json:"luggage"`      //行李
	Child        int64  `json:"child"`        //小孩
	Aldult       int64  `json:"aldult"`       //成人
	PageSize     int    `json:"pageSize"`     //偏移量
	Page         int    `json:"page"`         //当前页面
}

type GetCarReq struct {
	ID           int64  `json:"id"`
	StartTime    string `json:"startTime"`    //出发时间
	FromCityName string `json:"fromCityName"` //起点城市
	ToCityName   string `json:"toCityName"`   //到达城市
	ToLocation   string `json:"toLocation"`   //到达定位
	FromLocation string `json:"fromLocation"` //起点定位
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
