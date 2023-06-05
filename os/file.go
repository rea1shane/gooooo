package os

import (
	"github.com/morikuni/failure"
	"github.com/rea1shane/gooooo/data"
	"os"
)

// Load 读取文件并将数据加载到 model 中。
// 注意，这里需要传入的是 model 的地址。
func Load(path string, model any, format data.Format) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return failure.Wrap(err, failure.Message("read file"))
	}
	return failure.Wrap(data.UnmarshalBytes(file, model, format), failure.Message("unmarshal"))
}
