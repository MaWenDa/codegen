package graceful

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup
var ctx, cancel = context.WithCancel(context.Background())

type Graceful interface {
	GracefulStart(ctx context.Context)
}

func StartFunc(f func(ctx context.Context)) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f(ctx)
	}()
}

// Start 启动服务
func Start(srv Graceful) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.GracefulStart(ctx)
	}()
}

// Wait 阻塞等待信号
// Wait 函数监听操作系统的中断信号（如 Ctrl+C），一旦接收到信号，就会触发优雅地停止服务的过程。
// 它创建了一个信号通道 s，并通过 signal.Notify 函数将操作系统的中断信号通知发送到该通道。
// 一旦收到中断信号，Wait 函数会打印一条日志，然后调用 cancel 函数取消上下文中的操作，最后等待所有启动的任务完成。
func Wait() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-s
	log.Println("graceful stopping...")
	cancel()
	wg.Wait()
}
