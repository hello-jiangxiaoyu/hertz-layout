// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	oauth "hertz/demo/biz/router/oauth"
	user "hertz/demo/biz/router/user"
	ws "hertz/demo/biz/router/ws"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	ws.Register(r)

	oauth.Register(r)

	user.Register(r)
}
