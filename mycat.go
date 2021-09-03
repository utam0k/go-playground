package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func print_file(number *int, file io.Reader, is_output_lines bool) {
	*number += 1

	scanner := bufio.NewScanner(file)
	scanner.Split((bufio.ScanLines))
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for i, line := range text {
		if is_output_lines {
			fmt.Printf("%d: %s\n", i, line)
		} else {
			fmt.Printf("%s\n", line)
		}
	}
}

func main() {
	var is_output_lines = flag.Bool("n", false, "Number the output lines, starting at 1.")
	flag.Parse()
	file_names := flag.Args()
	fmt.Println(*is_output_lines)
	var number int
	for _, file_name := range file_names {
		fmt.Println(file_name)
		file, err := os.Open(file_name)
		if err != nil {
			log.Fatal("failed to open")
		}
		print_file(&number, file, *is_output_lines)
		file.Close()
	}
}
