// Hello app v2 is a slight modification of v1.
// We've exposed /metrics HTTP endpoint that shows Go runtime metrics.
package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/braintree/manners"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	addr := flag.String("http", "0.0.0.0:8000", "HTTP server address")
	flag.Parse()

	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		<-sigchan
		manners.Close()
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.Handle("/metrics", promhttp.Handler())
	if err := manners.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	status := doSomeWork()
	w.WriteHeader(status)
	w.Write([]byte("Hello, World!\n"))
}

func doSomeWork() int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	statusCodes := [...]int{
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusInternalServerError,
	}
	return statusCodes[rand.Intn(len(statusCodes))]
}
