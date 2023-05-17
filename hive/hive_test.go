package hive

import (
	"context"
	"github.com/beltran/gohive"
	"testing"
)

const (
	zookeeperQuorum = "localhost:2181"
	username        = "username"
	password        = "password"
	sql             = "SHOW DATABASES"
)

// TestNewClient 通过 ZooKeeper 配置创建客户端
func TestNewClient(t *testing.T) {
	config := gohive.NewConnectConfiguration()
	config.Username = username
	config.Password = password
	connection, err := gohive.ConnectZookeeper(zookeeperQuorum, "NONE", config)
	if err != nil {
		panic(err)
	}
	cursor := connection.Cursor()
	cursor.Exec(context.Background(), sql)
}
