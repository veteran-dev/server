package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/veteran-dev/server/global"
	cityCarCombinationReq "github.com/veteran-dev/server/model/cityCarCombination/request"
	carResp "github.com/veteran-dev/server/model/cityCarCombination/response"
	"github.com/veteran-dev/server/model/common/response"
	"github.com/veteran-dev/server/service"
	"github.com/veteran-dev/server/utils"
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

var carService = service.ServiceGroupApp.CityCarCombinationServiceGroup

// GetCityList Web获取车刑列表
// @Tags WebApi
// @Summary 获取车刑列表
// @accept application/json
// @Produce application/json
// @Param data query cityCarCombinationReq.GetCarReq true "用id查询订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/car/list [get]
func (wApi *WebApi) GetCarList(c *gin.Context) {

	var req cityCarCombinationReq.GetCarListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if req.From <= 0 && req.To <= 0 && req.From == req.To {
		response.FailWithMessage("请检查参数", c)
		return
	}

	if result, total, err := carService.GetModel(req); err != nil {
		global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取数据失败", c)
		return
	} else {
		var charge int
		if req.ArrivalTime != "" {
			charge = Charge(req.ArrivalTime)
		}

		var cars []carResp.Car
		if total > 0 {
			for _, v := range result {
				title := fmt.Sprintf("商务%d座", *v.CarCombination.Seats)
				cars = append(cars, carResp.Car{
					ID:      int64(v.ID),
					Title:   title,
					Level:   v.CarCombination.Level,
					Brand:   v.CarCombination.ModelName,
					Price:   int64(*v.BasePrice + charge), //不同的时间定价不一样
					Seats:   int64(*v.CarCombination.Seats),
					Luggage: int64(*v.CarCombination.Luggage),
				})
			}
		}
		response.OkWithData(response.PageResult{List: cars, Total: total, Page: req.Page, PageSize: req.PageSize}, c)
		return
	}
}

func Charge(timeStr string) int {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		global.GVA_LOG.Error("解析时间失败!", zap.Error(err))
		return 0
	}

	// 判断时间段并计算收费
	var charge int
	hour := t.Hour()
	minute := t.Minute()

	if hour >= 7 && hour < 22 {
		charge = 0
	} else if hour == 5 && minute >= 0 && minute <= 59 {
		charge = 100
	} else if hour == 6 && minute >= 0 && minute <= 59 {
		charge = 100
	} else if hour >= 22 && hour < 24 {
		charge = 100
	} else if hour >= 0 && hour < 5 {
		charge = 200
	}
	return charge
}

// CarDetail 车型详情
// @Tags WebApi
// @Summary 车型详情
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/car/detail [get]
func (wApi *WebApi) CarDetail(c *gin.Context) {

	var req cityCarCombinationReq.GetCarReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	if result, err := carService.ModelDetail(req); err != nil {
		global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取数据失败", c)
	} else {
		// var charge int
		// if req.StartTime != "" {
		// 	charge = Charge(req.StartTime)
		// }
		response.OkWithData(gin.H{"list": result}, c)
	}
}

// CarQuote 获取报价
// @Tags WebApi
// @Summary 获取报价
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/car/quote [get]
func (wApi *WebApi) CarQuote(c *gin.Context) {

	var req cityCarCombinationReq.GetCarListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	if result, _, err := carService.GetModel(req); err != nil {
		global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取数据失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}

// OrderComplete 提交订单
// @Tags WebApi
// @Summary 提交订单
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/order/complete [get]
func (wApi *WebApi) OrderComplete(c *gin.Context) {

	// var req cityCarCombinationReq.GetCarReq
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	// 	response.FailWithMessage("获取失败", c)
	// }

	// if result, _, err := carService.GetModel(req); err != nil {
	// 	global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
	// 	response.FailWithMessage("获取数据失败", c)
	// } else {
	// 	response.OkWithData(gin.H{"list": result}, c)
	// }
}

// OrderDetail 订单详情
// @Tags WebApi
// @Summary 订单详情
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/order/detail [get]
func (wApi *WebApi) OrderDetail(c *gin.Context) {

	// var req cityCarCombinationReq.GetCarReq
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	// 	response.FailWithMessage("获取失败", c)
	// }

	// if result, _, err := carService.GetModel(req); err != nil {
	// 	global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
	// 	response.FailWithMessage("获取数据失败", c)
	// } else {
	// 	response.OkWithData(gin.H{"list": result}, c)
	// }
}

// OrderUpdate 修改订单
// @Tags WebApi
// @Summary 修改订单
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/order/update [get]
func (wApi *WebApi) OrderUpdate(c *gin.Context) {

	// var req cityCarCombinationReq.GetCarReq
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	// 	response.FailWithMessage("获取失败", c)
	// }

	// if result, _, err := carService.GetModel(req); err != nil {
	// 	global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
	// 	response.FailWithMessage("获取数据失败", c)
	// } else {
	// 	response.OkWithData(gin.H{"list": result}, c)
	// }
}

// OrderCancel 取消订单
// @Tags WebApi
// @Summary 取消订单
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/order/cancel [get]
func (wApi *WebApi) OrderCancel(c *gin.Context) {

	// var req cityCarCombinationReq.GetCarReq
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	// 	response.FailWithMessage("获取失败", c)
	// }

	// if result, _, err := carService.GetModel(req); err != nil {
	// 	global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
	// 	response.FailWithMessage("获取数据失败", c)
	// } else {
	// 	response.OkWithData(gin.H{"list": result}, c)
	// }
}

// OrderCancel 登录Token
// @Tags WebApi
// @Summary 登录Token
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /web/login [get]
func (wApi *WebApi) Login(c *gin.Context) {
	var loginReq Login
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	url := "https://devcr.dachema.net/cmdcapp/api/auth/loginAliSmallByCode"
	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		global.GVA_LOG.Error("JSON编码失败!", zap.Error(err))
		return
	}

	// 创建一个请求体
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		global.GVA_LOG.Error("创建请求失败!", zap.Error(err))
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("请求发送失败!", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取失败!", zap.Error(err))
		return
	}
	var res LoginResp
	json.Unmarshal(body, res)
	if res.UserId == 0 {
		global.GVA_LOG.Error("获取USRID失败!", zap.Error(err))
		response.FailWithMessage("获取USRID失败", c)
		return
	}
	newJwt := utils.NewUserJWT()
	if result, err := newJwt.GenerateToken(int(res.UserId)); err != nil {
		global.GVA_LOG.Error("获取Token失败!", zap.Error(err))
		response.FailWithMessage("获取Token失败", c)
	} else {
		response.OkWithData(gin.H{"x-token": result}, c)
	}
}

type LoginResp struct {
	AccessToken string `json:"accessToken"`
	UserId      int64  `json:"userId"`
	IsNewUser   bool   `json:"isNewUser"`
	Expires     int64  `json:"expires"` //分钟
}

type Login struct {
	Code string `json:"code"`
}
