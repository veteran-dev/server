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
	gRouter := Router.Group("web")
	var wApi = v1.ApiGroupApp.WebApiGroup.WebApi
	{
		gRouter.GET("city/list", wApi.GetCityList) // 城市列表

	}
	// var cApi = v1.ApiGroupApp.CityApiGroup.CityDataApi
	// {
	// 	gRouter.GET("citys", cApi.Citys) // 城市列表
	// }
	// var carcApi = v1.ApiGroupApp.CarCombinationApiGroup.CarCombinationApi
	// {
	// 	gRouter.GET("getCarList", carcApi.GetCarList) //
	// }
}
