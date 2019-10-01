package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// routes is a collection of our routes and
// their respective handlers
var routes = map[string]http.HandlerFunc{
	"/":       HandleRoot(),
	"/prime":  HandleRoot(),
	"/prime/": HandlePrime(),
}

func HandleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request path: ", r.URL.Path)
		// extract the path from url and write it to response
		fmt.Fprint(w, strings.SplitAfterN(r.URL.Path, "/", 2)[1])
	}
}

func HandlePrime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request path: ", r.URL.Path)
		// extract the number string from url
		nthStr := strings.SplitAfterN(r.URL.Path, "/prime/", 2)[1]
		// convert the number string to uint
		nth, err := strconv.ParseUint(nthStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "ERROR: not a positive number")
			return
		}
		prime, err := NthPrime(uint(nth))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ERROR: %s", err)
			return
		}
		fmt.Fprintf(w, "%d", prime)
	}
}

func main() {
	addr := flag.String("listen", "0.0.0.0:3333", "<IP:PORT> string, default=0.0.0.0:3333")
	flag.Parse()

	mux := http.NewServeMux()
	for route, handler := range routes {
		mux.Handle(route, handler)
	}

	log.Println("listening at", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
