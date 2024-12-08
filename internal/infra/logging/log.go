package logging

import (
	"log/slog"
	"os"
)

func getLevel(key string) slog.Level {
	switch key {
	case "DEBUG":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "Error":
		return slog.LevelError
	default:
		return slog.LevelInfo

	}
}

func NewLogger(groupName string) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: getLevel(os.Getenv("LOG_LEVEL")),
			},
		),
	)
}
