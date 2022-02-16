package creature

import "fmt"

type Cat struct {
	Name   string
	Age    int
	IsFull bool
}

func (c *Cat) SayHello() error {
	fmt.Println("Meow~")
	return nil
}

func (c *Cat) Run() error {
	return fmt.Errorf("I don't care")
}

func (c *Cat) Eat() error {
	return fmt.Errorf("I don't care")
}
