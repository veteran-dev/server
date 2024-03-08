package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/veteran-dev/server/global"
)

type JwtCustomClaims struct {
	ID int
	jwt.RegisteredClaims
}

type UserJWT struct {
	SigningKey []byte
}

func (u *UserJWT) GenerateToken(id int) (string, error) {
	// 初始化
	ep, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	iJwtCustomClaims := JwtCustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,              // 签名的发行者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(u.SigningKey)
}

// ParseToken 解析token
func (u *UserJWT) ParseToken(tokenStr string) (JwtCustomClaims, error) {
	// 声明一个空的数据声明
	iJwtCustomClaims := JwtCustomClaims{}
	//ParseWithClaims是NewParser().ParseWithClaims()的快捷方式
	//第一个值是token ，
	//第二个值是我们之后需要把解析的数据放入的地方，
	//第三个值是Keyfunc将被Parse方法用作回调函数，以提供用于验证的键。函数接收已解析但未验证的令牌。
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return u.SigningKey, nil
	})

	// 判断 是否为空 或者是否无效只要两边有一处是错误 就返回无效token
	if err != nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return iJwtCustomClaims, err
}
func (u *UserJWT) IsTokenValid(tokenStr string) bool {
	_, err := u.ParseToken(tokenStr)
	return err == nil
}
func NewUserJWT() *UserJWT {
	return &UserJWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}
