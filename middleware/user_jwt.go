package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/veteran-dev/server/model/common/response"
	"github.com/veteran-dev/server/utils"
)

func UserJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		token = token[len("Bearer "):]
		j := utils.NewUserJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("UserID", claims.ID)
		c.Next()

		if claims.ID <= 0 {
			response.FailWithDetailed(gin.H{"reload": true}, "非法访问", c)
			c.Abort()
			return
		}
	}
}
