package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Envelope{STATUS_SUCCESS, getUserFromContext(c)})
}
