package tool

import (
	"fmt"
	"github.com/jenrain/OnlyID/library/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func QuitSignal(quitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	log.GetLogger().Info("server start success", zap.Any("pid", os.Getpid()))
	fmt.Printf("server start success pid:%d\n", os.Getpid())
	for s := range c {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			quitFunc()
			return
		default:
			return
		}
	}
}
