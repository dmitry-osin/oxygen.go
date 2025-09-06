package domain

import (
	"fmt"
	"log"
	"oxygenBlog/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	applicationConfig, err := config.ReadConfig()
	if err != nil {
		return nil
	}
	dbConfig := applicationConfig.Db

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.Name, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Post{}, &Tag{}, &User{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func paginate(pageable Page) func(db *gorm.DB) *gorm.DB {
	page := pageable.page
	limit := pageable.limit

	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		switch {
		case limit > 20:
			limit = 20
		case limit < 10:
			limit = 10
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

type Page struct {
	page  int
	limit int
}
