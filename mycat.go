package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func print_file(number *int, file_name string, is_output_lines bool, max_digits uint8) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal("failed to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split((bufio.ScanLines))
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, line := range text {
		*number += 1
		if is_output_lines {
			fmt.Printf("%0*d: %s\n", max_digits, *number, line)
		} else {
			fmt.Printf("%s\n", line)
		}
	}
}

func count_number_of_lines(file_name string) (int, error) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	number_of_lines, err := lineCounter(file)
	if err != nil {
		log.Fatal(err)
	}
	return number_of_lines, nil
}

func main() {
	var is_output_lines = flag.Bool("n", false, "Number the output lines, starting at 1.")
	flag.Parse()
	file_names := flag.Args()

	var total_number_of_lines int
	for _, file_name := range file_names {
		number_of_lines, err := count_number_of_lines(file_name)
		if err != nil {
			log.Fatal(err)
		}
		total_number_of_lines += number_of_lines
	}
	line_count_max_digit := uint8(math.Log10(float64(total_number_of_lines))) + 1

	var number int
	for _, file_name := range file_names {
		print_file(&number, file_name, *is_output_lines, line_count_max_digit)
	}
}
