package main

import (
	"encoding/json"
	grmon "github.com/bcicen/grmon/agent"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	grmon.Start()
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	go func() {
		obj := make(map[string]float64)
		if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
			ch <- http.StatusBadRequest
			return
		}
		// simulation of a time consuming process like writing logs into db
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
	}()

	select {
	case status := <-ch:
		w.WriteHeader(status)
	case <-time.After(200 * time.Millisecond):
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
