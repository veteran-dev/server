package response

type PriceResp struct {
	ToCityName   string `json:"toCityName"`   //出发城市
	FromCityName string `json:"fromCityName"` //到达城市
	StartAt      string `json:"startAt"`      //出发时间
	BasePrice    int64  `json:"basePrice"`    //基础价格
	SubPrice     int64  `json:"subPrice"`     //加价
	TotalPrice   int64  `json:"totalPrice"`   //总价
}
