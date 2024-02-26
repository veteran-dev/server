package city

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/veteran-dev/server/api/v1"
	"github.com/veteran-dev/server/middleware"
)

type CityDataRouter struct {
}

// InitCityDataRouter 初始化 城市数据 路由信息
func (s *CityDataRouter) InitCityDataRouter(Router *gin.RouterGroup) {
	cdRouter := Router.Group("cd").Use(middleware.OperationRecord())
	cdRouterWithoutRecord := Router.Group("cd")
	var cdApi = v1.ApiGroupApp.CityApiGroup.CityDataApi
	{
		cdRouter.POST("createCityData", cdApi.CreateCityData)             // 新建城市数据
		cdRouter.DELETE("deleteCityData", cdApi.DeleteCityData)           // 删除城市数据
		cdRouter.DELETE("deleteCityDataByIds", cdApi.DeleteCityDataByIds) // 批量删除城市数据
		cdRouter.PUT("updateCityData", cdApi.UpdateCityData)              // 更新城市数据
	}
	{
		cdRouterWithoutRecord.GET("findCityData", cdApi.FindCityData)       // 根据ID获取城市数据
		cdRouterWithoutRecord.GET("getCityDataList", cdApi.GetCityDataList) // 获取城市数据列表
	}
}
