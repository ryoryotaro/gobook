package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// 出力ファイルを追記モードで開く
	outputFile, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	start := time.Now()
	ch := make(chan string)
	/*
		for _, url := range os.Args[1:] {
			go fetch(url, ch) // groutine を開始
		}

		for range os.Args[1:] {
			fmt.Println(<-ch) // ch チャネルから受信
		}
	*/
	// 入力ファイルを開く
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイルオープンエラー: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var urls []string // URLを格納するスライス
	for scanner.Scan() {
		// ファイルスキャンが成功する限り、urlsというスライスに読み込んだテキストを追記
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ファイル読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	// 出力をファイルに書き込む
	fmt.Fprintf(outputFile, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
