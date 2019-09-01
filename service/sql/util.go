package sql

import (
	"crypto/sha256"
	"encoding/hex"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func hash(str string) string {
	sha := sha256.New()
	sha.Write([]byte(str))
	return hex.EncodeToString(sha.Sum(nil))
}
