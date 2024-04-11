package order

import (
	"github.com/gin-gonic/gin"
	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/common/response"
	"github.com/veteran-dev/server/model/order"
	orderReq "github.com/veteran-dev/server/model/order/request"
	"github.com/veteran-dev/server/service"
	"github.com/veteran-dev/server/utils"
	"go.uber.org/zap"
)

type OrderApi struct {
}

var oService = service.ServiceGroupApp.OrderServiceGroup.OrderService

// CreateOrder 创建订单
//	@Tags		Order
//	@Summary	创建订单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		order.Order	true	"创建订单"
//	@Success	200		{string}	string		"{"success":true,"data":{},"msg":"创建成功"}"
//	@Router		/o/createOrder [post]
func (oApi *OrderApi) CreateOrder(c *gin.Context) {
	var o order.Order
	err := c.ShouldBindJSON(&o)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	o.CreatedBy = utils.GetUserID(c)

	if err := oService.CreateOrder(&o); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除订单
//	@Tags		Order
//	@Summary	删除订单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		order.Order	true	"删除订单"
//	@Success	200		{string}	string		"{"success":true,"data":{},"msg":"删除成功"}"
//	@Router		/o/deleteOrder [delete]
func (oApi *OrderApi) DeleteOrder(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := oService.DeleteOrder(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除订单
//	@Tags		Order
//	@Summary	批量删除订单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Success	200	{string}	string	"{"success":true,"data":{},"msg":"批量删除成功"}"
//	@Router		/o/deleteOrderByIds [delete]
func (oApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := oService.DeleteOrderByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新订单
//	@Tags		Order
//	@Summary	更新订单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		order.Order	true	"更新订单"
//	@Success	200		{string}	string		"{"success":true,"data":{},"msg":"更新成功"}"
//	@Router		/o/updateOrder [put]
func (oApi *OrderApi) UpdateOrder(c *gin.Context) {
	var o order.Order
	err := c.ShouldBindJSON(&o)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	o.UpdatedBy = utils.GetUserID(c)

	if err := oService.UpdateOrder(o); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询订单
//	@Tags		Order
//	@Summary	用id查询订单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		order.Order	true	"用id查询订单"
//	@Success	200		{string}	string		"{"success":true,"data":{},"msg":"查询成功"}"
//	@Router		/o/findOrder [get]
func (oApi *OrderApi) FindOrder(c *gin.Context) {
	ID := c.Query("ID")
	if reo, err := oService.GetOrder(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reo": reo}, c)
	}
}

// GetOrderList 分页获取订单列表
//	@Tags		Order
//	@Summary	分页获取订单列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		orderReq.OrderSearch	true	"分页获取订单列表"
//	@Success	200		{string}	string					"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router		/o/getOrderList [get]
func (oApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo orderReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := oService.GetOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
