package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	g := gin.Default()
	middL := InitMiddleware()
	g.Use(middL.CORS())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/projects", nil)
	g.ServeHTTP(w, req)
	assert.Contains(t, w.HeaderMap["Access-Control-Allow-Origin"], "*")
	assert.Contains(t, w.HeaderMap["Access-Control-Allow-Credentials"], "true")
	assert.Contains(t, w.HeaderMap["Access-Control-Allow-Methods"], "GET")
}
