package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// URLが "http://" で始まっていない場合、それを追加する
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// io.Copy を使用して、レスポンスボディを直接 os.Stdout に書き込む
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Status: %s\n", resp.Status) // ステータスコードの出力

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
