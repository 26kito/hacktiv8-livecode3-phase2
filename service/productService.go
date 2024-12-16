package service

import (
	"hacktiv8-lc3-p2/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (ps *ProductService) GetProducts(c echo.Context) error {
	products, err := ps.ProductRepository.GetProducts()

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "Internal server error",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "Success",
		"data":    products,
	})
}

func (ps *ProductService) GetProductByID(c echo.Context) error {
	productID, _ := strconv.Atoi(c.Param("product_id"))
	product, err := ps.ProductRepository.GetProductByID(productID)

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "Internal server error",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "Success",
		"data":    product,
	})
}
