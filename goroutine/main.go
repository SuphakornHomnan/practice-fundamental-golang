package main

import (
	"fmt"
	"sync"
)

var i int
var mux sync.Mutex

type thing struct {
	i   int
	mux sync.Mutex
}

func main() {
	// go func() {
	// 	for {
	// 		fmt.Println(read())
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	write(i)
	// }

	ch := make(chan int)
	qCh := make(chan struct{})
	go fibonacci(ch, qCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	qCh <- struct{}{}
	<-qCh
	fmt.Println("end")
}

func write(n int) {
	mux.Lock()
	i = n
	mux.Unlock()
}

func read() int {
	mux.Lock()
	defer mux.Unlock()
	return i
}

func fibonacci(ch chan int, qCh chan struct{}) {
	a, b := 0, 1

	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-qCh:
			qCh <- struct{}{}
			return
		}

	}
}
