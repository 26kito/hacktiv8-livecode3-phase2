package repository

import (
	"hacktiv8-lc3-p2/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) GetProducts() ([]entity.Product, error) {
	var products []entity.Product

	if err := pr.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) GetProductByID(productID int) (*entity.Product, error) {
	var product entity.Product

	if err := pr.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
