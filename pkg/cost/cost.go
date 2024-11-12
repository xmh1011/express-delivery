package cost

import (
	"math"

	"github.com/xmh1011/express-delivery/pkg/variable"
)

// Calculate 计算快递费用
func Calculate(weight float64) float64 {
	// 规则：最大计费重量为100KG
	if weight > variable.MaxDeliveryWeight {
		weight = variable.MaxDeliveryWeight
	}

	// 1. 将实际重量不满1KG的部分进位到1KG
	roundedWeight := math.Ceil(weight)

	// 2. 计算基本运费：1KG以内18元
	cost := variable.BaseDeliveryCost
	if roundedWeight > 1 {
		// 每增加1KG，增加5元
		cost += (roundedWeight - 1) * variable.ExtraDeliveryCostPerKilo
	}

	// 3. 计算保险费（运费 + 保险费的1%）
	totalCost := cost * (1 + variable.InsuranceRate)

	// 4. 按照四舍五入取整
	return math.Round(totalCost)
}
