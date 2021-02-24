// Golang program to illustrate the usage of 
// After() function in Golang 

// Including main package 
package main 

// Importing fmt and time 
import ( 
	"fmt"
	"time"
) 

// Creating a channel 
// Using var keyword 
var ch chan int

// Main function 
func main() { 

	// For loop 
	for i := 1; i < 4; i++ { 

		// Prints these util loop stops 
		fmt.Printf("Keep displaying %s\n", time.Now()) 

		time.Sleep(time.Second)
	} 

	// Select statement 
	select { 

	// Using case statement to receive 
	// or send operation on channel and 
	// calling After() method with its 
	// parameter 
	case <-time.After(3 * time.Second): 

		// Printed when timed out 
		fmt.Println("Time Out!") 
	} 
} 

/*
Keep displaying 2021-02-10 16:20:55.242155 -0600 CST m=+0.000065684
Keep displaying 2021-02-10 16:20:56.242392 -0600 CST m=+1.000307980
Keep displaying 2021-02-10 16:20:57.244496 -0600 CST m=+2.002418440
Time Out!
*/
// between the last line and the the last second line, it has an obvious pause for 3 seconds