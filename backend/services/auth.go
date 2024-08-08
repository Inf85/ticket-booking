package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/Inf85/ticket-booking/models"
	"github.com/Inf85/ticket-booking/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type AuthService struct {
	repository models.AuthRepository
}

func (a *AuthService) LogIn(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User, error) {
	user, err := a.repository.GetUser(ctx, "email = ?", loginData.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("Invalid credentials")
		}
		return "", nil, err
	}

	if !models.MatchedHash(loginData.Password, user.Password) {
		return "", nil, fmt.Errorf("Invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", nil, err
	}

	return token, nil, nil
}

func (a *AuthService) Register(ctx context.Context, registerData *models.AuthCredentials) (string, *models.User, error) {
	if !models.IsEmailValid(registerData.Email) {
		return "", nil, fmt.Errorf("Please provide a valid Email")
	}

	if _, err := a.repository.GetUser(ctx, "email = ?", registerData.Email); errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, fmt.Errorf("Such user is exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", nil, err
	}

	registerData.Password = string(hashedPassword)
	user, err := a.repository.RegisterUser(ctx, registerData)

	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{repository: repository}
}
