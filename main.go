package main

import "fmt"

type Message struct {
	text   string
	author string
}

func main() {
	message := Message{}

	message.text = "Привет, Мир!"
	message.author = "Ivan Belmeshkin"

	fmt.Printf("Чувак %s сказал: %s", message.author, message.text)
}
