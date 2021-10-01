package auth

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/responserms/response/internal/config"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/ptr"

	"github.com/responserms/response/response"
	"github.com/volatiletech/authboss/v3"
)

type Storer interface {
	authboss.ServerStorer
	authboss.CreatingServerStorer
	authboss.OAuth2ServerStorer
}

type authbossServerStorer struct {
	core      *response.Core
	providers map[string]*config.OAuth2Provider
}

func newServerStorer(core *response.Core) Storer {
	storer := &authbossServerStorer{
		core:      core,
		providers: map[string]*config.OAuth2Provider{},
	}

	return storer
}

// findUserByID returns an authboss.User by their internal Response ID.
func (a *authbossServerStorer) findUserByEmail(ctx context.Context, email string) (authboss.User, error) {
	user, err := a.core.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, authboss.ErrUserNotFound
	}

	return userFromEnt(user), nil
}

// findUserByOAuthConnection returns an authboss.User using the providerKey
// and the providerUserID from the OAuth provider.
func (a *authbossServerStorer) findUserByOAuthConnection(ctx context.Context, providerKey string, providerUserID string) (authboss.User, error) {
	oauthConnection, err := a.core.GetOAuthConnectionByProvider(ctx, providerKey, providerUserID)
	if err != nil {
		return nil, fmt.Errorf("core.GetOAuthConnectionByProvider: %w", err)
	}

	if oauthConnection == nil {
		return nil, authboss.ErrUserNotFound
	}

	// Fetch the User's relationship
	user, err := oauthConnection.User(ctx)
	if err != nil {
		return nil, fmt.Errorf("oauthConnection.User: %w", err)
	}

	if user == nil {
		return nil, authboss.ErrUserFound
	}

	return userFromEnt(user), nil
}

// Load will look up the user based on the passed PrimaryID. Under
// normal circumstances this comes from GetPID() of the user.
//
// OAuth2 logins are special-cased to return an OAuth2 pid (combination of
// provider:oauth2uid)
func (a *authbossServerStorer) Load(ctx context.Context, key string) (authboss.User, error) {
	provider, oauthConnectionID, err := authboss.ParseOAuth2PID(key)
	if err != nil {
		// Since we have an error, we can assume this is a standard uid and not a
		// combination of provider and user id according to the Authboss docs.
		return a.findUserByEmail(ctx, key)
	}

	// We have both a provider and an oauthConnectionID which means this is an OAuth2
	// key and should instead get the user from the OAuthConnection itself.
	return a.findUserByOAuthConnection(ctx, provider, oauthConnectionID)
}

// Save persists the user in the database, this should never
// create a user and instead return ErrUserNotFound if the user
// does not exist.
func (a *authbossServerStorer) Save(ctx context.Context, user authboss.User) error {
	panic("implement me")
}

func (a *authbossServerStorer) New(ctx context.Context) authboss.User {
	return &User{}
}

func (a *authbossServerStorer) Create(ctx context.Context, rawUser authboss.User) error {
	user, ok := rawUser.(*User)
	if !ok {
		return errors.New("cannot convert User interface to User type")
	}

	_, err := a.core.CreateUser(ctx, ent.CreateUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: ptr.StringP(user.Password),
	})

	if err != nil {
		return fmt.Errorf("core.CreateUser: %w", err)
	}

	return nil
}

func (a *authbossServerStorer) NewFromOAuth2(ctx context.Context, providerKey string, details map[string]string) (authboss.OAuth2User, error) {
	connectionEntry, err := buildConnectionEntry(providerKey, a.providers[providerKey].UserinfoMapper, &UserinfoMapperData{User: details})
	if err != nil {
		return nil, fmt.Errorf("NewFromOAuth2: %w", err)
	}

	return &OAuth2User{
		Provider:       providerKey,
		ProviderUserID: connectionEntry.ProviderUserID,
		Entry:          &connectionEntry,
	}, nil
}

func (a *authbossServerStorer) SaveOAuth2(ctx context.Context, user authboss.OAuth2User) error {
	// TODO: Store user

	return nil
}

type UserinfoMapperData struct {
	User map[string]string
}

func buildConnectionEntry(providerKey string, mapper *config.UserinfoMapper, data *UserinfoMapperData) (response.OAuthConnectionEntry, error) {
	var providerUserId bytes.Buffer
	if err := mapper.ID.Execute(&providerUserId, data); err != nil {
		return response.OAuthConnectionEntry{}, fmt.Errorf("buildConnectionEntry providerUserId: %w", err)
	}

	var name bytes.Buffer
	if err := mapper.Name.Execute(&name, data); err != nil {
		return response.OAuthConnectionEntry{}, fmt.Errorf("buildConnectionEntry name: %w", err)
	}

	var email bytes.Buffer
	if err := mapper.Email.Execute(&email, data); err != nil {
		return response.OAuthConnectionEntry{}, fmt.Errorf("buildConnectionEntry email: %w", err)
	}

	var avatarURL bytes.Buffer
	if err := mapper.AvatarURL.Execute(&avatarURL, data); err != nil {
		return response.OAuthConnectionEntry{}, fmt.Errorf("buildConnectionEntry avatarURL: %w", err)
	}

	metadata := map[string]string{}
	for prop, t := range mapper.Metadata {
		var value bytes.Buffer
		if err := t.Execute(&value, data); err != nil {
			return response.OAuthConnectionEntry{}, fmt.Errorf("buildConnectionEntry metadata/%s: %w", prop, err)
		}

		metadata[prop] = value.String()
	}

	return response.OAuthConnectionEntry{
		ProviderKey:    providerKey,
		ProviderUserID: providerUserId.String(),
		Name:           name.String(),
		Email:          email.String(),
		AvatarURL:      avatarURL.String(),
		Metadata:       metadata,
	}, nil
}
