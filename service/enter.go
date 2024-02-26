package service

import (
	"github.com/veteran-dev/server/service/carCombination"
	"github.com/veteran-dev/server/service/city"
	"github.com/veteran-dev/server/service/cityCarCombination"
	"github.com/veteran-dev/server/service/order"
	"github.com/veteran-dev/server/service/system"
	"github.com/veteran-dev/server/service/user"
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
