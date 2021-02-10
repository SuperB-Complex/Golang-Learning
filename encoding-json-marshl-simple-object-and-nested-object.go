package main

import (
	"fmt"
	"encoding/json"
)

type JsonMessageOriginal map[string]interface{}
type JsonMessage = JsonMessageOriginal

type JsonObj struct {
	C string `json:"C"`
}

func main() {
    myJsonObj := JsonObj{C: "cVal"}
    // myJsonLiteralObj := {"C": "cVal"} invalid syntax
    // if you want to initiate a map you should use make(map[string]int)
    fmt.Printf("%T\n", myJsonObj) // main.JsonObj
    myJson := JsonMessage{"A": 1, "B": "bValue"} 
    myNestedJson := JsonMessage{"A": 1, "B": "bValue", "Generic": myJsonObj}
    
    myBody, _ := json.Marshal(myJson)
    fmt.Println(myBody) // [123 34 65 34 58 49 44 34 66 34 58 34 98 86 97 108 117 101 34 125]
    fmt.Println(string(myBody)) // {"A":1,"B":"bValue"}
    
    myNestedBodyWithOutAddressOperator, _ := json.Marshal(myNestedJson)
    fmt.Println(myNestedBodyWithOutAddressOperator) // [123 34 65 34 58 49 44 34 66 34 58 34 98 86 97 108 117 101 34 44 34 71 101 110 101 114 105 99 34 58 123 34 67 34 58 34 99 86 97 108 34 125 125]
    fmt.Println(string(myNestedBodyWithOutAddressOperator)) // {"A":1,"B":"bValue","Generic":{"C":"cVal"}}

    myNestedBodyWithAddressOperator, _ := json.Marshal(myNestedJson)
    fmt.Println(myNestedBodyWithAddressOperator) // [123 34 65 34 58 49 44 34 66 34 58 34 98 86 97 108 117 101 34 44 34 71 101 110 101 114 105 99 34 58 123 34 67 34 58 34 99 86 97 108 34 125 125]
    fmt.Println(string(myNestedBodyWithAddressOperator)) // {"A":1,"B":"bValue","Generic":{"C":"cVal"}}
}