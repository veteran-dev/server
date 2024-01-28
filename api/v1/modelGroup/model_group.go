package modelGroup

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/common/response"
	"github.com/5asp/gin-vue-admin/server/model/modelGroup"
	modelGroupReq "github.com/5asp/gin-vue-admin/server/model/modelGroup/request"
	"github.com/5asp/gin-vue-admin/server/service"
	"github.com/5asp/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ModelGroupApi struct {
}

var mgService = service.ServiceGroupApp.ModelGroupServiceGroup.ModelGroupService

// CreateModelGroup 创建车型组
// @Tags ModelGroup
// @Summary 创建车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelGroup.ModelGroup true "创建车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mg/createModelGroup [post]
func (mgApi *ModelGroupApi) CreateModelGroup(c *gin.Context) {
	var mg modelGroup.ModelGroup
	err := c.ShouldBindJSON(&mg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mg.CreatedBy = utils.GetUserID(c)

	if err := mgService.CreateModelGroup(&mg); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteModelGroup 删除车型组
// @Tags ModelGroup
// @Summary 删除车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelGroup.ModelGroup true "删除车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mg/deleteModelGroup [delete]
func (mgApi *ModelGroupApi) DeleteModelGroup(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := mgService.DeleteModelGroup(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteModelGroupByIds 批量删除车型组
// @Tags ModelGroup
// @Summary 批量删除车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mg/deleteModelGroupByIds [delete]
func (mgApi *ModelGroupApi) DeleteModelGroupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := mgService.DeleteModelGroupByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateModelGroup 更新车型组
// @Tags ModelGroup
// @Summary 更新车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelGroup.ModelGroup true "更新车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mg/updateModelGroup [put]
func (mgApi *ModelGroupApi) UpdateModelGroup(c *gin.Context) {
	var mg modelGroup.ModelGroup
	err := c.ShouldBindJSON(&mg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mg.UpdatedBy = utils.GetUserID(c)

	if err := mgService.UpdateModelGroup(mg); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindModelGroup 用id查询车型组
// @Tags ModelGroup
// @Summary 用id查询车型组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query modelGroup.ModelGroup true "用id查询车型组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mg/findModelGroup [get]
func (mgApi *ModelGroupApi) FindModelGroup(c *gin.Context) {
	ID := c.Query("ID")
	if remg, err := mgService.GetModelGroup(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remg": remg}, c)
	}
}

// GetModelGroupList 分页获取车型组列表
// @Tags ModelGroup
// @Summary 分页获取车型组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query modelGroupReq.ModelGroupSearch true "分页获取车型组列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mg/getModelGroupList [get]
func (mgApi *ModelGroupApi) GetModelGroupList(c *gin.Context) {
	var pageInfo modelGroupReq.ModelGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mgService.GetModelGroupInfoList(pageInfo); err != nil {
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
