package utils

import "github.com/gin-gonic/gin"

func ResponseSuccess(c *gin.Context, code int, msg string, err string, data interface{}) {
	if msg == "" {
		msg = "Successfully."
	}

	c.AbortWithStatusJSON(code, gin.H{"success": true, "code": code, "error": err, "msg": msg, "data": data})
}

func ResponseError(c *gin.Context, code int, msg string, err string, data interface{}) {
	if msg == "" {
		msg = "Something went wrong!"
	}

	c.AbortWithStatusJSON(code, gin.H{"success": false, "code": code, "error": err, "msg": msg, "data": data})
}
