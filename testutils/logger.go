package testutils

import (
	"log/slog"
	"os"

	"github.com/mpjunior92/eigensdk-go/logging"
)

func GetTestLogger() logging.Logger {
	return logging.NewTextSLogger(os.Stdout,
		&logging.SLoggerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	)
}
