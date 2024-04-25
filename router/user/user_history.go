package user

import (
	"github.com/veteran-dev/server/api/v1"
	"github.com/veteran-dev/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserHistoryRouter struct {
}

// InitUserHistoryRouter 初始化 用户历史 路由信息
func (s *UserHistoryRouter) InitUserHistoryRouter(Router *gin.RouterGroup) {
	uhRouter := Router.Group("uh").Use(middleware.OperationRecord())
	uhRouterWithoutRecord := Router.Group("uh")
	var uhApi = v1.ApiGroupApp.UserApiGroup.UserHistoryApi
	{
		uhRouter.POST("createUserHistory", uhApi.CreateUserHistory)   // 新建用户历史
		uhRouter.DELETE("deleteUserHistory", uhApi.DeleteUserHistory) // 删除用户历史
		uhRouter.DELETE("deleteUserHistoryByIds", uhApi.DeleteUserHistoryByIds) // 批量删除用户历史
		uhRouter.PUT("updateUserHistory", uhApi.UpdateUserHistory)    // 更新用户历史
	}
	{
		uhRouterWithoutRecord.GET("findUserHistory", uhApi.FindUserHistory)        // 根据ID获取用户历史
		uhRouterWithoutRecord.GET("getUserHistoryList", uhApi.GetUserHistoryList)  // 获取用户历史列表
	}
}
