package slogecs

import (
	"log/slog"
	"net/http"
	"time"
)

func HTTPLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		slog.Info("HTTP request",
			slog.Group("http",
				slog.Group("request", slog.String("method", r.Method)),
				slog.Group("response", slog.Int("status_code", wrapped.statusCode)),
			),
			slog.Group("url", slog.String("path", r.URL.Path)),
			slog.Group("event", slog.Int64("duration", duration.Nanoseconds())),
			slog.Group("client", slog.String("address", r.RemoteAddr)),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
