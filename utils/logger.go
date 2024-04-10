package utils

import (
	"log/slog"
	"os"
)

// Logger Level customize
type Leveler interface {
	Level() slog.Level
}

func Level(l string) *slog.HandlerOptions {
	var LogLevelMap = map[string]slog.Level{
		"TRACE":  slog.Level(-8),
		"NOTICE": slog.Level(2),
		"FETAL":  slog.Level(12),
	}

	var lS, dExist = LogLevelMap[l]
	if !dExist {
		panic("level not found")
	}

	opts := &slog.HandlerOptions{
		Level: lS,
	}

	return opts
}

type ConsoleLogger struct {
	Db *slog.Logger
}

func NewConsoleLogger(opts ...*slog.HandlerOptions) *ConsoleLogger {

	var opt *slog.HandlerOptions

	if len(opts) > 0 {
		opt = opts[0]
	}

	return &ConsoleLogger{
		Db: slog.New(slog.NewJSONHandler(os.Stdout, opt)),
	}
}
