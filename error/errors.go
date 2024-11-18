package error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func BadRequestError (c * gin.Context,  msg  string, details string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": msg, "details" : details})
}

func NotFoundError (c * gin.Context,  msg  string) {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
}

func UnAuthenticatedError(c * gin.Context,  msg  string) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
}
