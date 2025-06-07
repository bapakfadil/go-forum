package memberships

import (
	"context"
	"errors"

	"github.com/bapakfadil/fastcampus/internal/models/memberships"
	"github.com/bapakfadil/fastcampus/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user data.")
		return "", err
	}

	if user == nil {
		log.Warn().Str("email", req.Email).Msg("Email not found.")
		return "", errors.New("email not found")
	}

	log.Info().Str("submittedEmail", req.Email).Str("storedUsername", user.Username).Str("submittedPassword", req.Password).Msg("Comparing passwords")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("Password comparison failed")
		return "", errors.New("invalid password")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}

	return token, nil
}
