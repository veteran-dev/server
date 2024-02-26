package order

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/veteran-dev/server/api/v1"
	"github.com/veteran-dev/server/middleware"
)

type OrderRouter struct {
}

// InitOrderRouter 初始化 订单 路由信息
func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	oRouter := Router.Group("o").Use(middleware.OperationRecord())
	oRouterWithoutRecord := Router.Group("o")
	var oApi = v1.ApiGroupApp.OrderApiGroup.OrderApi
	{
		oRouter.POST("createOrder", oApi.CreateOrder)             // 新建订单
		oRouter.DELETE("deleteOrder", oApi.DeleteOrder)           // 删除订单
		oRouter.DELETE("deleteOrderByIds", oApi.DeleteOrderByIds) // 批量删除订单
		oRouter.PUT("updateOrder", oApi.UpdateOrder)              // 更新订单
	}
	{
		oRouterWithoutRecord.GET("findOrder", oApi.FindOrder)       // 根据ID获取订单
		oRouterWithoutRecord.GET("getOrderList", oApi.GetOrderList) // 获取订单列表
	}
}
