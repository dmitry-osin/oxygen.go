package domain

import "gorm.io/gorm"

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repository *PostRepository) FindAll(pageable Page) ([]Post, error) {
	var posts []Post
	err := repository.db.Scopes(paginate(pageable)).
		Preload("Author").
		Preload("Comments").
		Find(&posts).Error
	return posts, err
}

func (repository *PostRepository) FindById(id uint) (Post, error) {
	var post Post
	err := repository.db.Where("id = ?", id).
		Preload("Author").
		Preload("Comments").
		First(&post).Error
	return post, err
}

func (repository *PostRepository) FindByAuthor(author User, pageable Page) ([]Post, error) {
	var post []Post
	err := repository.db.Scopes(paginate(pageable)).
		Where("author = ?", author).
		Preload("Comments").
		Find(&post).Error
	return post, err
}

func (repository *PostRepository) Create(post Post) (uint, error) {
	err := repository.db.Create(&post).Error
	return post.ID, err
}

func (repository *PostRepository) Update(post Post) error {
	return repository.db.Save(&post).Error
}

func (repository *PostRepository) Delete(post Post) error {
	return repository.db.Delete(&post).Error
}

func (repository *PostRepository) FindByTag(tag Tag, pageable Page) ([]Post, error) {
	var posts []Post
	err := repository.db.Scopes(paginate(pageable)).
		Preload("Author").
		Preload("Comments").
		Where("tag = ?", tag).
		Find(&posts).Error
	return posts, err
}

func (repository *PostRepository) FindDraftPosts(pageable Page) ([]Post, error) {
	var posts []Post
	err := repository.db.Scopes(paginate(pageable)).
		Preload("Author").Preload("Comments").
		Where("is_published = false").
		Find(&posts).Error
	return posts, err
}
