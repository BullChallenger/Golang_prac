package myapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(response, request)

	assert.Equal(http.StatusOK, response.Code)

	data, _ := io.ReadAll(response.Body)
	assert.Equal("Hello World!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/bar?name=Go", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(response, request)

	assert.Equal(http.StatusOK, response.Code)

	data, _ := io.ReadAll(response.Body)
	assert.Equal("Hello Go!", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(response, request)

	assert.Equal(http.StatusBadRequest, response.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/foo",
		strings.NewReader(`{
							   "first_name":"seungbin", 
							   "last_name":"cho", 
							   "email":"test@test.com"
								}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(response, request)

	assert.Equal(http.StatusOK, response.Code)

	user := new(User)
	err := json.NewDecoder(response.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("seungbin", user.FirstName)
	assert.Equal("cho", user.LastName)
}
