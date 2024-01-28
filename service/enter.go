package service

import (
	"github.com/5asp/gin-vue-admin/server/service/carCombination"
	"github.com/5asp/gin-vue-admin/server/service/city"
	"github.com/5asp/gin-vue-admin/server/service/cityCarCombination"
	"github.com/5asp/gin-vue-admin/server/service/order"
	"github.com/5asp/gin-vue-admin/server/service/system"
	"github.com/5asp/gin-vue-admin/server/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup             system.ServiceGroup
	CityServiceGroup               city.ServiceGroup
	UserServiceGroup               user.ServiceGroup
	OrderServiceGroup              order.ServiceGroup
	CarCombinationServiceGroup     carCombination.ServiceGroup
	CityCarCombinationServiceGroup cityCarCombination.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
