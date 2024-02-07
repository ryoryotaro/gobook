package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filename)
			f.Close()
		}
	}
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%d\t%s\t", len(files), line)
			for filename, _ := range files {
				fmt.Printf("%s ", filename)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]bool)
		}
		counts[line][filename] = true
	}
}
