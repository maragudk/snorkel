// Package snorkel provides a [Logger] which logs wide events in JSON format.
package snorkel

import (
	"io"
	"log/slog"
	"math/rand/v2"
	"os"
)

type Logger struct {
	slogger *slog.Logger
	random  func() float32
}

type Options struct {
	NoTime       bool
	RandomSource rand.Source
	W            io.Writer
}

// New [Logger] with the given [Options].
func New(opts Options) *Logger {
	if opts.W == nil {
		opts.W = os.Stderr
	}

	random := rand.Float32
	if opts.RandomSource != nil {
		random = rand.New(opts.RandomSource).Float32
	}

	slogger := slog.New(slog.NewJSONHandler(opts.W, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "level":
				return slog.Attr{}
			case "time":
				if opts.NoTime {
					return slog.Attr{}
				}
				return a
			case "msg":
				a.Key = "name"
				return a
			default:
				return a
			}
		},
	}))

	return &Logger{
		slogger: slogger,
		random:  random,
	}
}

// Log an event with the given name, sample rate, and arguments.
func (l *Logger) Log(name string, rate float32, args ...any) {
	if rate <= l.random() {
		return
	}
	l.slogger.Info(name, args...)
}
