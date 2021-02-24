package main

import (
    "fmt"
	"time"
)

func main() {

    messages := make(chan string)

    go workOnWriteOnlyChannel(messages)

	fmt.Printf("start at %s\n", time.Now())
    msg := workOnReadOnlyChannel(message)
    fmt.Printf("get message : %s at $s\n", msg, time.Now())
}

func workOnWriteOnlyChannel(messages chan<- string) {
	time.Sleep(5 * time.Second)
    messages <- "done"
}

func workOnReadOnlyChannel(messag <-chan string) (string) {
    return <-message
}

// no outpus