package ginReq

import (
	"github.com/gin-gonic/gin"
)

//header auth get
func HeaderAuth(c *gin.Context) string {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return ""
	}
	return auth
}
