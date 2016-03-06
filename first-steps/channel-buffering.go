package main

import "fmt"

func main() {

	msgsToSend := [5]string{"first", "second", "third", "fourth", "fifth"}

	messages := make(chan string, len(msgsToSend))

	for _, msg := range msgsToSend {
		messages <- msg + " msg"
	}

	for i := 0; i < len(msgsToSend); i++ {
		fmt.Println(<-messages)
	}

}
