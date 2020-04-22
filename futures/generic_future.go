package futures

import (
	"reflect"
	"fmt"
	"time"
)

// Future is an interface that abstracts the implementation of the future methods
type Future interface {
	Get() (interface{}, error)
	GetWithTimeout(d time.Duration) (interface{}, error)
}

type future struct {
	value interface{}
	done bool
	err error
	resultChan chan interface{}
}

// NewFuture returns an object that implements the Future interface
func NewFuture(futureFunc interface{}, args... interface{}) Future {
	fn := reflect.ValueOf(futureFunc)
	fnArgs := make([]reflect.Value, len(args))

	for i, arg := range args {
		fnArgs[i] = reflect.ValueOf(arg)
	}

	resultChan := make(chan interface{}, 1)

	go func() {
		result := fn.Call(fnArgs)
		resultChan <- result[0].Interface()
	}()

	return &future{
		resultChan: resultChan,
	}
}

// Get will block till a value is available to be returned
func (f *future) Get() (interface{}, error) {
	if f.done {
		return f.value, f.err
	}
	f.value = <-f.resultChan
	f.done = true
	f.err = nil
	return f.value, f.err
}

// GetWithTimeout will block till either the given timeout has elapsed
// 	or till a value is available to be returned
func (f *future) GetWithTimeout(d time.Duration) (interface{}, error) {
	if f.done {
		return f.value, f.err
	}
	select {
	case <- time.After(d):
		f.value = nil
		f.err = fmt.Errorf("timeout")
		f.done = true
	case x := <-f.resultChan:
		f.value = x
		f.err = nil
		f.done = true
	}
	return f.value, f.err
}