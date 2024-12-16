package repository

import (
	"hacktiv8-lc3-p2/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (or *OrderRepository) GetOrders(userID int) ([]entity.OrderResponse, error) {
	var orders []entity.Order
	var orderResponses []entity.OrderResponse
	var orderItems []entity.OrderItemResponse

	// Fetch orders with preloaded User and OrderItems
	if err := or.DB.
		Preload("User").
		Preload("OrderItems").
		Where("user_id = ?", userID).
		Find(&orders).Error; err != nil {
		return nil, err
	}

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			orderItems = append(orderItems, entity.OrderItemResponse{
				OrderItemID: orderItem.OrderItemID,
				OrderID:     orderItem.OrderID,
				ProductID:   orderItem.ProductID,
				Quantity:    orderItem.Quantity,
				Price:       orderItem.Price,
			})
		}

		orderResponses = append(orderResponses, entity.OrderResponse{
			OrderID:    order.OrderID,
			UserID:     order.UserID,
			TotalPrice: order.TotalPrice,
			CreatedAt:  order.CreatedAt,
			User: entity.UserResponse{
				UserID: order.User.UserID,
				Name:   order.User.Name,
				Email:  order.User.Email,
			},
			OrderItems: orderItems,
		})
	}

	return orderResponses, nil
}

func (or *OrderRepository) CreateOrder(userID int) (*entity.OrderResponse, error) {
	// Step 1: Get the user's cart items
	var carts []entity.Cart
	if err := or.DB.
		Preload("Product"). // Preload the product details for each cart item
		Where("user_id = ?", userID).
		Find(&carts).Error; err != nil {
		return nil, err
	}

	// Step 2: Calculate the total price of the order
	var totalPrice float64
	for _, cart := range carts {
		totalPrice += float64(cart.Quantity) * cart.Product.Price
	}

	// Step 3: Create the order
	order := entity.Order{
		UserID:     uint(userID),
		TotalPrice: totalPrice,
	}

	// Insert the order into the Orders table
	if err := or.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	// Step 4: Insert the order items into the OrderItems table
	for _, cart := range carts {
		orderItem := entity.OrderItem{
			OrderID:   order.OrderID, // Reference to the newly created order
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.Product.Price,
		}

		if err := or.DB.Create(&orderItem).Error; err != nil {
			return nil, err
		}
	}

	// Step 5: Delete the cart items
	if err := or.DB.Delete(&carts).Error; err != nil {
		return nil, err
	}

	// Step 6: Return the order response
	orderResponse := entity.OrderResponse{
		OrderID:    order.OrderID,
		UserID:     order.UserID,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		// OrderItems: []entity.OrderItem{},
	}

	return &orderResponse, nil
}
