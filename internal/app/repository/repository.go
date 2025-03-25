package repository

import (
	"github.com/artyom-kalman/user-api-go/pkg/db"
)

type userRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) *userRepository {
	return &userRepository{
		db: db,
	}
}
