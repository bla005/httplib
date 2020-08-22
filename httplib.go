package httplib

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type httpResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (r *httpResponse) String() string {
	res, err := json.Marshal(r)
	if err != nil {
		return http.StatusText(http.StatusInternalServerError)
	}
	return string(res)
}

func newHttpResponse(code int, data interface{}) *httpResponse {
	return &httpResponse{
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
	httpResp1 := newHttpResponse(code, data)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(httpResp1); err != nil {
		httpResp1.Code = http.StatusInternalServerError
		httpResp1.Data = http.StatusText(http.StatusInternalServerError)
		http.Error(w, httpResp1.String(), http.StatusInternalServerError)
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
