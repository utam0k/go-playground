package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCountNumberOfLines(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "tmptest")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	expectedNumberOfLines := 10
	for i := 0; i < expectedNumberOfLines; i++ {
		_, err := tmpFile.WriteString("\n")
		if err != nil {
			log.Fatal(err)
		}
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
			if got, _ := countNumberOfLines(tt.args.file); got != tt.expect {
				t.Errorf("countNumberOfLines() = %v, want %v", got, tt.expect)
			}
		})
	}
}
