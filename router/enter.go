package router

import (
	"github.com/5asp/gin-vue-admin/server/router/app"
	"github.com/5asp/gin-vue-admin/server/router/carCombination"
	"github.com/5asp/gin-vue-admin/server/router/city"
	"github.com/5asp/gin-vue-admin/server/router/cityCarCombination"
	"github.com/5asp/gin-vue-admin/server/router/general"
	"github.com/5asp/gin-vue-admin/server/router/order"
	"github.com/5asp/gin-vue-admin/server/router/system"
	"github.com/5asp/gin-vue-admin/server/router/user"
)

type RouterGroup struct {
	System             system.RouterGroup
	City               city.RouterGroup
	User               user.RouterGroup
	Order              order.RouterGroup
	CarCombination     carCombination.RouterGroup
	CityCarCombination cityCarCombination.RouterGroup
	General            general.GeneralRouter
	App                app.AppRouter
}

var RouterGroupApp = new(RouterGroup)
