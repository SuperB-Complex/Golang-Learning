package main
 
import (
    "fmt"
    "encoding/json"
)

// The tags are represented as raw string values (wrapped within a pair of ``) and ignored by normal code execution.
type Employee struct {
    FirstName  string `json:"firstname"`
    LastName   string `json:"lastname"`
    City string `json:"city"`
}
 
func main() {
    json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`
 
    emp1 := new(Employee)
    json.Unmarshal([]byte(json_string), emp1)
    fmt.Println(emp1) // &{Rocky Sting London}
 
    emp2 := new(Employee)
    emp2.FirstName = "Ramesh"
    emp2.LastName = "Soni"
    emp2.City = "Mumbai"
    jsonStr, _ := json.Marshal(emp2)
    fmt.Printf("%s\n", jsonStr) // {"firstname":"Ramesh","lastname":"Soni","city":"Mumbai"}
}