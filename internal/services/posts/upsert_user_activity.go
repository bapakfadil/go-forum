package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/bapakfadil/fastcampus/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {

	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activity from database")
		return err
	}

	if userActivity == nil {
		// Menambahkan user activity
		if !request.IsLiked {
			return errors.New("anda belum pernah like sebelumnya")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		// Update user activity
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	// Error handling ketika gagal
	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity to database")
		return err
	}
	return nil
}
