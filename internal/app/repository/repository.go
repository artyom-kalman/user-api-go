package repository

import (
	"github.com/artyom-kalman/user-api-go/pkg/db"
)

type UserRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
