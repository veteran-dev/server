package cityCarCombination

import (
	"github.com/5asp/gin-vue-admin/server/global"
    "github.com/5asp/gin-vue-admin/server/model/cityCarCombination"
    cityCarCombinationReq "github.com/5asp/gin-vue-admin/server/model/cityCarCombination/request"
    "github.com/5asp/gin-vue-admin/server/model/common/response"
    "github.com/5asp/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/5asp/gin-vue-admin/server/utils"
)

type CityCarCombinationApi struct {
}

var cccService = service.ServiceGroupApp.CityCarCombinationServiceGroup.CityCarCombinationService


// CreateCityCarCombination 创建车型城市组合
// @Tags CityCarCombination
// @Summary 创建车型城市组合
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cityCarCombination.CityCarCombination true "创建车型城市组合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ccc/createCityCarCombination [post]
func (cccApi *CityCarCombinationApi) CreateCityCarCombination(c *gin.Context) {
	var ccc cityCarCombination.CityCarCombination
	err := c.ShouldBindJSON(&ccc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ccc.CreatedBy = utils.GetUserID(c)

	if err := cccService.CreateCityCarCombination(&ccc); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCityCarCombination 删除车型城市组合
// @Tags CityCarCombination
// @Summary 删除车型城市组合
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cityCarCombination.CityCarCombination true "删除车型城市组合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ccc/deleteCityCarCombination [delete]
func (cccApi *CityCarCombinationApi) DeleteCityCarCombination(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := cccService.DeleteCityCarCombination(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCityCarCombinationByIds 批量删除车型城市组合
// @Tags CityCarCombination
// @Summary 批量删除车型城市组合
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ccc/deleteCityCarCombinationByIds [delete]
func (cccApi *CityCarCombinationApi) DeleteCityCarCombinationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := cccService.DeleteCityCarCombinationByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCityCarCombination 更新车型城市组合
// @Tags CityCarCombination
// @Summary 更新车型城市组合
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cityCarCombination.CityCarCombination true "更新车型城市组合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ccc/updateCityCarCombination [put]
func (cccApi *CityCarCombinationApi) UpdateCityCarCombination(c *gin.Context) {
	var ccc cityCarCombination.CityCarCombination
	err := c.ShouldBindJSON(&ccc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ccc.UpdatedBy = utils.GetUserID(c)

	if err := cccService.UpdateCityCarCombination(ccc); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCityCarCombination 用id查询车型城市组合
// @Tags CityCarCombination
// @Summary 用id查询车型城市组合
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cityCarCombination.CityCarCombination true "用id查询车型城市组合"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ccc/findCityCarCombination [get]
func (cccApi *CityCarCombinationApi) FindCityCarCombination(c *gin.Context) {
	ID := c.Query("ID")
	if reccc, err := cccService.GetCityCarCombination(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reccc": reccc}, c)
	}
}

// GetCityCarCombinationList 分页获取车型城市组合列表
// @Tags CityCarCombination
// @Summary 分页获取车型城市组合列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cityCarCombinationReq.CityCarCombinationSearch true "分页获取车型城市组合列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ccc/getCityCarCombinationList [get]
func (cccApi *CityCarCombinationApi) GetCityCarCombinationList(c *gin.Context) {
	var pageInfo cityCarCombinationReq.CityCarCombinationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cccService.GetCityCarCombinationInfoList(pageInfo); err != nil {
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
