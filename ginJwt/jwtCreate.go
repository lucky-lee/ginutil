package ginJwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lucky-lee/gutil/gLog"
	"time"
)

//生成token
func CreateToken(subject int64) string {
	nowTs := time.Now().Unix()
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", subject),
		Issuer:    GetIssuer(),
		IssuedAt:  nowTs,
		NotBefore: nowTs,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(GetKey()))

	if err != nil {
		gLog.E("jwtTokenErr", err)
		gLog.E("jwtTokenObj", claims)
		return ""
	}

	return tokenStr
}
