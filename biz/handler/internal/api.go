package internal

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

var (
	ErrorSubIsNil       = errors.New("sub is nil")
	ErrorInvalidateType = errors.New("invalidate type")
	ErrorTooManyReq     = errors.New("too many req")
)

type Api struct {
	Ctx   *app.RequestContext
	Sub   *int64
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

func (a Api) SetReqWithSub(c *app.RequestContext, obj ...any) Api {
	a.Ctx = c
	if len(obj) > 1 {
		return a.setError(ErrorTooManyReq)
	}
	if len(obj) == 0 {
		return a
	}

	sub, ok := c.Get("sub")
	if !ok {
		return a.setError(ErrorSubIsNil)
	}
	int64Sub := new(int64)
	*int64Sub, ok = sub.(int64)
	if !ok {
		return a.setError(ErrorInvalidateType)
	}
	a.Sub = int64Sub
	if err := c.BindAndValidate(obj); err != nil {
		return a.setError(errors.WithMessage(err, "bind struct err"))
	}

	return a
}

func (a Api) setError(err error) Api {
	a.Error = err
	return a
}
