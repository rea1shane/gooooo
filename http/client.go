package http

import (
	"io"
	"net/http"

	"github.com/morikuni/failure"

	"github.com/rea1shane/gooooo/data"
)

// Load 执行请求并且解析返回体数据加载到 model 中。
func Load(client *http.Client, request *http.Request, model any, format data.Format) error {
	// 请求
	response, err := client.Do(request)
	if err != nil {
		return failure.Wrap(err, failure.Message("request"))
	}
	defer response.Body.Close()

	// 读取
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return failure.Wrap(err, failure.Message("read response body"))
	}

	// 解析
	return failure.Wrap(data.UnmarshalBytes(responseBody, model, format), failure.Message("unmarshal"))
}
