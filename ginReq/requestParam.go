package ginReq

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucky-lee/gutil/gStr"
	"net/http"
	"time"
)

//param uint8 value
//Deprecated:use ginReq/post.PostUint8 instead
func ParamUint8(c *gin.Context, key string) uint8 {
	val := c.PostForm(key)
	toVal := gStr.ToUint8(val)

	if toVal == 0 {
		renderAbort(c, key)
	}

	return toVal
}

//param uint8 value and have default value
//Deprecated:use ginReq/post.PostUint8Def instead
func ParamUint8Def(c *gin.Context, key string, defVal uint8) uint8 {
	val := c.PostForm(key)
	toVal := gStr.ToUint8(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param int value
//Deprecated:use ginReq/post.PostInt instead
func ParamInt(c *gin.Context, key string) int {
	val := c.PostForm(key)
	toVal := gStr.ToInt(val)

	if toVal == 0 {
		renderAbort(c, key)
	}
	return toVal
}

//param int value and have default value
//Deprecated:use ginReq/post.PostIntDef instead
func ParamIntDef(c *gin.Context, key string, defVal int) int {
	val := c.PostForm(key)
	toVal := gStr.ToInt(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param int64 value
//Deprecated:use ginReq/post.PostInt64 instead
func ParamInt64(c *gin.Context, key string) int64 {
	val := c.PostForm(key)
	toVal := gStr.ToInt64(val)

	if toVal == 0 {
		renderAbort(c, key)
	}
	return toVal
}

//param int64 value and have default value
//Deprecated:use ginReq/post.PostInt64Def instead
func ParamInt64Def(c *gin.Context, key string, defVal int64) int64 {
	val := c.PostForm(key)
	toVal := gStr.ToInt64(val)

	if toVal == 0 {
		toVal = defVal
	}

	return toVal
}

//param float value
//Deprecated:use ginReq/post.PostFloat64 instead
func ParamFloat64(c *gin.Context, key string) float64 {
	val := c.PostForm(key)

	if val == "" {
		renderAbort(c, key)
		return 0
	}

	toVal := gStr.ToFloat64(val)

	return toVal
}

//param string value
//Deprecated:use ginReq/post.PostStr instead
func ParamStr(c *gin.Context, key string) string {
	val := c.PostForm(key)

	if val == "" {
		renderAbort(c, key)
	}

	return val
}

//param string value and have default value
//Deprecated:use ginReq/post.PostStrDef instead
func ParamStrDef(c *gin.Context, key string, defVal string) string {
	val := c.PostForm(key)

	if val == "" {
		val = defVal
	}

	return val
}

//输出 abort信息
func renderAbort(c *gin.Context, key string) {
	if !c.IsAborted() {
		c.AbortWithStatusJSON(http.StatusOK, ginJson.RetBean(403, "param err", nil))
	}
	fmt.Println("time:", time.Now().Format("2006-01-02 15:04:05"), " param:", key, "no exist")
}
