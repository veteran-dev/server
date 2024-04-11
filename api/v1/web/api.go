package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/veteran-dev/server/global"
	cityReq "github.com/veteran-dev/server/model/city/request"
	cityCarCombinationReq "github.com/veteran-dev/server/model/cityCarCombination/request"
	carResp "github.com/veteran-dev/server/model/cityCarCombination/response"
	"github.com/veteran-dev/server/model/common/response"
	"github.com/veteran-dev/server/model/order"
	orderReq "github.com/veteran-dev/server/model/order/request"
	orderResp "github.com/veteran-dev/server/model/order/response"
	"github.com/veteran-dev/server/service"
	"github.com/veteran-dev/server/utils"
	"go.uber.org/zap"
)

type WebApi struct {
}

var cityService = service.ServiceGroupApp.CityServiceGroup.CityDataService
var orderService = service.ServiceGroupApp.OrderServiceGroup
var carService = service.ServiceGroupApp.CityCarCombinationServiceGroup
var carModelService = service.ServiceGroupApp.CarCombinationServiceGroup.CarCombinationService

// GetCityList Web获取城市列表
//
//	@Tags		WebApi
//	@Summary	获取城市列表
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		cityReq.CityDataReq	true	"用id查询城市"
//	@Success	200		{string}	string				"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router		/web/city/list [post]
func (wApi *WebApi) GetCityList(c *gin.Context) {
	var req cityReq.CityDataReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if result, err := cityService.City(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": result}, c)
	}
}

// GetCityList Web获取车型列表
//
//	@Tags		WebApi
//	@Summary	获取车型列表
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		cityCarCombinationReq.GetCarListReq	true	"用id查询订单"
//	@Success	200		{string}	string							"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router		/web/car/list [post]
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
		if req.StartTime != "" {
			charge = Charge(req.StartTime)
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

// @Tags		WebApi
// @Summary	车型详情
// @accept		application/json
// @Produce	application/json
// @Param		data	query		cityCarCombinationReq.GetCarReq	true	"用id查询订单"
// @Success	200	{string}	string	"{"success":true,"data":{},"msg":"获取成功"}"
// @Router		/web/car/detail [get]
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
		data := carResp.CarDetail{
			Title:        result.CarCombination.CombinationTitle,
			Luggage:      int64(*result.CarCombination.Luggage),
			LargeLuggage: int64(*result.CarCombination.LargeLuggage),
			Brand:        result.CarCombination.ModelName,
			ChildSeats:   int64(*result.CarCombination.ChildSeats),
			Seats:        int64(*result.CarCombination.Seats),
			Level:        result.CarCombination.Level,
			ID:           int64(result.ID),
		}

		quote := make(map[string]interface{})
		quote["child"] = req.Child
		quote["aldult"] = req.Aldult
		quote["startTime"] = req.StartTime
		quote["luggage"] = req.Luggage
		var charge int
		if req.StartTime != "" {
			charge = Charge(req.StartTime)
		}
		quote["price"] = int64(*result.BasePrice + charge)
		quote["from"] = cityService.GetParentCity(req.FromPID)
		quote["to"] = cityService.GetParentCity(req.ToPID)
		response.OkWithData(gin.H{"detail": data, "quote": quote}, c)
	}
}

// OrderCancel 提交订单

// @Tags		WebApi
// @Summary	提交订单
// @accept		application/json
// @Produce	application/json
// @Param		data	query	orderReq.OrderComplete	true	"插入数据"
// @Success 200 {object} orderResp.OrderCompleteResp "成功"
// @Router		/web/order/complete [post]
func (wApi *WebApi) OrderComplete(c *gin.Context) {
	var req orderReq.OrderComplete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	var fromCity, toCity, price, carModel int
	fromCity = int(req.FromID)
	toCity = int(req.ToID)
	price = int(req.Price)
	carModel = int(req.CarID)
	orderSerial := generateOrderNumber()

	// 定义时间字符串的格式
	layout := "2006-01-02T15:04:05Z"

	t, err := time.Parse(layout, req.StartTime)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	fmt.Println("转换后的时间:", t)
	orderCreate := &order.Order{
		Appointment:     &t, //出发时间
		FromCity:        &fromCity,
		ToCity:          &toCity,
		ToArea:          req.ToArea,
		FromArea:        req.FromArea,
		OrderSerial:     orderSerial,
		Passenger:       req.Passenger,
		PassengerMobile: req.PassengerMobile,
		Price:           &price,
		CarModel:        &carModel,
	}
	if err := orderService.CreateOrder(orderCreate); err != nil {
		global.GVA_LOG.Error("订单生成失败!", zap.Error(err))
		response.FailWithMessage("订单生成失败", c)
		return
	}

	url := "https://devcr.dachema.net/cmdcapp/api/aliPay/orderPay"
	payReq := orderReq.PayCode{
		OrderID:   time.Now().Unix(),
		Code:      req.Code,
		ReturnUrl: "/web/order/detail",
	}
	jsonData, err := json.Marshal(payReq)
	if err != nil {
		global.GVA_LOG.Error("JSON编码失败!", zap.Error(err))
		return
	}

	// 创建一个请求体
	resq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		global.GVA_LOG.Error("创建请求失败!", zap.Error(err))
		return
	}
	resq.Header.Set("Content-Type", "application/json")
	resq.Header.Set("channelCode", req.ChannelCode)
	client := &http.Client{}
	resp, err := client.Do(resq)
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
	var orderData orderResp.OrderResp
	json.Unmarshal(body, orderData)
	if orderData.TradeNo != "" {
		if err := orderService.UpdateOrderByOrderSerialStatus(orderSerial, 1, 0); err != nil {
			global.GVA_LOG.Error("订单状态失败!", zap.Error(err))
			response.FailWithMessage("订单生成失败", c)
			return
		}
		global.GVA_LOG.Info("订单状态成功!")
		response.OkWithData(&orderResp.OrderCompleteResp{OrderSerial: orderSerial}, c)
		return
	}
	global.GVA_LOG.Error("调用支付失败!", zap.Error(err))
	response.FailWithMessage("调用支付失败", c)
}

