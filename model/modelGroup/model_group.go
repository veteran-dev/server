// 自动生成模板ModelGroup
package modelGroup

import (
	"github.com/5asp/gin-vue-admin/server/global"

	"gorm.io/datatypes"
)

// 车型组 结构体  ModelGroup
type ModelGroup struct {
	global.GVA_MODEL
	Name         string         `json:"name" form:"name" gorm:"column:name;comment:车型组名称;"`                         //车型组名称
	Seat         *int           `json:"seat" form:"seat" gorm:"column:seat;comment:座位;"`                            //座位
	Baggage      *int           `json:"baggage" form:"baggage" gorm:"column:baggage;comment:行李数;"`                  //行李数
	Child        *int           `json:"child" form:"child" gorm:"column:child;comment:儿童座椅;"`                       //儿童座椅
	Desc         datatypes.JSON `json:"desc" form:"desc" gorm:"column:desc;comment:附加;"`                            //附加
	LargeBaggage *int           `json:"largeBaggage" form:"largeBaggage" gorm:"column:large_baggage;comment:最大行李;"` //最大行李
	Status       *int           `json:"status" form:"status" gorm:"column:status;comment:车型组状态;"`                   //车型组状态
	CreatedBy    uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 车型组 ModelGroup自定义表名 prev_model_group
func (ModelGroup) TableName() string {
	return "prev_model_group"
}
