package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// https://gist.github.com/miguelmota/2a0c0e96c22bccc8740819d5d64ff8d0
func main() {
	input := []byte{14, 255, 129, 4, 1, 2, 255, 130, 0, 1, 12, 1, 12, 0, 0, 12, 255, 130, 0, 1, 3, 102, 111, 111, 3, 98, 97, 114}
	buf := bytes.NewBuffer(input)
	dec := gob.NewDecoder(buf)

	m := make(map[string]string)

	if err := dec.Decode(&m); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m["foo"]) // "bar"
}