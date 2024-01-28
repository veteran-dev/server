package city

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/city"
	cityReq "github.com/5asp/gin-vue-admin/server/model/city/request"
	"github.com/5asp/gin-vue-admin/server/model/common/response"
	"github.com/5asp/gin-vue-admin/server/service"
	"github.com/5asp/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CityDataApi struct {
}

var cdService = service.ServiceGroupApp.CityServiceGroup.CityDataService

// CreateCityData 创建城市数据
// @Tags CityData
// @Summary 创建城市数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body city.CityData true "创建城市数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cd/createCityData [post]
func (cdApi *CityDataApi) CreateCityData(c *gin.Context) {
	var cd city.CityData
	err := c.ShouldBindJSON(&cd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cd.CreatedBy = utils.GetUserID(c)

	if err := cdService.CreateCityData(&cd); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCityData 删除城市数据
// @Tags CityData
// @Summary 删除城市数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body city.CityData true "删除城市数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cd/deleteCityData [delete]
func (cdApi *CityDataApi) DeleteCityData(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := cdService.DeleteCityData(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCityDataByIds 批量删除城市数据
// @Tags CityData
// @Summary 批量删除城市数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cd/deleteCityDataByIds [delete]
func (cdApi *CityDataApi) DeleteCityDataByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := cdService.DeleteCityDataByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCityData 更新城市数据
// @Tags CityData
// @Summary 更新城市数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body city.CityData true "更新城市数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cd/updateCityData [put]
func (cdApi *CityDataApi) UpdateCityData(c *gin.Context) {
	var cd city.CityData
	err := c.ShouldBindJSON(&cd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cd.UpdatedBy = utils.GetUserID(c)

	if err := cdService.UpdateCityData(cd); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCityData 用id查询城市数据
// @Tags CityData
// @Summary 用id查询城市数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query city.CityData true "用id查询城市数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cd/findCityData [get]
func (cdApi *CityDataApi) FindCityData(c *gin.Context) {
	ID := c.Query("ID")
	if recd, err := cdService.GetCityData(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recd": recd}, c)
	}
}

// GetCityDataList 分页获取城市数据列表
// @Tags CityData
// @Summary 分页获取城市数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cityReq.CityDataSearch true "分页获取城市数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cd/getCityDataList [get]
func (cdApi *CityDataApi) GetCityDataList(c *gin.Context) {
	var pageInfo cityReq.CityDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cdService.GetCityDataInfoList(pageInfo); err != nil {
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
