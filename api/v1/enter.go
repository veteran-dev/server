package v1

import (
	"github.com/5asp/gin-vue-admin/server/api/v1/carCombination"
	"github.com/5asp/gin-vue-admin/server/api/v1/city"
	"github.com/5asp/gin-vue-admin/server/api/v1/cityCarCombination"
	"github.com/5asp/gin-vue-admin/server/api/v1/order"
	"github.com/5asp/gin-vue-admin/server/api/v1/system"
	"github.com/5asp/gin-vue-admin/server/api/v1/user"
)

type ApiGroup struct {
	SystemApiGroup             system.ApiGroup
	CityApiGroup               city.ApiGroup
	UserApiGroup               user.ApiGroup
	OrderApiGroup              order.ApiGroup
	CarCombinationApiGroup     carCombination.ApiGroup
	CityCarCombinationApiGroup cityCarCombination.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
