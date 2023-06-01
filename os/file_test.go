package os

import (
	"fmt"
	"github.com/rea1shane/gooooo/data"
	"testing"
)

type JsonData struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func TestLoadJson(t *testing.T) {
	var jsonData JsonData
	if err := Load("test.json", &jsonData, data.JsonFormat); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", jsonData)
}

type YamlData struct {
	A string `yaml:"a"`
	B string `yaml:"b"`
	C string `yaml:"c"`
}

func TestLoadYaml(t *testing.T) {
	var yamlData YamlData
	if err := Load("test.yaml", &yamlData, data.YamlFormat); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", yamlData)
}
