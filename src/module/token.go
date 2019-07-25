package module

import (
	"math/rand"
	"time"
)

var local = newSection()

// Section 储存会话
type Section struct {
	Token  string
	Create time.Time
	MaxAge time.Duration
}

func (p Section) isOverTime() bool {
	return (time.Now().Sub(p.Create) > p.MaxAge)
}

func randSeq(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func newSection() Section {
	return Section{
		Token:  randSeq(16),
		Create: time.Now(),
		MaxAge: time.Minute * 20,
	}
}

// NewToken 根据是否超时创建 token
func NewToken() string {
	if local.isOverTime() {
		local = newSection()
	}
	return local.Token
}
