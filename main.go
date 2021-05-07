package main

import (
	"fmt"
	"github.com/ego008/sdb"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]+" "+sdb.B2s([]byte("test")))
}
