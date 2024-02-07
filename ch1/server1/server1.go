package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) //個々のリクエストに対してhandlerが呼ばれる
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handlerは、リクエストされたURL r のPath要素を返す
// *http.request の * はポインタを表す 
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
