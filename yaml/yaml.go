package yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Load 将 yaml 文件中的数据加载到 model 中
// 注意，这里需要传入的是 model 的地址，见测试用例 TestLoad
func Load(path string, model interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, model)
}
