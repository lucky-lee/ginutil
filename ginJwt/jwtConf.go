package ginJwt


var jwtKey string
var jwtExpireCode int

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
