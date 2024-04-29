package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/city"
	cityReq "github.com/veteran-dev/server/model/city/request"
	cityResp "github.com/veteran-dev/server/model/city/response"
	cityCarCombinationReq "github.com/veteran-dev/server/model/cityCarCombination/request"
	carResp "github.com/veteran-dev/server/model/cityCarCombination/response"
	"github.com/veteran-dev/server/model/common/response"
	"github.com/veteran-dev/server/model/order"
	orderReq "github.com/veteran-dev/server/model/order/request"
	orderResp "github.com/veteran-dev/server/model/order/response"
	priceReq "github.com/veteran-dev/server/model/price/request"
	priceResp "github.com/veteran-dev/server/model/price/response"
	"github.com/veteran-dev/server/model/user"
	"github.com/veteran-dev/server/service"
	"github.com/veteran-dev/server/utils"
	"go.uber.org/zap"
)

type WebApi struct {
}

var cityService = service.ServiceGroupApp.CityServiceGroup.CityDataService
var orderService = service.ServiceGroupApp.OrderServiceGroup
var carService = service.ServiceGroupApp.CityCarCombinationServiceGroup
var userService = service.ServiceGroupApp.UserServiceGroup

// GetCityList Web获取城市列表

// @Tags		WebApi
// @Summary	获取城市列表
// @accept		application/json
// @Produce	application/json
// @Success 200 {object} city.CityDataList "成功"
// @Router		/web/city/list [post]
func (wApi *WebApi) GetCityList(c *gin.Context) {
	result, err := cityService.City()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(result, c)
	}
}

// SearchCityItem Web关键词查询城市

// @Tags		WebApi
// @Summary	用关键词查询城市，不传参数则展示当前城市位置
// @accept		application/json
// @Produce	application/json
// @Param		data	query		cityReq.CitySearchReq	true	"用关键词查询城市，不传参数则展示当前城市位置"
// @Success 200 {object} city.City "成功"
// @Router		/web/city/search [post]
func (wApi *WebApi) SearchCityItem(c *gin.Context) {
	var req cityReq.CitySearchReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	clientIP := c.RemoteIP()
	result, err := cityService.SearchCity(clientIP, req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(result, c)
	}
}

// GetLocal Web获取当前所在城市POI

