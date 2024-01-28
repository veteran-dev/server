// 自动生成模板Combination
package combination

import (
	"github.com/5asp/gin-vue-admin/server/global"
)

// 城市车型价格 结构体  Combination
type Combination struct {
	global.GVA_MODEL
	From      *int `json:"from" form:"from" gorm:"column:from;comment:出发城市;"`                 //出发城市
	To        *int `json:"to" form:"to" gorm:"column:to;comment:到达城市;"`                       //到达城市
	Model     *int `json:"model" form:"model" gorm:"column:model;comment:车型;"`                //车型
	BasePrice *int `json:"basePrice" form:"basePrice" gorm:"column:base_price;comment:基础价格;"` //基础价格
	Status    *int `json:"status" form:"status" gorm:"column:status;comment:组合状态;"`           //组合状态
	CreatedBy uint `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 城市车型价格 Combination自定义表名 prev_combination
func (Combination) TableName() string {
	return "prev_combination"
}
