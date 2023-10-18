package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	secondImpl()
}

func secondImpl() {
	fmt.Println("starting main thread")

	var wg sync.WaitGroup
	wg.Add(5)

	firstChan := make(chan struct{}, 3)
	secondChan := make(chan struct{}, 3)

	for i := 1; i < 6; i++ {
		go func(i int) {
			fmt.Printf("starting %d goroutine\n", i)
			if i == 1 {
				defer func() {
					for i := 0; i < 3; i++ {
						firstChan <- struct{}{}
					}
				}()
			} else if i == 5 {
				for i := 0; i < 3; i++ {
					<-secondChan
				}
			} else {
				<-firstChan
				defer func() {
					secondChan <- struct{}{}
				}()
			}

			// sleepForUpToFiveSeconds()
			fmt.Printf("ending %d goroutine\n", i)
			wg.Done()
		}(i)

	}

	wg.Wait()

	fmt.Println("ending main thread")
}

func sleepForUpToFiveSeconds() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(5) + 1

	time.Sleep(time.Second * time.Duration(randomInt))
}
