package router

import (
	"github.com/5asp/gin-vue-admin/server/router/city"
	"github.com/5asp/gin-vue-admin/server/router/cityCarPrice"
	"github.com/5asp/gin-vue-admin/server/router/combination"
	"github.com/5asp/gin-vue-admin/server/router/model"
	"github.com/5asp/gin-vue-admin/server/router/modelGroup"
	"github.com/5asp/gin-vue-admin/server/router/order"
	"github.com/5asp/gin-vue-admin/server/router/price"
	"github.com/5asp/gin-vue-admin/server/router/system"
	"github.com/5asp/gin-vue-admin/server/router/user"
)

type RouterGroup struct {
	System       system.RouterGroup
	City         city.RouterGroup
	Model        model.RouterGroup
	Price        price.RouterGroup
	User         user.RouterGroup
	Order        order.RouterGroup
	ModelGroup   modelGroup.RouterGroup
	CityCarPrice cityCarPrice.RouterGroup
	Combination  combination.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
