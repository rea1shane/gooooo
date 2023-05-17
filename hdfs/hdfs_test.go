package hdfs

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"os"
	"testing"
)

const (
	nameNodeAddress = "localhost:19000"
	confPath        = "/etc/hadoop/conf/" // confPath 如果已经存在 HADOOP_CONF_DIR 或 HADOOP_HOME 环境变量，则将此值置空即可
	username        = "test"              // username 进行操作的 hdfs 用户名称

	dirPath = "/"
)

// TestNewClientUseSpecifiedNameNode1 使用指定的 NameNode 创建客户端
func TestNewClientUseSpecifiedNameNode1(t *testing.T) {
	client, err := hdfs.New(nameNodeAddress)
	if err != nil {
		panic(err)
	}
	readDir(client)
}

// TestNewClientUseSpecifiedNameNode2 使用指定的 NameNode 创建客户端
func TestNewClientUseSpecifiedNameNode2(t *testing.T) {
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: []string{nameNodeAddress},
		User:      username,
	})
	if err != nil {
		panic(err)
	}
	readDir(client)
}

// TestNewClientUseConfig 使用配置文件创建客户端
func TestNewClientUseConfig(t *testing.T) {
	hadoopConf := hdfs.LoadHadoopConf(confPath)
	namenodes, err := hadoopConf.Namenodes()
	if err != nil {
		panic(err)
	}
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: namenodes,
		User:      username,
	})
	if err != nil {
		panic(err)
	}
	readDir(client)
}

func readDir(client *hdfs.Client) {
	subDirs, err := client.ReadDir(dirPath)
	switch err.(type) {
	case *os.PathError: // 判断传入的路径是否错误
		fmt.Println("Path Error")
	case error:
		panic(err)
	default:
		// 遍历文件夹内容
		for _, dir := range subDirs {
			fmt.Println(dir.Name())
		}
	}
}
