package main

import (
	"fmt"

	"golandtest/creature"
	"golandtest/log"
)

func main() {
	sayHelloWord()

	cat := creature.Cat{}
	cat.SayHello()
}

func sayHelloWord() {
	index := 100
	log.Info("sayHalloWrod").Int("index", index).Send()

	for index := 0; index < 10; index++ {
		log.Info("say in loop").Int("index", index).Send()
	}
}

func newAnError(topic string) error {
	return fmt.Errorf(topic)
}
