package request

import (
	"time"

	"github.com/veteran-dev/server/model/common/request"
)

type OrderSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type OrderCompleteReq struct {
	OrderSerial     string `json:"orderSerial"`     //订单号
	Passenger       string `json:"passenger"`       //乘车人
	PassengerMobile string `json:"passengerMobile"` //乘车人联系
	Code            string `json:"code"`            //支付码
	ChannelCode     string `json:"channelCode"`     //通道码
}

type OrderCreateReq struct {
	StartTime    string `json:"startTime"`    //出发时间
	ToCityName   string `json:"toCityName"`   //到达城市
	FromCityName string `json:"fromCityName"` //起点城市
	CarID        int64  `json:"carID"`        //车型组ID
	FromLocation string `json:"fromLocation"` //起点定位
	ToLocation   string `json:"toLocation"`   //到达定位
}

type PayCode struct {
	OrderID   int64  `json:"orderId"`
	Code      string `json:"code"`
	ReturnUrl string `json:"returnUrl"`
}

type OrderDetail struct {
	OrderSerial string `json:"orderSerial"` //订单号
}

type OrderUpdate struct {
	Passenger       string `json:"passenger"`       //乘车人
	PassengerMobile string `json:"passengerMobile"` //乘车人联系
	OrderSerial     string `json:"orderSerial"`     //订单号
}
type OrderCancelReq struct {
	OrderSerial string `json:"orderSerial"`  //订单号
	CancelReson int    `json:"cancelReason"` //取消原因
}
