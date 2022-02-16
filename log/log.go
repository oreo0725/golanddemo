package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(env string) {
	log.Logger = log.With().Caller().Logger()

	switch env {
	case "local", "dev", "stag":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "prod":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func Disable() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func Debug(topic string) *zerolog.Event {
	return log.Debug().Str("topic", topic)
}

func Info(topic string) *zerolog.Event {
	return log.Info().Str("topic", topic)
}

func Warn(topic string) *zerolog.Event {
	return log.Warn().Str("topic", topic)
}

func Error(topic string) *zerolog.Event {
	return log.Error().Str("topic", topic)
}

func Fatal(topic string) *zerolog.Event {
	return log.Fatal().Str("topic", topic)
}
