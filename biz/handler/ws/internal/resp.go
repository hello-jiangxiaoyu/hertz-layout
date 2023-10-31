package internal

import (
	"github.com/bytedance/sonic"
	"github.com/hertz-contrib/websocket"
	"github.com/pkg/errors"
	"hertz/demo/pkg/response"
)

const (
	codeErrorSend = iota + 6000
	codeErrorKnown
)

func textResponse(conn *websocket.Conn, code int, errMsg string) error {
	body := response.Response{
		Code: code,
		Msg:  errMsg,
	}
	marshal, err := sonic.Marshal(&body)
	if err != nil {
		return errors.WithMessage(err, body.Msg+": marshal response failed")
	}

	if err = conn.WriteMessage(websocket.TextMessage, marshal); err != nil {
		return errors.WithMessage(err, body.Msg+": conn.WriteMessage err")
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

func PongResponse(conn *websocket.Conn) error {
	return conn.WriteMessage(websocket.PongMessage, nil)
}

func SuccessResponse(conn *websocket.Conn) error {
	return textResponse(conn, response.CodeSuccess, "")
}

func ErrorUnknown(conn *websocket.Conn, err error, msg string) error {
	msg = msg + ": " + err.Error()
	return textResponse(conn, codeErrorKnown, msg)
}

func ErrorMsg(conn *websocket.Conn, msg string) error {
	return textResponse(conn, codeErrorKnown, msg)
}
