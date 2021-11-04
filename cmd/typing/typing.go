package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	problems := [...]string{"Golang", "Rust", "Python", "Ruby", "Java", "Scala"}
	ch := make(chan string)
	timer := time.NewTimer(5 * time.Second)
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			line, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal("failed to read a rune.")
			}
			ch <- string(line)
		}
	}()

	for _, p := range problems {
		fmt.Println(p)
		select {
		case <-timer.C:
			fmt.Println("A Timer finished")
			return
		case answer, ok := <-ch:
			if !ok {
				return
			}
			if p == answer {
				fmt.Println("OK")
			} else {
				fmt.Println("NG")
			}
		}
	}
}
