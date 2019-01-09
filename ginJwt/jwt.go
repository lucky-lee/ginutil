package ginJwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lucky-lee/ginutil/ginJson"
	"github.com/lucky-lee/ginutil/ginReq"
	"github.com/lucky-lee/gutil/gLog"
	"github.com/lucky-lee/gutil/gStr"
	"net/http"
	"time"
)

//parse user id
func ParseUid(str string) int64 {
	t, err := jwt.ParseWithClaims(str, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetKey()), nil
	})

	if err != nil {
		gLog.E("tokenParseErr", err)
		return 0
	}

	if claims, ok := t.Claims.(*jwt.StandardClaims); ok && t.Valid {
		if claims.ExpiresAt > time.Now().Unix() {
			return gStr.ToInt64(claims.Subject)
		} else {
			return -1 //old
		}
	} else {
		gLog.E("tokenClaimsErr", "")
		return 0
	}
}

//get user id by token
func Uid(c *gin.Context) int64 {
	token := ginReq.HeaderAuth(c)

	if token == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "error")
		return 0
	}

	uid := ParseUid(token)

	if uid <= 0 {
		c.AbortWithStatusJSON(http.StatusOK, ginJson.RetBean(GetExpireCode(), "Login timeout, please log in again!", nil))
	}

	return uid
}

//get user id by token no check
func UidNoCheck(c *gin.Context) int64 {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		return 0
	}

	return ParseUid(token)
}
