package carCombination

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/veteran-dev/server/api/v1"
	"github.com/veteran-dev/server/middleware"
)

type CarCombinationRouter struct {
}

// InitCarCombinationRouter 初始化 车型组 路由信息
func (s *CarCombinationRouter) InitCarCombinationRouter(Router *gin.RouterGroup) {
	carcRouter := Router.Group("carc").Use(middleware.OperationRecord())
	carcRouterWithoutRecord := Router.Group("carc")
	var carcApi = v1.ApiGroupApp.CarCombinationApiGroup.CarCombinationApi
	{
		carcRouter.POST("createCarCombination", carcApi.CreateCarCombination)             // 新建车型组
		carcRouter.DELETE("deleteCarCombination", carcApi.DeleteCarCombination)           // 删除车型组
		carcRouter.DELETE("deleteCarCombinationByIds", carcApi.DeleteCarCombinationByIds) // 批量删除车型组
		carcRouter.PUT("updateCarCombination", carcApi.UpdateCarCombination)              // 更新车型组
	}
	{
		carcRouterWithoutRecord.GET("findCarCombination", carcApi.FindCarCombination)       // 根据ID获取车型组
		carcRouterWithoutRecord.GET("getCarCombinationList", carcApi.GetCarCombinationList) // 获取车型组列表
	}
}
