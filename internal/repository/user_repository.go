package repository

import (
	"context"
	"github.com/ihsanpraditya/gin-clean-1/internal/model"
	"github.com/ihsanpraditya/gin-clean-1/internal/query"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByEmail menggunakan query type-safe
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Where(query.User.Email.Eq(email)).
		Take(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID mencari user berdasarkan ID
func (r *UserRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).
		Preload("Roles").
		Where(query.User.ID.Eq(id)).
		Take(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create memasukkan data user baru
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindAll mengambil semua data user
func (r *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).
		Preload("Roles").
		Order(query.User.ID.Asc()).
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return err
		}

		if err := tx.Model(user).Association("Roles").Replace(user.Roles); err != nil {
			return err
		}

		if err := tx.Preload("Roles").First(user, user.ID).Error; err != nil {
			return err
		}
		
		return nil
	})

	return user, err
}

// Delete menghapus data user berdasarkan ID
func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Where(query.User.ID.Eq(id)).
		Delete(&model.User{}).Error
}

func (r *UserRepository) DeleteUsers(ctx context.Context, ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).
		Where(query.User.ID.In(ids...)).
		Delete(&model.User{}).Error
}
