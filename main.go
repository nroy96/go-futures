package main

import (
	"fmt"
	"go-futures/src"
	"time"
)

func slowIntegerFetch() int {
	time.Sleep(5 * time.Second)
	return 5
}

func main() {
	now := time.Now()
	intFuture := futures.CreateIntFuture(slowIntegerFetch)
	x, err := intFuture.GetWithTimeout(2*time.Second)
	fmt.Println(x)
	fmt.Println(err)
	fmt.Printf("%v\n", time.Since(now))
}