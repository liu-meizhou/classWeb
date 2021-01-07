package utils

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type TokenConfig struct {
	TokenSecrets []byte
	TokenExp     int // userType:  1.Admin, 2.学生, 3.老师, 4.系主任
}

var tokenConfig = &TokenConfig{}

type TokenInfo struct {
	LoginId  string `form:"loginId"`
	Password string `form:"password"`
	UserType int    `form:"userType"`
}

func NewTokenInfo(id, password string, userType int) *TokenInfo {
	return &TokenInfo{
		id,
		password,
		userType,
	}
}

func init() {
	tokenExp, err := web.AppConfig.Int("TokenExp")
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}
	tokenConfig.TokenExp = tokenExp
	str, err := web.AppConfig.String("TokenSecrets")
	if err != nil {
		logs.Error(err)
		os.Exit(-1)
	}
	tokenConfig.TokenSecrets = []byte(str)
}

func CheckTokenInfo(tokenInfo *TokenInfo) error {
	if tokenInfo.LoginId == "" || tokenInfo.Password == "" {
		return fmt.Errorf("账号或密码不可为空")
	}
	return nil
}

// CreateToken 根据登录用户和登录类型生成id
// userType:  1.Admin, 2.学生, 3.老师, 4.系主任
func CreateToken(tokenInfo *TokenInfo) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenConfig.TokenExp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = tokenInfo.LoginId
	claims["password"] = tokenInfo.Password
	claims["userType"] = tokenInfo.UserType
	token.Claims = claims
	tokenString, _ := token.SignedString(tokenConfig.TokenSecrets)
	return tokenString
}

// ParseToken 根据登录用户和登录类型生成id
// userType:  1.Admin, 2.学生, 3.老师, 4.系主任
func ParseToken(tokenString string) (*TokenInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenConfig.TokenSecrets, nil
	})
	if token != nil {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			return NewTokenInfo(claims["id"].(string), claims["password"].(string), claims["userType"].(int)), nil
		}
	}
	logs.Error(err)
	return nil, err
}
