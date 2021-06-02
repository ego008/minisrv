package main

import (
	"fmt"
	"github.com/ego008/sdb"
	"log"
	"net/http"
	"time"
	"os"
	"runtime"
)

var db *sdb.DB

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(runtime.Caller(1))
	
	var err error
	db, err = sdb.Open("lcdb", nil)
	if err != nil {
		log.Panicln(err)
	}

	http.HandleFunc("/", HelloServer)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	var ts string
	key := []byte("create_time")
	if rs := db.Hget("kv", key); rs.OK() {
		ts = rs.String()
	} else {
		ts = time.Now().String()
		_ = db.Hset("kv", key, sdb.S2b(ts))
	}
	_, _ = fmt.Fprintf(w, "Hello 2, %s!", r.URL.Path[1:]+" "+sdb.B2s([]byte("test"))+", sdb key created at "+ts)
}
