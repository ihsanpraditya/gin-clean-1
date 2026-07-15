package service

import (
	"context"
	"time"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/ihsanpraditya/gin-clean-1/internal/dto"
	"github.com/ihsanpraditya/gin-clean-1/internal/repository"
)

var  ErrInvalidCredentials = errors.New("invalid email or password")

type AuthService struct {
	repo      *repository.UserRepository
	jwtSecret []byte
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: []byte(os.Getenv("JWT_KEY")),
	}
}

func (a *AuthService) Login(ctx context.Context, email, password string) (string, dto.User, error) {
	user, err := a.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", dto.User{}, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", dto.User{}, ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // JWT berlaku 24 jam
	})

	tokenString, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", dto.User{}, err
	}

	return tokenString, toUserDTO(user), nil
}