package cityCarCombination

import (
	"github.com/5asp/gin-vue-admin/server/api/v1"
	"github.com/5asp/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CityCarCombinationRouter struct {
}

// InitCityCarCombinationRouter 初始化 车型城市组合 路由信息
func (s *CityCarCombinationRouter) InitCityCarCombinationRouter(Router *gin.RouterGroup) {
	cccRouter := Router.Group("ccc").Use(middleware.OperationRecord())
	cccRouterWithoutRecord := Router.Group("ccc")
	var cccApi = v1.ApiGroupApp.CityCarCombinationApiGroup.CityCarCombinationApi
	{
		cccRouter.POST("createCityCarCombination", cccApi.CreateCityCarCombination)   // 新建车型城市组合
		cccRouter.DELETE("deleteCityCarCombination", cccApi.DeleteCityCarCombination) // 删除车型城市组合
		cccRouter.DELETE("deleteCityCarCombinationByIds", cccApi.DeleteCityCarCombinationByIds) // 批量删除车型城市组合
		cccRouter.PUT("updateCityCarCombination", cccApi.UpdateCityCarCombination)    // 更新车型城市组合
	}
	{
		cccRouterWithoutRecord.GET("findCityCarCombination", cccApi.FindCityCarCombination)        // 根据ID获取车型城市组合
		cccRouterWithoutRecord.GET("getCityCarCombinationList", cccApi.GetCityCarCombinationList)  // 获取车型城市组合列表
	}
}
