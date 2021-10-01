package response

//var (
//	// ErrPasswordAuthNotAllowed is returned when attempting to verify credentials for a User who does not
//	// have a password set. This is commonly returned when a User registers with OAuth 2.0 and never sets
//	// a password.
//	ErrPasswordAuthNotAllowed = errors.New("password auth not allowed for this user")
//
//	// ErrPasswordAuthDisabled is returned when password-based authentication for users has been disabled
//	// by a Response administrator.
//	ErrPasswordAuthDisabled = errors.New("password auth is disabled")
//)

//func (c *Core) RegisterWithPassword(ctx context.Context, name, email, password string) (*ent.User, error) {
//	if c.config.AuthPasswordDisabled {
//		return nil, ErrPasswordAuthDisabled
//	}
//
//	hashedPassword, err := c.hashPassword(password)
//	if err != nil {
//		return nil, err
//	}
//
//	user, err := c.CreateUser(ctx, ent.CreateUser{
//		Name:     name,
//		Email:    email,
//		Password: ptr.StringP(hashedPassword),
//	})
//
//	if err != nil {
//		c.logger.Error().
//			Err(err).
//			Str("email", email).
//			Msg("Registration failed during RegisterWithPassword.")
//
//		return nil, fmt.Errorf("CreateUser: %w", err)
//	}
//
//	return user, nil
//}
