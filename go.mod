module ginutil

go 1.12

require github.com/lucky-lee/gutil v1.0.2

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/lucky-lee/ginutil v0.0.0-20190411040708-c9a96029b0cb
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net => github.com/golang/net v0.0.0-20190607181551-461777fb6f67
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190610200419-93c9922d18ae
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190610231749-f8d1dee965f7
)
