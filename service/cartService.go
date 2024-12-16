package service

import (
	"hacktiv8-lc3-p2/entity"
	"hacktiv8-lc3-p2/repository"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CartService struct {
	CartRepository *repository.CartRepository
}

func NewCartService(cartRepository *repository.CartRepository) *CartService {
	return &CartService{CartRepository: cartRepository}
}

func (cs *CartService) GetCarts(c echo.Context) error {
	// Get the user ID from the JWT claims
	userID := c.Get("user").(jwt.MapClaims)["user_id"].(float64)

	// Fetch the user's carts from the database
	carts, err := cs.CartRepository.GetCarts(int(userID))

	if err != nil {
		errCode, _ := strconv.Atoi(err.Error()[:3])
		errMessage := err.Error()[6:]

		return c.JSON(errCode, entity.ResponseError{
			Status:  errCode,
			Message: errMessage,
		})
	}

	return c.JSON(200, entity.ResponseOK{
		Status:  200,
		Message: "Carts fetched successfully",
		Data:    carts,
	})
}
