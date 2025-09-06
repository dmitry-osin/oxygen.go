package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64      `bun:",pk,autoincrement"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     *time.Time `bun:",soft_delete,nullzero"`
	Login         string     `bun:",type:varchar(64),unique,notnull"`
	Name          string     `bun:",type:varchar(128),notnull"`
	Surname       *string    `bun:",type:varchar(128)"`
	Bio           *string    `bun:",type:varchar(255)"`
	Email         string     `bun:",type:varchar(128),unique,notnull"`
	Password      string     `bun:",type:varchar(256),notnull"`
	Avatar        *string    `bun:",type:varchar(255)"`
}

type Tag struct {
	bun.BaseModel `bun:"table:tags,alias:t"`
	ID            int64      `bun:",pk,autoincrement"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     *time.Time `bun:",soft_delete,nullzero"`
	Name          string     `bun:",type:varchar(64),notnull"`
	Slug          string     `bun:",type:varchar(64),unique,notnull"`
	AuthorID      int64      `bun:",notnull"`
	Author        *User      `bun:"rel:belongs-to,join:author_id=id"`
	Posts         []*Post    `bun:"m2m:posts_tags"`
}

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`
	ID            int64      `bun:",pk,autoincrement"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     *time.Time `bun:",soft_delete,nullzero"`
	Title         string     `bun:",type:varchar(256),notnull"`
	Slug          string     `bun:",type:varchar(256),unique,notnull"`
	Description   string     `bun:",type:varchar(256),notnull"`
	Content       string     `bun:",type:text,notnull"`
	AuthorID      int64      `bun:",notnull"`
	Author        *User      `bun:"rel:belongs-to,join:author_id=id"`
	Tags          []*Tag     `bun:"m2m:posts_tags"`
}

type PostsTags struct {
	bun.BaseModel `bun:"table:posts_tags,alias:pt"`
	PostID        int64     `bun:"post_id,pk"`
	TagID         int64     `bun:"tag_id,pk"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Post          *Post     `bun:"rel:belongs-to,join:post_id=id"`
	Tag           *Tag      `bun:"rel:belongs-to,join:tag_id=id"`
}
