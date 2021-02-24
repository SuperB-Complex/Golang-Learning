package main

import (
    "fmt"
    "time"
)

func main() {

    message := make(chan string, 1)

    workOnWriteOnlyChannel(message)
    workOnReadOnlyChannel(message)
	/*
		start at 2009-11-10 23:00:05 +0000 UTC m=+5.000000001
		get message : done at 2009-11-10 23:00:05 +0000 UTC m=+5.000000001
	*/

	workOnWriteOnlyChannel(message)
    go workOnReadOnlyChannel(message)
	/*
		start at 2009-11-10 23:00:05 +0000 UTC m=+5.000000001
	*/

	go workOnWriteOnlyChannel(message)
    workOnReadOnlyChannel(message)
	/*
		start at 2009-11-10 23:00:05 +0000 UTC m=+5.000000001
		get message : done at 2009-11-10 23:00:05 +0000 UTC m=+5.000000001
	*/

	/*
	workOnReadOnlyChannel(message)
    workOnWriteOnlyChannel(message)
	*/
	// fatal error: all goroutines are asleep - deadlock!

	/*
	workOnReadOnlyChannel(message)
    go workOnWriteOnlyChannel(message)
	*/
	// fatal error: all goroutines are asleep - deadlock!

	go workOnReadOnlyChannel(message)
    workOnWriteOnlyChannel(message)
	/*
		start at 2009-11-10 23:00:05 +0000 UTC m=+5.000000001
	*/
}

func workOnWriteOnlyChannel(messages chan<- string) {
    time.Sleep(5 * time.Second)
    messages<- "done"
    defer fmt.Printf("start at %s\n", time.Now())
}

func workOnReadOnlyChannel(messages <-chan string) {
    fmt.Printf("get message : %s at %s\n", <-messages, time.Now())
}

