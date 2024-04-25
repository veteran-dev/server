// 自动生成模板UserHistory
package user

import (
	"github.com/veteran-dev/server/global"
)

// 用户历史 结构体  UserHistory
type UserHistory struct {
	global.GVA_MODEL
	UserId      *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"binding:"required"`                          //用户ID
	HistoryData string `json:"historyData" form:"historyData" gorm:"column:history_data;comment:历史数据;type:text;"binding:"required"` //历史数据
	HistoryType *int   `json:"historyType" form:"historyType" gorm:"column:history_type;comment:历史类型;"binding:"required"`           //历史类型
}

// TableName 用户历史 UserHistory自定义表名 prev_user_history
func (UserHistory) TableName() string {
	return "prev_user_history"
}
