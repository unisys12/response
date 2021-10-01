package contx

import (
	"context"

	"github.com/labstack/echo/v4"
)

var (
	ctxEchoContext = ctxKey("echoContext")
)

func SetEchoContext(ctx context.Context, c echo.Context) context.Context {
	return context.WithValue(ctx, ctxEchoContext, c)
}

func GetEchoContext(ctx context.Context) echo.Context {
	if val, ok := ctx.Value(ctxEchoContext).(echo.Context); ok {
		return val
	}

	return nil
}
