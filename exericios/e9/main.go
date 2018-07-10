package main

import (
	"fmt"
)

func main() {
	msg := make(chan string) // mensagem
	done := make(chan bool)  //verdadeira ou falsa

	go sendMessage(msg)          // go rotina de envio de mensagem
	go receiveMessage(msg, done) // go rotina de recebimento

	<-done
}

func sendMessage(msg chan string) {
	msg <- "SOMOS 3K"
}

func receiveMessage(msg chan string, done chan bool) {
	fmt.Println(<-msg)
	done <- true
}
