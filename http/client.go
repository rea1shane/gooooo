package http

import (
	"github.com/pkg/errors"
	"github.com/rea1shane/gooooo/data"
	"io"
	"net/http"
)

// DoAndUnmarshal 执行请求并且解析返回体数据加载到 model 中。
func DoAndUnmarshal(client *http.Client, request *http.Request, model any, format data.Format) error {
	// 请求
	response, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "request")
	}
	defer response.Body.Close()

	// 读取
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "read response body")
	}

	// 解析
	return errors.Wrap(data.UnmarshalBytes(responseBody, model, format), "unmarshal")
}
