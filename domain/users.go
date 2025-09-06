package domain

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) FindAll(pagable Page) ([]User, error) {
	var users []User
	err := repository.db.Scopes(paginate(pagable)).Find(&users).Error
	return users, err
}

func (repository *UserRepository) FindById(id uint) (User, error) {
	var user User
	err := repository.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (repository *UserRepository) FindByLogin(login string) (User, error) {
	var user User
	err := repository.db.Where("login = ?", login).First(&user).Error
	return user, err
}

func (repository *UserRepository) FindByEmail(email string) (User, error) {
	var user User
	err := repository.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (repository *UserRepository) Create(user User) (uint, error) {
	err := repository.db.Create(&user).Error
	return user.ID, err
}

func (repository *UserRepository) Update(user User) (User, error) {
	err := repository.db.Save(&user).Error
	return user, err
}

func (repository *UserRepository) Delete(user User) error {
	err := repository.db.Delete(&user).Error
	return err
}
