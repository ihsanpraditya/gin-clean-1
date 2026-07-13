package service

import (
	"context"
	"errors"

	"github.com/ihsanpraditya/gin-clean-1/internal/model"
	"github.com/ihsanpraditya/gin-clean-1/internal/dto"
	"github.com/ihsanpraditya/gin-clean-1/internal/repository"
)

var (
	ErrRoleNotFound = errors.New("role not found")
	ErrNameTaken = errors.New("name already taken")
)

type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) GetRoleByID(ctx context.Context, id uint) (dto.Role, error) {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.Role{}, err 
	}
	return toRoleDTO(role), nil
}

func (s *RoleService) CreateRole(ctx context.Context, name string) (dto.Role, error) {
	existing, _ := s.repo.FindByName(ctx, name)
	if existing != nil {
		return dto.Role{}, ErrNameTaken
	}

	newRole := &model.Role{Name: name}

	createdRole, err := s.repo.Create(ctx, newRole)
	if err != nil {
		return dto.Role{}, err
	}

	return toRoleDTO(createdRole), nil
}

func (s *RoleService) GetAllRoles(ctx context.Context) ([]dto.Role, error) {
	roles, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	
	rolesDto := make([]dto.Role, 0, len(roles))
	for _, role := range roles {
		rolesDto = append(rolesDto, toRoleDTO(&role))
	}
	return rolesDto, nil
}

func (s *RoleService) UpdateRole(ctx context.Context, id uint, name string) (dto.Role, error) {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.Role{}, ErrRoleNotFound
	}
	role.Name = name

	updatedRole, err := s.repo.Update(ctx, role)
	if err != nil {
		return dto.Role{}, err
	}
	
	return toRoleDTO(updatedRole), nil
}

func (s *RoleService) DeleteRole(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func toRoleDTO(r *model.Role) dto.Role {
	return dto.Role{
		ID:       r.ID,
		Name:     r.Name,
	}
}