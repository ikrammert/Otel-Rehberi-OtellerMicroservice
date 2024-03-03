package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"oteller-microservice/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateCommInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	controllers.CreateCommInfo()(c)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	return
}
