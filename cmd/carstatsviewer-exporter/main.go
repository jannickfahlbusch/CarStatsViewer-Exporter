package main

import (
	"flag"
	"github.com/jannickfahlbusch/CarStatsViewer-Exporter/pkg/data"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	liveDataHandler := &data.LiveDataHandler{}

	http.Handle("/live", liveDataHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Listening on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
