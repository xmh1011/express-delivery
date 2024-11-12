package service

import (
	"gorm.io/gorm"

	"github.com/xmh1011/express-delivery/model"
	"github.com/xmh1011/express-delivery/pkg/cost"
	"github.com/xmh1011/express-delivery/pkg/log"
)

// QueryUserOrders 查询指定用户的所有订单并计算总费用
func QueryUserOrders(db *gorm.DB, uid int64) (float64, error) {
	var orders []model.Order
	if err := db.Where("uid = ?", uid).Find(&orders).Error; err != nil {
		log.L.Errorf("Failed to query orders for %d, error: %v", uid, err)
		return 0, err
	}

	var totalCost float64
	for _, order := range orders {
		c := cost.Calculate(order.Weight)
		totalCost += c
	}

	return totalCost, nil
}
