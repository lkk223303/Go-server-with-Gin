package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 監聽全部訊號
func main() {
	c := make(chan os.Signal)

	signal.Notify(c,os.Kill,syscall.SIGHUP,syscall.SIGINT,syscall.SIGTERM, syscall.SIGUSR1,syscall.SIGUSR2)

	go func ()  {
		for s := range c {
			switch s {
			case os.Kill:
				fmt.Println("It's a kill")
			case syscall.SIGHUP,syscall.SIGINT,syscall.SIGTERM:
				fmt.Println("Exit...",s)
				ExitFunc()
			case syscall.SIGUSR1:
				fmt.Println("USR1... ",s)
			case syscall.SIGUSR2:
				fmt.Println("USR2...",s)
			default:
				fmt.Println("其他信號:",s)
			}
		
		}
	}()
	

	
	fmt.Println("App running...")
	sum := 0
	for{
		sum++
		fmt.Println("休眠了...",sum,"秒")
		time.Sleep(1*time.Second)
	}
}

func ExitFunc(){
	fmt.Println("開始退出...")
	fmt.Println("正在退出...")
	fmt.Println("結束退出...")
	os.Exit(0)
}