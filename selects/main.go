package main

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

func pinger(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func eai(canal chan string) {
	for {
		canal <- "Promoção caminhão do Faustão: Você ganhou um carro e uma casa. Ligue para 48 99856 5656"
	}
}

func blz(canal chan string) {
	for {
		canal <- "Também tô preso man... kkkkk"
	}
}

func impressora() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
		case msg = <-sms:
			fmt.Println("[SMS]", msg)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {

	go pinger(irc)
	go ponger(irc)

	go eai(sms)
	go blz(sms)

	go impressora()

	var entrada string
	fmt.Scanln(&entrada)
}
