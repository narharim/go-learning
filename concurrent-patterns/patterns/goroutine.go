package patterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// launch 5 goroutine wait for all goroutines to complete using with help of sync package
func Run() {
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("goroutine %d starting\n", id)
			doWork()
			fmt.Printf("goroutine %d done\n", id)

		}(i)
	}
	fmt.Println("waiting for all goroutines to to complete...")
	wg.Wait()
	fmt.Println("completed")

}

func doWork() {
	time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
}

// run with mutex for shared state protection
func RunWithMutex() {
	var wg sync.WaitGroup

	var mu sync.Mutex

	counter := 0

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			doWork()
			mu.Lock()
			counter++
			fmt.Printf("goroutine %d, counter %d\n", id, counter)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("final counter value: %d\n", counter)
}
