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

func (cs *CartService) AddToCart(c echo.Context) error {
	var request entity.AddToCartPayload

	c.Bind(&request)

	// Get the user ID from the JWT claims
	userID := c.Get("user").(jwt.MapClaims)["user_id"].(float64)

	// Add the product to the user's cart
	response, err := cs.CartRepository.AddToCart(int(userID), request.ProductID, request.Quantity)

	if err != nil {
		errCode, _ := strconv.Atoi(err.Error()[:3])
		errMessage := err.Error()[6:]

		return c.JSON(errCode, entity.ResponseError{
			Status:  errCode,
			Message: errMessage,
		})
	}

	return c.JSON(201, entity.ResponseOK{
		Status:  201,
		Message: "Product added to cart successfully",
		Data:    response,
	})
}

func (cs *CartService) DeleteCartByID(c echo.Context) error {
	// Get the user ID from the JWT claims
	userID := c.Get("user").(jwt.MapClaims)["user_id"].(float64)

	// Get the cart ID from the URL parameter
	cartID, err := strconv.Atoi(c.Param("cart_id"))

	if err != nil {
		return c.JSON(400, entity.ResponseError{
			Status:  400,
			Message: "Invalid cart ID",
		})
	}

	// Delete the cart from the user's cart
	err = cs.CartRepository.DeleteCartByID(int(userID), cartID)

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
		Message: "Cart deleted successfully",
	})
}
