package router

import (
	"github.com/5asp/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
