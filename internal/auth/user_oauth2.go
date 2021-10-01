package auth

import (
	"time"

	"github.com/responserms/response/response"

	"github.com/responserms/response/internal/ent"

	"github.com/volatiletech/authboss/v3"
)

var _ authboss.OAuth2User = (*OAuth2User)(nil)

type OAuth2User struct {
	conn              *ent.OAuthConnection
	user              *ent.User
	Entry             *response.OAuthConnectionEntry
	Provider          string
	ProviderUserID    string
	UserID            string
	OAuthAccessToken  string
	OAuthRefreshToken string
	OAuthExpiry       time.Time
}

func (u *OAuth2User) GetPID() (pid string) {
	return authboss.MakeOAuth2PID(u.Provider, u.ProviderUserID)
}

func (u *OAuth2User) PutPID(pid string) {
	provider, providerUserID := authboss.ParseOAuth2PIDP(pid)

	u.Provider = provider
	u.ProviderUserID = providerUserID
}

func (u *OAuth2User) IsOAuth2User() bool {
	return u.conn != nil
}

func (u *OAuth2User) GetOAuth2UID() (uid string) {
	return u.UserID
}

func (u *OAuth2User) GetOAuth2Provider() (provider string) {
	return u.Provider
}

func (u *OAuth2User) GetOAuth2AccessToken() (token string) {
	return u.OAuthAccessToken
}

func (u *OAuth2User) GetOAuth2RefreshToken() (refreshToken string) {
	return u.OAuthRefreshToken
}

func (u *OAuth2User) GetOAuth2Expiry() (expiry time.Time) {
	return u.OAuthExpiry
}

func (u *OAuth2User) PutOAuth2UID(uid string) {
	u.UserID = uid
}

func (u *OAuth2User) PutOAuth2Provider(provider string) {
	u.Provider = provider
}

func (u *OAuth2User) PutOAuth2AccessToken(token string) {
	u.OAuthAccessToken = token
}

func (u *OAuth2User) PutOAuth2RefreshToken(refreshToken string) {
	u.OAuthRefreshToken = refreshToken
}

func (u *OAuth2User) PutOAuth2Expiry(expiry time.Time) {
	u.OAuthExpiry = expiry
}
