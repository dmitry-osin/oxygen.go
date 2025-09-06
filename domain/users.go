package domain

import (
	"context"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(user *User) error {
	ctx := context.Background()

	_, err := repo.db.NewInsert().
		Model(user).
		Exec(ctx)

	return err
}

func (repo *UserRepository) FindByLogin(login string) (*User, error) {
	ctx := context.Background()

	user := new(User)
	err := repo.db.NewSelect().
		Model(user).
		Where("u.login = ?", login).
		Scan(ctx)

	return user, err
}
