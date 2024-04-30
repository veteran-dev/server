package response

import (
	"github.com/veteran-dev/server/model/cityCarCombination"
)

// 调用url的
type OrderResp struct {
	Code       string `json:"code"`
	SubCode    bool   `json:"sub_code"`
	SubMsg     bool   `json:"sub_msg"`
	Msg        string `json:"msg"`
	OutTradeNo string `json:"out_trade_no"`
	TradeNo    string `json:"trade_no"`
}

type OrderCreateResp struct {
	StartTime    string                                `json:"startTime"`    //出发时间
	CancelAt     string                                `json:"cancelAt"`     //取消截止日期
	ToCityName   string                                `json:"toCityName"`   //到达城市
	FromCityName string                                `json:"fromCityName"` //起点城市
	FromLocation string                                `json:"fromLocation"` //起点定位
	ToLocation   string                                `json:"toLocation"`   //到达定位
	Price        int64                                 `json:"price"`        //总价
	Car          cityCarCombination.CityCarCombination `json:"car"`          //车型组
	OrderSerial  string                                `json:"orderSerial"`  //订单编号
}

type OrderCompleteResp struct {
	OrderSerial string `json:"orderSerial"`
}

type OrderDetailResp struct {
	Car   cityCarCombination.CityCarCombination `json:"car"`
	Order Order                                 `json:"order"`
}

type Order struct {
	Appointment     string `json:"appointment"`     //预约时间
	StartTime       string `json:"startTime"`       //出发时间
	FromCity        string `json:"fromCity"`        //开始城市id
	ToCity          string `json:"toCity"`          //抵达城市ID
	FromArea        string `json:"fromArea"`        //开始地点
	ToArea          string `json:"toArea"`          //抵达区域
	OrderSerial     string `json:"orderSerial"`     //订单编号
	Passenger       string `json:"passenger"`       //乘车人
	PassengerMobile string `json:"passengerMobile"` //乘车人联系
	TotalPrice      *int   `json:"totalPrice"`      //订单金额
	SubPrice        *int   `json:"subPrice"`        //订单金额
	CarModel        *int   `json:"carModel"`        //车型组
	Status          *int   `json:"status"`          //订单状态
	CancelReason    *int   `json:"cancelReason"`    //取消原因
	CancelAt        string `json:"cancelAt"`        //取消截止日期
}

type ReasonResp struct {
	ReasonCode int `json:"reasonCode"`
	ReasonMsg  int `json:"reasonMsg"`
}
