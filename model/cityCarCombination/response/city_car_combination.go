package response

type Car struct {
	Title        string `json:"title"`        //商务级别
	Level        string `json:"level"`        //级别
	Seats        int64  `json:"seats"`        //座位数
	Luggage      int64  `json:"luggage"`      //行李数
	Price        int64  `json:"price"`        //单价根据时间段计算变动的价格
	Brand        string `json:"brand"`        //车型
	ID           int64  `json:"id"`           //车型ID
	FromCityName string `json:"fromCityName"` //起点城市
	FromAreaID   int64  `json:"fromAreaID"`   //起点区域ID
	ToCityName   string `json:"toCityName"`   //到达城市
	ToAreaID     int64  `json:"toAreaID"`     //到达区域ID
	StartTime    string `json:"startTime"`    //出发时间
}

type CarDetailResp struct {
	Detail       CarDetail `json:"detail"`
	FromCityName string    `json:"fromCityName"` //起点城市
	FromAreaID   int64     `json:"fromAreaID"`   //起点区域ID
	ToCityName   string    `json:"toCityName"`   //到达城市
	ToAreaID     int64     `json:"toAreaID"`     //到达区域ID
	StartTime    string    `json:"startTime"`    //出发时间
	FromLocation string    `json:"fromLocation"` //起点定位
	ToLocation   string    `json:"toLocation"`   //起点定位
}
type CarDetail struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"` //商务级别
	Seats        int64  `json:"seats"` //座位数
	Brand        string `json:"brand"` //车型
	Level        string `json:"level"`
	Luggage      int64  `json:"luggage"`      //行李数
	LargeLuggage int64  `json:"largeLuggage"` //24寸以上行李数
	ChildSeats   int64  `json:"childSeats"`   //儿童座椅数
	Price        int64  `json:"price"`
}
type PricePreviewResp struct {
	Price     int64  `json:"price"` //单价根据时间段计算变动的价格
	BasePrice int64  `json:"basePrice"`
	SubPrice  int64  `json:"subPrice"`
	StartTime string `json:"startTime"` //出发时间
	To        string `json:"to"`
	From      string `json:"from"`
}
