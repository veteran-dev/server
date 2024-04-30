package order

import (
	"github.com/veteran-dev/server/global"
	"github.com/veteran-dev/server/model/order"
	orderReq "github.com/veteran-dev/server/model/order/request"
	"gorm.io/gorm"
)

type OrderService struct {
}

// CreateOrder 创建订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) CreateOrder(o *order.Order) (err error) {
	err = global.GVA_DB.Debug().Create(o).Error
	return err
}

// DeleteOrder 删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) DeleteOrder(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.Order{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&order.Order{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteOrderByIds 批量删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) DeleteOrderByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.Order{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&order.Order{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateOrder 更新订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) UpdateOrder(o order.Order) (err error) {
	err = global.GVA_DB.Save(&o).Error
	return err
}

// UpdateOrder 更新订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) UpdateOrderByOrderSerialStatus(order string, status, reson int) (err error) {
	err = global.GVA_DB.Where("order_serial = ?", order).Updates(map[string]interface{}{"status": status, "cancel_reason": reson}).Error
	return err
}

// GetOrder 根据ID获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) GetOrder(ID string) (o order.Order, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&o).Error
	return
}

// GetOrder 根据ID获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) GetOrderByOrderSerial(ID string) (o order.Order, err error) {
	err = global.GVA_DB.Where("order_serial = ?", ID).First(&o).Error
	return
}

func (oService *OrderService) UpdateOrderByOrderSerial(order, user, mobile string, status int) (err error) {
	err = global.GVA_DB.Where("order_serial = ?", order).Updates(map[string]interface{}{"passenger": user, "passenger_mobile": mobile, "status": status}).Error
	return err

}

// GetOrderInfoList 分页获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (oService *OrderService) GetOrderInfoList(info orderReq.OrderSearch) (list []order.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&order.Order{})
	var os []order.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["appointment"] = true
	orderMap["from_city"] = true
	orderMap["to_city"] = true
	orderMap["order_serial"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&os).Error
	return os, total, err
}

func (oService *OrderService) ReasonList() (list []string) {
	return global.GVA_CONFIG.System.Reason
}
