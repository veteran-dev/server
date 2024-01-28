package general

import (
	v1 "github.com/5asp/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type GeneralRouter struct {
}

// InitGeneralRouter 初始化 通用接口信息
func (s *GeneralRouter) InitGeneralRouter(Router *gin.RouterGroup) {
	gRouter := Router.Group("general")
	var cApi = v1.ApiGroupApp.CityApiGroup.CityDataApi
	{
		gRouter.GET("cityList", cApi.GetCityList) // 城市列表
	}
	var carcApi = v1.ApiGroupApp.CarCombinationApiGroup.CarCombinationApi
	{
		gRouter.GET("getCarList", carcApi.GetCarList) // 获取车型组列表
	}
}
