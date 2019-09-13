package middleware

import (
	"errors"

	"github.com/arrebole/culaccino/service/session"
	"github.com/gin-gonic/gin"
)

// Session ...
func Session(ctx *gin.Context) (session.Body, error) {
	cookie, err := ctx.Cookie("user_session")
	if err != nil {
		return session.Body{}, errors.New("no user_session")
	}

	aSession, err := session.NewStore().Get(cookie)
	if err != nil {
		return session.Body{}, errors.New("not find session")
	}
	return aSession, nil
}
