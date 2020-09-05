package httplib

import (
	"net/http"

	"go.uber.org/zap"
)

func (l *HTTPLib) LoggerWithFields(r *http.Request) *zap.Logger {
	return l.Logger.With(
		zap.String("ip", l.ClientIP(r)),
		zap.String("method", r.Method),
		zap.String("scheme", r.URL.Scheme),
		zap.String("host", r.Host),
		zap.String("path", r.URL.Path),
	)
}
