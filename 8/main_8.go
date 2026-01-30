package main

type MyWaitGroup struct {
	semaphore chan struct{}
}

func NewMyWaitGroup(gorutinesCount int) *MyWaitGroup {
	wg := &MyWaitGroup{
		semaphore: make(chan struct{}, gorutinesCount),
	}
	for range gorutinesCount {
		wg.semaphore <- struct{}{}
	}
	return wg
}

func (wg *MyWaitGroup) Done() {
	<-wg.semaphore
}

func (wg *MyWaitGroup) Wait() {
	for i := 0; i < cap(wg.semaphore); i++ {
		wg.semaphore <- struct{}{}
	}
}