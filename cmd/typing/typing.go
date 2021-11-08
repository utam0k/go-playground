package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func readLineLoop(ch chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal("failed to read a rune.")
		}
		ch <- string(line)
	}

}

func main() {
	var timeoutSec = flag.Int("t", 10, "Timeout")
	flag.Parse()

	problems := [...]string{"Golang", "Rust", "Python", "Ruby", "Java", "Scala"}
	ch := make(chan string)
	timer := time.NewTimer(time.Duration(*timeoutSec) * time.Second)
	go readLineLoop(ch)

	var score int
	for {
		p := problems[rand.Intn(len(problems))]
		fmt.Println("problem: " + p)
		select {
		case <-timer.C:
			fmt.Println("TIME UP")
			fmt.Printf("score: %d\n", score)
			return
		case answer, ok := <-ch:
			if !ok {
				return
			}
			if p == answer {
				score += len(p)
				fmt.Println("OK")
			} else {
				fmt.Println("NG")
			}
		}
	}
}
