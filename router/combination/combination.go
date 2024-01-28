package combination

import (
	v1 "github.com/5asp/gin-vue-admin/server/api/v1"
	"github.com/5asp/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CombinationRouter struct {
}

// InitCombinationRouter 初始化 城市车型价格 路由信息
func (s *CombinationRouter) InitCombinationRouter(Router *gin.RouterGroup) {
	pcRouter := Router.Group("pc").Use(middleware.OperationRecord())
	pcRouterWithoutRecord := Router.Group("pc")
	var pcApi = v1.ApiGroupApp.CombinationApiGroup.CombinationApi
	{
		pcRouter.POST("createCombination", pcApi.CreateCombination)             // 新建城市车型价格
		pcRouter.DELETE("deleteCombination", pcApi.DeleteCombination)           // 删除城市车型价格
		pcRouter.DELETE("deleteCombinationByIds", pcApi.DeleteCombinationByIds) // 批量删除城市车型价格
		pcRouter.PUT("updateCombination", pcApi.UpdateCombination)              // 更新城市车型价格
	}
	{
		pcRouterWithoutRecord.GET("findCombination", pcApi.FindCombination)       // 根据ID获取城市车型价格
		pcRouterWithoutRecord.GET("getCombinationList", pcApi.GetCombinationList) // 获取城市车型价格列表
	}
}
