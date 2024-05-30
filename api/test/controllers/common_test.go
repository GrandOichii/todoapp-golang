package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func createTestContext(object interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	data, _ := json.Marshal(object)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	return c, w
}
