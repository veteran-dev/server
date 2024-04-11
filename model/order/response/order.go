package response

import (
	"github.com/veteran-dev/server/model/carCombination"
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

type OrderCompleteResp struct {
	OrderSerial string `json:"order_serial"`
}

type OrderDetailResp struct {
	Car   carCombination.CarCombination `json:"car"`
	Order order.Order                   `json:"order"`
}
