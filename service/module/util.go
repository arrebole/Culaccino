package module

import (
	"math/rand"
	"time"
)

// ChangeToken ...
func randString(length int) string {
	rand.Seed(time.Now().Unix())
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randRune := make([]rune, length)
	for i := range randRune {
		randRune[i] = letters[rand.Intn(len(letters))]
	}
	return string(randRune)
}
