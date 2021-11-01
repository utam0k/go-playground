package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	timer := time.NewTimer(2 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timer.C:
				fmt.Println("A Timer finished")
				wg.Done()
			case _, ok := <-ch:
				if !ok {
					return
				}
				log.Println("OK")
			}
		}
	}()
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
	}
	close(ch)
	wg.Wait()
}
