package slogecs

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

func NewHandler(level slog.Level) slog.Handler {
	base := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				a.Key = "@timestamp"
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.UTC().Format("2006-01-02T15:04:05.000Z"))
				}
			case slog.LevelKey:
				if level, ok := a.Value.Any().(slog.Level); ok {
					return slog.Group("log", slog.String("level", strings.ToLower(level.String())))
				}
			case slog.MessageKey:
				a.Key = "message"
			}
			return a
		},
	})

	return slog.New(base).With(slog.Group("ecs", slog.String("version", "1.6.0"))).Handler()
}
