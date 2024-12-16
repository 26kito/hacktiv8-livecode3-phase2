package entity

import "time"

type Order struct {
	OrderID    uint      `gorm:"primaryKey;column:order_id" json:"order_id"`
	UserID     uint      `gorm:"column:user_id" json:"user_id"`
	TotalPrice float64   `gorm:"type:decimal(10,2);column:total_price" json:"total_price"`
	CreatedAt  time.Time `gorm:"column:created_at;default:current_timestamp" json:"created_at"`

	// Associations
	User       User        `gorm:"foreignKey:UserID;references:user_id"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:order_id"`
}
