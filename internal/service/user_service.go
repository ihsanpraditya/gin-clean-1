package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ihsanpraditya/gin-clean-1/internal/model"
	"github.com/ihsanpraditya/gin-clean-1/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailTaken         = errors.New("email already taken")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type UserService struct {
	repo *repository.UserRepository
	// JWT Secret Key (Idealnya ditaruh di internal/config)
	jwtSecret []byte
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: []byte("super_secret_jwt_key_change_me"),
	}
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) RegisterUser(ctx context.Context, user *model.User) error {
	existing, _ := s.repo.FindByEmail(ctx, user.Email)
	if existing != nil {
		return ErrEmailTaken
	}

	// Hash password sebelum disimpan ke repository
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.repo.Create(ctx, user)
}

// Login memvalidasi kredensial dan mengembalikan JWT token jika sukses
func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Bandingkan password plain text dari request dengan hash dari DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Membuat klaim token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token kedaluwarsa dalam 3 hari
	})

	// Menandatangani token menggunakan secret key
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, name, email string) (*model.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	user.Name = name
	user.Email = email
	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
