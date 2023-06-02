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

// UnmarshalString 将字符串解析到指定结构体中，需要指定 Format。
// 注意，这里需要传入的是 model 的地址。
func UnmarshalString(s string, model any, format Format) error {
	return UnmarshalBytes([]byte(s), model, format)
}

// UnmarshalBytes 将 byte 数组解析到指定结构体中，需要指定 Format。
// 注意，这里需要传入的是 model 的地址。
func UnmarshalBytes(data []byte, model any, format Format) error {
	switch format {
	case JsonFormat:
		return json.Unmarshal(data, model)
	case YamlFormat:
		return yaml.Unmarshal(data, model)
	default:
		return errors.New("unsupported format")
	}
}
