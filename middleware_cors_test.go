package gotuna_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gotuna/gotuna"
	"github.com/gotuna/gotuna/test/assert"
)

func TestCORS(t *testing.T) {
	request := httptest.NewRequest(http.MethodOptions, "/sample", nil)
	response := httptest.NewRecorder()

	middleware := gotuna.App{}.Cors()

	middleware(http.NotFoundHandler()).ServeHTTP(response, request)

	assert.Equal(t, "*", response.HeaderMap.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "OPTIONS,HEAD,GET,POST,PUT,PATCH,DELETE", response.HeaderMap.Get("Access-Control-Allow-Methods"))
}
