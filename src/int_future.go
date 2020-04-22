package futures

import (
	"time"
	"fmt"
)

type IntFuture struct {
	outChan chan int
}

func CreateIntFuture(fetch func() int) *IntFuture {
	outChan := make(chan int, 1)
	future := &IntFuture{}
	future.outChan = outChan
	go func() {
		outChan <- fetch()
	}()
	return future
}

func (f *IntFuture) Get() (int, error) {
	x := <- f.outChan
	return x, nil
}

func (f *IntFuture) GetWithTimeout(timeout time.Duration) (int, error) {
	select {
		case x := <-f.outChan:
			return x, nil
		case <-time.After(timeout):
	}

	return -1, fmt.Errorf("future timed out")
}