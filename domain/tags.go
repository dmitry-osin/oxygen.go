package domain

import "gorm.io/gorm"

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) FindAll(pageable Page) ([]Tag, error) {
	var tags []Tag
	err := r.db.Scopes(paginate(pageable)).Find(&tags).Error
	return tags, err
}

func (r *TagRepository) FindByID(id string) (Tag, error) {
	var tag Tag
	err := r.db.Where("id = ?", id).First(&tag).Error
	return tag, err
}

func (r *TagRepository) Create(tag Tag) (uint, error) {
	err := r.db.Create(&tag).Error
	return tag.ID, err
}

func (r *TagRepository) Update(tag Tag) error {
	err := r.db.Save(&tag).Error
	return err
}

func (r *TagRepository) Delete(tag Tag) error {
	err := r.db.Where("id = ?", tag.ID).Delete(&Tag{}).Error
	return err
}
