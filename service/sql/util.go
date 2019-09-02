package sql

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func defaultPassWord() string {
	passwd := os.Getenv("PASSWORD")
	if passwd == "" {
		passwd = "root@culaccino"
	}
	return password(passwd)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func password(passwd string) string {
	return hash(hash(passwd))
}

func hash(str string) string {
	sha := sha256.New()
	sha.Write([]byte(str))
	return hex.EncodeToString(sha.Sum(nil))
}
