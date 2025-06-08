package posts

import (
	"context"

	"github.com/bapakfadil/fastcampus/internal/configs"
	"github.com/bapakfadil/fastcampus/internal/models/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	CountLikeByPostID(ctx context.Context, postID int64) (int, error)
}

type service struct {
	cfg            *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg      : cfg,
		postRepo : postRepo,
	}
}