package variable

const (
	// MaxDeliveryWeight 最大计费重量
	MaxDeliveryWeight = 100

	// BaseDeliveryCost 基本运费：1KG以内18元
	BaseDeliveryCost = 18.0

	// ExtraDeliveryCostPerKilo 每增加1KG, 增加5元
	ExtraDeliveryCostPerKilo = 5

	// InsuranceRate 保险费率
	InsuranceRate = 0.01
)

const (
	DefaultConfigType = "yaml"
)

const (
	OrderTableName = "orders"
)

// 定义全局版本变量
var (
	Version   string
	GitCommit string
)
