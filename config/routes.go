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

	e.POST("/users/register", userService.Register)
	e.POST("/users/login", userService.Login)

	e.GET("/users/carts", cartService.GetCarts, middleware.ValidateJWTMiddleware)
	e.POST("/users/carts", cartService.AddToCart, middleware.ValidateJWTMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
