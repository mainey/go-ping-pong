package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPing(t *testing.T) {
	expectedResponse := `{"resp":"pong, Golang!"}`

	r := SetupRouter()
	r.GET("/ping", returnPong)
	req, _ := http.NewRequest("GET", "/ping", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	responseData, _ := ioutil.ReadAll(rec.Body)
	assert.Equal(t, expectedResponse, string(responseData))
	assert.Equal(t, http.StatusOK, rec.Code)
}
