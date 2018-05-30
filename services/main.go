package main

import (
 "flag"
 "net/http"
 "fmt"
 "log"
 "os"

// "github.com/uber/jaeger-client-go"

 "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

    word := flag.String("word", "foo", "a string")
    port := flag.String("serve", ":3000", "server bind")
    flag.Parse()

    name, err := os.Hostname()
    if err != nil {
        panic(err)
    }

    http.Handle("/metrics", promhttp.Handler())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q, %q", *word, name)
    })

    log.Fatal(http.ListenAndServe(*port, nil))
}