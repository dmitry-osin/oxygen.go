package domain

import (
	"context"

	"github.com/uptrace/bun"
)

type PostRepository struct {
	db *bun.DB
}

func NewPostRepository(db *bun.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repo *PostRepository) Find(id int64) (*Post, error) {
	ctx := context.Background()

	post := new(Post)
	err := repo.db.NewSelect().
		Model(post).
		Relation("Author").
		Relation("Tags").
		Where("p.id = ?", id).
		Scan(ctx)

	return post, err
}

func (repo *PostRepository) FindBySlug(slug string) (*Post, error) {
	ctx := context.Background()
	post := new(Post)
	err := repo.db.NewSelect().
		Model(post).
		Relation("Author").
		Relation("Tags").
		Where("slug = ?", slug).
		Scan(ctx)

	return post, err
}

func (repo *PostRepository) Delete(post *Post) error {
	return repo.db.RunInTx(context.Background(), nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewDelete().
			Model(post).
			WherePK().
			Exec(ctx)

		return err
	})
}

func (repo *PostRepository) FindByUser(login string) ([]Post, error) {
	ctx := context.Background()

	var posts []Post
	err := repo.db.NewSelect().
		Model(&posts).
		Relation("Author").
		Relation("Tags").
		Where("author.login = ?", login).
		Scan(ctx)

	return posts, err
}
