package utils

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/client/retry"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
	"time"
)

var c *client.Client

func init() {
	var err error
	stateFunc := func(state config.HostClientState) {
		fmt.Printf("state=%v\n", state.ConnPoolState().Addr)
	}
	var customDialFunc network.DialFunc = func(addr string) (network.Conn, error) {
		return nil, nil
	}

	c, err = client.NewClient(client.WithDialTimeout(1*time.Second),
		client.WithMaxConnsPerHost(512),                // 最大连接数
		client.WithMaxIdleConnDuration(10*time.Second), // 最大空闲连接时间
		client.WithMaxConnDuration(10*time.Second),     // keep-alive 连接保持时长
		client.WithKeepAlive(true),                     // 是否启用keep-alive
		client.WithClientReadTimeout(1*time.Second),    // 响应读取时间
		client.WithDialer(netpoll.NewDialer()),
		client.WithDialFunc(customDialFunc, netpoll.NewDialer()),
		client.WithName("hertz-layout"), // 设置用户代理头中使用的客户端名称
		client.WithRetryConfig( // 重试配置
			retry.WithMaxAttemptTimes(1), // 重试次数，不重试
		),
		client.WithConnStateObserve(stateFunc, 10*time.Second),
	)
	if err != nil {
		panic(err)
	}
}

func performRequest(method string, url string, body any, data any) error {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI(url)
	req.SetMethod(method)
	if body != nil {
		bodyMash, err := sonic.Marshal(body)
		if err != nil {
			return errors.WithMessage(err, "json marshal err")
		}
		req.SetBody(bodyMash)
	}

	if err := c.Do(context.Background(), req, resp); err != nil {
		return errors.WithMessage(err, "acquire request err")
	}
	if resp.StatusCode() > consts.StatusAccepted {
		return errors.New("response status")
	}

	if err := sonic.Unmarshal(resp.Body(), data); err != nil {
		return errors.WithMessage(err, "unmarshal response body err")
	}

	return nil
}

func GET(url string, resp any) error {
	return performRequest(consts.MethodGet, url, nil, resp)
}

func POST(url string, body any, resp any) error {
	return performRequest(consts.MethodPost, url, body, resp)
}

func PUT(url string, body any, resp any) error {
	return performRequest(consts.MethodPut, url, body, resp)
}

func DELETE(url string, resp any) error {
	return performRequest(consts.MethodDelete, url, nil, resp)
}
