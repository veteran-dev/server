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

type UserApi struct {
}

var uService = service.ServiceGroupApp.UserServiceGroup.UserService


// CreateUser 创建用户
// @Tags User
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.User true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /u/createUser [post]
func (uApi *UserApi) CreateUser(c *gin.Context) {
	var u user.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := uService.CreateUser(&u); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUser 删除用户
// @Tags User
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.User true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /u/deleteUser [delete]
func (uApi *UserApi) DeleteUser(c *gin.Context) {
	ID := c.Query("ID")
	if err := uService.DeleteUser(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserByIds 批量删除用户
// @Tags User
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /u/deleteUserByIds [delete]
func (uApi *UserApi) DeleteUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := uService.DeleteUserByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUser 更新用户
// @Tags User
// @Summary 更新用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.User true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /u/updateUser [put]
func (uApi *UserApi) UpdateUser(c *gin.Context) {
	var u user.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := uService.UpdateUser(u); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUser 用id查询用户
// @Tags User
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.User true "用id查询用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /u/findUser [get]
func (uApi *UserApi) FindUser(c *gin.Context) {
	ID := c.Query("ID")
	if reu, err := uService.GetUser(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reu": reu}, c)
	}
}

// GetUserList 分页获取用户列表
// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.UserSearch true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /u/getUserList [get]
func (uApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo userReq.UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := uService.GetUserInfoList(pageInfo); err != nil {
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
