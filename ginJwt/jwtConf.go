package ginJwt

var jwtIssuer string
var jwtKey string
var jwtExpireCode int

func SetIssuer(s string) {
	jwtIssuer = s
}

func GetIssuer() string {
	return jwtIssuer
}

func SetKey(str string) {
	jwtKey = str
}

func GetKey() string {
	return jwtKey
}

func SetExpireCode(code int) {
	jwtExpireCode = code
}

func GetExpireCode() int {
	return jwtExpireCode
}
