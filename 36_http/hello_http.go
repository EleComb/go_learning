package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		_, _ = w.Write([]byte(timeStr))
	})
	_ = http.ListenAndServe(":8080", nil)
}
