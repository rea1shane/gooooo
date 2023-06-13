package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// Wait 阻塞程序直到接收到 os.Signal，并将其返回。
// 参考：https://juejin.cn/post/7038415550082449416
func Wait() os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	return <-signals
}
