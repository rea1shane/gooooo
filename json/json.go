package json

import (
	"github.com/goccy/go-json"
	"io/ioutil"
)

// Load 将 json 文件中的数据加载到 model 中
// 注意，这里需要传入的是 model 的地址，见测试用例 TestLoad
func Load(path string, model any) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, model)
}
