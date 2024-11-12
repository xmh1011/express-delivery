package model

import (
	"time"

	"github.com/xmh1011/express-delivery/pkg/variable"
)

type Order struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;comment:订单ID"`
	UID       int64     `gorm:"not null;comment:用户ID;column:uid"`
	Weight    float64   `gorm:"not null;comment:重量(kg);column:weight"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间;column:created_at"`
}

func (o *Order) TableName() string {
	return variable.OrderTableName
}
