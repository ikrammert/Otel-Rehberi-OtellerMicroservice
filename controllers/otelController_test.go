package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"oteller-microservice/controllers"
	"oteller-microservice/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateOtel(t *testing.T) {
	gin.SetMode(gin.TestMode)
	otel := models.Otel{
		Yetkililer:       []models.Yetkili{},
		Firma_unvan:      "",
		Iletisim_bilgisi: []models.IletisimBilgisi{},
	}
	otelJSON, _ := json.Marshal(otel)
	req, _ := http.NewRequest(http.MethodPost, "/create-otel", bytes.NewBuffer(otelJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	controllers.CreateOtel()(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	insertedID, ok := response["InsertedID"]
	assert.True(t, ok)
	assert.NotNil(t, insertedID)
	otelID, ok := response["OtelID"]
	assert.True(t, ok)
	assert.NotNil(t, otelID)
}
