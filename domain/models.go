package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey;autoIncrement:true;column:id;comment:User ID"`
	Login    string  `gorm:"type:varchar(64);unique;not null;column:login;comment:User login"`
	Name     string  `gorm:"type:varchar(128);not null;column:name;comment:User name"`
	Surname  *string `gorm:"type:varchar(128);column:surname;comment:User surname"`
	Bio      *string `gorm:"type:varchar(255);column:bio;comment:User bio"`
	Email    string  `gorm:"type:varchar(128);not null;column:email;comment:User email"`
	Password string  `gorm:"type:varchar(256);not null;column:password;comment:User hashed password"`
	Avatar   *string `gorm:"type:varchar(255);not null;column:avatar;comment:User avatar file location"`
}

type Tag struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey;autoIncrement:true;column:id;comment:Tag ID"`
	Name     string  `gorm:"type:varchar(64);not null;column:name;comment:Tag name"`
	Slug     string  `gorm:"type:varchar(64);not null;column:slug;comment:Tag slug"`
	AuthorID uint    `gorm:"column:author_id;comment:Author ID"`
	Author   *User   `gorm:"foreignKey:AuthorID;references:id"`
	Posts    []*Post `gorm:"many2many:posts_tags"`
}

type Post struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement:true;column:id;comment:Post ID"`
	Title       string `gorm:"type:varchar(256);not null;column:title;comment:Post title"`
	Slug        string `gorm:"type:varchar(256);not null;column:slug;comment:Post slug"`
	Description string `gorm:"type:varchar(256);not null;column:description;comment:Post description"`
	Content     string `gorm:"type:text;not null;column:content;comment:Post content"`
	AuthorID    uint   `gorm:"column:author_id;comment:Author ID"`
	Author      *User  `gorm:"foreignKey:AuthorID;references:id"`
	Tags        []*Tag `gorm:"many2many:posts_tags;"`
}
