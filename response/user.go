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

func (c *Core) GetUserByID(ctx context.Context, id int) (*ent.User, error) {
	res, err := c.ent.User.
		Get(ctx, id)

	if ent.MaskNotFound(err) != nil {
		return nil, err
	}

	return res, nil
}

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

func (c *Core) CreateUser(ctx context.Context, create ent.CreateUser) (*ent.User, error) {
	if errs.FieldsHaveValidationErrors(ctx, c.validator.StructCtx(ctx, create)) {
		return nil, nil
	}

	val, err := c.GetUserByEmail(ctx, create.Email)
	if err != nil {
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
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
		return nil, fmt.Errorf("User.Create: %w", err)
	}

	return val, nil
}

func (c *Core) UpdateUserByID(ctx context.Context, id int, create ent.UpdateUser) (*ent.User, error) {
	val, err := ent.FromContext(ctx).
		User.
		UpdateOneID(id).
		SetInput(create).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("User.UpdateOneID: %w", err)
	}

	return val, nil
}
