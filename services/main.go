package main

import (
 "flag"
 "net/http"
 "fmt"
 "os"

 log "github.com/sirupsen/logrus"
 "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

    // setup logging
    // Log as JSON instead of the default ASCII formatter.
    log.SetFormatter(&log.JSONFormatter{})

    // Output to stdout instead of the default stderr
    // Can be any io.Writer, see below for File example
    log.SetOutput(os.Stdout)

    // parse flags
    word := flag.String("word", "foo", "a string")
    port := flag.String("serve", ":3000", "server bind")
    badCanary := flag.Bool("badCanary", false, "Is this a dud?")
    flag.Parse()

    // get the hostname
    name, err := os.Hostname()
    if err != nil {
        panic(err)
    }

    // bind the prometheus metrics endpoint
    http.Handle("/metrics", promhttp.Handler())

    // healthcheck
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // send ok
        fmt.Fprintf(w, "OK")

        // log that an event happened
        log.Info("Healthcheck")
    })

    // create a simple root cath-all endpoint which returns Hello and the container information
    http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
        if (*badCanary == true) {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("500 - Oops something bad happened!"))
        } else {
            fmt.Fprintf(w, "Traefik, %q, %q", *word, name)
        }

        // log that an event happened
        log.WithFields(log.Fields{"word": word, "badCanary": badCanary}).Info("Request invoked....")
    })

    // fire it up!
    log.Fatal(http.ListenAndServe(*port, nil))
}