package service

import (
	"github.com/5asp/gin-vue-admin/server/service/city"
	"github.com/5asp/gin-vue-admin/server/service/cityCarPrice"
	"github.com/5asp/gin-vue-admin/server/service/combination"
	"github.com/5asp/gin-vue-admin/server/service/model"
	"github.com/5asp/gin-vue-admin/server/service/modelGroup"
	"github.com/5asp/gin-vue-admin/server/service/order"
	"github.com/5asp/gin-vue-admin/server/service/price"
	"github.com/5asp/gin-vue-admin/server/service/system"
	"github.com/5asp/gin-vue-admin/server/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	CityServiceGroup         city.ServiceGroup
	ModelServiceGroup        model.ServiceGroup
	PriceServiceGroup        price.ServiceGroup
	UserServiceGroup         user.ServiceGroup
	OrderServiceGroup        order.ServiceGroup
	ModelGroupServiceGroup   modelGroup.ServiceGroup
	CityCarPriceServiceGroup cityCarPrice.ServiceGroup
	CombinationServiceGroup  combination.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
