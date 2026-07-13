package service

import (
	"context"
	"errors"
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ihsanpraditya/gin-clean-1/internal/model"
	"github.com/ihsanpraditya/gin-clean-1/internal/dto"
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
	jwtSecret []byte
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: []byte(os.Getenv("JWT_KEY")),
	}
}

func (s *UserService) GetRoleByID(ctx context.Context, id uint) (dto.Role, error) {
	role, err := s.repo.FindRoleByID(ctx, id)
	if err != nil {
		return dto.Role{}, err
	}
	return dto.Role{
		ID:   role.ID,
		Name: role.Name,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (dto.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.User{}, err 
	}
	return toUserDTO(user), nil
}

func (s *UserService) CreateUser(ctx context.Context, input *dto.CreateUser) (dto.User, error) {
	existing, _ := s.repo.FindByEmail(ctx, input.Email)
	if existing != nil {
		return dto.User{}, ErrEmailTaken
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.User{}, err
	}
	
	newUser := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(ctx, newUser); err != nil {
		return dto.User{}, err
	}

	return toUserDTO(newUser), nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, dto.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", dto.User{}, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", dto.User{}, ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", dto.User{}, err
	}

	return tokenString, toUserDTO(user), nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]dto.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	
	usersDto := make([]dto.User, 0, len(users))
	for _, user := range users {
		usersDto = append(usersDto, toUserDTO(&user))
	}
	return usersDto, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, input dto.UpdateUser) (dto.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.User{}, ErrUserNotFound
	}
	user.Name = input.Name
	user.Email = input.Email
	user.IsActive = input.IsActive

	user.Roles = make([]model.Role, len(*input.Roles))
	for i, roleID := range *input.Roles {
		user.Roles[i] = model.Role{ID: roleID}
	}

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return dto.User{}, err
		}
		user.Password = string(hashedPassword)
	}

	updatedUser, err := s.repo.Update(ctx, user)
	if err != nil {
		return dto.User{}, err
	}
	
	return toUserDTO(updatedUser), nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func toUserDTO(u *model.User) dto.User {
	rolesDto := make([]dto.Role, len(u.Roles))
	for i := range u.Roles {
		rolesDto[i] = dto.Role{
			ID:   u.Roles[i].ID,
			Name: u.Roles[i].Name,
		}
	}

	return dto.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Roles:    rolesDto,
		IsActive: u.IsActive,
	}
}