// @Tags		WebApi
// @Summary	获取当前所在城市POI
// @accept		application/json
// @Produce	application/json
// @Param		data	query		cityReq.CityLocalReq	true	"用id查询城市"
// @Success 200 {object} []cityResp.GetLocalResp "成功"
// @Router		/web/city/local [post]
func (wApi *WebApi) GetLocal(c *gin.Context) {
	var req cityReq.CityLocalReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if req.CityID <= 0 {
		global.GVA_LOG.Error("非法城市ID!", zap.Error(err))
		response.FailWithMessage("非法城市ID", c)
		return
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	id := strconv.Itoa(int(req.CityID))
	result, err := cityService.GetCityData(id)
	if err != nil {
		global.GVA_LOG.Error("获取城市失败!", zap.Error(err))
		response.FailWithMessage("获取城市失败", c)
		return
	} else {
		locals, err := GetLocals(result.Name, &req, global.GVA_CONFIG.System.POI)
		if err != nil {
			global.GVA_LOG.Error("获取城市失败!", zap.Error(err))
			response.FailWithMessage("获取城市失败", c)
		}
		if len(locals) > 0 {
			areaNames := []string{}
			cityMap := make(map[string][]cityResp.Poi)
			for _, v := range locals {
				areaNames = append(areaNames, v.AreaName)
			}
			areaNames = removeDuplicates(areaNames)
			areas, _ := cityService.GetCityDataByNameAndPid(areaNames, id)
			areaMap := make(map[string]int64)
			for _, v := range areas {
				areaMap[v.Name] = int64(v.ID)
			}
			unique := make(map[string]bool)
			for _, v := range locals {
				typeItem := strings.Split(v.Type, ";")
				if !unique[typeItem[1]] {
					unique[typeItem[1]] = true
				}
				cityMap[typeItem[1]] = append(cityMap[typeItem[1]], cityResp.Poi{
					Location: v.Name,
					AreaID:   areaMap[v.AreaName],
					AreaName: v.AreaName,
				})
			}
			var data []cityResp.GetLocal
			for k, v := range cityMap {
				data = append(data, cityResp.GetLocal{
					CityID:   int64(result.ID),
					CityName: result.Name,
					Type:     k,
					Poi:      v,
				})
			}

			response.OkWithData(response.PageResult{List: &cityResp.GetLocalResp{Locals: data}, Total: locals[0].Count, Page: req.Page, PageSize: req.PageSize}, c)
			return
		}

		global.GVA_LOG.Error("所在城市没有参照物!", zap.Error(err))
		response.FailWithMessage("所在城市没有参照物！", c)
		return
	}
}

func removeDuplicates(strs []string) []string {
	unique := make(map[string]bool)
	result := []string{}

	for _, str := range strs {
		if !unique[str] {
			unique[str] = true
			result = append(result, str)
		}
	}

	return result
}

func GetLocals(cityName string, req *cityReq.CityLocalReq, poi []string) (area []city.LocalArea, err error) {
	area = make([]city.LocalArea, 0)
	gaodeReq := cityReq.GaodeReq{
		Key:      global.GVA_CONFIG.System.DituKey,
		Keywords: req.Address,
		Children: 1,
		City:     cityName,
		Offset:   req.PageSize,
		Page:     req.Page,
	}
	params := url.Values{}
	Url, err := url.Parse("https://restapi.amap.com/v3/place/text?parameters")
	if err != nil {
		global.GVA_LOG.Error("创建请求失败!", zap.Error(err))
		return nil, err
	}
	params.Set("key", gaodeReq.Key)
	types := ""
	if len(poi) != 0 {
		types = strings.Join(poi, "|")
	} else {
		global.GVA_LOG.Error("POI不能为空!", zap.Error(err))
		return nil, err
	}

	params.Set("types", types)
	params.Set("keywords", gaodeReq.Keywords)
	params.Set("children", strconv.Itoa(gaodeReq.Children))
	params.Set("city", gaodeReq.City)
	params.Set("offset", strconv.Itoa(gaodeReq.Offset))
	params.Set("page", strconv.Itoa(gaodeReq.Page))
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		global.GVA_LOG.Error("创建请求失败!", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取失败!", zap.Error(err))
		return nil, err
	}
	var gaode city.GaodeData
	json.Unmarshal(body, &gaode)
	counts, _ := strconv.Atoi(gaode.Count)
	if len(gaode.Pois) > 0 {
		for _, v := range gaode.Pois {
			area = append(area, city.LocalArea{
				AreaName: v.Adname,
				Name:     v.Name,
				Type:     v.Type,
				Address:  v.Address,
				Count:    int64(counts),
			})
		}

		return area, nil
	}
	return nil, errors.New("当前区域为空")
}

// GetCityList Web获取车型列表
// @Tags		WebApi
// @Summary	获取车型列表
// @accept		application/json
// @Produce	application/json
// @Param		data	query		cityCarCombinationReq.GetCarListReq	true	"用id查询订单"
// @Success 200 {object} []carResp.Car "成功"
// @Router		/web/car/list [post]
func (wApi *WebApi) GetCarList(c *gin.Context) {
	var req cityCarCombinationReq.GetCarListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if req.FromCityName == "" && req.FromAreaID <= 0 && req.ToCityName == "" && req.ToAreaID <= 0 {
		response.FailWithMessage("请检查参数", c)
		return
	}

	if result, total, err := carService.GetModel(req); err != nil {
		global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取数据失败", c)
		return
	} else {
		var charge int64
		if req.StartTime != "" {
			charge = Charge(req.StartTime)
		}

		var cars []carResp.Car
		if total > 0 {
			for _, v := range result {
				title := fmt.Sprintf("商务%d座", *v.CarCombination.Seats)
				cars = append(cars, carResp.Car{
					ID:           int64(v.ID),
					Title:        title,
					Level:        v.CarCombination.Level,
					Brand:        v.CarCombination.ModelName,
					Price:        int64(*v.BasePrice) + charge, //不同的时间定价不一样
					Seats:        int64(*v.CarCombination.Seats),
					Luggage:      int64(*v.CarCombination.Luggage),
					ToCityName:   req.ToCityName,
					FromCityName: req.FromCityName,
					ToAreaID:     req.ToAreaID,
					FromAreaID:   req.FromAreaID,
					StartTime:    req.StartTime,
				})
			}
		}
		response.OkWithData(response.PageResult{List: cars, Total: total, Page: req.Page, PageSize: req.PageSize}, c)
		return
	}
}

func Charge(timeStr string) int64 {
	t, err := time.ParseInLocation("01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		global.GVA_LOG.Error("解析时间失败!", zap.Error(err))
		return 0
	}

	// 判断时间段并计算收费
	var charge int64
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
// @Success 200 {object} carResp.CarDetailResp "成功"
// @Router		/web/car/detail [post]
func (wApi *WebApi) CarDetail(c *gin.Context) {
	var req cityCarCombinationReq.GetCarReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	result, err := carService.ModelDetail(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取数据失败", c)
		return
	} else {
		var charge int64
		if req.StartTime != "" {
			charge = Charge(req.StartTime)
		}
		data := carResp.CarDetail{
			Title:        result.CarCombination.CombinationTitle,
			Luggage:      int64(*result.CarCombination.Luggage),
			LargeLuggage: int64(*result.CarCombination.LargeLuggage),
			Brand:        result.CarCombination.ModelName,
			ChildSeats:   int64(*result.CarCombination.ChildSeats),
			Seats:        int64(*result.CarCombination.Seats),
			Level:        result.CarCombination.Level,
			ID:           int64(result.ID),
			Price:        charge + int64(*result.BasePrice),
		}
		response.OkWithData(carResp.CarDetailResp{
			Detail:       data,
			FromCityName: req.FromCityName,
			ToCityName:   req.ToCityName,
			FromAreaID:   int64(*result.From),
			ToAreaID:     int64(*result.To),
			StartTime:    req.StartTime,
			FromLocation: req.FromLocation,
			ToLocation:   req.ToLocation,
		}, c)
		return
	}
}

// OrderCancel 提交订单

// @Tags		WebApi
// @Summary	提交订单
// @accept		application/json
// @Produce	application/json
// @Param		data	query	orderReq.OrderCompleteReq	true	"插入数据"
// @Success 200 {object} orderResp.OrderCompleteResp "成功"
// @Router		/web/order/complete [post]
func (wApi *WebApi) OrderComplete(c *gin.Context) {
	var req orderReq.OrderCompleteReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	url := "https://devcr.dachema.net/cmdcapp/api/aliPay/orderPay"
	payReq := orderReq.PayCode{
		OrderID:   time.Now().Unix(),
		Code:      req.Code,
		ReturnUrl: "http://api.h5doc.com/web/order/detail",
	}
	jsonData, err := json.Marshal(payReq)
	if err != nil {
		global.GVA_LOG.Error("JSON编码失败!", zap.Error(err))
		return
	}

	// 创建一个请求体
	resq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData)) //ei
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
		if err := orderService.UpdateOrderByOrderSerial(req.OrderSerial, req.Passenger, req.PassengerMobile, 1); err != nil {
			global.GVA_LOG.Error("订单状态失败!", zap.Error(err))
			response.FailWithMessage("订单生成失败", c)
			return
		}
		global.GVA_LOG.Info("订单状态成功!")
		response.OkWithData(&orderResp.OrderCompleteResp{OrderSerial: req.OrderSerial}, c)
		return
	}
	global.GVA_LOG.Error("调用支付失败!", zap.Error(err))
	response.FailWithMessage("调用支付失败", c)
}

