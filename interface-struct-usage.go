package main

import (
	"fmt"
)

type Message interface {
  New() Message
  Get() string
}

type Entity struct {}
func (e *Entity) New() Message {
  return e
}
func (e Entity) Get() string {
  return "hi"
}

func msg(m Message) {
  fmt.Println(m.Get())
}

func main() {
	e := &Entity{}
	message := e.New()
	msg(message)
}