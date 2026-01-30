package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestNewMyWaitGroup() {
	wg2 := NewMyWaitGroup(5)
	assert.NotNil(s.T(), wg2)
	assert.Equal(s.T(), 5, cap(wg2.semaphore))
}

func (s *ServiceSuite) TestWaitGroupBasic() {
	const goroutinesCount = 3
	wg := NewMyWaitGroup(goroutinesCount)

	var counter int
	var mu sync.Mutex
	completed := make(chan bool, goroutinesCount)

	for i := 0; i < goroutinesCount; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			mu.Lock()
			counter++
			mu.Unlock()
			completed <- true
		}(i)
	}

	wg.Wait()

	assert.Equal(s.T(), goroutinesCount, counter)
	for i := 0; i < goroutinesCount; i++ {
		select {
		case <-completed:
			// OK
		default:
			s.T().Errorf("Не все горутины завершились")
		}
	}
}

func (s *ServiceSuite) TestWaitBlocksUntilAllDone() {
	const goroutinesCount = 4
	wg := NewMyWaitGroup(goroutinesCount)

	waitStarted := make(chan bool)
	allDone := make(chan bool)
	startTime := time.Now()

	for i := 0; i < goroutinesCount; i++ {
		go func(duration time.Duration) {
			defer wg.Done()
			time.Sleep(duration)
		}(time.Duration(i+1) * 50 * time.Millisecond)
	}

	go func() {
		waitStarted <- true
		wg.Wait()
		allDone <- true
	}()

	<-waitStarted

	select {
	case <-allDone:
		s.T().Error("Wait завершился слишком рано")
	case <-time.After(100 * time.Millisecond):
		
	}


	select {
	case <-allDone:
		elapsed := time.Since(startTime)
		
		assert.GreaterOrEqual(s.T(), elapsed.Milliseconds(), int64(200))
	case <-time.After(300 * time.Millisecond):
		s.T().Error("Wait не завершился вовремя")
	}
}