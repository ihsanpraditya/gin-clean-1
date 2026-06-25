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

// FindByEmail uses type-safe query fields
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	// query.User.Email.Eq(email) generates the safe SQL clause under the hood
	err := r.db.WithContext(ctx).
		Where(query.User.Email.Eq(email)).
		Take(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create inserts a new user record
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	// Since we aren't filtering by a specific field clause here,
	// we just pass the standard query target context.
	err := r.db.WithContext(ctx).
		Order(query.User.ID.Asc()). // Safely order by ID using GORM CLI field helper
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}
