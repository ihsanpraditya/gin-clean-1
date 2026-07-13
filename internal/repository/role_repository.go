package repository

import (
	"context"
	"github.com/ihsanpraditya/gin-clean-1/internal/model"
	"github.com/ihsanpraditya/gin-clean-1/internal/query"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) FindByID(ctx context.Context, id uint) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).
		Where(query.Role.ID.Eq(id)).
		Take(&role).Error

	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) FindByName(ctx context.Context, name string) (*model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).
		Where(query.Role.Name.Eq(name)).
		Take(&role).Error

	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) Create(ctx context.Context, role *model.Role) (*model.Role, error) {
	err := r.db.WithContext(ctx).Create(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) FindAll(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.WithContext(ctx).
		Order(query.Role.ID.Asc()).
		Find(&roles).Error

	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) Update(ctx context.Context, role *model.Role) (*model.Role, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(role).Error; err != nil {
			return err
		}
		
		return nil
	})

	return role, err
}

func (r *RoleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Where(query.Role.ID.Eq(id)).
		Delete(&model.Role{}).Error
}