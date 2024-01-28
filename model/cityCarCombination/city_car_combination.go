// 自动生成模板CityCarCombination
package cityCarCombination

import (
	"github.com/5asp/gin-vue-admin/server/global"
	
	
)

// 车型城市组合 结构体  CityCarCombination
type CityCarCombination struct {
 global.GVA_MODEL 
      From  *int `json:"from" form:"from" gorm:"column:from;comment:出发城市;"binding:"required"`  //出发城市 
      To  *int `json:"to" form:"to" gorm:"column:to;comment:到达城市;"binding:"required"`  //到达城市 
      Car  *int `json:"car" form:"car" gorm:"column:car;comment:车型ID;"binding:"required"`  //车型ID 
      BasePrice  *int `json:"basePrice" form:"basePrice" gorm:"column:base_price;comment:城市车型组合价格;"binding:"required"`  //城市车型组合价格 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:组合状态;"`  //组合状态 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 车型城市组合 CityCarCombination自定义表名 prev_city_car_combination
func (CityCarCombination) TableName() string {
  return "prev_city_car_combination"
}

