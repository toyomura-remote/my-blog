package usecase

import (
	"context"
	"errors"
	dto "my-blog-backend/DTO"
	"my-blog-backend/domain"
)

//

type PostUseCase interface {
	GetPosts(ctx context.Context) ([]*domain.Post, error)
	GetPostByDid(ctx context.Context, did dto.GetPostByDidInput) (*domain.Post, error)
	GetPostsByUserID(ctx context.Context, userID int64) ([]*domain.Post, error)
	CreatePost(ctx context.Context, createPostInput dto.CreatePostInput, userID int64) error
	UpdatePost(ctx context.Context, updatePostInput dto.UpdatePostInput, did string) error
	DeletePost(ctx context.Context, did string, currentUserID int64) error
}

type postUseCase struct {
	postRepo domain.PostRepository
}

func NewPostUseCase(postRepo domain.PostRepository) PostUseCase {
	return &postUseCase{postRepo: postRepo}
}

func (u *postUseCase) GetPosts(ctx context.Context) ([]*domain.Post, error) {
	posts, err := u.postRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func (u *postUseCase) GetPostByDid(ctx context.Context, input dto.GetPostByDidInput) (*domain.Post, error) {
	post, err := u.postRepo.GetByDid(ctx, input.Did)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (u *postUseCase) GetPostsByUserID(ctx context.Context, userID int64) ([]*domain.Post, error) {
	posts, err := u.postRepo.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (u *postUseCase) CreatePost(ctx context.Context, createPostInput dto.CreatePostInput, userID int64) error {

	newPost, err := domain.NewPost(createPostInput, userID)
	if err != nil {
		return err
	}

	if err := u.postRepo.Create(ctx, newPost); err != nil {
		return err
	}
	return nil
}

func (u *postUseCase) UpdatePost(ctx context.Context, updatePostInput dto.UpdatePostInput, did string) error {
	targetPost, err := u.postRepo.GetByDid(ctx, did)
	if err != nil {
		return err
	}

	targetPost.Title = updatePostInput.Title
	targetPost.Content = updatePostInput.Content

	if err := u.postRepo.Update(ctx, targetPost); err != nil {
		return err
	}

	return nil
}

func (u *postUseCase) DeletePost(ctx context.Context, did string, currentUserID int64) error {

	targetPost, err := u.postRepo.GetByDid(ctx, did)
	if err != nil {
		return err
	}

	if targetPost.UserID != currentUserID {
		return errors.New("unauthorized to delete this post")
	}

	targetPost.IsDeleted = true

	if err := u.postRepo.Delete(ctx, targetPost); err != nil {
		return err
	}

	return nil
}
