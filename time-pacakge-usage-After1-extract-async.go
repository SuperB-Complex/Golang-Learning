// Golang program to illustrate the usage of 
// After() function in Golang 
  
// Including main package 
package main 
  
// Importing fmt and time 
import ( 
    "fmt"
    "time"
) 
  
func asyncMethod(channel chan string) {
    select {
	case output := <- channel:
	    fmt.Printf("stopped after output gets %s from channel. %s\n", output, time.Now())
  	    
	case <- time.After(10 * time.Second):
		fmt.Printf("stopped after wating 10 seconds. %s\n", time.Now())
    }
}

// Main function 
func main() { 
  
    // Creating a channel 
    // Using make keyword 
    channel := make(chan string, 1) 
  
    fmt.Printf("%s at %s\n", "start", time.Now())
    go asyncMethod(channel)

    time.Sleep(20 * time.Second)
    channel<- "terminal "
    // channel <- "is "
    // channel <- "done!"
} 