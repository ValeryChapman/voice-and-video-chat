package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// error structure
type errorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// error response
func newErrorResponse(c *gin.Context, statusCode int, code int, message string) {
	logrus.Error(message)
	c.JSON(statusCode, errorResponse{"ERR", code, message})
}

// status structure
type statusResponse struct {
	Status string `json:"status"`
}

// success response
func newSuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{"OK"})
}
