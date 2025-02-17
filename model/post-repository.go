package model

import (
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) FindByID(id int) (*Post, error) {
	var post Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindAll() ([]Post, error) {
	var posts []Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(post *Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(post *Post) error {
	return r.db.Delete(post).Error
}

func (r *PostRepository) FindPublishedPosts() ([]Post, error) {
	var posts []Post
	if err := r.db.Where("is_published = ?", true).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
