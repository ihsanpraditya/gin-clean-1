package service

import (
	"context"
	"errors"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/model"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrEmailTaken    = errors.New("email already taken")
	ErrUserNotFound  = errors.New("user not found")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) RegisterUser(ctx context.Context, user *model.User) error {
	existing, _ := s.repo.FindByEmail(ctx, user.Email)
	if existing != nil {
		return ErrEmailTaken
	}

	// TODO: Lakukan hashing password di sini sebelum disimpan ke repo
	// user.Password = hashPassword(user.Password)

	return s.repo.Create(ctx, user)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, name string, email string) (*model.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Validasi jika email baru ternyata sudah digunakan oleh user lain
	if email != user.Email {
		existing, _ := s.repo.FindByEmail(ctx, email)
		if existing != nil {
			return nil, ErrEmailTaken
		}
	}

	user.Name = name
	user.Email = email

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return s.repo.Delete(ctx, id)
}
