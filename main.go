package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var n uint64

func main() {
	mux := http.NewServeMux()
	srv := &http.Server{Addr: ":8080", Handler: mux}
	mux.HandleFunc("/", index)
	go func() {
		usr1 := make(chan os.Signal, 1)
		signal.Notify(usr1, syscall.SIGUSR1)
		for {
			<-usr1
			atomic.AddUint64(&n, 1)
		}
	}()
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	fmt.Fprintf(w, "%d\n", atomic.LoadUint64(&n))
}
