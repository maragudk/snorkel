// Package snorkel provides a [Logger] which logs wide events in JSON format.
package snorkel

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

type Options struct {
	NoTime bool
	W      io.Writer
}

// New [Logger] with the given [Options].
func New(opts Options) *Logger {
	if opts.W == nil {
		opts.W = os.Stderr
	}

	return &Logger{logger: slog.New(slog.NewJSONHandler(opts.W, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "level":
				return slog.Attr{}
			case "time":
				if opts.NoTime {
					return slog.Attr{}
				}
				return a
			default:
				return a
			}
		},
	}))}
}

func (l *Logger) Log(msg string, args ...any) {
	l.logger.Info(msg, args...)
}
