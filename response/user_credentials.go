package response

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/ptr"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrPasswordAuthNotAllowed is returned when attempting to verify credentials for a User who does not
	// have a password set. This is commonly returned when a User registers with OAuth 2.0 and never sets
	// a password.
	ErrPasswordAuthNotAllowed = errors.New("password auth not allowed for this user")

	// ErrPasswordAuthDisabled is returned when password-based authentication for users has been disabled
	// by a Response administrator.
	ErrPasswordAuthDisabled = errors.New("password auth is disabled")
)

func (c *Core) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c.config.AuthPasswordBcryptCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	return string(bytes), nil
}

func (c *Core) RegisterWithPassword(ctx context.Context, name, email, password string) (*ent.User, error) {
	if c.config.AuthPasswordDisabled {
		return nil, ErrPasswordAuthDisabled
	}

	hashedPassword, err := c.hashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := c.CreateUser(ctx, ent.CreateUser{
		Name:     name,
		Email:    email,
		Password: ptr.StringP(hashedPassword),
	})

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("email", email).
			Msg("Registration failed during RegisterWithPassword.")

		return nil, fmt.Errorf("CreateUser: %w", err)
	}

	return user, nil
}

func (c *Core) VerifyPasswordCredentials(ctx context.Context, email, password string) (bool, error) {
	if c.config.AuthPasswordDisabled {
		return false, ErrPasswordAuthDisabled
	}

	user, err := c.GetUserByEmail(ctx, email)
	if err != nil {
		return false, fmt.Errorf("GetUserByEmail: %w", err)
	}

	if user.Password == nil || *user.Password == "" {
		return false, ErrPasswordAuthNotAllowed
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password)); err != nil {
		log.Error().
			Err(err).
			Str("email", email).
			Msg("Comparison of hash and password failed.")

		return false, nil
	}

	return true, nil
}
