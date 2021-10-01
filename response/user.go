package response

import (
	"context"
	"errors"
	"fmt"

	"github.com/responserms/response/internal/errs"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/ent/user"
)

var (
	ErrEmailExists = errors.New("email already exists")
)

// GetUserByID returns an ent.User by their unique ID. This method will return nil
// for both ent.User and error if the User does not exist.
func (c *Core) GetUserByID(ctx context.Context, id int) (*ent.User, error) {
	res, err := c.ent.User.
		Get(ctx, id)

	if ent.MaskNotFound(err) != nil {
		return nil, err
	}

	return res, nil
}

// GetUserByEmail returns an ent.User by their email address. This method will return
// nil for both ent.User and error if the User does not exist.
func (c *Core) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	res, err := c.ent.User.
		Query().
		Where(user.Email(email)).
		First(ctx)

	if ent.MaskNotFound(err) != nil {
		return nil, err
	}

	return res, nil
}

// CreateUser handles creation of the User in Response. The provided ent.CreateUser
// will be validated to ensure all requirements are met.
//
// This method will return nil for both values if validation fails and the user
// creation was not attempted.
//
// If the Email already exists ErrEmailExists is returned.
func (c *Core) CreateUser(ctx context.Context, create ent.CreateUser) (*ent.User, error) {
	if errs.FieldsHaveValidationErrors(ctx, c.validator.StructCtx(ctx, create)) {
		return nil, nil
	}

	val, err := c.GetUserByEmail(ctx, create.Email)
	if err != nil {
		return nil, fmt.Errorf("CreateUser GetUserByEmail: %w", err)
	}

	if val != nil {
		return nil, ErrEmailExists
	}

	val, err = ent.FromContext(ctx).
		User.
		Create().
		SetInput(create).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("CreateUser User.Create: %w", err)
	}

	return val, nil
}

// UpdateUserByID updates a User by their unique ID according to the ent.UpdateUser
// specifications.
func (c *Core) UpdateUserByID(ctx context.Context, id int, update ent.UpdateUser) (*ent.User, error) {
	val, err := ent.FromContext(ctx).
		User.
		UpdateOneID(id).
		SetInput(update).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("UpdateUserByID User.UpdateOneID: %w", err)
	}

	return val, nil
}
