package main

import (
    "fmt"
	"time"
)

func main() {

    messages := make(chan string)

    go work(messages)

	fmt.Printf("start at %s\n", time.Now())
    msg := <-messages
    fmt.Printf("get message : %s at $s\n", msg, time.Now())
}

func work(messages chan<- string) {
	time.Sleep(10 * time.Second)
    messages <- "done"
}
