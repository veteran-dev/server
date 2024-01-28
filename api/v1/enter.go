package v1

import (
	"github.com/5asp/gin-vue-admin/server/api/v1/city"
	"github.com/5asp/gin-vue-admin/server/api/v1/cityCarPrice"
	"github.com/5asp/gin-vue-admin/server/api/v1/combination"
	"github.com/5asp/gin-vue-admin/server/api/v1/model"
	"github.com/5asp/gin-vue-admin/server/api/v1/modelGroup"
	"github.com/5asp/gin-vue-admin/server/api/v1/order"
	"github.com/5asp/gin-vue-admin/server/api/v1/price"
	"github.com/5asp/gin-vue-admin/server/api/v1/system"
	"github.com/5asp/gin-vue-admin/server/api/v1/user"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	CityApiGroup         city.ApiGroup
	ModelApiGroup        model.ApiGroup
	PriceApiGroup        price.ApiGroup
	UserApiGroup         user.ApiGroup
	OrderApiGroup        order.ApiGroup
	ModelGroupApiGroup   modelGroup.ApiGroup
	CityCarPriceApiGroup cityCarPrice.ApiGroup
	CombinationApiGroup  combination.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
