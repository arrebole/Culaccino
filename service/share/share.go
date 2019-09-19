package share

import (
	"crypto/sha256"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// lettrs 用于生成随机字符串的模板
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init(){
	rand.Seed(time.Now().UnixNano())
}

// PassWord 两次hash计算产生密码
func PassWord(passwd string) string {
	return HashSha256(passwd)
}

// RandString 返回MD5编码随机字符串
func RandString() string {
	var bytes = make([]byte, 100)
	for i := range bytes {
		bytes[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return HashMD5(string(bytes))
}

// HashSha256 计算字符串hash,16进制编码
func HashSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// HashMD5 计算数据的hash，使用md5算法
func HashMD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
 }


