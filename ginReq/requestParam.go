package ginReq

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky-lee/gutil/gStr"
	"net/http"
)

//param uint8 value
func ParamUint8(c *gin.Context, key string) uint8 {
	val := c.PostForm(key)
	toVal := gStr.ToUint8(val)

	if toVal == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, "param err")
	}

	return toVal
}

//param uint8 value and have default value
func ParamUint8Def(c *gin.Context, key string, defVal uint8) uint8 {
	val := c.PostForm(key)
	toVal := gStr.ToUint8(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param int value
func ParamInt(c *gin.Context, key string) int {
	val := c.PostForm(key)
	toVal := gStr.ToInt(val)

	if toVal == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, "param err")
	}
	return toVal
}

//param int value and have default value
func ParamIntDef(c *gin.Context, key string, defVal int) int {
	val := c.PostForm(key)
	toVal := gStr.ToInt(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param int64 value
func ParamInt64(c *gin.Context, key string) int64 {
	val := c.PostForm(key)
	toVal := gStr.ToInt64(val)

	if toVal == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, "param err")
	}
	return toVal
}

//param int64 value and have default value
func ParamInt64Def(c *gin.Context, key string, defVal int64) int64 {
	val := c.PostForm(key)
	toVal := gStr.ToInt64(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param float value
func ParamFloat64(c *gin.Context, key string) float64 {
	val := c.PostForm(key)

	if val == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "param err")
		return 0
	}

	toVal := gStr.ToFloat64(val)

	return toVal
}

//param string value
func ParamStr(c *gin.Context, key string) string {
	val := c.PostForm(key)

	if val == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "param err")
	}

	return val
}

//param string value and have default value
func ParamStrDef(c *gin.Context, key string, defVal string) string {
	val := c.PostForm(key)

	if val == "" {
		val = defVal
	}

	return val
}
