package model

import (
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindByID(id int) (*Category, error) {
	var category Category
	if err := r.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) FindAll() ([]Category, error) {
	var categories []Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) Update(category *Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(category *Category) error {
	return r.db.Delete(category).Error
}

func (r *CategoryRepository) FindPostsByCategoryID(categoryID int) ([]Post, error) {
	var posts []Post
	if err := r.db.Joins("JOIN post_categories ON posts.id = post_categories.post_id").
		Where("post_categories.category_id = ?", categoryID).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
