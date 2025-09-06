package domain

import (
	"context"

	"github.com/uptrace/bun"
)

type TagRepository struct {
	db *bun.DB
}

func NewTagRepository(db *bun.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (repo *TagRepository) Create(tag *Tag) error {
	ctx := context.Background()
	_, err := repo.db.NewInsert().
		Model(tag).
		Exec(ctx)
	return err
}

func (repo TagRepository) FindByName(name string) (*Tag, error) {
	ctx := context.Background()

	tag := new(Tag)
	err := repo.db.NewSelect().
		Model(tag).
		Where("t.name = ?", name).
		Scan(ctx)

	return tag, err
}

func (repo TagRepository) FindByNameIncludePosts(name string, limit int) ([]Tag, error) {
	ctx := context.Background()

	var tags []Tag
	err := repo.db.NewSelect().
		Model(tags).
		Relation("Posts").
		Where("t.name = ?", name).
		Limit(limit).
		Scan(ctx)

	return tags, err
}
