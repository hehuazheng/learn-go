package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
	fmt.Println("all done")
}

func worker(i int) {
	fmt.Println(i)
	time.Sleep(time.Second * time.Duration(1))
	wg.Done()
}
