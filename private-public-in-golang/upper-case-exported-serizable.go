package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	type JsonMessage struct {
        A int `json:"A"`
        B string `json:"B,omitempty"`
        GenericObj interface{} 
    }

    type JsonObj struct {
        C string `json:"C"`
    }

    myJsonObj := JsonObj{C: "cVal"}
    myJson := JsonMessage{A: 1, B: "bValue", GenericObj: myJsonObj}

    body, _ := json.Marshal(myJson)
	fmt.Print(body) // [123 34 65 34 58 49 44 34 66 34 58 34 98 86 97 108 117 101 34 44 34 71 101 110 101 114 105 99 79 98 106 34 58 123 34 67 34 58 34 99 86 97 108 34 125 125]
    fmt.Print(string(body)) // {"A":1,"B":"bValue","GenericObj":{"C":"cVal"}}
}