package repository

import (
	"context"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/model"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/query"

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
		Order(query.User.ID.Asc()).
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update memperbarui data user
func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete menghapus data user berdasarkan ID
func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Where(query.User.ID.Eq(id)).
		Delete(&model.User{}).Error
}
