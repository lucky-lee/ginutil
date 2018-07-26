package ginReq

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//header auth get
func HeaderAuth(c *gin.Context) string {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "err")
		return ""
	}
	return auth
}
