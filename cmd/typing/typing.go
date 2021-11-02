package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan struct{})
	timer := time.NewTimer(2 * time.Second)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			ch <- struct{}{}
		}
	}()
	for {
		select {
		case <-timer.C:
			fmt.Println("A Timer finished")
			return
		case _, ok := <-ch:
			if !ok {
				return
			}
			log.Println("OK")
		}
	}
}