// OrderCancel 填写订单

// @Tags		WebApi
// @Summary	填写订单
// @accept		application/json
// @Produce	application/json
// @Param		data	query	orderReq.OrderCreateReq	true	"新建订单"
// @Success 200 {object} orderResp.OrderCreateResp "成功"
// @Router		/web/order/create [post]
func (wApi *WebApi) OrderCreate(c *gin.Context) {
	var req orderReq.OrderCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if req.CarID == 0 {
		global.GVA_LOG.Error("车型组参数非法!", zap.Error(err))
		response.FailWithMessage("车型组参数非法!", c)
		return
	}
	carModel, err := carService.ModelDetail(req.CarID)
	if err != nil {
		global.GVA_LOG.Error("获取车型组信息失败!", zap.Error(err))
		response.FailWithMessage("获取车型组信息失败", c)
		return
	}
	var charge int64
	var cancelAt string
	if req.StartTime != "" {
		charge = Charge(req.StartTime)
		cancelAt, err = ConvertBack(req.StartTime)
		if err != nil {
			global.GVA_LOG.Error("时间格式有误!", zap.Error(err))
			response.FailWithMessage("时间格式有误,2006-01-02 15:04:05", c)
		}
		inputFormat := "01-02 15:04:05"
		startTime, _ := time.Parse(inputFormat, req.StartTime)
		orderSerial := generateOrderNumber()
		carID := int(carModel.ID)
		price := charge + int64(*carModel.BasePrice)
		var ptr, pricePtr *int
		ptr = &carID
		priceConv := int(price)
		pricePtr = &priceConv
		createResult := orderService.CreateOrder(&order.Order{
			Appointment: &startTime,
			FromCity:    carModel.From,
			ToCity:      carModel.To,
			FromArea:    req.FromLocation,
			ToArea:      req.ToLocation,
			OrderSerial: orderSerial,
			CarModel:    ptr,
			Price:       pricePtr,
		})
		if createResult != nil {
			response.OkWithData(&orderResp.OrderCreateResp{
				StartTime:    req.StartTime,
				CancelAt:     cancelAt,
				ToCityName:   req.ToCityName,
				FromCityName: req.FromCityName,
				FromLocation: req.FromLocation,
				ToLocation:   req.ToLocation,
				Car:          carModel,
				Price:        price,
				OrderSerial:  orderSerial,
			}, c)
			return
		}
	}
	global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
	response.FailWithMessage("创建订单失败", c)
	return
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
// @Router		/web/order/detail [post]
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
	carinfo, err := carService.ModelDetail(int64(*orderData.CarModel))
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

	if err := orderService.UpdateOrderByOrderSerial(req.OrderSerial, req.Passenger, req.PassengerMobile, 1); err != nil {
		global.GVA_LOG.Error("更改失败!", zap.Error(err))
		response.FailWithMessage("更改失败", c)
		return
	}
	response.OkWithMessage("更改成功", c)
}

