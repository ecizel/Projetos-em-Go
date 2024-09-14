package main

import (
	"fmt"
	"time"
)

func main() {
	// funcao principal
	ping_pong()

	// parar: ENTER
	fmt.Scanln()
}

func ping_pong() {

	ping_pong := make(chan (string))

	// funcao ping
	go func() {
		for {
			time.Sleep(time.Second * 2)
			ping_pong <- "ping"
			fmt.Println(<- ping_pong)
		}
	}()

	// funcao pong
	go func() {
		for {
			fmt.Println(<- ping_pong)
			time.Sleep(time.Second * 2)
			ping_pong <- "pong"
		}
	}()
}