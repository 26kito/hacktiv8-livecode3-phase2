package entity

import (
	"time"
)

type Cart struct {
	CartID    uint      `gorm:"primaryKey;column:cart_id" json:"cart_id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	ProductID uint      `gorm:"column:product_id" json:"product_id"`
	Quantity  int       `gorm:"column:quantity" json:"quantity"`
	CreatedAt time.Time `gorm:"column:created_at;default:current_timestamp" json:"created_at"`

	// Associations
	User    User    `gorm:"foreignKey:UserID;references:user_id"`
	Product Product `gorm:"foreignKey:ProductID;references:product_id"`
}

type CartResponse struct {
	CartID    uint         `json:"cart_id"`
	UserID    uint         `json:"user_id"`
	ProductID uint         `json:"product_id"`
	Quantity  int          `json:"quantity"`
	CreatedAt time.Time    `json:"created_at"`
	User      UserResponse `json:"user"`
	Product   Product      `json:"product"`
}
