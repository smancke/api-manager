package access

import (
	"github.com/golang/mock/gomock"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	. "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CaddyHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	next := httpserver.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
		w.Write([]byte("next was called"))
		return 0, nil
	})

	r, _ := http.NewRequest("GET", "https://api.example.com/the_api_key/foo/bar", nil)
	r.Header.Set("Authorization", "Bearer secret")
	r.Header.Set("Referer", "http://referer.com")
	r.RemoteAddr = "127.0.0.1:9999"

	// success case
	storeMock := NewMockUsageStore(ctrl)
	storeMock.EXPECT().
		GetLimit("the_api_key", "secret", "127.0.0.1", "http://referer.com").
		Return(100, 200, uint64(424242), nil)
		//	storeMock.EXPECT().
		//		Log("the_api_key", "secret", "127.0.0.1", "http://referer.com")

	handler := &CaddyHandler{next: next, store: storeMock}
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, r)

	Equal(t, "next was called", recorder.Body.String())

	// limit exeeded case
	storeMock = NewMockUsageStore(ctrl)
	storeMock.EXPECT().
		GetLimit("the_api_key", "secret", "127.0.0.1", "http://referer.com").
		Return(0, 200, uint64(424242), nil)
		//	storeMock.EXPECT().
		//		Log("the_api_key", "secret", "127.0.0.1", "http://referer.com")

	handler = &CaddyHandler{next: next, store: storeMock}
	recorder = httptest.NewRecorder()

	handler.ServeHTTP(recorder, r)
	Equal(t, 429, recorder.Code)
	Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	Contains(t, recorder.Body.String(), "Api limit exceeded")
}

func Test_CaddyHandler_credentialsFromRequest(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://api.example.com/the_api_key/foo/bar", nil)

	handler := &CaddyHandler{}

	key, secret := handler.credentialsFromRequest(r)
	Equal(t, "the_api_key", key)
	Equal(t, "", secret)

	r.Header.Set("Authorization", "Bearer secret")
	key, secret = handler.credentialsFromRequest(r)
	Equal(t, "the_api_key", key)
	Equal(t, "secret", secret)
}

func Test_CaddyHandler_ipFromRequest(t *testing.T) {
	r := &http.Request{
		RemoteAddr: "127.0.0.1:9999",
		Header: http.Header{
			"X-Forwarded-For": {"127.0.0.2, 127.0.0.3"},
		},
	}

	handler := &CaddyHandler{}
	Equal(t, "127.0.0.1", handler.ipFromRequest(r))

	preferHeaderHandler := &CaddyHandler{preferIpFromHeader: true}
	Equal(t, "127.0.0.2", preferHeaderHandler.ipFromRequest(r))

	r.Header.Set("X-Real-Ip", "127.0.0.4")
	Equal(t, "127.0.0.4", preferHeaderHandler.ipFromRequest(r))
}
