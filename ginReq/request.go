package ginReq

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//header auth get
func HeaderAuth(c *gin.Context) string {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		if !c.IsAborted() {
			c.AbortWithStatusJSON(http.StatusForbidden, "err")
		}

		return ""
	}
	return auth
}
