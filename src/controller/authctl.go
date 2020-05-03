package controller

import (
	"net/http"
	"time"

	"github.com/arrebole/culaccino/src/model"
	"github.com/arrebole/culaccino/src/service"

	"github.com/dgrijalva/jwt-go"
)

// oauth2
// 通过用户名密码获取授权码         api/auth/authorize
// 通过授权码获取token            api/auth/token

// CreateToken 使用用户名和密码的登录接口， 返回访问权限的凭证
func CreateToken(w http.ResponseWriter, r *http.Request) {
	userPayload, err := BodyPaserLogin(r.Body)
	if err != nil || !service.UsersRepo.Check(userPayload.Username, userPayload.Password) {
		ErrorHandle(w, r)
		return
	}

	claim := jwt.MapClaims{
		"username": userPayload.Username,
		"nbf":      time.Now().Unix(), // 生效时间
		"iat":      time.Now().Unix(), // 签发时间
		"exp":      int64(time.Now().Unix() + 3600),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secretKey))
	if err != nil {
		ErrorHandle(w, r)
		return
	}

	response := &model.Token{
		AccessToken: token,
		TokenType:   "bearer",
		ExpiresAt:   int(time.Now().Unix() + 3600),
	}
	JSONResponse(w, response)
}

// CheckToken 检测token的合法性
func CheckToken(w http.ResponseWriter, r *http.Request) {
	if err := Authorization(r); err != nil {
		ErrorHandle(w, r)
		return
	}

	response := &model.Status{
		Code:        0,
		Description: "in effect",
	}

	JSONResponse(w, response)
}
