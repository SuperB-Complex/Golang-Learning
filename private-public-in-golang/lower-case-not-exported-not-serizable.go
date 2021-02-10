package main

import (
	"encoding/json"
	"fmt"
)

type JsonMessage struct {
    A int `json:"A"`
    B string `json:"B,omitempty"`
    genericObj interface{} 
}

type JsonObj struct {
	C string `json:"C"`
}
        	
func main() {
	myJsonObj := JsonObj{C: "cVal"}
    myJson := JsonMessage{A: 1, B: "bValue", genericObj: myJsonObj}

    body, _ := json.Marshal(myJson)
    fmt.Print(body) // [123 34 65 34 58 49 44 34 66 34 58 34 98 86 97 108 117 101 34 125]
	fmt.Print(string(body)) // {"A":1,"B":"bValue"}
}