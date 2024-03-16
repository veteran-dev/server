package web

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/veteran-dev/server/api/v1"
)

type WebRouter struct {
}

// TODO
// 所有城市列表接口
// 退订接口
// 订单详情接口 联系信息{乘客姓名+联系方式} ,订单明细{起点城市+终点城市+基础价格+出发时间+加价金额}，车型做信息{车型详情座位等}, 订单详情{下单时间+起始城市+开始结束地点}，订单状态{订单状态，订单编号}
// 提交订单接口
// https://devcr.dachema.net/cmdcapp
// InitAppRouter 初始化 前台接口信息
func (s *WebRouter) InitWebRouter(Router *gin.RouterGroup) {
	lRouter := Router.Group("web")
	var lApi = v1.ApiGroupApp.WebApiGroup.WebApi
	{
		lRouter.GET("login", lApi.Login)
	}

	gRouter := Router.Group("web")
	// Use(middleware.UserJWT())

	var wApi = v1.ApiGroupApp.WebApiGroup.WebApi
	{

		gRouter.POST("city/list", wApi.GetCityList) // 城市列表
		gRouter.POST("car/list", wApi.GetCarList)   // 选车
		// gRouter.POST("quote/pricing/rules", wApi.PriceRules) // 资费规则
		gRouter.POST("order/create", wApi.OrderComplete) //下单
		gRouter.GET("order/detail", wApi.OrderDetail)    //订单详情
		gRouter.POST("order/update", wApi.OrderUpdate)   //订单修改
		gRouter.POST("order/cancel", wApi.OrderCancel)   //订单取消
		gRouter.GET("car/detail", wApi.CarDetail)        //车型组详情
	}
}
