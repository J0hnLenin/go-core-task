package main

import "sync"

func Reduce(channels []chan int64) chan int64 {
	reduced := make(chan int64)

	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	go func() {
		for _, channel := range channels {
			go func(ch chan int64) {
				for value := range ch {
					reduced <- value
				}
				defer wg.Done()
			}(channel)
		}
		wg.Wait()
		close(reduced)
	}()
	
	return reduced
}