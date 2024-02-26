// 自动生成模板CarCombination
package carCombination

import (
	"github.com/veteran-dev/server/global"
)

// 车型组 结构体  CarCombination
type CarCombination struct {
	global.GVA_MODEL
	CombinationTitle string `json:"combinationTitle" form:"combinationTitle" gorm:"column:combination_title;comment:车型组标题;"binding:"required"` //车型组标题
	ModelName        string `json:"modelName" form:"modelName" gorm:"column:model_name;comment:汽车车型;"`                                         //汽车车型
	Level            string `json:"level" form:"level" gorm:"column:level;comment:车型级别;"`                                                      //车型级别
	Seats            *int   `json:"seats" form:"seats" gorm:"column:seats;comment:座位数量;"`                                                      //座位数量
	ChildSeats       *int   `json:"childSeats" form:"childSeats" gorm:"column:child_seats;comment:儿童座椅数目;"`                                    //儿童座椅数目
	LargeLuggage     *int   `json:"largeLuggage" form:"largeLuggage" gorm:"column:large_luggage;comment:24寸以上的箱包数量;"`                          //24寸以上的箱包数量
	Luggage          *int   `json:"luggage" form:"luggage" gorm:"column:luggage;comment:24寸以下箱包数目;"`                                           //24寸以下箱包数目
	Status           *int   `json:"status" form:"status" gorm:"column:status;comment:车型组状态;"`                                                  //车型组状态
	CreatedBy        uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 车型组 CarCombination自定义表名 prev_car_combination
func (CarCombination) TableName() string {
	return "prev_car_combination"
}
