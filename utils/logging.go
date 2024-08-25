package utils

import (
	"log/slog"
	"os"
)

// Initializes a new josn logger and sets it as the default logger
func GetInitLogger() *slog.Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	return l
}
