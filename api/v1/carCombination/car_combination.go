package carCombination

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/carCombination"
	carCombinationReq "github.com/5asp/gin-vue-admin/server/model/carCombination/request"
	"github.com/5asp/gin-vue-admin/server/model/common/response"
	"github.com/5asp/gin-vue-admin/server/service"
	"github.com/5asp/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CarCombinationApi struct {
}

var carcService = service.ServiceGroupApp.CarCombinationServiceGroup.CarCombinationService

// CreateCarCombination 创建车型组
// @Tags CarCombination
// @Summary 创建车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body carCombination.CarCombination true "创建车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /carc/createCarCombination [post]
func (carcApi *CarCombinationApi) CreateCarCombination(c *gin.Context) {
	var carc carCombination.CarCombination
	err := c.ShouldBindJSON(&carc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	carc.CreatedBy = utils.GetUserID(c)

	if err := carcService.CreateCarCombination(&carc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCarCombination 删除车型组
// @Tags CarCombination
// @Summary 删除车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body carCombination.CarCombination true "删除车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carc/deleteCarCombination [delete]
func (carcApi *CarCombinationApi) DeleteCarCombination(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := carcService.DeleteCarCombination(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCarCombinationByIds 批量删除车型组
// @Tags CarCombination
// @Summary 批量删除车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /carc/deleteCarCombinationByIds [delete]
func (carcApi *CarCombinationApi) DeleteCarCombinationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := carcService.DeleteCarCombinationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCarCombination 更新车型组
// @Tags CarCombination
// @Summary 更新车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body carCombination.CarCombination true "更新车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carc/updateCarCombination [put]
func (carcApi *CarCombinationApi) UpdateCarCombination(c *gin.Context) {
	var carc carCombination.CarCombination
	err := c.ShouldBindJSON(&carc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	carc.UpdatedBy = utils.GetUserID(c)

	if err := carcService.UpdateCarCombination(carc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCarCombination 用id查询车型组
// @Tags CarCombination
// @Summary 用id查询车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query carCombination.CarCombination true "用id查询车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carc/findCarCombination [get]
func (carcApi *CarCombinationApi) FindCarCombination(c *gin.Context) {
	ID := c.Query("ID")
	if recarc, err := carcService.GetCarCombination(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recarc": recarc}, c)
	}
}

// GetCarCombinationList 分页获取车型组列表
// @Tags CarCombination
// @Summary 分页获取车型组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query carCombinationReq.CarCombinationSearch true "分页获取车型组列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carc/getCarCombinationList [get]
func (carcApi *CarCombinationApi) GetCarCombinationList(c *gin.Context) {
	var pageInfo carCombinationReq.CarCombinationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := carcService.GetCarCombinationInfoList(pageInfo); err != nil {
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

// GetCarList 获取车型组数据列表
// @Tags GeneralData
// @Summary 获取车型组数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /general/carList [get]
func (carcApi *CarCombinationApi) GetCarList(c *gin.Context) {
	if result, err := carcService.GetCarList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}
