package service

import (
	"hacktiv8-lc3-p2/entity"
	"hacktiv8-lc3-p2/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderService struct {
	OrderRepository *repository.OrderRepository
}

func NewOrderService(orderRepository *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepository: orderRepository}
}

func (os *OrderService) GetOrders(c echo.Context) error {
	// Get the user ID from the JWT claims
	userID := c.Get("user").(jwt.MapClaims)["user_id"].(float64)

	orders, err := os.OrderRepository.GetOrders(int(userID))
	if err != nil {
		return c.JSON(500, entity.ResponseError{
			Status:  500,
			Message: "Internal server error",
		})
	}

	return c.JSON(200, entity.ResponseOK{
		Status:  200,
		Message: "Success",
		Data:    orders,
	})
}

func (os *OrderService) CreateOrder(c echo.Context) error {
	// Get the user ID from the JWT claims
	userID := c.Get("user").(jwt.MapClaims)["user_id"].(float64)

	// Save the new order to the database
	order, err := os.OrderRepository.CreateOrder(int(userID))
	if err != nil {
		return c.JSON(500, entity.ResponseError{
			Status:  500,
			Message: "Internal server error",
		})
	}

	return c.JSON(200, entity.ResponseOK{
		Status:  200,
		Message: "Success",
		Data:    order,
	})
}
