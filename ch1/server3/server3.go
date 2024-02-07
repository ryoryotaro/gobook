package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)                // "/" パスに対するリクエストをhandler関数で処理
	log.Fatal(http.ListenAndServe(":8080", nil)) // 8080ポートでサーバーを起動
}

// handlerは、HTTPリクエストの情報を返します
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	// URLパラメータ"param"の値を取得して表示
	paramValue := r.URL.Query().Get("param")
	fmt.Fprintf(w, "Parameter = %q\n", paramValue)

}
