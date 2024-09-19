// Package snorkel provides a [Logger] which logs wide events in JSON format.
package snorkel

import (
	"io"
	"log/slog"
	"math/rand/v2"
	"os"
	"runtime"
	"runtime/debug"
)

type Logger struct {
	discard bool
	groups  []slog.Attr
	random  func() float32
	slogger *slog.Logger
}

type Options struct {
	NoTime       bool
	RandomSource rand.Source
	W            io.Writer
}

// New [Logger] with the given [Options].
func New(opts Options) *Logger {
	if opts.W == io.Discard {
		return NewDiscard()
	}

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

	var groups []slog.Attr

	if bi, ok := debug.ReadBuildInfo(); ok {
		groups = append(groups, slog.Group("build",
			"goVersion", bi.GoVersion,
			"goOS", runtime.GOOS,
			"goArch", runtime.GOARCH,
			"path", bi.Main.Path,
			"version", bi.Main.Version,
			"sum", bi.Main.Sum,
		))
	}

	groups = append(groups, slog.Group("runtime"))

	return &Logger{
		groups:  groups,
		random:  random,
		slogger: slogger,
	}
}

// Event to log with the given name, sample rate, and arguments.
func (l *Logger) Event(name string, rate float32, args ...any) {
	if l.discard {
		return
	}

	if rate <= l.random() {
		return
	}

	var allArgs []any
	allArgs = append(allArgs, "rate", rate)

	allArgs = append(allArgs, args...)

	for _, g := range l.groups {
		allArgs = append(allArgs, g)
	}
	l.slogger.Info(name, allArgs...)
}

func NewDiscard() *Logger {
	return &Logger{discard: true}
}
