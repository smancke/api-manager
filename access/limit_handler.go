package access

import (
	"fmt"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"github.com/tarent/lib-compose/logging"
	"net"
	"net/http"
	"strings"
)

type CaddyHandler struct {
	next               httpserver.Handler
	store              UsageStore
	preferIpFromHeader bool
}

func (h *CaddyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	key, secret := h.credentialsFromRequest(r)
	ip := h.ipFromRequest(r)
	referer := r.Referer()

	remaining, limit, reset, err := h.store.GetLimit(key, secret, ip, referer)
	if err != nil {
		logging.Application(r.Header).WithError(err).WithField("api_key", key).Errorf("error fetching limits")
	}

	// allow acces in the case, that we could not get the limit
	if err != nil || remaining > 0 {
		return h.next.ServeHTTP(w, r)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTooManyRequests) // 429
	fmt.Fprintf(w, `{
  "result": "error",
  "resultDetail": "clientError",
  "message": "Api limit exceeded."
  "limit": %v,
  "remaining": %v,
  "reset": %v
}`, limit, remaining, reset)
	return 0, nil
}

func (h *CaddyHandler) credentialsFromRequest(r *http.Request) (key, secret string) {
	if r.Header.Get("Authorization") != "" {
		auth := strings.Split(r.Header.Get("Authorization"), " ")
		if auth[0] == "Bearer" && len(auth) == 2 {
			secret = auth[1]
		}
	}
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) > 1 {
		key = pathSegments[1]
	}

	return key, secret
}

func (h *CaddyHandler) ipFromRequest(r *http.Request) string {
	if h.preferIpFromHeader {
		if r.Header.Get("X-Real-Ip") != "" {
			return r.Header.Get("X-Real-Ip")
		}
		if r.Header.Get("X-Forwarded-For") != "" {
			return strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]
		}
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}
