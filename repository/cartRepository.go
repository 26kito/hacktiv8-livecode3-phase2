package repository

import (
	"hacktiv8-lc3-p2/entity"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (cr *CartRepository) GetCarts(userID int) ([]entity.CartResponse, error) {
	var carts []entity.Cart
	var cartResponses []entity.CartResponse

	// Fetch carts with preloaded User and Product
	if err := cr.DB.
		Preload("User").
		Preload("Product").
		Where("user_id = ?", userID).
		Find(&carts).Error; err != nil {
		return nil, err
	}

	// Map the results to CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, entity.CartResponse{
			CartID:    cart.CartID,
			UserID:    cart.UserID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			CreatedAt: cart.CreatedAt,
			User: entity.UserResponse{
				UserID: cart.User.UserID,
				Name:   cart.User.Name,
				Email:  cart.User.Email,
			},
			Product: cart.Product, // Assuming Product struct doesn't need modification
		})
	}

	return cartResponses, nil
}
