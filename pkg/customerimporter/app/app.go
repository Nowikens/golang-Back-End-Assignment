package app

import (
	"log/slog"
)

// App implements simple repository pattern, for injecting logs
type App struct {
	Logger *slog.Logger
}
