package share

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// PassWord 两次hash计算产生密码
func PassWord(passwd string) string {
	return HashHexStr(HashHexStr(passwd))
}

// HashHexStr 计算字符串hash,16进制编码
func HashHexStr(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// RandString 返回base64编码随机字符串
func RandString() string {
	cache := make([]byte, 100)
	rand.Seed(time.Now().Unix())
	for i := range cache {
		cache[i] = letters[rand.Intn(len(letters))]
	}
	return HashHexStr(string(cache))
}
