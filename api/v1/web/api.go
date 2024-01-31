package web

import (
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/model/common/response"
	"github.com/5asp/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebApi struct {
}

var cityService = service.ServiceGroupApp.CityServiceGroup.CityDataService

// GetCityList Web获取城市列表
// @Tags WebApi
// @Summary 获取城市列表
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/city/list [get]
func (wApi *WebApi) GetCityList(c *gin.Context) {
	if result, err := cityService.City(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}
