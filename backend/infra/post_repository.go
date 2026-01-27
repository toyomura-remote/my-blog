package infra

import (
	"context"
	"database/sql"
	"my-blog-backend/domain"
	"my-blog-backend/infra/model"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) domain.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetAll(ctx context.Context) ([]*domain.Post, error) {
	posts, err := model.Posts(
		qm.Load(model.PostRels.User),
		qm.OrderBy("created_at DESC"),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var items = make([]*domain.Post, len(posts))

	for i, post := range posts {
		println("users did:", post.R.User.Did)
		items[i] = &domain.Post{
			UserID:  post.UserID,
			Did:     post.Did,
			Title:   post.Title,
			Content: post.Content,
			User: &domain.User{
				Did:  post.R.User.Did,
				Name: post.R.User.Name,
			},
		}
	}

	return items, nil
}

func (r *postRepository) GetByDid(ctx context.Context, did string) (*domain.Post, error) {
	return nil, nil
}
func (r *postRepository) GetAllByUserID(ctx context.Context, userID int64) ([]*domain.Post, error) {
	posts, err := model.Posts(
		qm.Load(model.PostRels.User),
		qm.Where("user_id = ?", userID),
		qm.OrderBy("created_at DESC"),
	).All(ctx, r.db)

	if err != nil {
		return nil, err
	}

	var items = make([]*domain.Post, len(posts))

	for i, post := range posts {
		items[i] = &domain.Post{
			UserID:  post.UserID,
			Did:     post.Did,
			Title:   post.Title,
			Content: post.Content,
			User: &domain.User{
				ID:   post.R.User.ID,
				Did:  post.R.User.Did,
				Name: post.R.User.Name,
			},
		}
	}

	return items, nil

}

func (r *postRepository) Create(ctx context.Context, post *domain.Post) error {
	p := model.Post{
		UserID:  post.UserID,
		Did:     post.Did,
		Title:   post.Title,
		Content: post.Content,
	}

	if err := p.Insert(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (r *postRepository) Update(ctx context.Context, post *domain.Post) error {
	return nil
}
func (r *postRepository) Delete(ctx context.Context, postID int64) error {
	return nil
}
