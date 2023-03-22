package yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func Load(path string, model interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, model)
}
