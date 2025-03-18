package patterns

import (
	"fmt"
	"sync"
)

var greetings = []string{
	"hello",
	"hola",
	"bonjour",
	"ciao",
	"namaste",
}

var names = []string{
	"alice",
	"bob",
	"charlie",
	"david",
}

func getNames(names []string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, name := range names {
			c <- name
		}
	}()
	return c
}

func fanIn(channels ...<-chan string) <-chan string {
	c := make(chan string)

	var wg sync.WaitGroup
	//Interesting we need to start goroutine for each input channel
	for _, ch := range channels {
		//Do we need to parameter to goroutine or form a closure?
		//it is fixed in go 1.22 https://go.dev/blog/loopvar-preview
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				c <- msg
			}
		}(ch)
	}

	//Waiting for all goroutine to completes then close the channel
	go func() { wg.Wait(); close(c) }()
	return c
}

func processMessage(c <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for ch := range c {
			for _, greeting := range greetings {
				out <- fmt.Sprintf("%s %s ", ch, greeting)
			}
		}
	}()

	return out
}
func FanInFanOut() {

	namesChan := getNames(names)
	var workers = 3
	channels := make([]<-chan string, workers)

	for worker := range workers {
		channels[worker] = processMessage(namesChan)
	}

	messages := fanIn(channels...)
	for msg := range messages {
		fmt.Println(msg)
	}

	fmt.Println("all message received")
}
