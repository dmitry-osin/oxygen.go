package test

import (
	"testing"

	"oxygen.go/model"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestPostRepository_FindPublishedPosts(t *testing.T) {
	// Set up in-memory test database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Run migration to create the table
	err = db.AutoMigrate(&model.Post{})
	assert.NoError(t, err)

	repo := model.NewPostRepository(db)

	// Create test data
	testPosts := []model.Post{
		{Title: "Published Post 1", Content: "Content 1", IsPublished: true},
		{Title: "Draft Post", Content: "Content 2", IsPublished: false},
		{Title: "Published Post 2", Content: "Content 3", IsPublished: true},
	}

	// Save test posts
	for _, post := range testPosts {
		err := repo.Create(&post)
		assert.NoError(t, err)
	}

	// Execute the tested method
	publishedPosts, err := repo.FindPublishedPosts()

	// Check results
	assert.NoError(t, err)
	assert.Len(t, publishedPosts, 2)

	// Verify that all retrieved posts are published
	for _, post := range publishedPosts {
		assert.True(t, post.IsPublished)
	}
}
