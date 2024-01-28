package combination

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/combination"
	combinationReq "github.com/5asp/gin-vue-admin/server/model/combination/request"
	"github.com/5asp/gin-vue-admin/server/model/common/response"
	"github.com/5asp/gin-vue-admin/server/service"
	"github.com/5asp/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CombinationApi struct {
}

var pcService = service.ServiceGroupApp.CombinationServiceGroup.CombinationService

// CreateCombination 创建城市车型价格
// @Tags Combination
// @Summary 创建城市车型价格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body combination.Combination true "创建城市车型价格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pc/createCombination [post]
func (pcApi *CombinationApi) CreateCombination(c *gin.Context) {
	var pc combination.Combination
	err := c.ShouldBindJSON(&pc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pc.CreatedBy = utils.GetUserID(c)

	if err := pcService.CreateCombination(&pc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCombination 删除城市车型价格
// @Tags Combination
// @Summary 删除城市车型价格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body combination.Combination true "删除城市车型价格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pc/deleteCombination [delete]
func (pcApi *CombinationApi) DeleteCombination(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := pcService.DeleteCombination(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCombinationByIds 批量删除城市车型价格
// @Tags Combination
// @Summary 批量删除城市车型价格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pc/deleteCombinationByIds [delete]
func (pcApi *CombinationApi) DeleteCombinationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := pcService.DeleteCombinationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCombination 更新城市车型价格
// @Tags Combination
// @Summary 更新城市车型价格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body combination.Combination true "更新城市车型价格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pc/updateCombination [put]
func (pcApi *CombinationApi) UpdateCombination(c *gin.Context) {
	var pc combination.Combination
	err := c.ShouldBindJSON(&pc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pc.UpdatedBy = utils.GetUserID(c)

	if err := pcService.UpdateCombination(pc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCombination 用id查询城市车型价格
// @Tags Combination
// @Summary 用id查询城市车型价格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query combination.Combination true "用id查询城市车型价格"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pc/findCombination [get]
func (pcApi *CombinationApi) FindCombination(c *gin.Context) {
	ID := c.Query("ID")
	if repc, err := pcService.GetCombination(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repc": repc}, c)
	}
}

// GetCombinationList 分页获取城市车型价格列表
// @Tags Combination
// @Summary 分页获取城市车型价格列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query combinationReq.CombinationSearch true "分页获取城市车型价格列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pc/getCombinationList [get]
func (pcApi *CombinationApi) GetCombinationList(c *gin.Context) {
	var pageInfo combinationReq.CombinationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := pcService.GetCombinationInfoList(pageInfo); err != nil {
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
