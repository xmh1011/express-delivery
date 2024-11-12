package service

import (
	"math/rand"
	"time"

	"gorm.io/gorm"

	"github.com/xmh1011/express-delivery/model"
	"github.com/xmh1011/express-delivery/pkg/log"
)

// InsertData 插入数据
func InsertData(db *gorm.DB, userCount, orderCount int64) {
	userIDs := generateUserIDs(userCount)
	generateOrders(db, userIDs, orderCount)
}

// 生成1000个用户ID
func generateUserIDs(count int64) []int64 {
	userIDs := make([]int64, count)
	var i int64
	for i = 0; i < count; i++ {
		userIDs[i] = i + 1 // 用户ID从1到1000
	}
	return userIDs
}

// 生成100,000条订单数据并插入数据库
func generateOrders(db *gorm.DB, userIDs []int64, orderCount int64) {
	rand.Seed(time.Now().UnixNano())

	var i int64
	for i = 0; i < orderCount; i++ {
		uid := userIDs[rand.Intn(len(userIDs))] // 随机选择UID
		weight := generateWeight()              // 按照1/W权重生成重量
		order := model.Order{
			ID:        i + 1,
			UID:       uid,
			Weight:    weight,
			CreatedAt: time.Now(),
		}
		db.Create(&order)
	}

	log.L.Debugf("Successfully inserted %d orders", orderCount)
}

// 按照1/W权重生成重量
func generateWeight() float64 {
	weights := []float64{1, 2, 3, 4, 5, 10, 20, 50, 100} // 示例权重
	probabilities := make([]float64, len(weights))

	// 计算每个重量的概率 1/W
	for i, w := range weights {
		probabilities[i] = 1 / w
	}

	// 选择权重
	return weightedRandom(weights, probabilities)
}

// 根据权重返回随机值
func weightedRandom(values []float64, weights []float64) float64 {
	sum := 0.0
	for _, w := range weights {
		sum += w
	}

	r := rand.Float64() * sum
	for i, w := range weights {
		if r < w {
			return values[i]
		}
		r -= w
	}

	return values[len(values)-1]
}
