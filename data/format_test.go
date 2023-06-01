package data

import (
	"fmt"
	"testing"
)

type JsonData struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

const jsonString = `{"a": "1", "b": "2", "c": "3"}`

func TestUnmarshalJsonString(t *testing.T) {
	var model JsonData
	if err := UnmarshalString(jsonString, &model, JsonFormat); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", model)
}

type YamlData struct {
	A string `yaml:"a"`
	B string `yaml:"b"`
	C string `yaml:"c"`
}

const yamlString = `a: 4
b: 5
c: 6`

func TestUnmarshalYamlString(t *testing.T) {
	var model YamlData
	if err := UnmarshalString(yamlString, &model, YamlFormat); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", model)
}
