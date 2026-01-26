package domain

import (
	"context"
)

type PostRepository interface {
	GetAll(ctx context.Context) ([]*Post, error)
	GetByDid(ctx context.Context, did string) (*Post, error)
	GetAllByUserID(ctx context.Context, userID int64) ([]*Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, postID int64) error
}
