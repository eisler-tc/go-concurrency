// esieve implements a Sieve of Eratosthenes
// as a series of channels connected together
// by goroutines
package main

import "fmt"

func sieve(mine int, inch chan int) {
	start := true                                 // First-number switch
	ouch := make(chan int)                        // Output channel for this instance
	fmt.Printf("%v\n", mine)                      // Print this instance's prime
	for next := <-inch; next > 0; next = <-inch { // Read input channel
		fmt.Printf("%v <- %v\n", mine, next) // (Trace)
		if (next % mine) > 0 {               // Divisible by my prime?
			if start { // No; is it the first number through?
				go sieve(next, ouch) // First number - create instance for it
				start = false        // First time done
			} else { // Not first time
				ouch <- next // Pass it to the next instance
			}
		}
	}
}

func main() {
	lim := 30                                 // Let's do up to 30
	fmt.Printf("%v\n", 2)                     // Treat 2 as a special case
	ouch := make(chan int)                    // Create the first segment of the pipe
	go sieve(3, ouch)                         // Create the instance for '3'
	for prime := 3; prime < lim; prime += 2 { // Generate 3, 5, ...
		fmt.Printf("Send %v\n", prime) // Trace
		ouch <- prime                  // Send it down the pipe
	}
}