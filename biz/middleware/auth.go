package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"hertz/demo/pkg/conf"
	"hertz/demo/pkg/response"
	"hertz/demo/pkg/utils"
	"strconv"
)

// AuthHandler cookie鉴权
func AuthHandler(_ context.Context, c *app.RequestContext) {
	idToken := c.Cookie(conf.DefaultCookieKey)
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

	// set user sub
	sub, err := claim.Claims.GetSubject()
	if err != nil {
		response.ErrorUnknown(c, err, "get sub err")
		return
	}
	int64Sub, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		response.ErrorUnknown(c, err, "sub to integer err")
		return
	}
	c.Set("sub", int64Sub)
}
