package authold

import "github.com/rs/zerolog"

type authLogger struct {
	logger zerolog.Logger
}

func (a authLogger) Logf(format string, args ...interface{}) {
	a.logger.Log().
		Msgf(format, args...)
}
