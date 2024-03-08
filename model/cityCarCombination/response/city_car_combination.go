package response

type Car struct {
	Title   string `json:"title"`   //商务级别
	Level   string `json:"level"`   //级别
	Seats   int64  `json:"seats"`   //座位数
	Luggage int64  `json:"luggage"` //行李数
	Price   int64  `json:"price"`   //单价根据时间段计算变动的价格
	Brand   string `json:"brand"`   //车型
	ID      int64  `json:"id"`      //车型ID
}

type CarDetail struct {
	Title        string `json:"title"` //商务级别
	Seats        int64  `json:"seats"` //座位数
	Brand        string `json:"brand"` //车型
	Level        string `json:"level"`
	Luggage      int64  `json:"luggage"`      //行李数
	LargeLuggage int64  `json:"largeLuggage"` //24寸以上行李数
	ChildSeats   int64  `json:"childSeats"`   //儿童座椅数

	Price int64 `json:"price"` //单价根据时间段计算变动的价格
	To int64 
}
