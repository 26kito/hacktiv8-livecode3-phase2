package entity

type Product struct {
	ProductID   uint    `gorm:"primaryKey;column:product_id" json:"product_id"`
	Name        string  `gorm:"type:varchar(100)" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
}
