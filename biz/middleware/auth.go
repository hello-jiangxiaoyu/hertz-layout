package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"hertz/demo/internal/response"
	"hertz/demo/internal/utils"
)

// AuthHandler cookie鉴权
func AuthHandler(_ context.Context, c *app.RequestContext) {
	idToken := c.Cookie("hertz-layout")
	if len(idToken) == 0 {
		response.ErrorNoLogin(c)
		return
	}

	key, err := utils.GetJWKPrivate()
	if err != nil {
		response.ErrorUnknown(c, err, "GetJWKPrivate err")
		return
	}
	claim := jwt.New(jwt.SigningMethodRS256)
	token, err := jwt.ParseWithClaims(string(idToken), claim.Claims, func(token *jwt.Token) (interface{}, error) {
		return key.Public(), nil
	})

	if err != nil {
		response.ErrorUnknown(c, err, "parse claim err")
		return
	}
	if !token.Valid {
		response.ErrorForbidden(c, "invalid token")
		return
	}
	sub, err := claim.Claims.GetSubject()
	if err != nil {
		response.ErrorUnknown(c, err, "get sub err")
		return
	}
	c.Set("sub", sub)
}
