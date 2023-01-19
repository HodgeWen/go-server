/*
 * @Author: whj
 * @Date: 2023-01-12 16:22:03
 * @LastEditors: whj
 * @LastEditTime: 2023-01-19 10:40:37
 * @FilePath: /test-go/main.go
 *
 */
// package main

// import (
// 	"net"
// 	"net/http"
// 	"os"
// 	"sync"
// 	_ "test-go/api"
// )

// func main() {}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal)
	signal.Notify(
		sigs,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
	)

	// 监听所有信号
	log.Println("listen sig")
	signal.Notify(sigs)

	// 打印进程id
	log.Println("PID:", os.Getppid())

	<-sigs
}
