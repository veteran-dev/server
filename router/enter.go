package router

import (
	"github.com/veteran-dev/server/router/carCombination"
	"github.com/veteran-dev/server/router/city"
	"github.com/veteran-dev/server/router/cityCarCombination"
	"github.com/veteran-dev/server/router/general"
	"github.com/veteran-dev/server/router/order"
	"github.com/veteran-dev/server/router/system"
	"github.com/veteran-dev/server/router/user"
	"github.com/veteran-dev/server/router/web"
)

type RouterGroup struct {
	System             system.RouterGroup
	City               city.RouterGroup
	User               user.RouterGroup
	Order              order.RouterGroup
	CarCombination     carCombination.RouterGroup
	CityCarCombination cityCarCombination.RouterGroup
	General            general.GeneralRouter
	Web                web.WebRouter
}

var RouterGroupApp = new(RouterGroup)
