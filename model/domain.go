package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primaryKey"`
	Username  string         `gorm:"type:varchar(64);not null"`
	Email     string         `gorm:"type:varchar(255);not null"`
	FirstName string         `gorm:"type:varchar(64);not null"`
	LastName  string         `gorm:"type:varchar(64);not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Post struct {
	ID          int            `gorm:"primaryKey"`
	Title       string         `gorm:"type:varchar(255);not null"`
	Content     string         `gorm:"type:text;not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Author      User           `gorm:"foreignKey:AuthorID"`
	AuthorID    int            `gorm:"not null"`
	IsPublished bool           `gorm:"not null"`
}

type Category struct {
	ID        int            `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PostCategory struct {
	PostID     int `gorm:"primaryKey"`
	CategoryID int `gorm:"primaryKey"`
}
