// 自动生成模板Order
package order

import (
	"time"

	"github.com/veteran-dev/server/global"
)

// 订单 结构体  Order
type Order struct {
	global.GVA_MODEL
	Appointment     *time.Time `json:"appointment" form:"appointment" gorm:"column:appointment;comment:预约时间;"`               //预约时间
	FromCity        *int       `json:"fromCity" form:"fromCity" gorm:"column:from_city;comment:开始城市id;"`                     //开始城市id
	ToCity          *int       `json:"toCity" form:"toCity" gorm:"column:to_city;comment:抵达城市ID;"`                           //抵达城市ID
	FromArea        string     `json:"fromArea" form:"fromArea" gorm:"column:from_area;comment:开始地点;"`                       //开始地点
	ToArea          string     `json:"toArea" form:"toArea" gorm:"column:to_area;comment:抵达区域;"`                             //抵达区域
	OrderSerial     string     `json:"orderSerial" form:"orderSerial" gorm:"column:order_serial;comment:订单编号;"`              //订单编号
	Passenger       string     `json:"passenger" form:"passenger" gorm:"column:passenger;comment:乘车人;"`                      //乘车人
	PassengerMobile string     `json:"passengerMobile" form:"passengerMobile" gorm:"column:passenger_mobile;comment:乘车人联系;"` //乘车人联系
	Price           *int       `json:"price" form:"price" gorm:"column:price;comment:订单金额;"`                                 //订单金额
	CarModel        *int       `json:"carModel" form:"carModel" gorm:"column:car_model;comment:车型组;"`                        //车型组
	Status          *int       `json:"status" form:"status" gorm:"column:status;comment:订单状态;"`                              //订单状态
	CancelReason    *int       `json:"cancelReason" form:"cancelReason" gorm:"column:cancel_reason;comment:取消原因;"`           //取消原因
	CreatedBy       uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 订单 Order自定义表名 prev_order
func (Order) TableName() string {
	return "prev_order"
}
