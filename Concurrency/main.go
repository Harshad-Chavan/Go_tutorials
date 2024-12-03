package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

// agg 'go' in front of the methods to make them run as go routines
// but this will not give any ouptut

// channel can be used as a communication well
// pass values strings
// if we dont use channal the fucntions which are dispatched as go routines will run in backgorud
// but the main will finish it will not wait for the fucntiosn to be finish
// for that pass somethig in channel and wait for completion
func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	// output of channel
	// this will give random outputs becoz go does not wait till all are finished
	// as soon as any one of them finishes it will end the main
	//<-done

	//one way of avoiding this is to write <- done as many time as the number of go routines
	// <-done
	// <-done
	// <-done
	// <-done

	// another way is to loop over the channal
	// but this gives an error
	// fatal error: all goroutines are asleep - deadlock!/
	// to fic this close the channel in the logest running go routine
	for range done {

	}

}
