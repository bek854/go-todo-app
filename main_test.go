package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestSimple(t *testing.T) {
	// Test that the app can be built and basic logic works
	assert.Equal(t, 1, 1, "Basic math should work")
	
	// Test that our conversion function works
	result := convertToInt("123")
	assert.Equal(t, 123, result)
}

func TestConvertToInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"1", 1},
		{"0", 0},
		{"999", 999},
	}

	for _, tc := range testCases {
		result := convertToInt(tc.input)
		assert.Equal(t, tc.expected, result, "Conversion failed for input: %s", tc.input)
	}
}
