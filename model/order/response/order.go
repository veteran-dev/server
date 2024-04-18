package response

import (
	"github.com/veteran-dev/server/model/cityCarCombination"
	"github.com/veteran-dev/server/model/order"
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
	OrderSerial string `json:"order_serial"`
}

type OrderDetailResp struct {
	Car   cityCarCombination.CityCarCombination `json:"car"`
	Order order.Order                           `json:"order"`
}

type ReasonResp struct {
	ReasonCode int `json:"reasonCode"`
	ReasonMsg  int `json:"reasonMsg"`
}
