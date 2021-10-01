package contx

import "context"

const (
	// ctxUserID provides the key for setting/getting a UserID from the context.
	ctxUserID = ctxKey("UserID")
)

// SetUserID sets a UserID to be passed through the Context.
func SetUserID(ctx context.Context, userID int) context.Context {
	if userID == 0 {
		return ctx
	}

	return context.WithValue(ctx, ctxUserID, userID)
}

// GetUserID returns the UserID from the Context.
func GetUserID(ctx context.Context) int {
	ctxVal := ctx.Value(ctxUserID)

	if ctxVal == nil {
		return 0
	}

	if userID, ok := ctxVal.(int); ok {
		return userID
	}

	return 0
}
