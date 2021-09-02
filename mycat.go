package main

import (
	"flag"
	"fmt"
)

func main() {
	var is_output_lines = flag.Bool("n", false, "Number the output lines, starting at 1.")
	flag.Parse()
	fmt.Println(*is_output_lines)
}
