package response

type OrderResp struct {
	Code       string `json:"code"`
	SubCode    bool   `json:"sub_code"`
	SubMsg     bool   `json:"sub_msg"`
	Msg        string `json:"msg"`
	OutTradeNo string `json:"out_trade_no"`
	TradeNo    string `json:"trade_no"`
}
