package entity

func (OrderItem) TableName() string {
	return "orderitems" // Matches the exact table name in your database
}

type OrderItem struct {
	OrderItemID uint    `gorm:"primaryKey;column:order_item_id" json:"order_item_id"`
	OrderID     uint    `gorm:"column:order_id" json:"order_id"`
	ProductID   uint    `gorm:"column:product_id" json:"product_id"`
	Quantity    int     `gorm:"column:quantity" json:"quantity"`
	Price       float64 `gorm:"type:decimal(10,2);column:price" json:"price"`

	// Associations
	Order   Order   `gorm:"foreignKey:OrderID;references:order_id"`
	Product Product `gorm:"foreignKey:ProductID;references:product_id"`
}

type OrderItemResponse struct {
	OrderItemID uint    `json:"order_item_id"`
	OrderID     uint    `json:"order_id"`
	ProductID   uint    `json:"product_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
