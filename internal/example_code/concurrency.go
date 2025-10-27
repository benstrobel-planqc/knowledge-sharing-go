package example

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func DemoRoutines() {
	go say("world")
	say("hello")
}

func DemoChannels() {
	_ = make(chan int)
	bch := make(chan int, 2)

	writeNumbersToChannel := func(ch chan int) {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}

	go writeNumbersToChannel(bch)
	for v := range bch {
		fmt.Println(v)
	}
}

func DemoMutex() {
	counter := 0
	var mu sync.Mutex
	increment := func() {
		mu.Lock()
		counter++
		fmt.Println(counter)
		mu.Unlock()
	}

	iterate := func() {
		for range 5 {
			increment()
			time.Sleep(100 * time.Millisecond)
		}
	}

	go iterate()
	iterate()
}