func generateOrderNumber() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	random := rand.Intn(1000)
	orderNumber := fmt.Sprintf("%d%d", timestamp, random)
	return orderNumber
}

// OrderCancel 订单详情

// @Tags		WebApi
// @Summary	订单详情
// @accept		application/json
// @Produce	application/json
// @Param		data	query	orderReq.OrderDetail	true	"订单号查询"
// @Success 200 {object} orderResp.OrderDetailResp "成功"
// @Router		/web/order/detail [get]
func (wApi *WebApi) OrderDetail(c *gin.Context) {
	var req orderReq.OrderDetail
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	orderData, err := orderService.GetOrderByOrderSerial(req.OrderSerial)
	if err != nil {
		global.GVA_LOG.Error("获取Order数据失败!", zap.Error(err))
		response.FailWithMessage("获取Order数据失败", c)
	}

	carID := strconv.Itoa(*orderData.CarModel)
	carinfo, err := carModelService.GetCarCombination(carID)
	if err != nil {
		global.GVA_LOG.Error("获取Car数据失败!", zap.Error(err))
		response.FailWithMessage("获取Car数据失败", c)
	}
	global.GVA_LOG.Info("订单详情成功!")
	response.OkWithData(&orderResp.OrderDetailResp{Car: carinfo, Order: orderData}, c)
	return
}

// OrderUpdate 修改订单

// @Tags		WebApi
// @Summary	修改订单
// @accept		application/json
// @Produce	application/json
// @Param		data	query	orderReq.OrderUpdate	true	"更改数据"
// @Success	200	{string}	string	"{"success":true,"data":{},"msg":"取消成功"}"
// @Router		/web/order/update [post]
func (wApi *WebApi) OrderUpdate(c *gin.Context) {
	var req orderReq.OrderUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	if req.OrderSerial != "" {
		global.GVA_LOG.Error("非法参数!", zap.Error(err))
		response.FailWithMessage("非法参数", c)
	}

	if err := orderService.UpdateOrderByOrderSerial(req.OrderSerial, req.Passenger, req.PassengerMobile); err != nil {
		global.GVA_LOG.Error("更改失败!", zap.Error(err))
		response.FailWithMessage("更改失败", c)
		return
	}
	response.OkWithMessage("更改成功", c)
}

// OrderCancel 取消订单
//
//	@Tags		WebApi
//	@Summary	取消订单
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query	orderReq.OrderCancelReq	true	"插入数据"
//	@Success	200	{string}	string	"{"success":true,"data":{},"msg":"取消成功"}"
//	@Router		/web/order/cancel [post]
func (wApi *WebApi) OrderCancel(c *gin.Context) {
	var req orderReq.OrderCancelReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	if req.OrderSerial != "" {
		global.GVA_LOG.Error("非法参数!", zap.Error(err))
		response.FailWithMessage("非法参数", c)
	}

	if err := orderService.UpdateOrderByOrderSerialStatus(req.OrderSerial, 0, req.CancelReson); err != nil {
		global.GVA_LOG.Error("订单状态取消失败!", zap.Error(err))
		response.FailWithMessage("订单取消失败", c)
		return
	}
	global.GVA_LOG.Info("订单状态更改成功!")
	response.OkWithMessage("更改成功", c)
	return
}

// OrderCancel 登录Token

// @Tags		WebApi
// @Summary	登录Token
// @accept		application/json
// @Produce	application/json
// @Param		data	query	Login	true	"小程序授权Code"
// @Success 200 {object} RespLogin "成功"
// @Router		/web/login [get]
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
	global.GVA_LOG.Debug("打印登录请求返回")
	fmt.Println(body)
	global.GVA_LOG.Debug("-------------===Login---==============")
	var res LoginResp
	json.Unmarshal(body, res)
	if res.UserId == 0 {
		global.GVA_LOG.Error("获取USRID失败!", zap.Error(err))
		response.FailWithMessage("获取USRID失败", c)
		return
	}
	newJwt := utils.NewUserJWT()
	result, err := newJwt.GenerateToken(int(res.UserId))
	if err != nil {
		global.GVA_LOG.Error("获取Token失败!", zap.Error(err))
		response.FailWithMessage("获取Token失败", c)
	} else {
		response.OkWithData(&RespLogin{XToken: result}, c)
	}
}

type RespLogin struct {
	XToken string `json:"x-token"` //请求Token
}
type LoginResp struct {
	AccessToken string `json:"accessToken"`
	UserId      int64  `json:"userId"`
	IsNewUser   bool   `json:"isNewUser"`
	Expires     int64  `json:"expires"` //分钟
}

type Login struct {
	Code string `json:"code"` //授权Code
}