// PriceRules 计费规则

// @Tags		WebApi
// @Summary	计费规则
// @accept		application/json
// @Produce	application/json
// @Param		data	query	priceReq.PriceRulesReq	true	"计费规则"
// @Success 200 {object} priceResp.PriceResp "成功"
// @Router		/web/price/rules [post]
func (wApi *WebApi) PriceRules(c *gin.Context) {
	var req priceReq.PriceRulesReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	carData, err := carService.GetCityCarCombination(strconv.Itoa(int(req.CarModelID)))
	if err != nil {
		global.GVA_LOG.Error("获取当地车型组失败!", zap.Error(err))
		response.FailWithMessage("获取当地车型组失败", c)
		return
	}

	var charge int64
	var at string
	if req.StartTime != "" {
		charge = Charge(req.StartTime)
		at, err = ConvertToAt(req.StartTime)
		if err != nil {
			global.GVA_LOG.Error("转换时间格式失败!", zap.Error(err))
			response.FailWithMessage("转换时间格式失败", c)
			return
		}
	}
	basePrice := int64(*carData.BasePrice)
	response.OkWithData(priceResp.PriceResp{
		ToCityName:   req.ToCityName,
		FromCityName: req.FromCityName,
		BasePrice:    basePrice,
		SubPrice:     charge,
		TotalPrice:   charge + basePrice,
		StartAt:      at,
	}, c)
}

