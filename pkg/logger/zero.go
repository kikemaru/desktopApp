package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	log zerolog.Logger
}

func NewZerologLogger() (*ZerologLogger, error) {
	zerolog.TimeFieldFormat = time.RFC3339

	core := zerolog.NewConsoleWriter()
	core.Out = os.Stdout
	core.TimeFormat = "Jan 2 15:04:05"

	logger := zerolog.New(core).With().Logger()

	return &ZerologLogger{log: logger}, nil
}

func (t *ZerologLogger) Debug(msgs ...interface{}) {
	t.log.Debug().Msgf("%v", msgs)
}

func (t *ZerologLogger) Info(msgs ...interface{}) {
	t.log.Info().Msgf("%v", msgs)
}

func (t *ZerologLogger) Warn(msgs ...interface{}) {
	t.log.Warn().Msgf("%v", msgs)
}

func (t *ZerologLogger) Error(msgs ...interface{}) {
	t.log.Error().Msgf("%v", msgs)
}

func (t *ZerologLogger) Fatal(msgs ...interface{}) {
	t.log.Fatal().Msgf("%v", msgs)
}

func (t *ZerologLogger) GetSugaredZapInstance() *zerolog.Logger {
	return &t.log
}
