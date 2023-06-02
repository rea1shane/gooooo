package data

import (
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Format int

const (
	_ Format = iota
	JsonFormat
	YamlFormat
	XmlFormat
)

func (f Format) String() string {
	switch f {
	case JsonFormat:
		return "json format"
	case YamlFormat:
		return "yaml format"
	case XmlFormat:
		return "xml format"
	default:
		return "unknown format"
	}
}

// UnmarshalString 将字符串解析到指定结构体中，需要指定 Format。
// 注意，这里需要传入的是 model 的地址。
func UnmarshalString(s string, model any, format Format) error {
	return UnmarshalBytes([]byte(s), model, format)
}

// UnmarshalBytes 将 byte 数组解析到指定结构体中，需要指定 Format。
// 注意，这里需要传入的是 model 的地址。
func UnmarshalBytes(data []byte, model any, format Format) (err error) {
	switch format {
	case JsonFormat:
		err = json.Unmarshal(data, model)
	case YamlFormat:
		err = yaml.Unmarshal(data, model)
	default:
		err = errors.New("unsupported format")
	}
	return errors.Wrap(err, format.String())
}
