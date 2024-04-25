package user

import (
	"github.com/veteran-dev/server/api/v1"
	"github.com/veteran-dev/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// InitUserRouter 初始化 用户 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	uRouter := Router.Group("u").Use(middleware.OperationRecord())
	uRouterWithoutRecord := Router.Group("u")
	var uApi = v1.ApiGroupApp.UserApiGroup.UserApi
	{
		uRouter.POST("createUser", uApi.CreateUser)   // 新建用户
		uRouter.DELETE("deleteUser", uApi.DeleteUser) // 删除用户
		uRouter.DELETE("deleteUserByIds", uApi.DeleteUserByIds) // 批量删除用户
		uRouter.PUT("updateUser", uApi.UpdateUser)    // 更新用户
	}
	{
		uRouterWithoutRecord.GET("findUser", uApi.FindUser)        // 根据ID获取用户
		uRouterWithoutRecord.GET("getUserList", uApi.GetUserList)  // 获取用户列表
	}
}
