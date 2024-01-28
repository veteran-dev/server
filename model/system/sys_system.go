package system

import (
	"github.com/5asp/gin-vue-admin/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
