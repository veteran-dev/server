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

type OrderComplete struct {
	ToID            int64  `json:"toId"`
	ToPID           int64  `json:"toPid"`
	FromID          int64  `json:"fromId"`          //一级父id
	FromPID         int64  `json:"fromPid"`         //一级父id
	StartTime       string `json:"startTime"`       //出发时间
	Luggage         int64  `json:"luggage"`         //行李
	Child           int64  `json:"child"`           //小孩
	Aldult          int64  `json:"aldult"`          //成人
	Passenger       string `json:"passenger"`       //乘车人
	PassengerMobile string `json:"passengerMobile"` //乘车人联系
	CarID           int64  `json:"carId"`
	Price           int64  `json:"price"` //价格
	Code            string `json:"code"`
	ChannelCode     string `json:"channelCode"`
}

type PayCode struct {
	OrderID   int64  `json:"orderId"`
	Code      string `json:"code"`
	ReturnUrl string `json:"returnUrl"`
}

