package main

import (
	"flag"
	"log"
	"net/http"
	"pkg.jf-projects.de/owntracks/client"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/data"
)

var (
	addr     = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	otUser   = flag.String("owntracks-user", "csv-exporter", "Value the user should be set to for Owntracks locations")
	otDevice = flag.String("owntracks-device", "car", "Value the device should be set to for Owntracks locations")
	otURL    = flag.String("owntracks-url", "", "URL of the Owntracks Server. This enables/disables the Owntracks integration")
)

func main() {
	flag.Parse()

	liveDataHandler := &data.LiveDataHandler{}

	if otURL != nil && *otURL != "" {
		log.Printf("Owntracks integration enabled. Will report locations to %s with user %s and device %s\n", *otURL, *otUser, *otDevice)
		otClient := client.New(*otURL)

		liveDataHandler.Owntracks = &data.Owntracks{
			Client: otClient,
			User:   *otUser,
			Device: *otDevice,
		}
	}

	http.Handle("/live", liveDataHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/data", &data.Proxyhandler{})

	log.Printf("Listening on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
