package main

import (
	"bufio"
	"fmt"
	tempconv "gobook/ch2/tempconv0"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		// 引数が与えられた場合は引数を処理する
		for _, arg := range os.Args[1:] {
			processInput(arg)
		}
	} else {
		fmt.Println("Enter temperature(e.g., 32)")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			if input == "" {
				break // 空行が入力されたら終了
			}
			processInput(input)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "reading standard input: %v\n", err)
		}
	}
}

func processInput(input string) {
	t, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: parsing input: %v\n", err)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
