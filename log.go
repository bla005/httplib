package httplib

import (
	"net/http"

	"go.uber.org/zap"
)

type event struct {
	id      int
	message string
}

type handlerLogger struct {
	*zap.Logger
}

var (
	badInput       = &event{1, "bad input"}
	badCredentials = &event{2, "bad credentials"}
	funcErr        = &event{3, "function error"}
)

func (l *HTTPLib) NewHandlerLogger(r *http.Request) *handlerLogger {
	return &handlerLogger{
		l.Logger.With(
			zap.String("ip", l.ClientIP(r)),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("user_agent", r.UserAgent()),
		),
	}
}
func (l *handlerLogger) BadInput() {
	l.Info(badInput.message)
}
func (l *handlerLogger) BadCredentials() {
	l.Info(badCredentials.message)
}
func (l *handlerLogger) FuncErr(err error) {
	l.Error(funcErr.message, zap.Error(err))
}
