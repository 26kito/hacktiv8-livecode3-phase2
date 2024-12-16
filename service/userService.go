package service

import (
	"fmt"
	"hacktiv8-lc3-p2/entity"
	getJWTSecret "hacktiv8-lc3-p2/middleware"
	"hacktiv8-lc3-p2/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (us *UserService) Login(c echo.Context) error {
	var request entity.UserLoginPayload

	c.Bind(&request)

	if err := validateUserLoginPayload(request); err != nil {
		errCode, _ := strconv.Atoi(err.Error()[:3])
		errMessage := err.Error()[6:]

		return c.JSON(errCode, entity.ResponseError{
			Status:  errCode,
			Message: errMessage,
		})
	}

	// Generate JWT token

	response, err := us.UserRepository.Login(request)

	if err != nil {
		errCode, _ := strconv.Atoi(err.Error()[:3])
		errMessage := err.Error()[6:]

		return c.JSON(int(errCode), entity.ResponseError{
			Status:  errCode,
			Message: errMessage,
		})
	}

	jwtToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": response.UserID,
		"email":   request.Email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}).SignedString(getJWTSecret.JWTSecret)

	return c.JSON(200, entity.ResponseOK{
		Status:  200,
		Message: "User logged in successfully",
		Data: map[string]string{
			"name":      response.Name,
			"email":     response.Email,
			"jwt_token": jwtToken,
		},
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

func validateUserLoginPayload(request entity.UserLoginPayload) error {
	if request.Email == "" {
		return fmt.Errorf("400 | Email is required")
	}

	if request.Password == "" {
		return fmt.Errorf("400 | Password is required")
	}

	return nil
}
