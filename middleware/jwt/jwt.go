package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

// JWT 设置token
// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		// token 这里要求包含在请求url中, 例如http://127.0.0.1:8000/api/v1/articles?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA5OGY2YmNkNDYyMWQzNzNjYWRlNGU4MzI2MjdiNGY2IiwicGFzc3dvcmQiOiJjYzAzZTc0N2E2YWZiYmNiZjhiZTc2NjhhY2ZlYmVlNSIsImV4cCI6MTY1OTA2MjE1MiwiaXNzIjoiZ2luLWJsb2cifQ.YLR3qFxfNSKyUXR6Y13gyZPzmOmR0kbicpEuQJZGrGE
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS //如果token为空，这里就会返回INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
