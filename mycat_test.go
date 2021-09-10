package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCountNumberOfLines(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("", "tmptest")
	defer os.Remove(tmpFile.Name())
	expectedNumberOfLines := 10
	for i := 0; i < expectedNumberOfLines; i++ {
		tmpFile.WriteString("\n")
	}

	type args struct {
		file string
	}
	tests := []struct {
		name   string
		args   args
		expect int
	}{
		{
			name:   "happy path",
			args:   args{file: tmpFile.Name()},
			expect: expectedNumberOfLines,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := count_number_of_lines(tt.args.file); got != tt.expect {
				t.Errorf("count_number_of_lines() = %v, want %v", got, tt.expect)
			}
		})
	}
}
