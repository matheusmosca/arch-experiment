package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/matheusmosca/arch-experiment/extensions/xlog"
	"go.uber.org/zap"
)

// type Response struct {
// 	Status int
// 	header http.Header
// 	Error  error
// 	Body   any
// }

// func Handle(handler func(r *http.Request) Response) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		response := handler(r)
// 		if response.Error != nil {
// 			// slog.From
// 		}
// 	}
// }

func LoggerToContext(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(xlog.WithLogger(r.Context(), logger))

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

var _excludeURLs = []string{"/healtcheck", "/health", "/metrics", "/swagger"}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, d := range _excludeURLs {
			if strings.Contains(r.URL.Path, d) {
				next.ServeHTTP(w, r)
				return
			}
		}

		t := time.Now()
		rw := &responseWriter{ResponseWriter: w}

		next.ServeHTTP(rw, r)

		logger := xlog.FromContext(r.Context()).With(
			zap.Int("status", rw.statusCode),
			zap.String("method", r.Method),
			zap.Duration("duration", time.Duration(time.Since(t))),
			zap.String("user_agent", r.UserAgent()),
			zap.String("path", r.URL.Path),
		)

		if rw.statusCode >= 400 {
			logger.Error("failed to process request")
			return
		}

		logger.Info("request succedeed")
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
