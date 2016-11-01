// Hello app v1 exposes /hello HTTP endpoint which returns a random status code.
// It also adds less than a second latency to imitate a work load.
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
