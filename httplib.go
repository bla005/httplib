package httplib

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type response struct {
	Code int         `json:"-"`
	Data interface{} `json:"data"`
}

func newResponse(code int, data interface{}) *response {
	return &response{
		Code: code,
		Data: data,
	}
}

// HTTPLib ...
type HTTPLib struct {
	Logger *zap.Logger
}

// New creates a new instance
func New(logger *zap.Logger) *HTTPLib {
	return &HTTPLib{
		Logger: logger,
	}
}

// JSON responds to a request in json format
func (l *HTTPLib) JSON(w http.ResponseWriter, code int, data interface{}) {
	resp := newResponse(code, data)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"data":"internal server error"}`))
		// http.Error(w, `{"data":"internal server error"}`, http.StatusInternalServerError)
		return
	}
}

// RemoveCookie removes cookie by name
func (l *HTTPLib) RemoveCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		MaxAge: -1,
	})
}

// NewCookie creates a new cookie
func (l *HTTPLib) NewCookie(w http.ResponseWriter, name, value, path, domain string, expires time.Time, secure, httpOnly bool, sameSite http.SameSite) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		Expires:  expires,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: sameSite,
	})
}

func (l *HTTPLib) ClientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return ""
	}
	return ip.String()
}
