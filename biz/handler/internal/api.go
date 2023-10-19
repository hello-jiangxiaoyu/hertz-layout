package internal

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

var (
	ErrorContext    = errors.New("context is nil")
	ErrorTooManyReq = errors.New("too many req")
)

type Api struct {
	Ctx   *app.RequestContext
	Sub   int64
	Error error
}

func (a Api) SetReq(c *app.RequestContext, obj ...any) Api {
	a.Ctx = c
	if len(obj) > 1 {
		return a.setError(ErrorTooManyReq)
	}
	if len(obj) == 0 {
		return a
	}
	if err := c.BindAndValidate(obj); err != nil {
		return a.setError(errors.WithMessage(err, "bind struct err"))
	}

	return a
}

func (a Api) setError(err error) Api {
	a.Error = err
	return a
}
