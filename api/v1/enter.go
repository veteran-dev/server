package v1

import (
	"github.com/veteran-dev/server/api/v1/carCombination"
	"github.com/veteran-dev/server/api/v1/city"
	"github.com/veteran-dev/server/api/v1/cityCarCombination"
	"github.com/veteran-dev/server/api/v1/order"
	"github.com/veteran-dev/server/api/v1/system"
	"github.com/veteran-dev/server/api/v1/user"
	"github.com/veteran-dev/server/api/v1/web"
)

type ApiGroup struct {
	WebApiGroup                web.ApiGroup
	SystemApiGroup             system.ApiGroup
	CityApiGroup               city.ApiGroup
	UserApiGroup               user.ApiGroup
	OrderApiGroup              order.ApiGroup
	CarCombinationApiGroup     carCombination.ApiGroup
	CityCarCombinationApiGroup cityCarCombination.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
