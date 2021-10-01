package auth

import (
	"strconv"

	"github.com/responserms/response/internal/ptr"

	"github.com/volatiletech/authboss/v3"

	"github.com/responserms/response/internal/ent"
)

var _ authboss.User = (*User)(nil)
var _ authboss.UserValuer = (*User)(nil)
var _ authboss.AuthableUser = (*User)(nil)
var _ authboss.ArbitraryUser = (*User)(nil)
var _ authboss.ArbitraryValuer = (*User)(nil)

type User struct {
	loaded   bool
	ID       string
	Name     string
	Email    string
	Password string
}

func userFromEnt(user *ent.User) *User {
	return &User{
		loaded:   true,
		ID:       strconv.Itoa(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: ptr.String(user.Password),
	}
}

// GetPID returns the User's ID
func (u *User) GetPID() (pid string) {
	return u.ID
}

// PutPID sets the ID on the User object
func (u *User) PutPID(pid string) {
	u.ID = pid
}

func (u *User) Validate() []error {
	return nil
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) PutPassword(password string) {
	u.Password = password
}

func (u *User) GetArbitrary() (arbitrary map[string]string) {
	panic("implement me")
}

func (u *User) PutArbitrary(arbitrary map[string]string) {
	panic("implement me")
}

func (u *User) GetValues() map[string]string {
	panic("implement me")
}
