package response

import (
	"context"
	"fmt"
	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/ent/oauthconnection"
	"github.com/responserms/response/internal/ptr"
)

type OAuthConnectionEntry struct {
	ProviderKey    string
	ProviderUserID string
	UpdateProfile  bool
	Name           string
	Email          string
	AvatarURL      string
	Metadata       map[string]string
}

// GetOAuthConnectionByProvider returns an OAuthConnection if it exists with the given providerKey and providerUserID. The result
// will be nil if the OAuthConnection does not exist.
//
// If the OAuthConnection does not have a User, this is considered a data problem and the OAuthConnection will not be returned
// as it is not a valid OAuthConnection.
func (c *Core) GetOAuthConnectionByProvider(ctx context.Context, providerKey, providerUserID string) (*ent.OAuthConnection, error) {
	conn, err := c.ent.OAuthConnection.Query().
		Where(
			oauthconnection.Provider(providerKey),
			oauthconnection.ProviderUserID(providerUserID),
			oauthconnection.HasUser(),
		).
		First(ctx)

	if ent.MaskNotFound(err) != nil {
		return nil, fmt.Errorf("OAuthConnection.Query: %w", err)
	}

	return conn, nil
}

// UpsertOAuthConnection creates a new OAuthConnection unless the connection already exists. If the OAuthConnection exists, the
// existing OAuthConnection is returned instead.
func (c *Core) UpsertOAuthConnection(ctx context.Context, create ent.CreateOAuthConnection) (*ent.OAuthConnection, error) {
	conn, err := c.GetOAuthConnectionByProvider(ctx, create.Provider, create.ProviderUserID)
	if err != nil {
		return nil, err
	}

	if conn != nil {
		conn, err = conn.Update().
			SetName(create.Name).
			Save(ctx)

		if err != nil {
			return nil, fmt.Errorf("Update.SetName: %w", err)
		}

		return conn, nil
	}

	return ent.FromContext(ctx).OAuthConnection.
		Create().
		SetInput(create).
		Save(ctx)
}

// EnsureOAuthConnection checks for an existing OAuthConnection for the given ProviderKey and ProviderUserID,
// returning the existing OAuthConnection/User if it exists. Otherwise, a OAuthConnection/User is created if
// it did not already exist.
func (c *Core) EnsureOAuthConnection(ctx context.Context, entry OAuthConnectionEntry) (*ent.OAuthConnection, error) {
	conn, err := c.UpsertOAuthConnection(ctx, ent.CreateOAuthConnection{
		Provider:       entry.ProviderKey,
		ProviderUserID: entry.ProviderUserID,
		Name:           entry.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("CreateOAuthConnection: %w", err)
	}

	user, err := conn.User(ctx)
	if ent.MaskNotFound(err) != nil {
		return nil, fmt.Errorf("OAuthConnection.User: %w", err)
	}

	if user != nil {
		_, err = c.UpdateUserByID(ctx, user.ID, ent.UpdateUser{
			Name:      ptr.StringP(entry.Name),
			AvatarURL: ptr.StringP(entry.AvatarURL),
		})

		if err != nil {
			return nil, fmt.Errorf("OAuthConnection UpdateUserByID: %w", err)
		}

		return conn, nil
	}

	user, err = c.CreateUser(ctx, ent.CreateUser{
		Name:      entry.Name,
		Email:     entry.Email,
		AvatarURL: ptr.StringP(entry.AvatarURL),
	})

	if err != nil {
		return nil, fmt.Errorf("CreateUser: %w", err)
	}

	conn, err = conn.Update().
		SetUser(user).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("conn.Update: %w", err)
	}

	return conn, nil
}
