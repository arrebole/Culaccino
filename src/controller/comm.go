package controller

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/arrebole/culaccino/src/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// URLQueryInt 从URL中的query获取数据
// f("/api/article?limit=100", "limit", 0) -> 100
func URLQueryInt(url *url.URL, key string, defaul int) int {
	result, err := strconv.Atoi(url.Query().Get(key))
	if err != nil {
		return defaul
	}
	return result
}

// JSONResponse 序列化实例对象为json bytes
func JSONResponse(w http.ResponseWriter, data interface{}) (n int, err error) {
	result, _ := json.Marshal(data)
	w.Header().Add("Content-type", "application/json")
	return w.Write(result)
}

// BodyPaserArticle ...
func BodyPaserArticle(body io.ReadCloser) (*model.Article, error) {
	var result = &model.Article{}
	decoder := json.NewDecoder(body)
	return result, decoder.Decode(result)
}

// BodyPaserLogin ...
func BodyPaserLogin(body io.ReadCloser) (*model.User, error) {
	var result = &model.User{}
	decoder := json.NewDecoder(body)
	return result, decoder.Decode(result)
}

// ErrorHandle 处理错误的返回
func ErrorHandle(w http.ResponseWriter, r *http.Request) {
	response := model.Status{
		Code:        -1,
		Description: "Fail",
	}
	JSONResponse(w, response)
}

// Params 从url中解析parsms
func Params(ctx context.Context, key string) string {
	return httprouter.ParamsFromContext(ctx).ByName(key)
}

// Authorization 认证登录
func Authorization(r *http.Request) error {
	_, err := parseToken(r.Header.Get("Authorization"))
	return err
}

// ParseToken
func parseToken(authorization string) (jwt.MapClaims, error) {
	if len(strings.Split(authorization, " ")) < 2 {
		return nil, errors.New("Not JWT")
	}

	token, err := jwt.Parse(strings.Split(authorization, " ")[1], secret)

	if err != nil || !token.Valid {
		return nil, errors.New("token is invalid")
	}
	claim, _ := token.Claims.(jwt.MapClaims)
	return claim, nil
}

const secretKey = "shd31db123d231dq"

func secret(token *jwt.Token) (interface{}, error) {
	return []byte(secretKey), nil
}
