package main

import "fmt"

func sendChannel(sent chan<- string, msg string) {
	sent <- msg
}

func dualChannel(received <-chan string, sent chan<- string) {
	msg := <-received
	sent <- msg
}

func main() {
	msgs := []string{"first msg", "second msg", "third msg"}

	toSend := make(chan string, len(msgs))
	toReceive := make(chan string, len(msgs))

	for _, msg := range msgs {
		sendChannel(toSend, msg)
		dualChannel(toSend, toReceive)
		fmt.Println(<-toReceive)
	}

}
