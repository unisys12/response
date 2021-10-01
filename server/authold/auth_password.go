package authold

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-pkgz/auth/provider"
)

func (h *Handler) initPasswordAuthProvider() error {
	if !h.config.AuthPasswordDisabled {
		h.auth.AddDirectProviderWithUserIDFunc("password", provider.CredCheckerFunc(h.handleCheckPassword), provider.UserIDFunc(h.handleUserID))
	}

	return nil
}

func (h *Handler) handleCheckPassword(email, password string) (bool, error) {
	ctx := h.core.ApplyContext(context.Background())
	valid, err := h.core.VerifyPasswordCredentials(ctx, email, password)
	if err != nil {
		return false, err
	}

	return valid, nil
}

func (h *Handler) handleUserID(email string, r *http.Request) string {
	ctx := h.core.ApplyContext(context.Background())
	user, err := h.core.GetUserByEmail(ctx, email)
	if err != nil {
		return ""
	}

	return strconv.Itoa(user.ID)
}
