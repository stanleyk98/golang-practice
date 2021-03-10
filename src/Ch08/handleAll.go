package main

import(
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(signal os.Signal){
	fmt.Println("받은 신호 : " signal)
}

func main(){
	sigs := make(chan os.Signal, 1)
	
	// sigs 에 특정 시그널 선택안했기 때문에 모든 시그널에 해당함
	signal.Notify(sigs)

	go func(){
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt :
				handle(sig)
			case syscall.SIGTERM :
				handle(sig)
				os.Exit(0)
			case syscall.SIGUSR2 :
				fmt.Println("Hadling syscall.SIGUSR2~!!!")

			default :
			fmt.Println("Ignoring : ", sig)
			}
		}
	}()

	for {
		fmt.Printf(("."))
		time.Sleep(30 * time.Second)
	}
}

// 컴파일해서 다른 터미널에서 구동해야함
// go build handleAll.go
// 나는 패스