func ConvertToAt(str string) (at string, err error) {
	// 定义日期时间字符串的格式
	inputFormat := "2006-01-02 15:04"
	outputFormat := "15:04"
	// 解析日期时间字符串
	dateTime, err := time.Parse(inputFormat, str)
	if err != nil {
		fmt.Println("解析错误:", err)
		return "", err
	}
	// 格式化日期时间字符串
	return dateTime.Format(outputFormat), nil
}

func ConvertBack(str string) (back string, err error) {
	format := "2006-01-02 15:04"

	// 解析日期时间字符串
	dateTime, err := time.Parse(format, str)
	if err != nil {
		return "", nil
	}
	// 获取昨天的日期
	yesterday := dateTime.AddDate(0, 0, -1)
	// 格式化为指定输出格式
	formattedYesterday := yesterday.Format(format)
	return formattedYesterday, nil
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

// ReasonList 取消原因

// @Tags		WebApi
// @Summary	取消原因
// @accept		application/json
// @Produce	application/json
// @Success 200 {object} []string "成功"
// @Router		/web/reason/list [post]
func (wApi *WebApi) ReasonList(c *gin.Context) {
	result := orderService.ReasonList()
	response.OkWithData(result, c)
}

// Login 登录Token
// @Tags		WebApi
// @Summary	登录Token
// @accept		application/json
// @Produce	application/json
// @Param		data	query	Login	true	"小程序授权Code"
// @Success 200 {object} RespLogin "成功"
// @Router		/web/login [post]
func (wApi *WebApi) Login(c *gin.Context) {
	var loginReq Login
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	baseURL := "https://devcr.dachema.net/cmdcapp/api/auth/loginAliSmallByCode"
	urlParams := url.Values{}
	urlParams.Set("channelCode", "CmdcAli")
	requestURL, _ := url.Parse(baseURL)
	requestURL.RawQuery = urlParams.Encode()

	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		global.GVA_LOG.Error("JSON编码失败!", zap.Error(err))
		return
	}
	log.Print(requestURL.String())
	// 创建一个请求体
	req, err := http.NewRequest("POST", requestURL.String(), bytes.NewBuffer(jsonData))
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
	var res RemoteAck
	json.Unmarshal(body, &res)
	log.Print(res)
	if !res.Success {
		global.GVA_LOG.Error("数据格式解析失败!", zap.Error(err))
		response.FailWithMessage("数据格式非法", c)
		return
	}
	userID := strconv.Itoa(int(res.Data.UserID))
	userData := &user.User{
		Token:  res.Data.AccessToken,
		UserId: userID,
	}
	err = userService.FindOrCreateUser(userData)
	if err != nil {
		global.GVA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败", c)
		return
	}

	newJwt := utils.NewUserJWT()
	result, err := newJwt.GenerateToken(int(userData.ID))
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

type Login struct {
	Code string `json:"code"` //授权Code
}
type RemoteAck struct {
	Data    RemoteAckResp `json:"data"`
	Success bool          `json:"success"`
	Code    int64         `json:"code"`
	Message string        `json:"message"`
}

type RemoteAckResp struct {
	AccessToken string `json:"accessToken"`
	UserID      int64  `json:"userId"`
	IsNewUser   bool   `json:"isNewUser"`
	Expires     int64  `json:"expires"`
}
