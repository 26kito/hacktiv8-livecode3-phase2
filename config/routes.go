package config

import (
	"hacktiv8-lc3-p2/repository"
	"hacktiv8-lc3-p2/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) {
	e := echo.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	e.POST("/users/register", userService.Register)
	e.POST("/users/login", userService.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
