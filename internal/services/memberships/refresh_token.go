package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/bapakfadil/fastcampus/internal/models/memberships"
	"github.com/bapakfadil/fastcampus/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	// Token validator to check the requested token with database token match or not
	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}
	return token, nil
}
