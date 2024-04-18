package request

type PriceRulesReq struct {
	FromCityName string `json:"fromCityName"` //起点城市
	FromAreaID   int64  `json:"fromAreaID"`   //起点区域ID
	ToCityName   string `json:"toCityName"`   //到达城市
	ToAreaID     int64  `json:"toAreaID"`     //到达区域ID
	StartTime    string `json:"startTime"`    //出发时间
	CarModelID   int64  `json:"carModelID"`   //当前车型组ID
}
