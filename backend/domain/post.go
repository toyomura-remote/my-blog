package domain

import (
	dto "my-blog-backend/DTO"

	"github.com/google/uuid"
)

type Post struct {
	UserID  int64
	Did     string
	Title   string
	Content string
	User    *User
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
