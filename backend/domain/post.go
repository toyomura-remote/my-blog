package domain

import (
	dto "my-blog-backend/DTO"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        int64     `json:"-"`
	UserID    int64     `json:"user_id"`
	Did       string    `json:"did"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsDeleted bool      `json:"is_deleted"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPost(createPostInput dto.CreatePostInput, userID int64) (*Post, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	post := Post{
		UserID:  userID,
		Did:     uuid.String(),
		Title:   createPostInput.Title,
		Content: createPostInput.Content,
	}

	return &post, nil
}
