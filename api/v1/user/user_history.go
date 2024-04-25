package user

import (
	"github.com/veteran-dev/server/global"
    "github.com/veteran-dev/server/model/user"
    userReq "github.com/veteran-dev/server/model/user/request"
    "github.com/veteran-dev/server/model/common/response"
    "github.com/veteran-dev/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserHistoryApi struct {
}

var uhService = service.ServiceGroupApp.UserServiceGroup.UserHistoryService


// CreateUserHistory 创建用户历史
// @Tags UserHistory
// @Summary 创建用户历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserHistory true "创建用户历史"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /uh/createUserHistory [post]
func (uhApi *UserHistoryApi) CreateUserHistory(c *gin.Context) {
	var uh user.UserHistory
	err := c.ShouldBindJSON(&uh)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := uhService.CreateUserHistory(&uh); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserHistory 删除用户历史
// @Tags UserHistory
// @Summary 删除用户历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserHistory true "删除用户历史"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /uh/deleteUserHistory [delete]
func (uhApi *UserHistoryApi) DeleteUserHistory(c *gin.Context) {
	ID := c.Query("ID")
	if err := uhService.DeleteUserHistory(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserHistoryByIds 批量删除用户历史
// @Tags UserHistory
// @Summary 批量删除用户历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /uh/deleteUserHistoryByIds [delete]
func (uhApi *UserHistoryApi) DeleteUserHistoryByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := uhService.DeleteUserHistoryByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserHistory 更新用户历史
// @Tags UserHistory
// @Summary 更新用户历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserHistory true "更新用户历史"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /uh/updateUserHistory [put]
func (uhApi *UserHistoryApi) UpdateUserHistory(c *gin.Context) {
	var uh user.UserHistory
	err := c.ShouldBindJSON(&uh)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := uhService.UpdateUserHistory(uh); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserHistory 用id查询用户历史
// @Tags UserHistory
// @Summary 用id查询用户历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.UserHistory true "用id查询用户历史"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /uh/findUserHistory [get]
func (uhApi *UserHistoryApi) FindUserHistory(c *gin.Context) {
	ID := c.Query("ID")
	if reuh, err := uhService.GetUserHistory(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuh": reuh}, c)
	}
}

// GetUserHistoryList 分页获取用户历史列表
// @Tags UserHistory
// @Summary 分页获取用户历史列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.UserHistorySearch true "分页获取用户历史列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /uh/getUserHistoryList [get]
func (uhApi *UserHistoryApi) GetUserHistoryList(c *gin.Context) {
	var pageInfo userReq.UserHistorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := uhService.GetUserHistoryInfoList(pageInfo); err != nil {
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
