package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMycat(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("", "tmptest")
	defer os.Remove(tmpFile.Name())

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
			args:   args{file: "Makefile"},
			expect: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := count_number_of_lines(tt.args.file); got != tt.expect {
				t.Errorf("add() = %v, want %v", got, tt.expect)
			}
		})
	}
}
