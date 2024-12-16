package service

import (
	"fmt"
	"hacktiv8-lc3-p2/entity"
	"hacktiv8-lc3-p2/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (us *UserService) Register(c echo.Context) error {
	var request entity.UserRegisterPayload

	c.Bind(&request)

	if err := validateUserRegisterPayload(request); err != nil {
		errCode, _ := strconv.Atoi(err.Error()[:3])
		errMessage := err.Error()[6:]

		return c.JSON(errCode, entity.ResponseError{
			Status:  errCode,
			Message: errMessage,
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(hashedPassword)

	response, err := us.UserRepository.Register(request)

	if err != nil {
		errCode, _ := strconv.Atoi(err.Error()[:3])
		errMessage := err.Error()[6:]

		return c.JSON(int(errCode), entity.ResponseError{
			Status:  errCode,
			Message: errMessage,
		})
	}

	return c.JSON(201, entity.ResponseOK{
		Status:  201,
		Message: "User registered successfully",
		Data:    response,
	})
}

func validateUserRegisterPayload(request entity.UserRegisterPayload) error {
	if request.Name == "" {
		return fmt.Errorf("400 | Name is required")
	}

	if request.Email == "" {
		return fmt.Errorf("400 | Email is required")
	}

	if request.Password == "" {
		return fmt.Errorf("400 | Password is required")
	}

	return nil
}
