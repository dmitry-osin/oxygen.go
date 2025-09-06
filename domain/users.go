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
	return repo.db.RunInTx(context.Background(), nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().
			Model(user).
			Exec(ctx)

		return err
	})
}

func (repo *UserRepository) Update(user *User) error {
	return repo.db.RunInTx(context.Background(), nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewUpdate().
			Model(user).
			Column("name", "surname", "email", "bio", "password", "avatar").
			WherePK().
			Exec(ctx)

		return err
	})
}

func (repo *UserRepository) Delete(user *User) error {
	return repo.db.RunInTx(context.Background(), nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewDelete().
			Model(user).
			WherePK().
			Exec(ctx)

		return err
	})
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
