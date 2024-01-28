package modelGroup

import (
	v1 "github.com/5asp/gin-vue-admin/server/api/v1"
	"github.com/5asp/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ModelGroupRouter struct {
}

// InitModelGroupRouter 初始化 车型组 路由信息
func (s *ModelGroupRouter) InitModelGroupRouter(Router *gin.RouterGroup) {
	mgRouter := Router.Group("mg").Use(middleware.OperationRecord())
	mgRouterWithoutRecord := Router.Group("mg")
	var mgApi = v1.ApiGroupApp.ModelGroupApiGroup.ModelGroupApi
	{
		mgRouter.POST("createModelGroup", mgApi.CreateModelGroup)             // 新建车型组
		mgRouter.DELETE("deleteModelGroup", mgApi.DeleteModelGroup)           // 删除车型组
		mgRouter.DELETE("deleteModelGroupByIds", mgApi.DeleteModelGroupByIds) // 批量删除车型组
		mgRouter.PUT("updateModelGroup", mgApi.UpdateModelGroup)              // 更新车型组
	}
	{
		mgRouterWithoutRecord.GET("findModelGroup", mgApi.FindModelGroup)       // 根据ID获取车型组
		mgRouterWithoutRecord.GET("getModelGroupList", mgApi.GetModelGroupList) // 获取车型组列表
	}
}
