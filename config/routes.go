package config

import (
	"hacktiv8-lc3-p2/middleware"
	"hacktiv8-lc3-p2/repository"
	"hacktiv8-lc3-p2/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) {
	e := echo.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	e.POST("/users/register", userService.Register)
	e.POST("/users/login", userService.Login)

	e.GET("/users/carts", cartService.GetCarts, middleware.ValidateJWTMiddleware)
	e.POST("/users/carts", cartService.AddToCart, middleware.ValidateJWTMiddleware)
	e.DELETE("/users/carts/:cart_id", cartService.DeleteCartByID, middleware.ValidateJWTMiddleware)
	e.GET("/users/orders", orderService.GetOrders, middleware.ValidateJWTMiddleware)
	e.POST("/users/orders", orderService.CreateOrder, middleware.ValidateJWTMiddleware)
	e.GET("/products", productService.GetProducts)
	e.GET("/products/:product_id", productService.GetProductByID)

	e.Logger.Fatal(e.Start(":8080"))
}
