package main

import (
	"fmt"
	"go-futures/futures"
	"time"
)

func slowIntegerFetch() int {
	time.Sleep(1 * time.Second)
	return 5
}

func slowStringFetch(d time.Duration) string {
	time.Sleep(d)
	return "X"
}

func main() {
	now := time.Now()
	f := futures.NewFuture(slowIntegerFetch)
	x, err := f.Get()
	fmt.Printf("value: %v\n", x)
	fmt.Printf("error: %v\n", err)
	fmt.Printf("Time taken to get value the first time: %v\n", time.Since(now))
	
	now = time.Now()
	x, err = f.Get()
	fmt.Printf("value: %v\n", x)
	fmt.Printf("error: %v\n", err)
	fmt.Printf("Time taken to get value the second time: %v\n", time.Since(now))

	now = time.Now()
	f = futures.NewFuture(slowStringFetch, 3*time.Second)
	x, err = f.GetWithTimeout(1 * time.Second)
	fmt.Printf("value: %v\n", x)
	fmt.Printf("error: %v\n", err)
	fmt.Printf("Time taken to get value with timeout the first time: %v\n", time.Since(now))

	now = time.Now()
	x, err = f.GetWithTimeout(5 * time.Second)
	fmt.Printf("value: %v\n", x)
	fmt.Printf("error: %v\n", err)
	fmt.Printf("Time taken to get value with timeout the first time: %v\n", time.Since(now))
}
