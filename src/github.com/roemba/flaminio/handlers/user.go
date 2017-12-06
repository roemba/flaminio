package handlers

import "github.com/gin-gonic/gin"

func UserHandler(c *gin.Context) {
	user := getUserFromContext(c)

	jsonResponse(Response{STATUS_SUCCESS,user}, c.Writer)
}
