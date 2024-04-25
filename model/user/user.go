// 自动生成模板User
package user

import (
	"github.com/veteran-dev/server/global"
)

// 用户 结构体  User
type User struct {
	global.GVA_MODEL
	UserId      string `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"binding:"required"`     //用户ID
	Token       string `json:"token" form:"token" gorm:"column:token;comment:凭证;type:text;"binding:"required"` //凭证
	UserHistory []UserHistory
}

// TableName 用户 User自定义表名 prev_users
func (User) TableName() string {
	return "prev_users"
}
